import logging

from bmpa.errors import IncompatibleHardwareMethodError

LOG = logging.getLogger(__name__)

RAID_INDEX = {
    "RAID0": "0",
    "RAID1": "1",
    "RAID5": "5",
    "RAID10": "10"
}


class BaseRaid(object):

    root_device = None

    def execute_raid_config(self, config):
        raid_level = config.raid_level.lower()
        if raid_level == 'noraid':
            return self.noraid(config)
        elif raid_level == 'raid0':
            return self.raid0(config)
        elif raid_level == 'raid1':
            return self.raid1(config)
        elif raid_level == 'raid5':
            return self.raid5(config)
        elif raid_level == 'raid10':
            return self.raid10(config)
        else:
            raise ValueError(f'Unsupported RAID level: {raid_level}')

    def list_adapter(self):
        """Get adapters info.

        """
        return None

    def get_adapter_ids(self):
        raise IncompatibleHardwareMethodError(
            'Not support get adapter ids method.')

    def display(self, adapter_id: int = 0):
        """Get physical disks.

        """
        return None

    def get_disks(self, adapter_id: str = '0'):
        raise IncompatibleHardwareMethodError('Not support get disk method.')

    def noraid(self, config):
        raise IncompatibleHardwareMethodError('Not support NORAID type.')

    def raid0(self, config):
        raise IncompatibleHardwareMethodError('Not support RAID0 type.')

    def raid1(self, config):
        raise IncompatibleHardwareMethodError('Not support RAID1 type.')

    def raid5(self, config):
        raise IncompatibleHardwareMethodError('Not support RAID5 type.')

    def raid10(self, config):
        raise IncompatibleHardwareMethodError('Not support RAID10 type.')

    def clean_raid(self, adapter_id: str = '0', delay_time=0):
        raise IncompatibleHardwareMethodError(
            'Not support clean raid config method.')
