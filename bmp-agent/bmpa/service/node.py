import logging

LOG = logging.getLogger(__name__)

NODE = dict()
NODE['blockdev_id_to_devpath'] = dict()


def get_cached_node() -> dict:
    """Return NODE."""
    return NODE


def update_node(node):
    NODE.update(node)
    LOG.info('node:%s', NODE)


def set_grub_ttys(grub_ttys):
    NODE['grub_ttys'] = grub_ttys


def get_grub_ttys():
    return NODE['grub_ttys']


def set_target(target):
    NODE['target'] = target


def get_target():
    return NODE['target']


def set_temp_directory(temp_directory):
    NODE['temp_directory'] = temp_directory


def get_temp_directory():
    return NODE['temp_directory']


def set_boot_mode(boot_mode):
    NODE['boot_mode'] = boot_mode


def get_boot_mode():
    return NODE.get('boot_mode', None)


def set_os_version(os_version: str):
    NODE['os_version'] = os_version


def get_os_version() -> str:
    return NODE.get('os_version', None)


def set_os_type(os_type: str):
    NODE['os_type'] = os_type


def get_os_type() -> str:
    return NODE.get('os_type', None)


def set_system_vendor_info(vendor: any):
    NODE['system_vendor_info'] = vendor


def get_system_vendor_info() -> any:
    return NODE.get('system_vendor_info', None)


def set_serial_number(serial_number: str):
    NODE['serial_number'] = serial_number


def get_serial_number() -> str:
    return NODE.get('serial_number', None)


def set_root_devcie_name(device_name: str):
    """Set root device name.

    device_name: like /dev/sda
    """
    NODE['root_devcie_name'] = device_name
    LOG.info('Root device name be set:%s', device_name)


def get_root_devcie_name() -> str:
    return NODE.get('root_devcie_name', None)


def set_data_devcie_name(device_name: str):
    NODE['data_devcie_name'] = device_name
    LOG.info('Data device name be set:%s', device_name)


def get_data_devcie_name() -> str:
    return NODE.get('data_devcie_name', None)


def set_root_partition(root_partition):
    NODE['root_partition'] = root_partition
    LOG.info("Root partition be set: %s", root_partition)


def get_root_partition():
    return NODE.get('root_partition', None)


def set_mounts(mounts):
    NODE['mounts'] = mounts


def get_mounts():
    return NODE.get('mounts', None)


def set_auto_create_efi(auto_create_efi: bool):
    NODE['auto_create_efi'] = auto_create_efi
    LOG.info('Auto create efi:%s', auto_create_efi)


def get_auto_create_efi():
    return NODE.get('auto_create_efi', True)


def set_efi_partition_id(id: str):
    NODE['efi_partition_id'] = id


def get_efi_partition_id():
    return NODE.get('efi_partition_id', None)


def blockdev_id_to_devpath_dict_add_item(id, path):
    NODE['blockdev_id_to_devpath'][id] = path
    LOG.info('Add blockdev_id_to_devpath %s:%s', id, path)


def blockdev_id_to_devpath(id):
    devpath = NODE['blockdev_id_to_devpath'].get(id, None)
    LOG.info('blockdev_id_to_devpath: %s:%s', id, devpath)
    return devpath


def blockdev_id_to_devpath_dict():
    devpath_dict = NODE.get('blockdev_id_to_devpath', None)
    LOG.info('blockdev_id_to_devpath_dict:%s', devpath_dict)
    return devpath_dict
