import logging

from bmpda.command.common import Command
from bmpda.utils import executor
from bmpda.utils import validates

LOG = logging.getLogger(__name__)


class Ping(Command):
    """Detect mq communication.

       Avoid high concurrency and the issue of message loss
       when establishing a connection for the first time.

    """

    def __init__(self, ip, **kwargs):
        super(Ping, self).__init__(**kwargs)
        self.ip = ip

        validates.is_required(self.password, 'password')

    def execute(self):
        ip = self.ip
        cmd = ["ping", "-c", "3", self.ip]
        executor.execute(*cmd)
        LOG.info("ping %s successfully", ip)
