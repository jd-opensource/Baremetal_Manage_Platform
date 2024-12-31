import os
import re

from bmpa import constants


def _read_params_from_file(filepath: str):
    """Extract key=value pairs from a file.

    :param filepath: path to a file containing key=value pairs separated by
                     whitespace or newlines.
    :returns: a dictionary representing the content of the file
    """
    with open(filepath) as f:
        cmdline = f.read()

    options = cmdline.split()
    params = {}
    for option in options:
        if '=' not in option:
            continue
        k, v = option.split('=', 1)
        params[k] = v

    return params


def get_agent_params(cmdline_path: str = '/proc/cmdline'):
    """Gets parameters passed to the agent via kernel cmdline."""
    return _read_params_from_file(cmdline_path)


def hump2underline(hunp_str):
    pattern = re.compile(r'([a-z]|\d)([A-Z])')
    sub = re.sub(pattern, r'\1_\2', hunp_str).lower()
    return sub


def json_hump2underline(hump_json_str):
    pattern = re.compile(r'"\s*(\w+)\s*"\s*:')
    sub = re.sub(pattern, lambda x: '"' + hump2underline(x.group(1)) + '" :',
                 hump_json_str)
    return sub


def is_uefi_bootable():
    return os.path.exists('/sys/firmware/efi') is True


def merge_list(list1: list, list2: list) -> list:
    if list1 and list2:
        res = list1 + list2
    elif list1:
        res = list1
    elif list2:
        res = list2
    else:
        res = None
    return res


def get_boot_mode(boot_mode) -> str:
    if boot_mode:
        return boot_mode
    return constants.UEFI if is_uefi_bootable() else constants.BIOS


def convert_disk_size(size, size_unit, capacity=1024):
    if size_unit == 'B':
        size = size / capacity**3
        size_unit = 'GB'
    elif size_unit == 'KB':
        size = size / capacity**2
        size_unit = 'GB'
    elif size_unit == 'MB':
        size = size / capacity
        size_unit = 'GB'

    if size_unit == 'GB' and size > capacity:
        size = size / capacity
        size_unit = 'TB'

    return size, size_unit
