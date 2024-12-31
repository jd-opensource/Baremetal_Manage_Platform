from contextlib import contextmanager
import errno
import json
import logging
import os
import platform
import re
import shlex
import stat
import sys
import time

import netaddr
import pint
import psutil
import pyudev
import tenacity

from bmpa import constants
from bmpa.errors import BlockDeviceError, BmpError, BlockDeviceEraseError
from bmpa.serialize import SerializableComparable
from bmpa.utils import executor
from bmpa.utils.executor import ProcessExecutionError, UnknownArgumentError
from bmpa.utils import multipath
from bmpa.utils import system
from bmpa.utils import utils


LOG = logging.getLogger(__name__)

EFI_SYSTEM_PARTITION_SIZE = 512
BISO_BOOT_PARTITION_SIZE = 1

NVME_CLI_FORMAT_SUPPORTED_FLAG = 0b10
NVME_CLI_CRYPTO_FORMAT_SUPPORTED_FLAG = 0b100

GPT_SIZE_SECTORS = 33

SECTOR_SIZE_BYTES = 512


class BlockDevice(SerializableComparable):
    serializable_fields = ('name', 'model', 'size', 'rotational', 'wwn',
                           'serial', 'vendor', 'wwn_with_extension',
                           'wwn_vendor_extension', 'hctl', 'by_path')

    def __init__(self,
                 name,
                 model,
                 size,
                 rotational,
                 wwn=None,
                 serial=None,
                 vendor=None,
                 wwn_with_extension=None,
                 wwn_vendor_extension=None,
                 hctl=None,
                 by_path=None):
        self.name = name
        self.model = model
        self.size = size
        self.rotational = rotational
        self.wwn = wwn
        self.serial = serial
        self.vendor = vendor
        self.wwn_with_extension = wwn_with_extension
        self.wwn_vendor_extension = wwn_vendor_extension
        self.hctl = hctl
        self.by_path = by_path

    def __repr__(self):
        return "BlockDevice %s" % self.serialize()


def list_all_block_devices(block_type='disk'):
    """List all physical block devices

    The switches we use for lsblk: P for KEY="value" output, b for size output
    in bytes, d to exclude dependent devices (like md or dm devices), i to
    ensure ascii characters only, and o to specify the fields/columns we need.

    Broken out as its own function to facilitate custom hardware managers that
    don't need to subclass GenericHardwareManager.

    :param block_type: Type of block device to find
    :return: A list of BlockDevices
    """
    udev_settle()

    # map device names to /dev/disk/by-path symbolic links that points to it

    by_path_mapping = {}

    disk_by_path_dir = '/dev/disk/by-path'

    try:
        paths = os.listdir(disk_by_path_dir)

        for path in paths:
            path = os.path.join(disk_by_path_dir, path)
            # Turn possibly relative symbolic link into absolute
            devname = os.path.join(disk_by_path_dir, os.readlink(path))
            devname = os.path.abspath(devname)
            by_path_mapping[devname] = path

    except OSError as e:
        LOG.warning(
            "Path %(path)s is inaccessible, /dev/disk/by-path/* "
            "version of block device name is unavailable "
            "Cause: %(error)s", {
                'path': disk_by_path_dir,
                'error': e
            })

    columns = ['KNAME', 'MODEL', 'SIZE', 'ROTA', 'TYPE']
    report = executor.execute('lsblk',
                              '-Pbdi',
                              '-o{}'.format(','.join(columns)),
                              check_exit_code=[0])[0]
    lines = report.split('\n')
    context = pyudev.Context()

    devices = []
    for line in lines:
        if line == '':
            continue
        device = {}
        # Split into KEY=VAL pairs
        vals = shlex.split(line)
        for key, val in (v.split('=', 1) for v in vals):
            device[key] = val.strip()
        # Ignore block types not specified
        if device.get('TYPE') != block_type:
            LOG.warn(
                "TYPE did not match. Wanted: {!r} but found: {!r}",
                block_type, line)
            continue

        # Ensure all required columns are at least present, even if blank
        missing = set(columns) - set(device)
        if missing:
            raise BlockDeviceError(
                '%s must be returned by lsblk.' % ', '.join(sorted(missing)))

        name = os.path.join('/dev', device['KNAME'])

        try:
            udev = pyudev.Device.from_device_file(context, name)
        # pyudev started raising another error in 0.18
        except (ValueError, EnvironmentError, pyudev.DeviceNotFoundError) as e:
            LOG.warning(
                "Device %(dev)s is inaccessible, skipping... "
                "Error: %(error)s", {
                    'dev': name,
                    'error': e
                })
            extra = {}
        else:
            # Since lsblk only supports
            # returning the short serial we are using
            # ID_SERIAL_SHORT here to keep compatibility with the
            # bash deploy ramdisk
            extra = {
                key: udev.get('ID_%s' % udev_key)
                for key, udev_key in [('wwn', 'WWN'), (
                    'serial', 'SERIAL_SHORT'
                ), ('wwn_with_extension', 'WWN_WITH_EXTENSION'
                    ), ('wwn_vendor_extension', 'WWN_VENDOR_EXTENSION')]
            }

        # Newer versions of the lsblk tool supports
        # HCTL as a parameter but let's get it from sysfs to avoid breaking
        # old distros.
        try:
            extra['hctl'] = os.listdir('/sys/block/%s/device/scsi_device' %
                                       device['KNAME'])[0]
        except (OSError, IndexError):
            LOG.warning(
                'Could not find the SCSI address (HCTL) for '
                'device %s. Skipping', name)

        # Not all /dev entries are pointed to from /dev/disk/by-path
        by_path_name = by_path_mapping.get(name)

        devices.append(
            BlockDevice(name=name,
                        model=device['MODEL'],
                        size=int(device['SIZE']),
                        rotational=bool(int(device['ROTA'])),
                        vendor=get_device_info(
                            device['KNAME'], 'block', 'vendor'),
                        by_path=by_path_name,
                        **extra))
    return devices


