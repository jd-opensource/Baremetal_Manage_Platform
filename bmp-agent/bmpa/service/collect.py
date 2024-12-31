import json
import logging
import os
import time

from bmpa.log import log_complete
from bmpa.model.disk import Disk
from bmpa.raid.factory import get_raid_driver
from bmpa.utils import executor
from bmpa.utils import hardware
from bmpa.utils import interfaceutils
from bmpa.utils import utils

LOG = logging.getLogger(__name__)


def list_hardware_info():
    _patch_i40e_enable_lldp()
    hardware_info = {}
    hardware_info['interfaces'] = interfaceutils.list_network_interfaces_meta()
    hardware_info['cpu'] = hardware.get_cpus()
    hardware_info['disks'] = hardware.list_all_block_devices()
    hardware_info['memory'] = hardware.get_memory()
    hardware_info['bmc_address'] = hardware.get_bmc_address()
    hardware_info['system_vendor'] = hardware.get_system_vendor_info()
    hardware_info['boot'] = hardware.get_boot_info()
    hardware_info['platform'] = hardware.get_platform_info()
    return hardware_info


@log_complete
def collect_disks(raid_driver: str) -> dict:
    raid = get_raid_driver(raid_driver)
    adapter_ids = raid.get_adapter_ids()
    controllers = None
    if isinstance(adapter_ids, list):
        controllers = []
        for adapter_id in adapter_ids:
            data = {}
            disks = raid.get_disks(adapter_id)
            data['adapter_id'] = int(adapter_id)
            data['disks'] = disks
            controllers.append(data)
    disk_data = {}
    disk_data['controllers'] = controllers
    disk_data['nvme'] = _get_nvme()
    disk_data['devices'] = _parse_devices(hardware.list_all_block_devices())
    return disk_data


def _patch_i40e_enable_lldp():
    if not os.path.exists("/sys/module/i40e"):
        LOG.warn('Not load i40e module.')
        return
    i40e_dir = "/sys/kernel/debug/i40e"
    if not os.path.exists(i40e_dir):
        LOG.warn("Not load i40e module.")
        return
    LOG.info("Ready to stop disabling embedded lldp in x710.")
    i40_pci_dirs = os.listdir(i40e_dir)
    for i40e_pci_dir in i40_pci_dirs:
        i40e_pci_path = os.path.join(i40e_dir, i40e_pci_dir)
        try:
            os.system('echo "lldp stop" > %s/command' % i40e_pci_path)
            LOG.info('echo "lldp stop" > %s/command' % i40e_pci_path)
            # wait i40e firmware apply
            time.sleep(30)
        except Exception:
            LOG.error('Send lldp stop to command failed.')


def _get_nvme():
    nvme_list = _get_nvme_list()
    return _parse_nvme(nvme_list)


def _parse_nvme(nvme_list: str):
    nvme = None
    if nvme_list:
        nvme_list = utils.json_hump2underline(nvme_list)
        nvme_list = json.loads(nvme_list)
        devices = nvme_list['devices']
        disks = []
        for device in devices:
            disks.append(Disk(
                serial=device['serial_number'],
                size=device['physical_size'],
                size_unit='B',
                pd_type='NVME',
                media_type='SSD',
                device_path=device['device_path']
            ).__dict__)
        nvme = {}
        nvme['devices'] = disks

    return nvme


def _parse_devices(devices: list):
    disks = []
    for device in devices:
        media_type = "HDD" if device.rotational else "SSD"
        disks.append(Disk(
            serial=device.serial,
            size=device.size,
            size_unit='B',
            media_type=media_type,
            device_path=device.name
        ).__dict__)
    return disks


def _get_nvme_list() -> str:
    cmd = ('nvme', "list", "-o", "json")
    stdout, stderr = executor.execute(*cmd)
    LOG.info("List nvme: %s : %s.", stdout, stderr)
    return stdout
