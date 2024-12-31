import logging
import os

from bmpa import config
from bmpa.errors import ImageDownloadError, ImageChecksumError
from bmpa.log import log_complete
from bmpa.osystem.factory import get_osystem
from bmpa.service import node
from bmpa.utils import clear_holders
from bmpa.utils import device_hints
from bmpa.utils import executor
from bmpa.utils import file
from bmpa.utils import hardware
from bmpa.utils import system
from bmpa.utils import urlutils

LOG = logging.getLogger(__name__)


def write_image(url: str, format: str, hash: str, filename: str):
    root_device = node.get_root_devcie_name()
    clear_holders.clear_holders(root_device)
    hardware.destroy_disk_metadata(root_device)
    local_file_dir = node.get_temp_directory()
    url = _parse_url(url)
    image_file_path = _download_image(url, hash, local_file_dir)
    write_image_script = system.get_relative_path(
        'shell/write_image.sh')
    cmd = ('/bin/bash', write_image_script, image_file_path, root_device)
    LOG.info('Writing image with command: %s.', ' '.join(cmd))
    stdout, stderr = executor.execute(*cmd)
    LOG.info('Write image to completely. stdout : %s, stderr : %s.',
             stdout, stderr)

    hardware.fix_gpt_partition(root_device)

    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.write_image_setup_boot()
    root_partition = node.get_root_partition()
    if not root_partition:
        root_partition = osystem.get_root_partition(root_device)
        node.set_root_partition(root_partition)
    res = dict()
    device = device_hints.get_device_by_hints(
        {'name': root_device})
    res['root_device_hints'] = device
    return res


@log_complete
def write_image_tar(url: str, format: str, hash: str, filename: str):
    local_file_dir = node.get_temp_directory()
    target = node.get_target()
    root_device = node.get_root_devcie_name()
    root_partition = node.get_root_partition()
    grub_ttys = node.get_grub_ttys()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    boot_mode = node.get_boot_mode()
    mounts = node.get_mounts()
    url = _parse_url(url)
    image_file_path = _download_image(url, hash, local_file_dir)
    _write_image(image_file_path, target)

    osystem = get_osystem(os_type=os_type, os_version=os_version)
    LOG.info('OS:%s', osystem.__class__.__name__)
    osystem.setup_boot(root_device=root_device, target=target, root_partition=root_partition,
                       boot_mode=boot_mode, mounts=mounts, grub_ttys=grub_ttys)
    res = dict()
    device = device_hints.get_device_by_hints(
        {'name': root_device})
    res['root_device_hints'] = device
    return res


def _parse_url(url):
    host = "http://%(ip)s:%(port)s" % {
        "ip": config.image_host,
        "port": config.image_port
    }
    return url.replace("{host}", host)


def _write_image(image_file_path, target):
    cmd = ('tar', 'xf', image_file_path, '-C',
                                         target)
    executor.execute(*cmd)

    LOG.info('Write image:%s success.', image_file_path)


def _download_image(url, hash, local_file_dir):
    local_filename = url.split('/')[-1]
    local_file_path = os.path.join(local_file_dir, local_filename)
    try:
        LOG.info('Dowload image start.')
        urlutils.download(url, local_file_path)
    except Exception as e:
        raise ImageDownloadError(url, str(e))
    _check_image_hash(url, local_file_path, hash)
    LOG.info('Dowload image finish.')
    return local_file_path


def _check_image_hash(image_id, local_file_path, hash):
    LOG.info('Check image hash start')
    real_hash = file.md5sum(local_file_path)
    if hash != real_hash:
        raise ImageChecksumError(image_id, local_file_path, hash, real_hash)
    LOG.info('Check image hash success.')