def udev_settle():
    """Wait for the udev event queue to settle.

    Wait for the udev event queue to settle to make sure all devices
    are detected once the machine boots up.

    :return: True on success, False otherwise.
    """
    LOG.info('Waiting until udev event queue is empty')
    try:
        executor.execute('udevadm', 'settle')
    except ProcessExecutionError as e:
        LOG.warning(
            'Something went wrong when waiting for udev '
            'to settle. Error: %s', e)
        return False
    else:
        return True


def _lsblock(args=None):
    """Get lsblock data as dict."""
    # lsblk  --help | sed -n '/Available/,/^$/p' |
    #     sed -e 1d -e '$d' -e 's,^[ ]\+,,' -e 's, .*,,' | sort
    keys = ['ALIGNMENT', 'DISC-ALN', 'DISC-GRAN', 'DISC-MAX', 'DISC-ZERO',
            'FSTYPE', 'GROUP', 'KNAME', 'LABEL', 'LOG-SEC', 'MAJ:MIN',
            'MIN-IO', 'MODE', 'MODEL', 'MOUNTPOINT', 'NAME', 'OPT-IO', 'OWNER',
            'PHY-SEC', 'RM', 'RO', 'ROTA', 'RQ-SIZE', 'SCHED', 'SIZE', 'STATE',
            'TYPE', 'UUID']
    if args is None:
        args = []
    args = [x.replace('!', '/') for x in args]

    # in order to avoid a very odd error with '-o' and all output fields above
    # we just drop one.  doesn't really matter which one.
    keys.remove('SCHED')
    basecmd = ['lsblk', '--noheadings', '--bytes', '--pairs',
               '--output=' + ','.join(keys)]
    out, _ = executor.execute(basecmd + list(args))
    out = out.replace('!', '/')
    return _lsblock_pairs_to_dict(out)


def _lsblock_pairs_to_dict(lines):
    """Parse lsblock output and convert to dict."""
    ret = {}
    for line in lines.splitlines():
        toks = shlex.split(line)
        cur = {}
        for tok in toks:
            k, v = tok.split("=", 1)
            if k == 'MAJ_MIN':
                k = 'MAJ:MIN'
            else:
                k = k.replace('_', '-')
            cur[k] = v
        # use KNAME, as NAME may include spaces and other info,
        # for example, lvm decices may show 'dm0 lvm1'
        cur['device_path'] = get_dev_name_entry(cur['KNAME'])[1]
        ret[cur['KNAME']] = cur
    return ret


def get_dev_name_entry(devname):
    """Convert device name to path in /dev."""
    bname = devname.split('/dev/')[-1]
    return (bname, "/dev/" + bname)


def get_unused_blockdev_info():
    """return a list of unused block devices.

    These are devices that do not have anything mounted on them.
    """

    # get a list of top level block devices, then iterate over it to get
    # devices dependent on those.  If the lsblk call for that specific
    # call has nothing 'MOUNTED", then this is an unused block device
    bdinfo = _lsblock(['--nodeps'])
    unused = {}
    for devname, data in bdinfo.items():
        cur = _lsblock([data['device_path']])
        mountpoints = [x for x in cur if cur[x].get('MOUNTPOINT')]
        if len(mountpoints) == 0:
            unused[devname] = data
    return unused


def rescan_block_devices(devices=None, warn_on_fail=True):
    """Run 'blockdev --rereadpt' for all block devices not currently mounted."""
    if not devices:
        unused = get_unused_blockdev_info()
        devices = []
        for devname, data in unused.items():
            if data.get('RM') == "1":
                continue
            if data.get('RO') != "0" or data.get('TYPE') != "disk":
                continue
            devices.append(data['device_path'])

    if not devices:
        LOG.info("no devices found to rescan")
        return

    # blockdev needs /dev/ parameters, convert if needed
    cmd = ['blockdev', '--rereadpt'] + [dev if dev.startswith('/dev/')
                                        else sysfs_to_devpath(dev)
                                        for dev in devices]
    try:
        executor.execute(*cmd)
    except ProcessExecutionError as e:
        if warn_on_fail:
            # FIXME: its less than ideal to swallow this error, but until
            # we fix LP: #1489521 we kind of need to.
            LOG.warn(
                "Error rescanning devices, possibly known issue LP: #1489521")
            # Reformatting the exception output so as to not trigger
            # vmtest scanning for Unexepected errors in install logfile
            LOG.warn("cmd: %s\nstdout:%s\nstderr:%s\nexit_code:%s", e.cmd,
                     e.stdout, e.stderr, e.exit_code)

    udev_settle()

    return


class SystemVendorInfo(SerializableComparable):

    serializable_fields = ('product_name', 'serial_number', 'manufacturer')

    def __init__(self, product_name, serial_number, manufacturer):
        self.product_name = product_name
        self.serial_number = serial_number
        self.manufacturer = manufacturer

    def __repr__(self):
        return "SystemVendorInfo %s" % self.serialize()


def get_system_vendor_info():
    try:
        cmd = ['dmidecode', '-t', 'system']
        lines = executor.execute(*cmd)[0]
        system_vendor_info = {
            k.strip().lower(): v.strip()
            for k, v in (line.split(':', 1) for line in lines.split('\n')
                         if (line.strip() and line.count(":") == 1))
        }
    except (ProcessExecutionError, OSError, ValueError) as e:
        LOG.warning('Could not retrieve vendor info from dmidecode: %e', e)
        system_vendor_info = {}
    return SystemVendorInfo(
        product_name=system_vendor_info.get('product name', ''),
        serial_number=system_vendor_info.get('serial number', ''),
        manufacturer=system_vendor_info.get('manufacturer', ''))


def shred_block_device(device_name):
    """Erase a block device using shred.

    :param node: Ironic node info.
    :param block_device: a BlockDevice object to be erased
    :returns: True if the erase succeeds, False if it fails for any reason
    """
    try:
        if is_nvme_device(device_name):
            _nvme_erase(device_name)
            return True

    except BlockDeviceEraseError:
        pass

    try:
        executor.execute('shred', '--force', '--zero', '--verbose',
                         '--iterations=0', '--size=100M',
                         device_name)
    except (ProcessExecutionError, OSError) as e:
        msg = "Erasing block device %(device_name)s failed with error %(err)s"
        LOG.error(msg, {'device_name': device_name, 'err': e})
        return False

    return True


