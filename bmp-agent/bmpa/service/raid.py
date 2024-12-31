import logging
import time

from bmpa.errors import BlockDeviceError
from bmpa.log import log_complete
from bmpa.model.raid import RaidConfig
from bmpa.raid.factory import get_raid_driver
from bmpa.service import node
from bmpa.utils import hardware


LOG = logging.getLogger(__name__)


@log_complete
def make_raid(raid_datas):
    res_raid_datas = []
    res_volumes = []
    for raid_data in raid_datas:
        raid_driver = raid_data.get('raid_driver')
        adapter_id = raid_data.get('adapter_id')
        raid = get_raid_driver(raid_driver)
        raid.clean_raid(adapter_id)
        time.sleep(20)
        for volume in raid_data['volumes']:
            is_root_device, is_data_device = _get_is_root_is_data(volume)
            devices_old = hardware.list_all_block_devices()

            raid_config = RaidConfig(
                raid_data['adapter_id'], volume['raid_level'], volume['physical_disks'])

            raid.execute_raid_config(raid_config)
            time.sleep(20)
            devices_new = hardware.list_all_block_devices()
            if not devices_new or len(devices_new) == 0:
                raise BlockDeviceError('Devices get None.')
            device_new = _get_device_new(devices_old, devices_new)
            if is_root_device:
                node.set_root_devcie_name(device_new.name)
                LOG.info('Set root device:%s', device_new.name)
            if is_data_device:
                node.set_data_devcie_name(device_new.name)
                LOG.info('Set data device:%s', device_new.name)
            res_volumes_item = {}
            res_volumes_item['volume_id'] = volume['volume_id']
            res_volumes_item['disk_hints'] = device_new
            res_volumes.append(res_volumes_item)
        res_raid_datas_item = {}
        res_raid_datas_item['adapter_id'] = adapter_id
        res_raid_datas_item['volumes'] = res_volumes
        res_raid_datas.append(res_raid_datas_item)
    time.sleep(10)
    return res_raid_datas


def clean_raid(raid_datas):
    for raid_data in raid_datas:
        raid_driver = raid_data.get('raid_driver')
        raid = get_raid_driver(raid_driver)
        adapter_id = raid_data.get('adapter_id')
        raid.clean_raid(adapter_id)


def _get_is_root_is_data(volume: dict):
    is_root_device = False
    is_data_device = False
    if volume.get('is_root_device', False):
        is_root_device = True
    if volume.get('is_data_device', False):
        is_data_device = True
    return is_root_device, is_data_device


def _get_device_new(devices_old, devices_new):
    LOG.info("Devices old:%s.", devices_old)
    LOG.info("Devices new:%s.", devices_new)
    device_new = None
    if not devices_old or len(devices_old) == 0:
        return devices_new[0]
    devic_names = [device.name for device in devices_old]
    for item in devices_new:
        if item.name in devic_names:
            LOG.info("Device name:%s in.", item.name)
            continue
        else:
            LOG.info("New device name:%s.")
            device_new = item
            break
    else:
        BlockDeviceError('Can not find new device.')
    return device_new
