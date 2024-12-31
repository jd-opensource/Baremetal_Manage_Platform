import logging

from bmpa import constants
from bmpa.osystem.default import DefaultOsystem
from bmpa.utils import mountutils
from bmpa.utils import paths
from bmpa.utils import system

LOG = logging.getLogger(__name__)


class WindowsBase(DefaultOsystem):
    def _root_lable(self):
        return constants.WINDOWS_ROOT_LABEL

    def set_password(self, target, root_partition, password, username='Administrator'):
        with mountutils.mount(root_partition, target, ['-t', 'ntfs']):
            passworddata_dict = dict()
            if username == 'root':
                username = 'Administrator'
            passworddata_dict["name"] = username
            passworddata_dict["admin_pass"] = password
            path = 'Program Files/Cloudbase Solutions/Cloudbase-Init/openstack/latest/password_data.json'
            super(WindowsBase, self)._write_data_to_target(
                target=target, file_path=path, data=passworddata_dict)

    def set_meta_data(self, target, root_partition, data):
        with mountutils.mount(root_partition, target, ['-t', 'ntfs']):
            super(WindowsBase, self).set_meta_data(
                target, root_partition, data)

    def set_user_data(self, target, root_partition, data):
        with mountutils.mount(root_partition, target, ['-t', 'ntfs']):
            super(WindowsBase, self).set_user_data(
                target, root_partition, data)

    def set_network(self, target, root_partition, data):
        with mountutils.mount(root_partition, target, ['-t', 'ntfs']):
            super(WindowsBase, self).set_network(target, root_partition, data)

    def _user_data_path(self):
        return 'Program Files/Cloudbase Solutions/Cloudbase-Init/openstack/latest/user_data'

    def _meta_data_path(self):
        return 'Program Files/Cloudbase Solutions/Cloudbase-Init/openstack/latest/meta_data.json'

    def _network_path(self):
        return 'Program Files/Cloudbase Solutions/Cloudbase-Init/openstack/latest/network_data.json'

    def _parse_meta_data(self, data):
        if 'local-hostname' in data:
            data['hostname'] = data.pop('local-hostname')
        if 'instance-id' in data:
            data['uuid'] = data.pop('instance-id')
        return data

    def set_cloudbase_conf(self, target, root_partition, only_set_network=False):
        with mountutils.mount(root_partition, target, ['-t', 'ntfs']):
            cloudbase_conf_path = 'Program Files/Cloudbase Solutions/Cloudbase-Init/conf/cloudbase-init.conf'
            absolute_cloudbase_conf_file_path = paths.target_path(
                target, cloudbase_conf_path)

            cloudinit_conf_template_file_path = system.get_relative_path(
                "template/cloudinit/cloudbase-init.tpl")
            cloudbase_config = system.render_template(
                cloudinit_conf_template_file_path,
                {"only_set_network": only_set_network})

            with open(absolute_cloudbase_conf_file_path, 'w') as json_file:
                json_file.write(cloudbase_config)

            LOG.info("Write  %s to %s", cloudbase_config,
                     absolute_cloudbase_conf_file_path)

    def set_cloudinit_conf(self, target, root_partition, ssh_pwauth, disable_root):
        pass


class Windows10(WindowsBase):
    pass


class Windows2012(WindowsBase):
    pass


class Windows2016(Windows2012):
    pass


class Windows2019(Windows2016):
    pass