def erase_block_device(device_name):
    if shred_block_device(device_name):
        return

    msg = 'Unable to erase block device %(device_name)s: device is unsupported.' % {
        'device_name': device_name
    }
    LOG.error(msg)
    raise BmpError(msg)


def _nvme_erase(device_name):
    """Attempt to clean the NVMe using the most secure supported method

        :param block_device: a BlockDevice object
        :return: True if cleaning operation succeeded, False if it failed
        :raises: BlockDeviceEraseError
        """

    # check if crypto format is supported
    try:
        LOG.info("Attempting to fetch NVMe capabilities for device %s",
                 device_name)
        nvme_info, _e = executor.execute('nvme', 'id-ctrl',
                                         device_name, '-o', 'json')
        nvme_info = json.loads(nvme_info)

    except ProcessExecutionError as e:
        msg = (("Failed to fetch NVMe capabilities for device {}: {}").format(
            device_name, e))
        LOG.error(msg)
        raise BlockDeviceEraseError(msg)

    # execute format with crypto option (ses=2) if supported
    # if crypto is unsupported use user-data erase (ses=1)
    if nvme_info:
        # Check if the device supports NVMe format at all. This info
        # is in "oacs" section of nvme-cli id-ctrl output. If it does,
        # set format mode to 1 (this is passed as -s <mode> parameter
        # to nvme-cli later)
        fmt_caps = nvme_info['oacs']
        if fmt_caps & NVME_CLI_FORMAT_SUPPORTED_FLAG:
            # Given the device supports format, check if crypto
            # erase format mode is supported and pass it to nvme-cli
            # instead
            crypto_caps = nvme_info['fna']
            if crypto_caps & NVME_CLI_CRYPTO_FORMAT_SUPPORTED_FLAG:
                format_mode = 2  # crypto erase
            else:
                format_mode = 1  # user-data erase
        else:
            msg = ('nvme-cli did not return any supported format modes '
                   'for device: {device}').format(device=device_name)
            LOG.error(msg)
            raise BlockDeviceEraseError(msg)
    else:
        # If nvme-cli output is empty, raise an exception
        msg = ('nvme-cli did not return any information '
               'for device: {device}').format(device=device_name)
        LOG.error(msg)
        raise BlockDeviceEraseError(msg)

    try:
        LOG.debug(
            "Attempting to nvme-format %s using secure format mode "
            "(ses) %s", device_name, format_mode)
        executor.execute('nvme', 'format', device_name, '-s',
                         format_mode)
        LOG.info(
            "nvme-cli format for device %s (ses= %s ) completed "
            "successfully.", device_name, format_mode)
        return True

    except ProcessExecutionError as e:
        msg = (("Failed to nvme format device {}: {}").format(device_name, e))
        raise BlockDeviceEraseError(msg)


def wait_for_disk_to_become_available(device):
    """Wait for a disk device to become available.

    Waits for a disk device to become available for use by
    waiting until all process locks on the device have been
    released.

    Timeout and iteration settings come from the configuration
    options used by the in-library disk_partitioner:
    ``check_device_interval`` and ``check_device_max_retries``.

    :params device: The path to the device.
    :raises: IronicException If the disk fails to become
        available.
    """
    pids = ['']
    stderr = ['']
    interval = 1
    max_retries = 20

    def _wait_for_disk():
        # A regex is likely overkill here, but variations in fuser
        # means we should likely use it.
        fuser_pids_re = re.compile(r'\d+')

        # There are 'psmisc' and 'busybox' versions of the 'fuser' program. The
        # 'fuser' programs differ in how they output data to stderr.  The
        # busybox version does not output the filename to stderr, while the
        # standard 'psmisc' version does output the filename to stderr.  How
        # they output to stdout is almost identical in that only the PIDs are
        # output to stdout, with the 'psmisc' version adding a leading space
        # character to the list of PIDs.
        try:
            # fuser returns a non-zero return code if none of
            # the specified files is accessed.
            # user does not report LVM devices as in use
            # unless the LVM device-mapper device is the
            # device that is directly polled.
            # The -m flag allows fuser to reveal data about
            # mounted filesystems, which should be considered
            # busy/locked. That being said, it is not used
            # because busybox fuser has a different behavior.
            # fuser outputs a list of found PIDs to stdout.
            # All other text is returned via stderr, and the
            # output to a terminal is merged as a result.
            out, err = executor.execute('fuser',
                                        device,
                                        check_exit_code=[0, 1],
                                        run_as_root=True)

            if not out and not err:
                return True

            stderr[0] = err
            # findall() returns a list of matches, or an empty list if no matches
            pids[0] = fuser_pids_re.findall(out)

        except ProcessExecutionError as exc:
            LOG.warning(
                'Failed to check the device %(device)s with fuser:'
                ' %(err)s', {
                    'device': device,
                    'err': exc
                })
        return False

    retry = tenacity.retry(retry=tenacity.retry_if_result(lambda r: not r),
                           stop=tenacity.stop_after_attempt(max_retries),
                           wait=tenacity.wait_fixed(interval),
                           reraise=True)
    try:
        retry(_wait_for_disk)()
    except tenacity.RetryError:
        if pids[0]:
            raise BmpError(
                ('Processes with the following PIDs are holding '
                 'device %(device)s: %(pids)s. '
                 'Timed out waiting for completion.') % {
                     'device': device,
                     'pids': ', '.join(pids[0])
                })
        else:
            raise BmpError(
                ('Fuser exited with "%(fuser_err)s" while checking '
                 'locks for device %(device)s. Timed out waiting for '
                 'completion.') % {
                     'device': device,
                     'fuser_err': stderr[0]
                })


def get_dev_block_size(dev):
    """Get the device size in 512 byte sectors."""
    block_sz, _ = executor.execute('blockdev',
                                   '--getsz',
                                   dev,
                                   run_as_root=True,
                                   check_exit_code=[0])
    return int(block_sz)


