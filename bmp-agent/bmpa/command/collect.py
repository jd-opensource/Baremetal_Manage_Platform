import logging

from bmpa.command.common import Command
from bmpa.service import collect

LOG = logging.getLogger(__name__)


class CollectHardwareInfo(Command):
    def __init__(self, **kwargs):
        super(CollectHardwareInfo, self).__init__(**kwargs)
        self.hardware_info = None

    def execute_before(self):
        pass

    def execute(self):
        self.hardware_info = collect.list_hardware_info()

    def execute_after(self):
        self.data = self.hardware_info


class CollectDiskLocations(Command):
    """Collecting Disk Information."""

    def __init__(self, raid_driver=None, **kwargs):
        super(CollectDiskLocations, self).__init__(**kwargs)
        self.disk_data = None
        self.raid_driver = raid_driver

    def execute(self):
        self.disk_data = collect.collect_disks(self.raid_driver)

    def execute_after(self):
        self.data = self.disk_data
