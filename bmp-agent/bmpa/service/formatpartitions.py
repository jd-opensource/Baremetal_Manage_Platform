

import logging

from bmpa import constants
from bmpa.log import log_complete
from bmpa.osystem.factory import get_osystem
from bmpa.service import node
from bmpa.utils import utils

LOG = logging.getLogger(__name__)


@log_complete
def formatpartitions(partitions):
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    LOG.info('OS:%s', osystem.__class__.__name__)
    partitions = _parse_partitions(partitions)
    root_partition_path = osystem.formatpartitions(partitions)
    if root_partition_path:
        node.set_root_partition(root_partition_path)


def _parse_partitions(partitions):
    for index, partition in enumerate(partitions):
        volume = partition['volume']
        device_path = node.blockdev_id_to_devpath(volume)
        partitions[index]['device_path'] = device_path
    boot_mode = utils.get_boot_mode(node.get_boot_mode())
    LOG.info('boot mode:%s', boot_mode)
    auto_create_efi = node.get_auto_create_efi()
    LOG.info('auto create efi:%s', auto_create_efi)
    if auto_create_efi and boot_mode == constants.UEFI:
        efi_partition = dict()
        efi_partition['label'] = constants.EFI_LABLE
        efi_partition['fs_type'] = constants.EFI_FSTYPE
        efi_partition_id = node.get_efi_partition_id()
        device_path = node.blockdev_id_to_devpath(efi_partition_id)
        efi_partition['device_path'] = device_path
        partitions = [efi_partition] + partitions
        LOG.info('Format auto efi partition %s.', partitions)
    return partitions
