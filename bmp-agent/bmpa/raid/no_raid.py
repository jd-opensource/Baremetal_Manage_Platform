import logging

from bmpa.raid import BaseRaid

LOG = logging.getLogger(__name__)


class NoRaid(BaseRaid):

    def __init__(self, **kwargs):
        super(NoRaid, self).__init__(**kwargs)

    def get_disks(self, adapter_id: str = '0'):
        LOG.warn('No raid adapter.')
        return None

    def get_adapter_ids(self):
        LOG.warn('Nothing operation to do.')
        return None

    def noraid(self, config):
        LOG.warn('Nothing raid operation to do.')

    def raid0(self, config):
        LOG.warn('Nothing raid operation to do.')

    def raid1(self, config):
        LOG.warn('Nothing raid operation to do.')

    def raid5(self, config):
        LOG.warn('Nothing raid operation to do.')

    def raid10(self, config):
        LOG.warn('Nothing raid operation to do.')

    def clean_raid(self, adapter_id: str = '0', delay_time=0):
        LOG.warn('Nothing raid operation to do.')
