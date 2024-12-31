import logging
import os
import time

from bmpa.utils import executor
from bmpa.utils import hardware
from bmpa.utils import lvm
from bmpa.utils import system
from bmpa.utils import udev


LOG = logging.getLogger(__name__)


def clear_holders(base_paths):
    """Clear all storage layers depending on the devices specified in 'base_paths'

    A single device or list of devices can be specified.
    Device paths can be specified either as paths in /dev or /sys/block
    Will throw OSError if any holders could not be shut down
    """
    # handle single path
    if not isinstance(base_paths, (list, tuple)):
        base_paths = [base_paths]
    LOG.info('Generating device storage trees for path(s): %s', base_paths)

    # get current holders and plan how to shut them down
    holder_trees = [gen_holders_tree(path) for path in base_paths]
    LOG.info('Current device storage tree:\n%s',
             '\n'.join(_format_holders_tree(tree) for tree in holder_trees))
    ordered_devs = _plan_shutdown_holder_trees(holder_trees)
    LOG.info('Shutdown Plan:\n%s', "\n".join(map(str, ordered_devs)))

    # run shutdown functions
    for dev_info in ordered_devs:
        dev_type = DEV_TYPES.get(dev_info['dev_type'])
        shutdown_function = dev_type.get('shutdown')
        if not shutdown_function:
            continue

        if os.path.exists(dev_info['device']):
            LOG.info("shutdown running on holder type: '%s' syspath: '%s'",
                     dev_info['dev_type'], dev_info['device'])
            shutdown_function(dev_info['device'])


def gen_holders_tree(device):
    """Generate a tree representing the current storage hirearchy above 'device'.

    """

    device = hardware.sys_block_path(device)
    kname = hardware.path_to_kname(device)

    # the holders for a device should consist of the devices in the holders/
    # dir in sysfs and any partitions on the device. this ensures that a
    # storage tree starting from a disk will include all devices holding the
    # disk's partitions
    holders = hardware.get_holders(device)
    holder_paths = ([hardware.sys_block_path(h) for h in holders] +
                    hardware.get_sysfs_partitions(device))
    # the DEV_TYPE registry contains a function under the key 'ident' for each
    # device type entry that returns true if the device passed to it is of the
    # correct type. there should never be a situation in which multiple
    # identify functions return true. therefore, it will always work to take
    # the device type with the first identify function that returns true as the
    # device type for the current device. in the event that no identify
    # functions return true, the device will be treated as a disk
    # (DEFAULT_DEV_TYPE). the identify function for disk never returns true.
    # the next() builtin in python will not raise a StopIteration exception if
    # there is a default value defined

    dev_type = next((k for k, v in DEV_TYPES.items() if v['ident'](device)),
                    DEFAULT_DEV_TYPE)
    return {
        'device': device, 'dev_type': dev_type, 'name': kname,
        'holders': [gen_holders_tree(h) for h in holder_paths],
    }


def is_mpath_partition(devpath, info=None):
    """Check if a device is a multipath partition, returns boolean. """
    result = False
    if devpath.startswith('/dev/dm-'):
        if not info:
            info = udev.udevadm_info(devpath)
        if 'DM_PART' in info and 'DM_MPATH' in info:
            result = True

    LOG.debug("%s is multipath device partition? %s", devpath, result)
    return result


def get_dmsetup_uuid(device):
    """Get the dm uuid for a specified dmsetup device

    """
    blockdev = hardware.sysfs_to_devpath(device)
    cmd = ('dmsetup', 'info', blockdev, '-C', '-o', 'uuid',
           '--noheadings')
    stdout, _ = executor.execute(*cmd)
    return stdout.strip()


def _define_handlers_registry():
    """Returns instantiated dev_types

    """
    return {
        'lvm': {'shutdown': _shutdown_lvm, 'ident': _identify_lvm},
        'disk': {'shutdown': wipe_superblock, 'ident': lambda x: False}
    }


def _identify_lvm(device):
    """Determine if specified device is a lvm device.

    """
    return (hardware.path_to_kname(device).startswith('dm') and
            get_dmsetup_uuid(device).startswith('LVM'))


def _format_holders_tree(holders_tree):
    """Draw a nice dirgram of the holders tree.

    """
    # spacer styles based on output of 'tree --charset=ascii'
    spacers = (('`-- ', ' ' * 4), ('|-- ', '|' + ' ' * 3))

    def format_tree(tree):
        """Format entry and any subentries.

        """
        result = [tree['name']]
        holders = tree['holders']
        for (holder_no, holder) in enumerate(holders):
            spacer_style = spacers[min(len(holders) - (holder_no + 1), 1)]
            subtree_lines = format_tree(holder)
            for (line_no, line) in enumerate(subtree_lines):
                result.append(spacer_style[min(line_no, 1)] + line)
        return result

    return '\n'.join(format_tree(holders_tree))