def destroy_disk_metadata(dev):
    """Destroy metadata structures on node's disk.

    Ensure that node's disk magic strings are wiped without zeroing the
    entire drive. To do this we use the wipefs tool from util-linux.
    :param dev: Path for the device to work on.
    """
    LOG.debug("Start destroy disk metadata .")
    try:
        executor.execute('wipefs',
                         '--force',
                         '--all',
                         dev,
                         run_as_root=True,
                         use_standard_locale=True)
    except ProcessExecutionError as e:
        # Check if --force option is supported for wipefs,
        # if not, we should try without it.
        if '--force' in str(e):
            executor.execute('wipefs', '--all', dev)

    # Overwrite the Primary GPT, catch very small partitions (like EBRs)
    dd_device = 'of=%s' % dev
    dd_count = 'count=%s' % GPT_SIZE_SECTORS
    dev_size = get_dev_block_size(dev)
    LOG.info("dev size %s", dev_size)
    if dev_size < GPT_SIZE_SECTORS:
        dd_count = 'count=%s' % dev_size
    executor.execute('dd',
                     'bs=512',
                     'if=/dev/zero',
                     dd_device,
                     dd_count,
                     run_as_root=True,
                     use_standard_locale=True)

    # Overwrite the Secondary GPT, do this only if there could be one
    if dev_size > GPT_SIZE_SECTORS:
        gpt_backup = dev_size - GPT_SIZE_SECTORS
        dd_seek = 'seek=%i' % gpt_backup
        dd_count = 'count=%s' % GPT_SIZE_SECTORS
        executor.execute('dd',
                         'bs=512',
                         'if=/dev/zero',
                         dd_device,
                         dd_count,
                         dd_seek,
                         run_as_root=True,
                         use_standard_locale=True)

    # Go ahead and let sgdisk run as well.
    executor.execute('sgdisk',
                     '-Z',
                     dev,
                     run_as_root=True,
                     use_standard_locale=True)

    try:
        wait_for_disk_to_become_available(dev)
    except BmpError as e:
        raise BmpError(
            ('Destroying metadata failed on device %(device)s. '
             'Error: %(error)s') % {
                 'device': dev,
                 'error': e
            })

    LOG.info("Disk metadata on %(dev)s successfully destroyed ",
             {'dev': dev})


def is_nvme_device(device_name):
    """Check whether the device path belongs to an NVMe drive."""
    return "/dev/nvme" in device_name


def is_block_device(dev):
    """Check whether a device is block or not."""
    attempts = 3
    for attempt in range(attempts):
        try:
            s = os.stat(dev)
        except OSError as e:
            LOG.warn(
                "Unable to stat device %(dev)s. Attempt %(attempt)d "
                "out of %(total)d. Error: %(err)s", {
                    "dev": dev,
                    "attempt": attempt + 1,
                    "total": attempts,
                    "err": e
                })
            time.sleep(1)
        else:
            return stat.S_ISBLK(s.st_mode)
    msg = ("Unable to stat device %(dev)s after attempting to verify "
           "%(attempts)d times.") % {
               'dev': dev,
               'attempts': attempts
    }
    LOG.error(msg)
    raise BmpError(msg)


def partprobe(device, attempts=None):
    """Probe partitions on the given device.

    :param device: The block device containing paritions that is attempting
                   to be updated.
    :param attempts: Number of attempts to run partprobe, the default is read
                     from the configuration.
    :return: True on success, False otherwise.
    """
    if attempts is None:
        attempts = 3

    try:
        executor.execute('partprobe',
                         device,
                         run_as_root=True,
                         attempts_times=attempts)
    except (UnknownArgumentError, ProcessExecutionError, OSError) as e:
        LOG.warning(
            "Unable to probe for partitions on device %(device)s, "
            "the partitioning table may be broken. Error: %(error)s", {
                'device': device,
                'error': e
            })
        return False
    else:
        return True


def trigger_device_rescan(device, attempts=None):
    """Sync and trigger device rescan.

    Disk partition performed via parted, when performed on a ramdisk
    do not have to honor the fsync mechanism. In essence, fsync is used
    on the file representing the block device, which falls to the kernel
    filesystem layer to trigger a sync event. On a ramdisk using ramfs,
    this is an explicit non-operation.

    As a result of this, we need to trigger a system wide sync operation
    which will trigger cache to flush to disk, after which partition changes
    should be visible upon re-scan.

    When ramdisks are not in use, this also helps ensure that data has
    been safely flushed across the wire, such as on iscsi connections.

    :param device: The block device containing paritions that is attempting
                   to be updated.
    :param attempts: Number of attempts to run partprobe, the default is read
                     from the configuration.
    :return: True on success, False otherwise.
    """
    LOG.info('Explicitly calling sync to force buffer/cache flush')
    executor.execute('sync')
    # Make sure any additions to the partitioning are reflected in the
    # kernel.
    udev_settle()
    partprobe(device, attempts=attempts)
    try:
        # Also verify that the partitioning is correct now.
        executor.execute('sgdisk', '-v', device, run_as_root=True)
    except ProcessExecutionError as exc:
        LOG.warning(
            'Failed to verify partition tables on device %(dev)s: '
            '%(err)s', {
                'dev': device,
                'err': exc
            })
        return False
    else:
        return True


def get_device_info(dev, devclass, field):
    """Get the device info according to device class and field."""
    try:
        devname = os.path.basename(dev)
        with open('/sys/class/%s/%s/device/%s' % (devclass, devname, field),
                  'r') as f:
            return f.read().strip()
    except IOError:
        LOG.warning("Can't find field %(field)s for "
                    "device %(dev)s in device class %(class)s",
                    {'field': field, 'dev': dev, 'class': devclass})


