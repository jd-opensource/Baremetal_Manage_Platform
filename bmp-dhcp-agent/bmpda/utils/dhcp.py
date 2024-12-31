
import logging

from bmpda import config
from bmpda.utils import executor

LOG = logging.getLogger(__name__)


def dhcp_ctl(option='start'):
    """Control dhcp status

    option: start stop restart status
    """
    dhcp_control_bin = config.dhcp_control_bin
    cmd = (dhcp_control_bin, option)
    try:
        stdout, stderr = executor.execute(*cmd,
                                          attempts_times=3)
        LOG.info("DHCP %s .", option)
        if stderr is not None:
            message = "DHCP %(option)s error:%(err)s, exit code:(exit_code)s." % {
                'option': option,
                'err': stdout}
            LOG.error(message)
            return False
    except Exception as e:
        message = "DHCP %(option)s error:%(err)s." % {
            'option': option,
            'err': str(e)
        }
        LOG.error(message)
        return False
    return True
