import logging
import platform

from bmpa import constants
from bmpa.errors import BmpError
from bmpa.osystem.default import DefaultOsystem
from bmpa.utils import chroot
from bmpa.utils.executor import ProcessExecutionError

LOG = logging.getLogger(__name__)


class Ubuntu18X(DefaultOsystem):
    def _distro(self):
        return 'ubuntu'

    def _get_cloudinit_conf_template_file(self):
        return "cloud_ubuntu.tpl"

    def _update_initramfs(self, target):
        cmd = ('update-initramfs', '-u')
        try:
            with chroot.ChrootTarget(target) as inchroot:
                inchroot.execute(*cmd)
        except ProcessExecutionError as e:
            error_msg = ('repair initram fs'
                         'failed with %(err)s' % {
                             'err': e
                         })
            LOG.error(error_msg)
            raise BmpError(error_msg)

    def _apply_swap_options(self):
        if platform.machine() == constants.MACHINE_AARCH64:
            return ('-p', '4096')
        return None


class Ubuntu20X(Ubuntu18X):
    pass