def get_blockdev_for_partition(devpath, strict=True):
    """Find the parent device for a partition.

    returns a tuple of the parent block device and the partition number
    if device is not a partition, None will be returned for partition number
    """
    # normalize path
    rpath = os.path.realpath(devpath)

    # convert an entry in /dev/ to parent disk and partition number
    # if devpath is a block device and not a partition, return (devpath, None)
    base = '/sys/class/block'

    # input of /dev/vdb, /dev/disk/by-label/foo, /sys/block/foo,
    # /sys/block/class/foo, or just foo
    syspath = os.path.join(base, path_to_kname(devpath))

    # don't need to try out multiple sysfs paths as path_to_kname handles cciss
    if strict and not os.path.exists(syspath):
        raise OSError("%s had no syspath (%s)" % (devpath, syspath))

    if rpath.startswith('/dev/dm-'):
        parent_info = multipath.mpath_partition_to_mpath_id_and_partnumber(
            rpath)
        if parent_info is not None:
            mpath_id, ptnum = parent_info
            return os.path.realpath('/dev/mapper/' + mpath_id), ptnum

    ptpath = os.path.join(syspath, "partition")
    if not os.path.exists(ptpath):
        return (rpath, None)

    ptnum = system.load_file(ptpath).rstrip()

    # for a partition, real syspath is something like:
    # /sys/devices/pci0000:00/0000:00:04.0/virtio1/block/vda/vda1
    rsyspath = os.path.realpath(syspath)
    disksyspath = os.path.dirname(rsyspath)

    diskmajmin = system.load_file(os.path.join(disksyspath, "dev")).rstrip()
    diskdevpath = os.path.realpath("/dev/block/%s" % diskmajmin)

    # diskdevpath has something like 253:0
    # and udev has put links in /dev/block/253:0 to the device name in /dev/
    return (diskdevpath, ptnum)


def get_sysfs_partitions(device):
    """Get a list of sysfs paths for partitions.

    under a block device accepts input as a device kname,
    sysfs path, or dev path
    returns empty list if no partitions available
    """
    sysfs_path = sys_block_path(device)
    LOG.debug("sysfs_path:%s", sysfs_path)
    return [sys_block_path(kname) for kname in os.listdir(sysfs_path)
            if os.path.exists(os.path.join(sysfs_path, kname, 'partition'))]


def sysfs_partition_data(blockdev=None, sysfs_path=None):
    # given block device or sysfs_path, return a list of tuples
    # of (kernel_name, number, offset, size)
    if blockdev:
        blockdev = os.path.normpath(blockdev)
        sysfs_path = sys_block_path(blockdev)
    elif sysfs_path:
        # use normpath to ensure that paths with trailing slash work
        sysfs_path = os.path.normpath(sysfs_path)
        blockdev = os.path.join('/dev', os.path.basename(sysfs_path))
    else:
        raise ValueError("Blockdev and sysfs_path cannot both be None")

    sysfs_prefix = sysfs_path
    (parent, partnum) = get_blockdev_for_partition(blockdev)
    if partnum:
        sysfs_prefix = sys_block_path(parent)
        partnum = int(partnum)

    ptdata = []
    for part_sysfs in get_sysfs_partitions(sysfs_prefix):
        data = {}
        for sfile in ('partition', 'start', 'size'):
            dfile = os.path.join(part_sysfs, sfile)
            if not os.path.isfile(dfile):
                continue
            data[sfile] = int(system.load_file(dfile))
        if partnum is None or data['partition'] == partnum:
            ptdata.append((
                path_to_kname(part_sysfs),
                data['partition'],
                data['start'] * SECTOR_SIZE_BYTES,
                data['size'] * SECTOR_SIZE_BYTES,
            ))

    return ptdata


def kname_to_path(kname):
    """Converts a kname to a path in /dev.

    taking special devices and unusual naming schemes into account
    """
    # if given something that is already a dev path, return it
    if os.path.exists(kname) and is_block_device(kname):
        return os.path.realpath(kname)
    # adding '/dev' to path is not sufficient to handle cciss devices and
    # possibly other special devices which have not been encountered yet
    path = os.path.realpath(os.sep.join(['/dev'] + kname.split('!')))
    # make sure path we get is correct
    if not (os.path.exists(path) and is_block_device(path)):
        raise OSError(
            'Could not get path to dev from kname: {}.'.format(kname))
    return path


def sysfs_to_devpath(sysfs_path):
    """Convert a path in /sys/class/block to a path in /dev."""
    path = kname_to_path(path_to_kname(sysfs_path))
    if not is_block_device(path):
        raise ValueError('could not find blockdev for sys path: {}'
                         .format(sysfs_path))
    return path


def dev_path(devname):
    """Convert device name to path in /dev."""
    if devname.startswith('/dev/'):
        # it could be something like /dev/mapper/mpatha-part2
        return os.path.realpath(devname)
    else:
        return '/dev/' + devname


def is_online(device):
    """Check if device is online."""
    sys_path = sys_block_path(device)
    device_size = system.load_file(
        os.path.join(sys_path, 'size'))
    # a block device should have non-zero size to be usable
    return int(device_size) > 0


def wipe_volume(path, mode="superblock", exclusive=True):
    """Wipe a volume/block device."""
    if mode == "superblock":
        quick_zero(path, partitions=False, exclusive=exclusive)
    elif mode == "superblock-recursive":
        quick_zero(path, partitions=True, exclusive=exclusive)
    else:
        raise ValueError("wipe mode %s not supported" % mode)


def quick_zero(path, partitions=True, exclusive=True):
    """Zero 1M at front, 1M at end, and 1M at front

    if this is a block device and partitions is true, then
    zero 1M at front and end of each partition.
    """
    buflen = 1024
    count = 1024
    zero_size = buflen * count
    offsets = [0, -zero_size]
    is_block = is_block_device(path)
    LOG.info('%s is block: %s', path, is_block)
    if not (is_block or os.path.isfile(path)):
        raise ValueError("%s: not an existing file or block device.", path)

    pt_names = []
    if partitions and is_block:
        ptdata = sysfs_partition_data(path)
        for kname, ptnum, start, size in ptdata:
            pt_names.append((dev_path(kname), kname, ptnum))
        pt_names.reverse()

    for (pt, kname, ptnum) in pt_names:
        LOG.info('Wiping path: dev:%s kname:%s partnum:%s',
                 pt, kname, ptnum)
        quick_zero(pt, partitions=False)

    LOG.info("Wiping 1M on %s at offsets %s.", path, offsets)
    return zero_file_at_offsets(path,
                                offsets,
                                buflen=buflen,
                                count=count,
                                exclusive=exclusive)


