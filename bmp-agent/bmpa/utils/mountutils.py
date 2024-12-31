
from contextlib import contextmanager
import os

from bmpa.utils import executor
from bmpa.utils import system


@contextmanager
def mount(src, target, opts=None):
    do_mount(src, target, opts)
    yield
    do_umount(target)


def do_mount(src, target, opts=None):
    """Mount src at target with opts.

    if already mounted, return False
    """
    if opts is None:
        opts = []
    if isinstance(opts, str):
        opts = [opts]

    if is_mounted(target):
        return False

    system.ensure_dir(target)
    cmd = ['mount'] + opts + [src, target]
    executor.execute(*cmd)
    return True


def do_umount(mountpoint, recursive=False):
    """Unmount mountpoint.

    """
    mp = os.path.abspath(mountpoint)

    ret = False
    mountpoints = [
        line.split()[1]
        for line in system.load_file("/proc/mounts", decode=True).splitlines()]

    for curmp in reversed(mountpoints):
        if curmp == mp or (recursive and curmp.startswith(mp + os.path.sep)):
            executor.execute('umount', curmp, attempts_times=10,
                             delay_on_retry=True)
        if curmp == mp:
            ret = True
    return ret


def is_mounted(target):
    """Return whether or not target was mounted .

    """
    mounts = ""
    with open("/proc/mounts", "r") as fp:
        mounts = fp.read()

    for line in mounts.splitlines():
        if line.split()[1] == os.path.abspath(target):
            return True
    return False


def unmount_by_device(device):
    mounts = ""
    with open("/proc/mounts", "r") as fp:
        mounts = fp.read()
    mount_lines = mounts.splitlines()
    mount_info = [line.split() for line in mount_lines]

    for mount_point, mounted_device in mount_info:
        if mounted_device == device:
            executor.execute(*('umount', mount_point))
