import logging

import bmpda.errors as errors
from bmpda.serialize import Serializable
import bmpda.utils.validates as validates


LOG = logging.getLogger(__name__)

OK = 200

OK_STRING = "OK"
ERROR_STRING = "ERROR"

COMMAND_MAP = {
    "Heart": "bmpda.command.common.Heart",
    "PingHost": "bmpda.command.ping.PingHost",
    "DHCPConfigAddHost": "bmpda.command.dhcp_config.DHCPConfigAddHost",
    "DHCPConfigDelHost": "bmpda.command.dhcp_config.DHCPConfigDelHost",
    "DHCPConfigAddSubnet": "bmpda.command.dhcp_config.DHCPConfigAddSubnet",
    "DHCPConfigDelSubnet": "bmpda.command.dhcp_config.DHCPConfigDelSubnet",
    "TFTPConfigAddGrub": "bmpda.command.grub_config.TFTPConfigAddGrub",
}


class Command(Serializable):

    serializable_fields = ('sn', 'action', 'code', 'type', 'message',
                           'details', 'data', 'request_id', 'status')

    def __init__(self,
                 sn: str = None,
                 action: str = None,
                 code: int = OK,
                 message: str = None,
                 request_id: str = None,
                 status=OK_STRING,
                 **kwargs):

        self.sn = sn  # required
        self.action = action  # required
        self.code = code
        self.message = message
        self.type = kwargs.get('type', None)
        self.details = kwargs.get('details', None)
        self.request_id = request_id
        self.data = kwargs.get('data', None)
        self.status = status

    def validate_common_param(self):
        validates.is_required(self.sn, 'Sn')
        validates.is_required(self.action, 'Action')

    def validate(self):
        pass

    def run(self):
        try:
            try:
                self.validate_common_param()
                self.validate()
                self.execute_before()
                self.execute()
                self.__handle_success()
            except Exception as e:
                self.__handle_exception(e)
            finally:
                self.data = self.data or dict()
                self.execute_after()
        except Exception as e:
            self.__handle_exception(e)

    def execute_before(self):
        pass

    def execute(self):
        raise errors.CommandExecutionError("{}not impl execute method.".format(
            self.__class__.__name__))

    def execute_after(self):
        pass

    def __handle_success(self):
        message = "Run {} success.".format(type(self).__name__)
        self.status = OK_STRING
        self.message = message

        LOG.info(message)

    def __handle_exception(self, exec):
        self.status = ERROR_STRING
        self.code = getattr(exec, 'code', 500)
        self.type = exec.__class__.__name__
        self.message = getattr(exec, 'message', None)
        self.details = getattr(exec, 'details', None)
        LOG.error("Run {} failed:{}.".format(
            type(self).__name__, str(exec)),
            exc_info=1)
        raise exec


class SystemCommandError(Command):
    """agent to scheduler command"""

    def __init__(self, sn: str, action: str, **kwargs):
        super(SystemCommandError, self).__init__(sn=sn,
                                                 action=action,
                                                 status=ERROR_STRING,
                                                 **kwargs)


class Heart(Command):
    """agent to scheduler command"""

    def __init__(self, **kwargs):
        super(Heart, self).__init__(action=self.__class__.__name__, **kwargs)
