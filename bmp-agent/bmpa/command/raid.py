import logging

from bmpa.command.common import Command
from bmpa.service import raid
import bmpa.utils.validates as validates

LOG = logging.getLogger(__name__)


class MakeRaid(Command):

    def __init__(self, raid_datas, **kwargs):
        super(MakeRaid, self).__init__(**kwargs)
        self.raid_datas = raid_datas
        self.res_raid_datas = None
        validates.is_required(raid_datas, 'raid_datas')
        for raid_data in raid_datas:
            validates.check_field(raid_data, 'adapter_id')
            validates.check_field(raid_data, 'volumes')

    def execute(self):
        self.res_raid_datas = raid.make_raid(self.raid_datas)

    def execute_after(self):
        self.data["raid_datas"] = self.res_raid_datas


class CleanRaid(Command):

    def __init__(self, raid_datas, **kwargs):
        super(CleanRaid, self).__init__(**kwargs)
        self.raid_datas = raid_datas
        validates.is_required(raid_datas, 'raid_datas')
        for raid_data in raid_datas:
            validates.check_field(raid_data, 'adapter_id')

    def execute(self):
        raid.clean_raid(self.raid_datas)
