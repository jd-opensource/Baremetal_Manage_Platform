from bmpa.osystem.default import DefaultOsystem


class OpenEuler20X(DefaultOsystem):
    def _distro(self):
        return 'openEuler'

    def _kernel_install(self, target):
        super(OpenEuler20X, self)._kernel_install_execute(target)
