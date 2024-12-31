from bmpa.log import log_complete
from bmpa.osystem.factory import get_osystem
from bmpa.service import node


@log_complete
def set_meta_data(data):
    target = node.get_target()
    root_partition = node.get_root_partition()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_meta_data(target=target, root_partition=root_partition,
                          data=data)


def set_user_data(data):
    target = node.get_target()
    root_partition = node.get_root_partition()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_user_data(target=target, root_partition=root_partition,
                          data=data)


def set_vendor_data(data):
    target = node.get_target()
    root_partition = node.get_root_partition()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_vendor_data(target=target, root_partition=root_partition,
                            data=data)


def set_network(data):
    target = node.get_target()
    root_partition = node.get_root_partition()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_network(target=target, root_partition=root_partition,
                        data=data)


def set_cloudinit_conf(ssh_pwauth, disable_root):
    target = node.get_target()
    root_partition = node.get_root_partition()
    os_type = node.get_os_type()
    os_version = node.get_os_version()
    osystem = get_osystem(os_type=os_type, os_version=os_version)
    osystem.set_cloudinit_conf(target=target, root_partition=root_partition, ssh_pwauth=ssh_pwauth,
                               disable_root=disable_root)
