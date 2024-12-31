import logging

from bmpa.command.common import Command
import bmpa.service.hostname as service_hostname
from bmpa.utils import validates


LOG = logging.getLogger(__name__)


class SetHostname(Command):

    def __init__(self, hostname, **kwargs):
        super(SetHostname, self).__init__(**kwargs)
        self.hostname = hostname

    def validate(self):
        validates.is_required(self.hostname)

    def execute(self):
        hostname = self.hostname

        service_hostname.set_hostname(hostname)
        LOG.info("Set hostname=%(hostname)s successfully.",
                 {"hostname": hostname})

    def execute_after(self):
        self.data["hostname"] = self.hostname
