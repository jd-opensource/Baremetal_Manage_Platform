import logging

LOG = logging.getLogger(__name__)

NODE = dict()


def set_node_id(serial_number: str):
    NODE['serial_number'] = serial_number


def get_node_id() -> str:
    return NODE.get('serial_number', None)


def set_need_restart_dhcp(need_restart_dhcp: bool):
    NODE['need_restart_dhcp'] = need_restart_dhcp


def get_need_restart_dhcp() -> bool:
    return NODE.get('need_restart_dhcp', None)


def set_dhcp_config_path(dhcp_config_path: str):
    NODE['dhcp_config_path'] = dhcp_config_path


def get_dhcp_config_path() -> str:
    return NODE.get('dhcp_config_path', None)


def set_dhcp_config_lock_path(dhcp_config_lock_path: str):
    NODE['dhcp_config_lock_path'] = dhcp_config_lock_path


def get_dhcp_config_lock_path() -> str:
    return NODE.get('dhcp_config_lock_path', None)