def zero_file_at_offsets(path,
                         offsets,
                         buflen=1024,
                         count=1024,
                         strict=False,
                         exclusive=True):
    """Write zeros to file at specified offsets."""

    bmsg = "{path} (size={size}): "
    m_short = bmsg + "{tot} bytes from {offset} > size."
    m_badoff = bmsg + "invalid offset {offset}."
    if not strict:
        m_short += " Shortened to {wsize} bytes."
        m_badoff += " Skipping."

    buf = b'\0' * buflen
    tot = buflen * count
    msg_vals = {'path': path, 'tot': buflen * count}

    # allow caller to control if we require exclusive open
    with exclusive_open(path, exclusive=exclusive) as fp:
        # get the size by seeking to end.
        fp.seek(0, 2)
        size = fp.tell()
        msg_vals['size'] = size

        for offset in offsets:
            if offset < 0:
                pos = size + offset
            else:
                pos = offset
            msg_vals['offset'] = offset
            msg_vals['pos'] = pos
            if pos > size or pos < 0:
                if strict:
                    raise ValueError(m_badoff.format(**msg_vals))
                else:
                    LOG.debug(m_badoff.format(**msg_vals))
                    continue

            msg_vals['wsize'] = size - pos
            if pos + tot > size:
                if strict:
                    raise ValueError(m_short.format(**msg_vals))
                else:
                    LOG.debug(m_short.format(**msg_vals))
            fp.seek(pos)
            for i in range(count):
                pos = fp.tell()
                if pos + buflen > size:
                    fp.write(buf[0:size - pos])
                else:
                    fp.write(buf)


@contextmanager
def exclusive_open(path, exclusive=True):
    """Obtain an exclusive file-handle to the file/device specified unless

    caller specifics exclusive=False.
    """
    mode = 'rb+'
    fd = None
    if not os.path.exists(path):
        raise ValueError("No such file at path: %s" % path)

    flags = os.O_RDWR
    if exclusive:
        flags += os.O_EXCL
    try:
        fd = os.open(path, flags)
        try:
            fd_needs_closing = True
            with os.fdopen(fd, mode) as fo:
                yield fo
            fd_needs_closing = False
        except OSError:
            LOG.exception("Failed to create file-object from fd")
            raise
        finally:
            # python2 leaves fd open if there os.fdopen fails
            if fd_needs_closing and sys.version_info.major == 2:
                os.close(fd)
    except OSError:
        LOG.error("Failed to exclusively open path: %s", path)
        holders = get_holders(path)
        LOG.error('Device holders with exclusive access: %s', holders)
        mount_points = list_device_mounts(path)
        LOG.error('Device mounts: %s', mount_points)
        raise


def list_device_mounts(device):
    # return mount entry if device is in /proc/mounts
    mounts = ""
    with open("/proc/mounts", "r") as fp:
        mounts = fp.read()

    dev_mounts = []
    for line in mounts.splitlines():
        if line.split()[0] == device:
            dev_mounts.append(line)
    return dev_mounts


def get_holders(device_name):
    """Look up any block device holders, return list of knames.

    """
    # block.sys_block_path works when given a /sys or /dev path
    sysfs_path = sys_block_path(device_name)
    # get holders
    holders_path = os.path.join(sysfs_path, 'holders')
    holders = os.listdir(holders_path)
    LOG.info("The device '%s' had holders: %s.", device_name, holders)
    return holders


def path_to_kname(path):
    """Converts a path in /dev or a path in /sys/block to the device kname.

    """
    if os.path.sep in path:
        path = os.path.realpath(path)
    device_kname = os.path.basename(path)
    return device_kname


def sys_block_path(device_name, additional_path=None, strict=True):
    """Get path to device in /sys/class/block.

    """
    path_components = ['/sys/class/block']
    device_name = os.path.normpath(device_name)
    if device_name.startswith('/dev/') and not os.path.exists(device_name):
        LOG.warning('The devname %s does not exist.', device_name)

    path_components.append(path_to_kname(device_name))

    if additional_path is not None:
        path_components.append(additional_path)
    path = os.sep.join(path_components)
    if strict and not os.path.exists(path):
        err = OSError(
            "The devname '{}' did not have existing syspath '{}'.".format(
                device_name, path))
        err.errno = errno.ENOENT
        raise err

    return os.path.normpath(path)


def mkfs(fs, path, label=None, options=None):
    """Format a file or block device.

    :param fs: Filesystem type (examples include 'swap', 'ext3', 'ext4'
               'btrfs', etc.)
    :param path: Path to file or block device to format
    :param label: Volume label to use
    """
    if fs == 'fat32':
        fs = 'vfat'

    if fs == 'swap':
        args = ['mkswap']
    else:
        args = ['mkfs', '-t', fs]

    if options:
        args.extend(options)
    # add -F to force no interactive execute on non-block device.
    if fs in ('ext3', 'ext4'):
        args.extend(['-F'])
    if fs in ('xfs'):
        args.extend(['-f'])

    if label:
        if fs in ('msdos', 'vfat'):
            label_opt = '-n'
        else:
            label_opt = '-L'
        args.extend([label_opt, label])
    args.append(path)
    try:
        executor.execute(*args,
                         run_as_root=True,
                         use_standard_locale=True)
        executor.execute('sync')
    except ProcessExecutionError as e:
        if os.strerror(errno.ENOENT) in e.stdout:
            LOG.exception(
                'Failed to make file system. '
                'File system %s is not supported.', fs)
            raise BmpError(("Failed to create a file system. "
                            "File system %s is not supported.") % fs)
        else:
            LOG.exception(
                'Failed to create a file system '
                'in %(path)s. Error: %(error)s', {
                    'path': path,
                    'error': e
                })


