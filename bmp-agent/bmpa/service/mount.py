import logging
import operator
import os

from bmpa import constants
from bmpa.errors import BmpError, DeviceNotFound
from bmpa.log import log_complete
from bmpa.service import node
from bmpa.utils import device_hints
from bmpa.utils.executor import ProcessExecutionError
from bmpa.utils import hardware
from bmpa.utils import mountutils
from bmpa.utils import utils

LOG = logging.getLogger(__name__)

EFI_LABEL = 'l_efi'
EFI_MOUNTPOINT = '/boot/efi'


@log_complete
def mount_partitions(auto_mount_efi, mounts):
    target = node.get_target()
    mounts = _parse_mounts(auto_mount_efi, mounts)
    # Mounting needs to be done according to the hierarchy of mountpoint,
    # with '/' being mounted first。
    mounts = sorted(mounts, key=operator.itemgetter('mountpoint'))
    LOG.info('Mounts:%s', mounts)
    node.set_mounts(mounts)
    try:
        for mount in mounts:
            if mount['mountpoint'] == 'swap':
                continue
            mount_point = mount['mountpoint']
            partition_path = mount['partition_path']
            partition_mount_point = os.path.join(target,
                                                 mount_point.lstrip("/"))

            if not os.path.exists(partition_mount_point):
                os.makedirs(partition_mount_point)

            mountutils.do_mount(src=partition_path,
                                target=partition_mount_point)

    except ProcessExecutionError as e:
        error_msg = ('Mount partition error, failed with %(err)s.' % {
            'err': e
        })
        LOG.error(error_msg)
        raise BmpError(error_msg)


def _parse_mounts(auto_mount_efi, mounts):
    blkid = hardware.blkid()
    for index, mount in enumerate(mounts):
        device_name = _get_device_name(mount)
        if not device_name:
            LOG.error("Device not found, volume:%s", mount)
            raise DeviceNotFound()
        label = mount['label']
        LOG.info('Mount info, device name:%s, lable:%s', device_name, label)
        partition_path = hardware.get_labelled_partition(
            device_name, label, attempts_times=10)
        LOG.info('Partition path:%s.', partition_path)
        mounts[index]['partition_path'] = partition_path
        mounts[index]['fs_type'] = blkid[partition_path]['TYPE']

    boot_mode = node.get_boot_mode()
    boot_mode = utils.get_boot_mode(boot_mode)
    LOG.info('Parse mounts boot_mode:%s.', boot_mode)
    LOG.info('Parse mounts auto_mount_efi:%s.', auto_mount_efi)
    if auto_mount_efi and boot_mode == constants.UEFI:
        LOG.info('Parse auto mount efi.')
        mount = dict()
        mount['label'] = EFI_LABEL
        mount['is_root_device'] = True
        mount['mountpoint'] = EFI_MOUNTPOINT
        device_name = node.get_root_devcie_name()
        LOG.info('Mount info, device anme:%s, lable:%s',
                 device_name, mount['label'])
        partition_path = hardware.get_labelled_partition(
            device_name, mount['label'], attempts_times=10)
        mount['partition_path'] = partition_path
        mount['fs_type'] = constants.EFI_FSTYPE
        mounts.append(mount)
    return mounts


def _get_device_name(volume) -> str:
    if volume.get('is_root_device', False):
        return node.get_root_devcie_name()
    if volume.get('is_data_device', False):
        return node.get_data_devcie_name()
    disk_hints = volume.get('disk_hints', None)
    device = device_hints.get_device_by_hints(disk_hints)
    LOG.info("Device name:%s", device.name)
    return device.name
