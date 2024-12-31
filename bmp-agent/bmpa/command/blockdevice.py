import logging

from bmpa.command.common import Command
from bmpa.service import blockdevice

LOG = logging.getLogger(__name__)


class CleanBlockDevice(Command):

    def __init__(self, **kwargs):
        super(CleanBlockDevice, self).__init__(**kwargs)

    def execute(self):
        blockdevice.clean_blockdevice()
