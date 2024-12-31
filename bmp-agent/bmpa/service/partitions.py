import logging
import uuid

from bmpa import constants
from bmpa.errors import DeviceNotFound
from bmpa.log import log_complete
from bmpa.service import node
from bmpa.utils import clear_holders
from bmpa.utils import device_hints
from bmpa.utils import disk_partitioner
from bmpa.utils import hardware
from bmpa.utils import utils

LOG = logging.getLogger(__name__)


@log_complete
def make_partition(volumes: list, boot_mode: str,
                   auto_create_efi: bool, auto_create_bios_grub: bool) -> dict:
    boot_mode = utils.get_boot_mode(boot_mode)
    node.set_boot_mode(boot_mode)
    res = []
    for volume in volumes:
        device_name = _get_device_name(volume)
        if not device_name:
            raise DeviceNotFound('Partition device not found.')
        keep_data = volume.get('keep_data', True)
        if keep_data:
            continue
        clear_holders.clear_holders(device_name)
        hardware.destroy_disk_metadata(device_name)
        hardware.erase_block_device(device_name)
        partition = _parse_partition(
            volume, boot_mode, auto_create_efi, auto_create_bios_grub)
        part_dict = disk_partitioner.make_partitions(
            device_name=partition['device_name'], partitions=partition['partitions'],
            boot_mode=partition['boot_mode'], disk_label=partition['disk_label'])
        for key, value in part_dict.items():
            node.blockdev_id_to_devpath_dict_add_item(
                key, value)
        disk_hints = device_hints.get_device_by_hints(
            {'name': partition['device_name']})
        res_volume = dict()
        res_volume['disk_hints'] = disk_hints
        res.append(res_volume)
    return res


def _get_device_name(volume) -> str:
    if volume.get('is_root_device', False):
        return node.get_root_devcie_name()
    if volume.get('is_data_device', False):
        return node.get_data_devcie_name()
    disk_hints = volume.get('disk_hints', None)
    device = device_hints.get_device_by_hints(disk_hints)
    return device.name


def _parse_partition(volume: dict, boot_mode: str,
                     auto_create_efi: bool, auto_create_bios_grub: bool) -> list:
    partitions = []
    is_root_device = volume.get('is_root_device', None)
    disk_label = volume.get('disk_label', None)
    device_name = _get_device_name(volume)
    if is_root_device:
        if auto_create_efi and boot_mode == constants.UEFI:
            LOG.info('Auto create efi.')
            partition = dict()
            partition['id'] = uuid.uuid4().hex
            partition['size'] = hardware.EFI_SYSTEM_PARTITION_SIZE
            partition['fs_type'] = 'fat32'
            partition['boot_flag'] = 'boot'
            partitions.append(partition)
            node.set_auto_create_efi(auto_create_efi)
            node.set_efi_partition_id(partition['id'])

        if auto_create_bios_grub and boot_mode == constants.BIOS and disk_label == constants.GPT:
            LOG.info('Auto crete bios grub.')
            partition = dict()
            partition['id'] = uuid.uuid4().hex
            partition['size'] = hardware.BISO_BOOT_PARTITION_SIZE
            partition['boot_flag'] = 'bios_grub'
            partitions.append(partition)
    partitions.extend(volume['partitions'])
    disk_partitions = dict()
    disk_partitions['device_name'] = device_name
    disk_partitions['partitions'] = partitions
    disk_partitions['boot_mode'] = boot_mode
    disk_partitions['disk_label'] = disk_label
    LOG.info('Parse partitions:%s', disk_partitions)
    return disk_partitions
