import logging

from bmpa.log import log_complete
from bmpa.utils import clear_holders
from bmpa.utils import hardware


LOG = logging.getLogger(__name__)


@log_complete
def clean_blockdevice():
    block_devices = hardware.list_all_block_devices()
    LOG.info('Clean blocks:%s', block_devices)
    for bd in block_devices:
        clear_holders.clear_holders(bd.name)
        hardware.destroy_disk_metadata(bd.name)
        hardware.erase_block_device(bd.name)
