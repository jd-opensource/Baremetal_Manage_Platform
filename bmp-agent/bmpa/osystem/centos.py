import logging
import platform

from bmpa import constants
from bmpa.osystem.default import DefaultOsystem
from bmpa.utils import chroot
from bmpa.utils import executor


LOG = logging.getLogger(__name__)


class CentOS7X(DefaultOsystem):

    def _distro(self):
        return 'centos'

    def _set_network_service(self, target):
        with chroot.ChrootTarget(target) as inchroot:
            inchroot.execute('systemctl', 'enable', 'NetworkManager')


class CentOS8X(CentOS7X):

    def _set_network_service(self, target):
        pass

    def _parse_user_data(self, data):
        return data.replace("#!/usr/bin/env python",
                            "#!/usr/bin/env python3")

    def _install_grub_to_device(self, target, root_device):
        # grub2-install has added UEFI Secure Boot verification, causing an error.
        # We now utilize the corresponding modules from the image, eliminating the need for grub2-install.
        pass

    def _kernel_install(self, target):
        super(CentOS8X, self)._kernel_install_execute(target)

    def _configure_efi_boot_entry(self, root_device):
        super(CentOS8X, self)._delete_efi_boot(constants.EFIBOOTMGR_LABEL)
        distro = self._distro()
        efi_file = '"\EFI\%(distro)s\grubx64.efi"' % {'distro': distro}
        if platform.machine() == constants.MACHINE_AARCH64:
            efi_file = '"\EFI\%(distro)s\grubaa64.efi"' % {'distro': distro}
        cmd = ('efibootmgr', '-c', '-w', '-L', constants.EFIBOOTMGR_LABEL, '-d', root_device, '-p', '1',
               '--write-signature', '-l', efi_file)
        executor.execute(*cmd)

    def _efi_grub_config_path(self, target):
        return '/boot/efi/EFI/centos/grub.cfg'