def blkid(devs=None, cache=True):
    """Get data about block devices from blkid and convert to dict.

    """
    if devs is None:
        devs = []

    # 14.04 blkid reads undocumented /dev/.blkid.tab
    # man pages mention /run/blkid.tab and /etc/blkid.tab
    if not cache:
        cfiles = ("/run/blkid/blkid.tab", "/dev/.blkid.tab", "/etc/blkid.tab")
        for cachefile in cfiles:
            if os.path.exists(cachefile):
                os.unlink(cachefile)

    cmd = ['blkid', '-o', 'full']
    cmd.extend(devs)
    # blkid output is <device_path>: KEY=VALUE
    # where KEY is TYPE, UUID, PARTUUID, LABEL
    out, _ = executor.execute(*cmd)
    data = {}
    for line in out.splitlines():
        curdev, curdata = line.split(":", 1)
        data[curdev] = dict(tok.split('=', 1)
                            for tok in shlex.split(curdata))
    LOG.info('Blkid info:%s', data)
    return data


def get_labelled_partition(device_path, label, attempts_times=1):
    """Check and return if partition with given label exists

    :param device_path: The device path.
    :param label: Partition label
    :raises: InstanceDeployFailure, if any disk partitioning related
        commands fail.
    :returns: block device file for partition if it exists; otherwise it
              returns None.
    """
    LOG.info("Get device path:%s by %s", device_path, label)
    temp_attempts_times = attempts_times
    while attempts_times > 0:
        temp_attempts_times -= 1
        partprobe(device_path)
        wait_for_disk_to_become_available(device_path)
        try:
            output, _ = executor.execute('lsblk',
                                         '-Po',
                                         'name,label',
                                         device_path,
                                         check_exit_code=[0, 1],
                                         use_standard_locale=True,
                                         run_as_root=True)

        except (UnknownArgumentError, ProcessExecutionError, OSError) as e:
            msg = (('Failed to retrieve partition labels on disk %(disk)s '
                    'Error: %(error)s') % {
                        'disk': device_path,
                        'error': e
            })
            LOG.error(msg)
            raise BmpError(msg)

        found_part = None
        if output:
            for device in output.split('\n'):
                dev = {
                    key: value
                    for key, value in (v.split('=', 1)
                                       for v in shlex.split(device))
                }
                if not dev:
                    continue
                if dev['LABEL'].upper() == label.upper():
                    if found_part:
                        found_2 = '/dev/%(part)s' % {
                            'part': dev['NAME'].strip()}
                        found = [found_part, found_2]
                        raise BmpError(
                            ('More than one partition with label "%(label)s" '
                                'exists on device %(device)s: '
                                '%(found)s.') % {
                                    'label': label,
                                    'device': device_path,
                                    'found': ' and '.join(found)
                            })
                    found_part = '/dev/%(part)s' % {
                        'part': dev['NAME'].strip()}

        if found_part:
            LOG.info("Get labelled partition, found part%s", found_part)
            return found_part
        else:
            attempts_period = 0.5 * 2**(attempts_times -
                                        temp_attempts_times)
            LOG.info("Retry to run after %.2f seconds." % attempts_period)
            time.sleep(attempts_period)
            continue


def _fix_gpt_structs(device):
    """Checks backup GPT data structures and moves them to end of the device

    :param device: The device path.
    :raises: InstanceDeployFailure, if any disk partitioning related
        commands fail.
    """
    try:
        output, _err = executor.execute('sgdisk',
                                        '-v',
                                        device,
                                        run_as_root=True)

        search_str = "it doesn't reside\nat the end of the disk"
        if search_str in output:
            executor.execute('sgdisk',
                             '-e',
                             device,
                             run_as_root=True)
    except (UnknownArgumentError, ProcessExecutionError, OSError) as e:
        msg = (('Failed to fix GPT data structures on disk %(disk)s '
                'Error: %(error)s') % {
                    'disk': device,
                    'error': e
        })
        LOG.error(msg)
        raise BmpError(msg)


def fix_gpt_partition(device):
    """Fix GPT partition

    Fix GPT table information when image is written to a disk which
    has a bigger extend (e.g. 30GB image written on a 60Gb physical disk).

    :param device: The device path.
    :raises: InstanceDeployFailure if exception is caught.
    """
    try:
        disk_is_gpt_partitioned = _is_disk_gpt_partitioned(device)
        if disk_is_gpt_partitioned:
            _fix_gpt_structs(device)
    except Exception as e:
        msg = (('Failed to fix GPT partition on disk %(disk)s '
                'Error: %(error)s') % {
                    'disk': device,
                    'error': e
        })
        LOG.error(msg)
        raise BmpError(msg)


def _is_disk_gpt_partitioned(device):
    """Checks if the disk is GPT partitioned

    :param device: The device path.
    :raises: InstanceDeployFailure, if any disk partitioning related
        commands fail.
    :returns: Boolean. Returns True if disk is GPT partitioned
    """
    try:
        stdout, _ = executor.execute(
            'blkid',
            '-p',
            '-o',
            'value',
            '-s',
            'PTTYPE',
            device,
            use_standard_locale=True,
            run_as_root=True)
    except (UnknownArgumentError, ProcessExecutionError, OSError) as e:
        msg = (('Failed to retrieve partition table type for disk %(disk)s '
                'Error: %(error)s') % {
                    'disk': device,
                    'error': e
        })
        LOG.error(msg)
        raise BmpError(msg)

    return (stdout.lower().strip() == 'gpt')


class CPU(SerializableComparable):

    serializable_fields = ('model_name', 'frequency', 'count', 'architecture',
                           'flags')

    def __init__(self, model_name, frequency, count, architecture, flags=None):
        self.model_name = model_name
        self.frequency = frequency
        self.count = count
        self.architecture = architecture
        self.flags = flags or []

    def __repr__(self):
        return "CPU %s" % self.serialize()


class Memory(SerializableComparable):

    serializable_fields = ('total', 'physical_mb')

    # physical = total + kernel binary + reserved space

    def __init__(self, total, physical_mb=None):
        self.total = total
        self.physical_mb = physical_mb

    def __repr__(self):
        return "Memory %s" % self.serialize()


def is_device(interface_name):
    device_path = '/sys/class/net/{}/device'.format(interface_name)
    return os.path.exists(device_path)


