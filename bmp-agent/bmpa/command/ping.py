import logging

from bmpa.command.common import Command

LOG = logging.getLogger(__name__)


class Ping(Command):
    """Detect mq communication.

       Avoid high concurrency and the issue of message loss
       when establishing a connection for the first time.

    """

    def __init__(self, **kwargs):
        super(Ping, self).__init__(**kwargs)

    def execute(self):
        LOG.info("Ping completely.")