def _plan_shutdown_holder_trees(holders_trees):
    """Plan best order to shut down holders in, taking into account high level

    storage layers that may have many devices below them

    returns a sorted list of descriptions of storage config entries including
    their path in /sys/block and their dev type

    can accept either a single storage tree or a list of storage trees assumed
    to start at an equal place in storage hirearchy (i.e. a list of trees
    starting from disk)
    """
    # holds a temporary registry of holders to allow cross references
    # key = device sysfs path, value = {} of priority level, shutdown function
    reg = {}

    # normalize to list of trees
    if not isinstance(holders_trees, (list, tuple)):
        holders_trees = [holders_trees]

    # sort the trees to ensure we generate a consistent plan
    holders_trees = sorted(holders_trees, key=lambda x: x['device'])

    def htree_level(tree):
        if len(tree['holders']) == 0:
            return 0
        return 1 + sum(htree_level(holder) for holder in tree['holders'])

    def flatten_holders_tree(tree, level=0):
        """Add entries from holders tree to registry with level key corresponding

        to how many layers from raw disks the current device is at
        """
        device = tree['device']
        device_level = htree_level(tree)

        # always go with highest level if current device has been
        # encountered already. since the device and everything above it is
        # re-added to the registry it ensures that any increase of level
        # required here will propagate down the tree
        # this handles a scenario like mdadm + bcache, where the backing
        # device for bcache is a 3rd level item like mdadm, but the cache
        # device is 1st level (disk) or second level (partition), ensuring
        # that the bcache item is always considered higher level than
        # anything else regardless of whether it was added to the tree via
        # the cache device or backing device first
        if device in reg:
            level = max(reg[device]['level'], level) + 1

        else:
            # first time device to registry, assume the larger value of the
            # current level or the length of its dependencies.
            level = max(device_level, level)

        reg[device] = {'level': level, 'device': device,
                       'dev_type': tree['dev_type']}

        # handle holders above this level
        for holder in tree['holders']:
            flatten_holders_tree(holder, level=level + 1)

    # flatten the holders tree into the registry
    for holders_tree in holders_trees:
        flatten_holders_tree(holders_tree)

    def devtype_order(dtype):
        """Return the order in which we want to clear device types, higher

         value should be cleared first.
        :param: dtype: string. A device types name from the holders registry,
                see _define_handlers_registry()
        :returns: integer
        """
        dev_type_order = [
            'disk', 'lvm']
        return 1 + dev_type_order.index(dtype)

    # return list of entry dicts with greatest htree depth. The 'level' value
    # indicates the number of additional devices that are "below" this device.
    # Devices must be cleared in descending 'level' value.  For devices which
    # have the same 'level' value, we sort within the 'level' by devtype order.
    return [reg[k]
            for k in sorted(reg, reverse=True,
                            key=lambda x: (reg[x]['level'],
                                           devtype_order(reg[x]['dev_type'])))]


def _shutdown_lvm(device):
    """Shutdown specified lvm device.

    """
    device = hardware.sys_block_path(device)
    # lvm devices have a dm directory that containes a file 'name' containing
    # '{volume group}-{logical volume}'. The volume can be freed using lvremove
    name_file = os.path.join(device, 'dm', 'name')
    lvm_name = system.load_file(name_file).strip()
    (vg_name, lv_name) = lvm.split_lvm_name(lvm_name)
    vg_lv_name = "%s/%s" % (vg_name, lv_name)
    devname = "/dev/" + vg_lv_name

    # wipe contents of the logical volume first
    LOG.info('Wiping lvm logical volume: %s', devname)
    hardware.quick_zero(devname)

    # remove the logical volume
    LOG.debug('using "lvremove" on %s', vg_lv_name)
    executor.execute(*['lvremove', '--force', '--force', vg_lv_name])

    # if that was the last lvol in the volgroup, get rid of volgroup
    if len(lvm.get_lvols_in_volgroup(vg_name)) == 0:
        pvols = lvm.get_pvols_in_volgroup(vg_name)
        executor.execute(*['vgremove', '--force', '--force',
                         vg_name], check_exit_code=[0, 5])

        # wipe the underlying physical volumes
        for pv in pvols:
            LOG.info('Wiping lvm physical volume: %s', pv)
            hardware.quick_zero(pv)


def wipe_superblock(device):
    """Wrapper for block.wipe_volume compatible with shutdown function interface."""
    blockdev = hardware.sysfs_to_devpath(device)
    if not hardware.is_online(blockdev):
        LOG.debug("Device is not online (size=0), so skipping:"
                  " '%s'", blockdev)
        return

    # gather any partitions
    partitions = hardware.get_sysfs_partitions(device)

    _wipe_superblock(blockdev)

    # if we had partitions, make sure they've been removed
    if partitions:
        LOG.debug('%s had partitions, issuing partition reread', device)
        retries = [.5, .5, 1, 2, 5, 7]
        for attempt, wait in enumerate(retries):
            try:
                # only rereadpt on wiped device
                hardware.rescan_block_devices(devices=[blockdev])
                # may raise IOError, OSError due to wiped partition table
                curparts = hardware.get_sysfs_partitions(device)
                if len(curparts) == 0:
                    return
            except (IOError, OSError):
                if attempt + 1 >= len(retries):
                    raise

            LOG.debug("%s partitions still present, rereading pt"
                      " (%s/%s).  sleeping %ss before retry",
                      device, attempt + 1, len(retries), wait)
            time.sleep(wait)


def _wipe_superblock(blockdev, exclusive=True):
    """No checks, just call wipe_volume"""

    retries = [1, 3, 5, 7]
    LOG.info('wiping superblock on %s', blockdev)
    for attempt, wait in enumerate(retries):
        LOG.debug('wiping %s attempt %s/%s',
                  blockdev, attempt + 1, len(retries))
        try:
            hardware.wipe_volume(blockdev, mode='superblock', exclusive=exclusive)
            LOG.debug('successfully wiped device %s on attempt %s/%s',
                      blockdev, attempt + 1, len(retries))
            return
        except OSError:
            if attempt + 1 >= len(retries):
                raise
            else:
                LOG.debug("wiping device '%s' failed on attempt"
                          " %s/%s.  sleeping %ss before retry",
                          blockdev, attempt + 1, len(retries), wait)
                time.sleep(wait)


DEFAULT_DEV_TYPE = 'disk'
DATA_DESTROYING_HANDLERS = [lambda:None]
DEV_TYPES = _define_handlers_registry()
