import logging

from bmpa.command.common import Command
from bmpa.service import metadata
from bmpa.utils import validates


LOG = logging.getLogger(__name__)


class SetMetaData(Command):

    def __init__(self, meta_data, **kwargs):
        super(SetMetaData, self).__init__(**kwargs)
        self.data = meta_data

    def validate(self):
        validates.is_required(self.data, 'meta_data')

    def execute(self):

        metadata.set_meta_data(self.data)


class SetUserData(Command):

    def __init__(self, user_data='', **kwargs):
        super(SetUserData, self).__init__(**kwargs)
        self.data = user_data

    def execute(self):

        metadata.set_user_data(self.data)


class SetVendorData(Command):

    def __init__(self, vendor_data, **kwargs):
        super(SetVendorData, self).__init__(**kwargs)
        self.data = vendor_data

    def validate(self):
        validates.is_required(self.data)

    def execute(self):

        metadata.set_vendor_data(self.data)


class SetNetwork(Command):

    def __init__(self, network, **kwargs):
        super(SetNetwork, self).__init__(**kwargs)
        self.data = network

    def validate(self):
        validates.is_required(self.data, 'network')

    def execute(self):

        metadata.set_network(self.data)


class SetNetworkWindows(Command):

    def __init__(self, network, **kwargs):
        super(SetNetworkWindows, self).__init__(**kwargs)
        self.data = network

    def validate(self):
        validates.is_required(self.data, 'network')

    def execute(self):

        metadata.set_network(self.data)


class SetCloudinitConf(Command):

    def __init__(self, ssh_pwauth="no", disable_root=False, **kwargs):
        super(SetCloudinitConf, self).__init__(**kwargs)
        self.ssh_pwauth = ssh_pwauth
        self.disable_root = disable_root

    def execute(self):

        metadata.set_cloudinit_conf(self.ssh_pwauth, self.disable_root)
