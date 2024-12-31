import logging
import os

from filelock import FileLock

from bmpda import config
from bmpda.errors import BmpError
from bmpda.utils import file
from bmpda.utils import paths
from bmpda.utils import system


LOG = logging.getLogger(__name__)

boot_modes = {"uefi", "bios"}
archs = {"x86_64", "aarch64", "loongarch64"}

grub_config_maps = {
    "bios": {
        "x86_64": {
            "config_dir": "pxelinux.cfg",
            "image_dir": ""
        }
    },
    "uefi": {
        "x86_64": {
            "config_dir": "uefi/x86_64",
            "image_dir": ""
        },
        "aarch64": {
            "config_dir": "uefi/arm64",
            "image_dir": "images/arm64"
        },
        "loongarch64": {
            "config_dir": "uefi/loongarch64",
            "image_dir": "images/loongarch64"
        }
    }
}


def get_lock_file(mac):
    lock_dir = os.path.join(config.tftp_config_dir, 'lock')
    if not os.path.exists(lock_dir):
        file.ensure_dir(lock_dir)

    return os.path.join(lock_dir, f"{mac}.lock")


def add_grub(mac, boot_mode, arch, kernel_name, initramfs_name):
    lock_file = get_lock_file(mac)
    lock = FileLock(lock_file, timeout=3)
    with lock:
        try:
            boot_mode_set = {boot_mode} if boot_mode is not None else boot_modes
            for item_boot_mode in boot_mode_set:
                arch_set = {arch} if arch is not None else grub_config_maps.get(
                    item_boot_mode).keys()
                LOG.info(f"Arch set:{arch_set}")
                for item_arch in arch_set:
                    write_grub_config(mac=mac, boot_mode=item_boot_mode, arch=item_arch,
                                      kernel_name=kernel_name, initramfs_name=initramfs_name)
        except Exception as e:
            LOG.error("TFTP add grub config mac:%s error:%s.", mac, e)
            raise BmpError("TFTP add grub config mac:{mac} error: {e}.")
    os.remove(lock_file)
    LOG.info(f"TFTP config add  mac:{mac}  completely.")


def del_grub(mac):
    lock_file = get_lock_file(mac)
    lock = FileLock(lock_file, timeout=3)
    with lock:
        try:
            boot_mode_set = boot_modes
            for item_boot_mode in boot_mode_set:
                arch_set = grub_config_maps.get(item_boot_mode).keys()
                for item_arch in arch_set:
                    file_path = get_absolute_grub_config_path(
                        mac=mac, boot_mode=item_boot_mode, arch=item_arch)
                    if os.path.exists(file_path):
                        os.remove(file_path)
                        LOG.info(f"{file_path} was remove.")

        except Exception as e:
            LOG.error("TFTP delete grub config mac:%s error:%s.", mac, e)
            raise BmpError("TFTP delete grub config mac:{mac} error: {e}.")
    os.remove(lock_file)
    LOG.info(f"TFTP config delete grub  mac:{mac}  completely.")


def write_grub_config(mac, boot_mode, arch, kernel_name, initramfs_name):
    LOG.info(f"Add file {mac} {boot_mode} {arch} "
             f"{kernel_name} {initramfs_name}.")
    config_data = get_grub_config_data(
        boot_mode=boot_mode, arch=arch, kernel_name=kernel_name, initramfs_name=initramfs_name)
    absolute_config_path = get_absolute_grub_config_path(mac, boot_mode, arch)
    with open(absolute_config_path, "w") as json_file:
        json_file.write(config_data)
    LOG.info(f"Add file {absolute_config_path} .")


def get_absolute_grub_config_path(mac, boot_mode, arch):
    def get_grub_file_name(boot_mode, mac):
        mac_formatted = mac.replace(":", "-")
        if boot_mode == "bios":
            return f"01-{mac_formatted}"
        return f"grub.cfg-01-{mac_formatted}"

    grub_confg_map = get_grub_config(boot_mode, arch)
    config_dir = grub_confg_map.get("config_dir")
    grub_file_name = get_grub_file_name(boot_mode, mac)
    grub_file_path = os.path.join(config_dir.lstrip("/"), grub_file_name)
    absolute_config_path = paths.target_path(
        config.tftp_config_dir, grub_file_path)
    return absolute_config_path


def get_grub_config_data(boot_mode, arch, kernel_name, initramfs_name):
    if boot_mode == "bios":
        tpl_path = 'template/tftp/default.tpl'
    else:
        tpl_path = 'template/tftp/grub.cfg.tpl'
    grub_conf_template_file_path = system.get_relative_path(tpl_path)
    grub_confg_map = get_grub_config(boot_mode, arch)
    image_dir = grub_confg_map.get('image_dir')
    kernel_path = os.path.join(image_dir, kernel_name)
    initramfs_path = os.path.join(image_dir, initramfs_name)
    grub_config = system.render_template(
        grub_conf_template_file_path, {
            'kernel_path': kernel_path,
            'initramfs_path': initramfs_path,
            'bmp_mq_host': config.tftp_mq_host,
            'bmp_mq_port': config.tftp_mq_port,
            'bmp_mq_user': config.tftp_mq_user,
            'bmp_mq_password': config.tftp_mq_password,
            'bmp_mq_vhost': config.tftp_mq_vhost,
            'bmp_mq_exchange_routing_key': config.tftp_mq_exchange_routing_key,
            'bmp_image_host': config.tftp_image_host,
            'bmp_image_port': config.tftp_image_port,
            'bmp_rsyslog_host': config.tftp_rsyslog_host,
            'bmp_rsyslog_port': config.tftp_rsyslog_port,
        })
    return grub_config


def get_grub_config(boot_mode, arch):
    LOG.info(f"Boot mode:{boot_mode} arch:{arch}")
    return grub_config_maps.get(boot_mode).get(arch)
