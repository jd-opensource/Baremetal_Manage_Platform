import logging

from bmpa.command.common import Command
from bmpa.service import formatpartitions
import bmpa.utils.validates as validates

LOG = logging.getLogger(__name__)


class FormatPartitions(Command):
    def __init__(self, partitions, **kwargs):
        self.partitions = partitions
        validates.is_required(partitions, 'partitions')

        super(FormatPartitions, self).__init__(**kwargs)

    def execute(self):
        formatpartitions.formatpartitions(self.partitions)
