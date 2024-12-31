import logging

from bmpa.errors import BmpError
from bmpa.osystem.default import DefaultOsystem
from bmpa.utils import chroot
from bmpa.utils.executor import ProcessExecutionError

LOG = logging.getLogger(__name__)


class Debian10x(DefaultOsystem):
    def _distro(self):
        return 'debian'

    def _get_cloudinit_conf_template_file(self):
        return "cloud_debian.tpl"

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
