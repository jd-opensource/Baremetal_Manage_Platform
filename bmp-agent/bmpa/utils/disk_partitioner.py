
import logging

from bmpa.errors import BmpError
from bmpa.utils import executor
from bmpa.utils import hardware

LOG = logging.getLogger(__name__)

ALL_UNALLOCATED_SPACE = 999999999


class DiskPartitioner(object):
    def __init__(self,
                 device,
                 disk_label='msdos',
                 alignment='optimal',
                 start=1):
        """A convenient wrapper around the parted tool.

        :param device: The device path.
        :param disk_label: The type of the partition table. Valid types are:
                           "bsd", "dvh", "gpt", "loop", "mac", "msdos",
                           "pc98", or "sun".
        :param alignment: Set alignment for newly created partitions.
                          Valid types are: none, cylinder, minimal and
                          optimal.

        """
        self._device = device
        self._disk_label = disk_label
        self._alignment = alignment
        self._partitions = []
        self.start = start

    def _exec(self, *args):
        executor.execute('parted',
                         '-a',
                         self._alignment,
                         '-s',
                         self._device,
                         '--',
                         'unit',
                         'MiB',
                         *args,
                         check_exit_code=[0],
                         use_standard_locale=True,
                         run_as_root=True)

    def add_partition(self,
                      size,
                      part_type='primary',
                      fs_type='',
                      boot_flag=None,
                      extra_flags=None):
        """Add a partition.

        :param size: The size of the partition in MiB.
        :param part_type: The type of the partition. Valid values are:
                          primary, logical, or extended.
        :param fs_type: The filesystem type. Valid types are: ext2, fat32,
                        fat16, HFS, linux-swap, NTFS, reiserfs, ufs.
                        If blank (''), it will create a Linux native
                        partition (83).
        :param boot_flag: Boot flag that needs to be configured on the
                          partition. Ignored if None. It can take values
                          'bios_grub', 'boot'.
        :param extra_flags: List of flags to set on the partition. Ignored
                            if None.
        :returns: The partition number.

        """
        if part_type is None:
            part_type = 'primary'
        if fs_type is None:
            fs_type = ''
        self._partitions.append({
            'size': size,
            'type': part_type,
            'fs_type': fs_type,
            'boot_flag': boot_flag,
            'extra_flags': extra_flags
        })
        return len(self._partitions)

    def get_partitions(self):
        """Get the partitioning layout.

        :returns: An iterator with the partition number and the
                  partition layout.

        """
        return enumerate(self._partitions, 1)

    def commit(self):
        """Write to the disk."""
        LOG.debug("Committing partitions to disk.")
        if self._disk_label is not None:
            cmd_args = ['mklabel', self._disk_label]
        else:
            cmd_args = []
        # Lead in with 1MiB to allow room for the
        # partition table itself.
        start = self.start
        for num, part in self.get_partitions():
            if part['size'] == ALL_UNALLOCATED_SPACE:
                end = "100%"
            else:
                end = start + part['size']
            cmd_args.extend([
                'mkpart', part['type'], part['fs_type'],
                str(start),
                str(end)
            ])
            if part['boot_flag']:
                cmd_args.extend(['set', str(num), part['boot_flag'], 'on'])
            if part['extra_flags']:
                for flag in part['extra_flags']:
                    cmd_args.extend(['set', str(num), flag, 'on'])
            start = end

        self._exec(*cmd_args)

        try:
            hardware.wait_for_disk_to_become_available(
                self._device)
        except BmpError as e:
            raise BmpError(
                ('Disk partitioning failed on device %(device)s. '
                 'Error: %(error)s') % {
                     'device': self._device,
                     'error': e
                })


def make_partitions(device_name,
                    partitions,
                    boot_mode="bios",
                    disk_label=None,
                    start=1,
                    exist_part_num=0):
    """Partition the disk device.

    Create partitions for root, boot and system partitions on a disk device.

    :param device_name: Path for the device to work on.
    :param partitions: partitions.
    :param boot_mode: bios or uefi.
    param disk_label: gpt or None.
    param start: partition start Mib.
    :returns: A dictionary containing the partition type as Key and partition
        path as Value for the partitions created by this method.

    """
    LOG.debug("Starting to partition the disk device: %(device_name)s ",
              {'device_name': device_name})
    part_dict = {}

    # prevent resetting partitions when adding new partitions to existing ones
    if disk_label is None and exist_part_num == 0:
        disk_label = 'gpt' if boot_mode == 'uefi' else 'msdos'

    dp = DiskPartitioner(device_name,
                         disk_label=disk_label,
                         start=start)

    for partion in partitions:
        LOG.debug("Add partition (%(size)d MiB) to device: %(dev)s ", {
            'dev': device_name,
            'size': partion['size']
        })
        part_type = partion.get('partion', None)
        fs_type = partion.get('fs_type', None)
        if fs_type == 'swap':
            fs_type = 'linux-swap'
        boot_flag = partion.get('boot_flag', None)
        extra_flags = partion.get('extra_flags', None)
        part_num = dp.add_partition(
            partion['size'], part_type, fs_type, boot_flag, extra_flags) + exist_part_num
        part_dict[partion['id']] = partition_index_to_path(
            device_name, part_num)

    dp.commit()
    hardware.trigger_device_rescan(device_name)
    return part_dict


def partition_index_to_path(device_name: str, index: int):
    """Guess a partition path based on its device and index.

    :param device_name: Device path.
    :param index: Partition index.
    """
    if is_last_char_digit(device_name):
        part_template = '%sp%d'
    else:
        part_template = '%s%d'
    return part_template % (device_name, index)


def is_last_char_digit(device_name):
    """check whether device name ends with a digit"""
    if len(device_name) >= 1:
        return device_name[-1].isdigit()
    return False
