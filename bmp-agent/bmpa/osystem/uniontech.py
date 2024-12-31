import platform

from bmpa import constants
from bmpa.osystem.default import DefaultOsystem


class UnionTech20X(DefaultOsystem):
    def _distro(self):
        if platform.machine() == constants.MACHINE_AARCH64:
            return 'UnionTech'
        return 'uos'

    def _efi_grub_config_path(self, target):
        if platform.machine() == constants.MACHINE_AARCH64:
            return super(UnionTech20X, self)._efi_grub_config_path(target)
        return '/boot/efi/EFI/%s/grub.cfg' % self._distro()
