import logging

from bmpda.command.common import Command
from bmpda.service import grub_config
from bmpda.utils import validates

LOG = logging.getLogger(__name__)


class TFTPConfigAddGrub(Command):
    """scheduler to agent command"""

    def __init__(self, mac, kernel_name, initramfs_name, boot_mode=None, arch=None, **kwargs):
        super(TFTPConfigAddGrub, self).__init__(**kwargs)

        self.mac = mac
        self.kernel_name = kernel_name
        self.initramfs_name = initramfs_name
        self.boot_mode = boot_mode
        self.arch = arch
        validates.is_required(self.mac, "mac")
        validates.is_required(self.kernel_name, "kernel_name")
        validates.is_required(self.initramfs_name, "initramfs_name")
        if boot_mode is not None:
            if boot_mode not in grub_config.boot_modes:
                raise ValueError("boot_mode must be one of %s" %
                                 grub_config.boot_modes)

        if arch is not None:
            if arch not in grub_config.archs:
                raise ValueError("arch must be one of %s" % grub_config.arch)

    def execute_before(self):
        pass

    def execute(self):
        grub_config.add_grub(mac=self.mac, kernel_name=self.kernel_name,
                             initramfs_name=self.initramfs_name, boot_mode=self.boot_mode, arch=self.arch)

    def execute_after(self):
        self.data['mac'] = self.mac


class TFTPConfigDelGrub(Command):
    """scheduler to agent command"""

    def __init__(self, mac, **kwargs):
        super(TFTPConfigDelGrub, self).__init__(**kwargs)
        self.mac = mac
        validates.is_required(self.mac, "mac")

    def execute_before(self):
        pass

    def execute(self):
        grub_config.del_grub(mac=self.mac)

    def execute_after(self):
        self.data['mac'] = self.mac
