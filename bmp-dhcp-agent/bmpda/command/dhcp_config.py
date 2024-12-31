import logging

from bmpda.command.common import Command
from bmpda.service import dhcp_config
from bmpda.utils import validates

LOG = logging.getLogger(__name__)


class DHCPConfigAddHost(Command):
    """scheduler to agent command"""

    serializable_fields = Command.serializable_fields + ('ip', 'mac')

    def __init__(self, ip, mac, **kwargs):
        super(DHCPConfigAddHost, self).__init__(**kwargs)

        self.ip = ip
        self.mac = mac
        validates.is_required(self.ip, "ip")
        validates.is_required(self.mac, "mac")

    def execute_before(self):
        pass

    def execute(self):
        dhcp_config.add_host(ip=self.ip, mac=self.mac)

    def execute_after(self):
        self.data['mac'] = self.mac
        self.data['ip'] = self.ip


class DHCPConfigDelHost(Command):
    """scheduler to agent command"""

    def __init__(self, mac, ip=None, **kwargs):
        super(DHCPConfigDelHost, self).__init__(**kwargs)
        self.mac = mac
        validates.is_required(self.mac, "mac")

    def execute_before(self):
        pass

    def execute(self):
        dhcp_config.del_host(mac=self.mac)

    def execute_after(self):
        self.data['mac'] = self.mac


class DHCPConfigAddSubnet(Command):
    """scheduler to agent command"""

    def __init__(self, subnet, subnet_mask, routes, **kwargs):
        super(DHCPConfigAddSubnet, self).__init__(**kwargs)

        self.subnet = subnet
        self.subnet_mask = subnet_mask
        self.routes = routes

        validates.is_required(self.subnet, "subnet")
        validates.is_required(self.subnet_mask, "subnet_mask")
        validates.is_required(self.routes, "routes")

    def execute_before(self):
        pass

    def execute(self):
        dhcp_config.add_subnet(
            subnet=self.subnet, subnet_mask=self.subnet_mask, routes=self.routes)

    def execute_after(self):
        self.data['subnet'] = self.subnet
        self.data['subnet_mask'] = self.subnet_mask
        self.data['routes'] = self.routes


class DHCPConfigDelSubnet(Command):
    """scheduler to agent command"""

    def __init__(self, subnet, **kwargs):
        super(DHCPConfigDelSubnet, self).__init__(**kwargs)

        self.subnet = subnet
        validates.is_required(self.subnet, "subnet")

    def execute_before(self):
        pass

    def execute(self):
        dhcp_config.del_subnet(subnet=self.subnet)

    def execute_after(self):
        self.data['subnet'] = self.subnet
