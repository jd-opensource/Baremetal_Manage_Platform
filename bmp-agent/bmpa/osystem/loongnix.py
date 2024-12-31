from bmpa import constants
from bmpa.osystem.default import DefaultOsystem
from bmpa.utils import executor


class Loongnix8X(DefaultOsystem):
    def _distro(self):
        return 'loongnix-server'

    def _install_grub_to_device(self, target, root_device):
        # grub2-install has added UEFI Secure Boot verification, causing an error.
        # We now utilize the corresponding modules from the image, eliminating the need for grub2-install.
        pass

    def _kernel_install(self, target):
        super(Loongnix8X, self)._kernel_install_execute(target)

    def _configure_efi_boot_entry(self, root_device):
        super(Loongnix8X, self)._delete_efi_boot(constants.EFIBOOTMGR_LABEL)
        distro = self._distro()
        efi_file = '"\EFI\%(distro)s\BOOTLOONGARCH64.efi"' % {'distro': distro}
        cmd = ('efibootmgr', '-c', '-w', '-L', constants.EFIBOOTMGR_LABEL, '-d', root_device, '-p', '1',
               '--write-signature', '-l', efi_file)
        executor.execute(*cmd)

    def _efi_grub_config_path(self, target):
        distro = self._distro()
        return '/boot/efi/EFI/%s/grub.cfg' % distro
