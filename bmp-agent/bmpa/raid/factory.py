import logging

from bmpa.raid.arcconf import Arcconf
from bmpa.raid.megacli64 import MegaCli64
from bmpa.raid.no_raid import NoRaid
from bmpa.raid.perccli64 import Perccli64
from bmpa.raid.sas2ircu import Sas2ircu
from bmpa.raid.sas3ircu import Sas3ircu
from bmpa.raid.storcli64 import Storcli64

LOG = logging.getLogger(__name__)

RAID_MAPS = {'storcli64': Storcli64, 'megacli64': MegaCli64,
             'perccli64': Perccli64, 'sas3ircu': Sas3ircu,
             'sas2ircu': Sas2ircu, 'arcconf': Arcconf, 'no_raid': NoRaid}
RAIDS = [Sas3ircu, Sas2ircu, Arcconf, Storcli64, MegaCli64, Perccli64, NoRaid]


def get_raid_driver(raid_driver: str = None):
    if raid_driver:
        raid_driver = raid_driver.lower()
        LOG.info('Raid driver:%s.', raid_driver)
    if raid_driver in RAID_MAPS.keys():
        LOG.info('Raid driver:%s in map.', raid_driver)
        return RAID_MAPS.get(raid_driver)()
    return _guess_raid_driver()


def _guess_raid_driver():
    LOG.info('Guess raid dirver.')
    for raid_driver in RAIDS:
        try:
            driver = raid_driver()
            driver.get_adapter_ids()
        except Exception as e:
            LOG.info('%s can not get raid adapter:%s.',
                     raid_driver.__name__, str(e))
        else:
            return driver
    return NoRaid()
