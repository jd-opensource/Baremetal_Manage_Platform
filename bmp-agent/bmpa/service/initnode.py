import logging
import operator
import os
import tempfile


from bmpa.device.factory import get_device
from bmpa.errors import BmpError
from bmpa.service import node
from bmpa.utils import device_hints
from bmpa.utils import hardware
from bmpa.utils import mountutils

LOG = logging.getLogger(__name__)


def init_node(node_data):
    mountutils.do_umount('/tmp', recursive=True)
    node.update_node(node_data)
    root_device = _init_root_device(node_data)
    data_device = _init_data_device(node_data)
    _init_dir()
    _init_grub_ttys()
    res_node = dict()
    res_node['root_device_hints'] = root_device
    if data_device is not None:
        res_node['data_device_hints'] = data_device
    return res_node


def _init_root_device(node_data) -> hardware.BlockDevice:
    root_device_hints = node_data.get('root_device_hints')
    if root_device_hints:
        device = device_hints.get_device_by_hints(root_device_hints)
        node.set_root_devcie_name(device['name'])
    root_device_name = node.get_root_devcie_name()
    if root_device_name:
        return device_hints.get_device_by_hints(
            {'name': root_device_name})

    block_devices = hardware.list_all_block_devices()
    if not block_devices or len(block_devices) <= 0:
        raise BmpError("No block device found.")
    block_devices = sorted(block_devices, key=operator.attrgetter('size'))
    node.set_root_devcie_name(block_devices[0].name)
    return block_devices[0]


def _init_data_device(node_data) -> hardware.BlockDevice:
    data_device_hints = node_data.get('data_device_hints')
    if data_device_hints:
        device = device_hints.get_device_by_hints(data_device_hints)
        node.set_data_devcie_name(device['name'])
    data_device_name = node.get_data_devcie_name()
    if data_device_name:
        return device_hints.get_device_by_hints(
            {'name': data_device_name})

    block_devices = hardware.list_all_block_devices()
    if block_devices and len(block_devices) > 1:
        block_devices = sorted(block_devices, key=operator.attrgetter('size'))
        node.set_data_devcie_name(block_devices[1].name)
        return block_devices[1]
    return None


def _init_dir():
    # work directory
    work_directory = tempfile.mkdtemp()
    LOG.info('Use work directory : %s.', work_directory)

    # target directory
    target = os.path.join(work_directory, "chroot")
    if not os.path.exists(target):
        os.makedirs(target)
    LOG.info('Use target directory : %s', target)

    # temp directory
    temp_directory = os.path.join(work_directory, "temp")
    if not os.path.exists(temp_directory):
        os.makedirs(temp_directory)
    LOG.info('Use temp directory : %s.', temp_directory)

    node.set_target(target)
    node.set_temp_directory(temp_directory)


def _init_grub_ttys():
    system_vendor_info = node.get_system_vendor_info()
    manufacturer = system_vendor_info.manufacturer
    product_name = system_vendor_info.product_name
    device = get_device(manufacturer, product_name)
    grub_ttys = device.grub_ttys()
    node.set_grub_ttys(grub_ttys)