def get_bios_given_nic_name(interface_name):
    """Collect the BIOS given NICs name.

    This function uses the biosdevname utility to collect the BIOS given
    name of network interfaces.
    The collected data is added to the network interface inventory with an
    extra field named ``biosdevname``.
    :param interface_name: list of names of node's interfaces.
    :return: the BIOS given NIC name of node's interfaces or default
             as None.
    """
    try:

        stdout, _ = executor.execute('biosdevname', '-i',
                                     interface_name)
        return stdout.rstrip('\n')
    except OSError:
        LOG.warning("Executable 'biosdevname' not found")
        return interface_name
    except ProcessExecutionError as e:
        LOG.warning('Biosdevname get error:%s', str(e))
        return interface_name


def get_system_lshw_dict():
    """Get a dict representation of the system from lshw

    Retrieves a json representation of the system from lshw and converts
    it to a python dict

    :return: A python dict from the lshw json output
    """
    out, _ = executor.execute('lshw',
                              '-quiet',
                              '-json')
    res = json.loads(out)
    if isinstance(res, list):
        return res[0]
    else:
        return res


def get_cpus():
    lines = executor.execute('lscpu')[0]
    cpu_info = {
        k.strip().lower(): v.strip()
        for k, v in (line.split(':', 1) for line in lines.split('\n')
                     if line.strip())
    }
    # Current CPU frequency can be different from maximum one on modern
    # processors
    freq = cpu_info.get('cpu max mhz', cpu_info.get('cpu mhz'))
    flags = []
    colum_grep = '^flags'
    if platform.machine() == constants.MACHINE_AARCH64:
        colum_grep = '^Features'

    out = executor.try_execute('grep', '-Em1', colum_grep,
                               '/proc/cpuinfo')
    if out:
        try:
            # Example output (much longer for a real system):
            # flags           : fpu vme de pse
            flags = out[0].strip().split(':', 1)[1].strip().split()
        except (IndexError, ValueError):
            LOG.warning('Malformed CPU flags information: %s', out)
    else:
        LOG.warning('Failed to get CPU flags')
    return CPU(
        model_name=cpu_info.get('model name'),
        frequency=freq,
        # this includes hyperthreading cores
        count=int(cpu_info.get('cpu(s)')),
        architecture=cpu_info.get('architecture'),
        flags=flags)


def get_memory():
    # psutil returns a long, so we force it to an int
    try:
        total = int(psutil.virtual_memory().total)
    except Exception:
        # This is explicitly catching all exceptions. We want to catch any
        # situation where a newly upgraded psutil would fail, and instead
        # print an error instead of blowing up the stack on IPA.
        total = None
        LOG.exception(("Cannot fetch total memory size using psutil "
                       "version %s"), psutil.version_info[0])
    sys_dict = None
    try:
        sys_dict = get_system_lshw_dict()
    except (ProcessExecutionError, OSError, ValueError) as e:
        LOG.warning('Could not get real physical RAM from lshw: %s', e)
        physical = None
    else:
        physical = _calc_memory(sys_dict)
        if not physical:
            LOG.warning('Did not find any physical RAM')
    return Memory(total=total, physical_mb=physical)


_MEMORY_ID_RE = re.compile(r'^memory(:\d+)?$')
UNIT_CONVERTER = pint.UnitRegistry(filename=None)
UNIT_CONVERTER.define('bytes = []')
UNIT_CONVERTER.define('MB = 1048576 bytes')


def _calc_memory(sys_dict):
    physical = 0
    for sys_child in sys_dict['children']:
        if sys_child['id'] != 'core':
            continue
        for core_child in sys_child['children']:
            if not _MEMORY_ID_RE.match(core_child['id']):
                continue
            if core_child.get('size'):
                value = ("%(size)s %(units)s" % core_child)
                physical += int(UNIT_CONVERTER(value).to('MB').magnitude)
            else:
                for bank in core_child.get('children', ()):
                    if bank.get('size'):
                        value = ("%(size)s %(units)s" % bank)
                        physical += int(
                            UNIT_CONVERTER(value).to('MB').magnitude)
    return physical


def get_bmc_address():
    """Attempt to detect BMC IP address

    :return: IP address of lan channel or 0.0.0.0 in case none of them is
             configured properly
    """
    # These modules are rarely loaded automatically
    executor.try_execute('modprobe', 'ipmi_msghandler')
    executor.try_execute('modprobe', 'ipmi_devintf')
    executor.try_execute('modprobe', 'ipmi_si',
                                     'unload_when_empty=0')

    try:
        # From all the channels 0-15, only 1-7 can be assigned to different
        # types of communication media and protocols and effectively used
        for channel in range(1, 8):
            try:
                out, _ = executor.execute(
                    "ipmitool lan print %(channel)s |grep -E \"IP Address              : \""
                    % {"channel": channel},
                    shell=True)
            except Exception:
                continue
            if out.startswith('Invalid channel'):
                continue
            if ":" not in out:
                continue
            out = out.split(':')[1].strip()

            try:
                netaddr.IPAddress(out)
            except ValueError:
                LOG.warning('Invalid IP address: %s', out)
                continue
            except netaddr.AddrFormatError:
                LOG.warning('Invalid IP address: %s', out)
                continue

            # In case we get 0.0.0.0 on a valid channel, we need to keep
            # querying
            if out == '0.0.0.0':
                continue
            return out

    except (ProcessExecutionError, OSError) as e:
        # Not error, because it's normal in virtual environment
        LOG.warning('Cannot get BMC address: %s', e)
        return

    return '0.0.0.0'


class BootInfo(SerializableComparable):

    serializable_fields = ('current_boot_mode', 'pxe_interface')

    def __init__(self, current_boot_mode, pxe_interface=None):
        self.current_boot_mode = current_boot_mode
        self.pxe_interface = pxe_interface

    def __repr__(self):
        return "BootInfo %s" % self.serialize()


def get_boot_info():
    boot_mode = constants.UEFI if utils.is_uefi_bootable() else constants.BIOS
    LOG.debug('The current boot mode is %s', boot_mode)
    return BootInfo(current_boot_mode=boot_mode)


class PlatformInfo(SerializableComparable):

    serializable_fields = ('machine', )

    def __init__(self, machine):
        self.machine = machine

    def __repr__(self):
        return "PlatformInfo %s" % self.serialize()


def get_platform_info():
    machine = platform.machine()
    return PlatformInfo(machine=machine)
