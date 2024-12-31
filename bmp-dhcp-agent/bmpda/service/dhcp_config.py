import logging

from filelock import FileLock

from bmpda import config
from bmpda.errors import InvalidArgumentError
from bmpda.service import node
import bmpda.utils.dhcp_editor
import bmpda.utils.executor
import bmpda.utils.omapi

LOG = logging.getLogger(__name__)


def add_host(ip, mac):
    lock = FileLock(node.get_dhcp_config_lock_path(), timeout=3)
    omapi_key = config.omapi_key
    host = config.omapi_host
    port = config.omapi_port
    omapi = bmpda.utils.omapi.OmapiClient(omapi_key=omapi_key,
                                          host=host,
                                          port=port)
    with lock:
        if omapi.exists_host(mac):
            omapi.delete_host(mac)
            message = "DHCP config add host, mac:%(mac)s or ip:%(ip)s is exist." % {
                'mac': mac,
                'ip': ip
            }
            LOG.warn(message)
        try:
            omapi.add_host(ip, mac)
            LOG.info("DHCP add host success.")
        except Exception as e:
            LOG.error("DHCP add host error:%s." % str(e))

    LOG.info("DHCP config add host ip:%(ip)s, mac:%(mac)s  completely.",
             {
                 'ip': ip,
                 'mac': mac
             })


def del_host(mac):
    lock = FileLock(node.get_dhcp_config_lock_path(), timeout=3)
    omapi_key = config.omapi_key
    host = config.omapi_host
    port = config.omapi_port
    omapi = bmpda.utils.omapi.OmapiClient(omapi_key=omapi_key,
                                          host=host,
                                          port=port)
    with lock:
        try:
            if omapi.exists_host(mac) is False:
                message = "DHCP config del host, mac:%(mac)s is not exist." % {
                    'mac': mac
                }
                LOG.error(message)
                raise InvalidArgumentError(message)
            omapi.delete_host(mac)

            if omapi.exists_host(mac):
                message = "DHCP config del host, mac:%(mac)s error." % {
                    'mac': mac
                }
                LOG.error(message)
                raise InvalidArgumentError(message)
            else:
                LOG.info("DHCP add host success.")
        except Exception as e:
            message = "DHCP config del host, mac:%(mac)s error:%(err)s." % {
                'mac': mac,
                'err': str(e)
            }
            LOG.error(message)

    LOG.info("DHCP config del host  mac:%(mac)s  completely.",
             {'mac': mac})


def add_subnet(subnet, subnet_mask, routes):
    dhcp_config_path = node.get_dhcp_config_path()
    dhcpeditor = bmpda.utils.dhcp_editor.DHCPEditor(dhcp_config_path)
    dhcp_config_lock_path = node.get_dhcp_config_lock_path()
    lock = FileLock(dhcp_config_lock_path, timeout=10)
    with lock:
        if dhcpeditor.exists_subnet(subnet):
            message = "Subnet:%(subnet)s is exist." % {'subnet': subnet}
            LOG.error(message)
            dhcpeditor.delete_subnet(subnet)
        try:
            dhcpeditor.add_subnet(subnet, subnet_mask, routes)
            node.set_need_restart_dhcp(True)
            LOG.info("DHCP config add subnet %s success.", subnet)
        except Exception as e:
            message = "DHCP config add subnet, subnet:%(subnet)s error:%(err)s." % {
                'subnet': subnet,
                'err': str(e)
            }
            LOG.error(message)

    LOG.info("DHCP config add subnet  subnet:%(subnet)s  completely.",
             {'subnet': subnet})


def del_subnet(subnet):
    dhcp_config_path = node.get_dhcp_config_path()
    dhcpeditor = bmpda.utils.dhcp_editor.DHCPEditor(dhcp_config_path)
    dhcp_config_lock_path = node.get_dhcp_config_lock_path()
    lock = FileLock(dhcp_config_lock_path, timeout=10)
    with lock:
        try:
            if dhcpeditor.exists_subnet(subnet) is False:
                message = "DHCP config delete subnet, subnet:%(subnet)s is not exist." % {
                    'subnet': subnet
                }
                LOG.error(message)
                raise InvalidArgumentError(message)
            dhcpeditor.delete_subnet(subnet)

            if dhcpeditor.exists_subnet(subnet):
                message = "DHCP config delete subnet, subnet:%(subnet)s error." % {
                    'subnet': subnet
                }
                LOG.error(message)
                raise InvalidArgumentError(message)
            node.set_need_restart_dhcp(True)
        except Exception as e:
            message = "DHCP config delete subnet, subnet:%(subnet)s error:%(err)s." % {
                'subnet': subnet,
                'err': str(e)
            }
            LOG.error(message)

        LOG.info("DHCP config delete subnet  subnet:%(subnet)s  completely.",
                 {'subnet': subnet})
