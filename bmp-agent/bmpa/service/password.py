from bmpa.log import log_complete
from bmpa.osystem.factory import get_osystem
from bmpa.service import node


@log_complete
def set_password(password: str, username: str):
    target = node.get_target()
    root_device = node.get_root_devcie_name()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    root_partition = node.get_root_partition()
    if not root_partition:
        root_partition = osystem.get_root_partition(root_device)
        node.set_root_partition(root_partition)
    osystem.set_password(target=target, root_partition=root_partition,
                         password=password, username=username)
