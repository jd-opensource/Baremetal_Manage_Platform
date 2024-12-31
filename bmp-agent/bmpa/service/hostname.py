from bmpa.log import log_complete
from bmpa.osystem.factory import get_osystem
from bmpa.service import node


@log_complete
def set_host_name(hostname: str):
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_hostname(hostname)
