
import base64
import json
import logging
import operator
import os
import re

from bmpa import constants
from bmpa.errors import BmpError
from bmpa.utils import chroot
from bmpa.utils import executor
from bmpa.utils.executor import ProcessExecutionError
from bmpa.utils import hardware
from bmpa.utils import mountutils
from bmpa.utils import paths
from bmpa.utils import strutils
from bmpa.utils import system
from bmpa.utils import utils

LOG = logging.getLogger(__name__)


class DefaultOsystem(object):

    def write_image_setup_boot(self):
        pass

    def setup_boot(self, target, root_partition, root_device, boot_mode, mounts, grub_ttys):
        self._update_fstab(target=target, mounts=mounts)
        self._install_grub(target=target, root_device=root_device,
                           boot_mode=boot_mode, grub_ttys=grub_ttys)
        self._update_initramfs(target=target)

    def set_password(self, target, root_partition, password, username='root'):
        if mountutils.is_mounted(target=target):
            self._set_password(target, password, username)
        else:
            mountutils.do_mount(src=root_partition, target=target)
            self._set_password(target, password, username)

        self._open_password_auth(target)
        os.sync()

    def _set_password(self, target, password, username='root'):
        with chroot.ChrootTarget(target) as inchroot:
            inchroot.execute('passwd',
                             username,
                             std_input='%(password)s\n%(password)s\n' %
                             {'password': password})

    def set_hostname(self, target, root_partition, hostname):
        path = paths.target_path(target, '/etc/hostname')
        executor.execute('echo %(hostname)s > %(path)s' %
                         {'hostname': hostname, 'path': path}, shell=True)

    def set_meta_data(self, target, root_partition, data):
        data = self._parse_meta_data(data)
        file_path = self._meta_data_path()
        self._write_data_to_target(target, file_path, data)
        self.set_cloudbase_conf(target, root_partition)

    def set_network(self, target, root_partition, data):
        file_path = self._network_path()
        self._write_data_to_target(target, file_path, data)
        self._set_network_service(target)

    def set_user_data(self, target, root_partition, data):
        if data != '':
            data = base64.b64decode(data)
            if isinstance(data, bytes):
                data = data.decode('utf-8')
            data = self._parse_user_data(data)
        file_path = self._user_data_path()
        self._write_data_to_target(target, file_path, data)

    def set_vendor_data(self, target, root_partition, data):
        file_path = 'var/lib/cloud/seed/nocloud/vendor-data'
        self._write_data_to_target(target, file_path, data)

    def get_root_partition(self, root_device):
        root_partition = None
        root_lable = self._root_lable()
        root_partition = hardware.get_labelled_partition(
            root_device, root_lable, attempts_times=10)

        if root_partition is not None:
            return root_partition
        LOG.warn('Not fund root partition by label:%s', root_lable)
        return self._guess_root_partition(root_device)

    def _root_lable(self):
        return constants.ROOT_LABEL

    def _guess_root_partition(self, root_device):
        partition_template = '%(root_device)s%(device_index)d'
        if hardware.is_nvme_device(root_device):
            partition_template = '%(root_device)sp%(device_index)d'
        root_partition = partition_template % {
            'root_device': root_device,
            'device_index': 2
        }
        return root_partition

    def set_cloudbase_conf(self, target, root_partition, only_set_network=False):
        pass

    def formatpartitions(self, partitions):
        root_partition_path = None
        partitions = self._parse_partitions(partitions)
        for partition in partitions:
            fs = partition['fs_type']
            label = partition['label']
            device_path = partition['device_path']
            options = partition.get('options')
            LOG.info('Format partitions fs:%(fs)s, path:%(path)s lable:%(label)s', {
                'fs': fs,
                'path': device_path,
                'label': label
            })
            hardware.mkfs(fs=fs, path=device_path,
                          label=label, options=options)
            if partition.get('is_root', False):
                root_partition_path = device_path
        return root_partition_path

    def _open_password_auth(self, target):
        self._set_password_auth(target, 'yes')

    def _install_grub(self, target, root_device, boot_mode, grub_ttys):
        try:
            self._update_default_grub(target=target, grub_ttys=grub_ttys)
            if boot_mode == constants.UEFI:
                self._install_uefi_grub(target, root_device)
            else:
                self._install_bios_grub(target, root_device)
            self._kernel_install(target)
        except ProcessLookupError as e:
            error_msg = ('Install grub' 'failed with %(err)s.' % {'err': e})
            LOG.error(error_msg)
            raise BmpError(error_msg)

    def _install_bios_grub(self, target, root_device):
        binary_name = 'grub'
        if os.path.exists(paths.target_path(target, 'usr/sbin/grub2-install')):
            binary_name = 'grub2'

        grub_install = '%s-install' % binary_name
        grub_mkconfig = '%s-mkconfig' % binary_name
        grub_cfg_path = '/boot/%s/grub.cfg' % binary_name
        with chroot.ChrootTarget(target) as inchroot:
            cmd = (grub_install, root_device)
            LOG.info('Installing GRUB2 on disk %s.', root_device)
            stdout, _ = inchroot.execute(*cmd)
            LOG.info('GRUB2 successfully installed on device:%s, %s.',
                     root_device, stdout)
            cmd = (grub_mkconfig, '-o', grub_cfg_path)
            stdout, _ = inchroot.execute(*cmd)
            LOG.info('Grub mkconfig successfully: %s.', stdout)

    def _install_uefi_grub(self, target, root_device):
        self._install_grub_to_device(target, root_device)
        self._generate_efi_grub_config(target)
        self._configure_efi_boot_entry(root_device)
        LOG.info('Install uefi grub success.')

    def _delete_efi_boot(self, efi_boot_label=None):

        delete_efi_boot_script = system.get_relative_path(
            'shell/delete_efi_boot.sh')
        cmd = ['/bin/bash', delete_efi_boot_script]
        if efi_boot_label is not None:
            cmd.append(efi_boot_label)
        LOG.info('Delete efi boot with command: %s', ' '.join(cmd))
        try:
            stdout, stderr = executor.execute(
                *cmd, check_exit_code=[0])
            LOG.info(
                'Delete efi boot completely. stdout : %s, stderr : %s.',
                stdout, stderr)
        except ProcessExecutionError as e:
            LOG.warn('Delete efi boot error :%s.', e)

    def _distro(self):
        return 'linux'

    def _update_dracut_conf(self, target):
        dracut_path = paths.target_path(target, '/etc/dracut.conf')
        with open(dracut_path, 'r') as g:
            contents = g.read()
        with open(dracut_path, 'w') as g:
            g.write(
                re.sub(r'#filesystems.*', r'filesystems="xfs ext4 "',
                       contents))

    def _update_initramfs(self, target):
        self._update_dracut_conf(target)
        linux_version = self._get_kernel_version(target)
        cmd = ('dracut', '--force', '--fstab', '--early-microcode',
               '--kmoddir', '/lib/modules/%s' % linux_version,
               '--kver=%s' % linux_version)
        try:
            with chroot.ChrootTarget(target) as inchroot:
                inchroot.execute(*cmd)
        except ProcessExecutionError as e:
            error_msg = ('repair initram fs'
                         'failed with %(err)s' % {
                             'err': e
                         })
            LOG.error(error_msg)
            raise BmpError(error_msg)

    def _get_kernel_version(self, target):
        return system.get_kernel_version_by_dir(target)

    def _update_fstab(self, target, mounts, mode='w'):
        LOG.info('Start update fstab mounts: %s.', mounts)
        mounts = self._parse_fstab(mounts)
        if mounts is not None and len(mounts) > 0:
            fstab_tpl_path = system.get_relative_path(
                'template/fstab/fstab.tpl')
            fstab_config = system.render_template(
                fstab_tpl_path, {'mounts': mounts})

            fstab_path = paths.target_path(target, 'etc/fstab')
            system.ensure_dir(os.path.dirname(fstab_path))
            with open(fstab_path, mode) as f:
                f.write(fstab_config)

            LOG.info('Update fstab write %s to %s.', fstab_config,
                     fstab_path)

    def _parse_fstab(self, mounts):
        ret = []
        mounts = sorted(mounts, key=operator.itemgetter('mountpoint'))
        for mount in mounts:
            temp_mount = dict()
            options = mount.get('options', '')
            if options == '':
                options = 'defaults'
            temp_mount['label'] = mount['label']
            temp_mount['mountpoint'] = mount['mountpoint']
            temp_mount['fs_type'] = mount['fs_type']
            temp_mount['options'] = options
            ret.append(temp_mount)

        return ret

    def _kernel_install(self, target):
        pass

    def _kernel_install_execute(self, target):
        """Incorporate to address Linux path generation errors in /boot/loader/entries.

        """
        kernel_version = self._get_kernel_version(target)
        kernel_modules_vmlinuz_path = '/lib/modules/%s/vmlinuz' % kernel_version

        try:
            cmd = ('kernel-install', 'add', kernel_version,
                   kernel_modules_vmlinuz_path)
            with chroot.ChrootTarget(target) as inchroot:
                inchroot.execute(*cmd)
        except Exception as e:
            # An exception may be raised here: 'No such file or directory,
            # but it doesn't affect the update of the /boot/loader/entries content.
            LOG.warn('%s', str(e))

    def _install_grub_to_device(self, target, root_device):
        binary_name = 'grub'
        if os.path.exists(paths.target_path(target, 'usr/sbin/grub2-install')):
            binary_name = 'grub2'
        grub_install = '%s-install' % binary_name
        cmd = (grub_install, root_device)
        with chroot.ChrootTarget(target) as inchroot:
            stdout, _ = inchroot.execute(*cmd)
            LOG.info('Grub install successfully: %s.', stdout)

    def _generate_efi_grub_config(self, target):
        binary_name = 'grub'
        if os.path.exists(paths.target_path(target, 'usr/sbin/grub2-install')):
            binary_name = 'grub2'
        grub_mkconfig = '%s-mkconfig' % binary_name
        grub_cfg_path = self._efi_grub_config_path(target)
        cmd = (grub_mkconfig, '-o', grub_cfg_path)
        with chroot.ChrootTarget(target) as inchroot:
            stdout, _ = inchroot.execute(*cmd)
            LOG.info('Generate efi grub config successfully: %s.', stdout)

    def _efi_grub_config_path(self, target):
        binary_name = 'grub'
        if os.path.exists(paths.target_path(target, 'usr/sbin/grub2-install')):
            binary_name = 'grub2'
        efi_grub_cfg_path = '/boot/%(binary_name)s/grub.cfg' % {
            'binary_name': binary_name}
        return efi_grub_cfg_path

    def _configure_efi_boot_entry(self, root_device):
        pass

    def _update_default_grub(self, target: str, grub_ttys: str):
        grub_cmdline_linux = 'console=tty0 console=%s,115200n8 crashkernel=auto' % grub_ttys
        LOG.info('Setting default grub biosedevname parameters: %s.',
                 grub_cmdline_linux)
        default_grub_path = paths.target_path(target, '/etc/default/grub')
        with open(default_grub_path, 'r') as g:
            contents = g.read()
        with open(default_grub_path, 'w') as g:
            g.write(
                re.sub(r'GRUB_CMDLINE_LINUX="(.*)"',
                       r'GRUB_CMDLINE_LINUX="%s"' % grub_cmdline_linux,
                       contents))

    def _set_password_auth(self, target, status):
        """Set password auth.

        parm: status: yes no
        """
        path = paths.target_path(target, '/etc/ssh/sshd_config')
        executor.execute('sed', '-i', '/PasswordAuthentication/d', path)
        executor.execute('echo PasswordAuthentication %(status)s  >> %(path)s' % {
                         'status': status, 'path': path}, shell=True)
        LOG.info('Set password auth %s', status)

    def _set_network_service(self, target):
        pass

    def _parse_user_data(self, data):
        return data

    def _parse_meta_data(self, data):
        return data

    def _user_data_path(self):
        return 'var/lib/cloud/seed/nocloud/user-data'

    def _network_path(self):
        return 'var/lib/cloud/seed/nocloud/network-config'

    def _meta_data_path(self):
        return 'var/lib/cloud/seed/nocloud/meta-data'

    def _vendor_data_path(self):
        return 'var/lib/cloud/seed/nocloud/vendor-data'

    def _write_data_to_target(self, target, file_path, data):
        absolute_file_path = paths.target_path(target, file_path)
        file_dir = os.path.dirname(absolute_file_path)
        system.ensure_dir(file_dir, 0o777)
        if isinstance(data, (dict, list)):
            data = json.dumps(data)
        with open(absolute_file_path, 'w') as data_file:
            data_file.write(data)
            data_file.flush()
        # Write the file from the buffer to the disk,
        # otherwise it may result in the file not being written to the disk.
        os.sync()
        LOG.info("Write  %s to %s", strutils.mask_password(
            data), absolute_file_path)

    def set_cloudinit_conf(self, target, root_partition, ssh_pwauth, disable_root):
        cloudinit_conf_path = 'etc/cloud/cloud.cfg'
        absolute_cloudinit_conf_file_path = paths.target_path(
            target, cloudinit_conf_path)

        cloudinit_conf_tpl_dir = system.get_relative_path('template/cloudinit')

        cloudinit_conf_tpl_file = self._get_cloudinit_conf_template_file()
        cloudinit_conf_template_file_path = paths.target_path(
            cloudinit_conf_tpl_dir, cloudinit_conf_tpl_file)
        cloudinit_config = system.render_template(
            cloudinit_conf_template_file_path, {
                'ssh_pwauth': ssh_pwauth,
                'disable_root': disable_root
            })

        with open(absolute_cloudinit_conf_file_path, 'w') as json_file:
            json_file.write(cloudinit_config)

        LOG.info('Write  %s to %s.', cloudinit_config,
                 absolute_cloudinit_conf_file_path)
        self._enable_cloudinit_service(target)
        self._remove_bonding_modules(target)

    def _enable_cloudinit_service(self, target):
        with chroot.ChrootTarget(target) as inchroot:
            inchroot.execute('systemctl', 'enable', 'cloud-init-local.service')
            inchroot.execute('systemctl', 'enable', 'cloud-init.service')
            inchroot.execute('systemctl', 'enable', 'cloud-config.service')
            inchroot.execute('systemctl', 'enable', 'cloud-final.service')
        LOG.info('Enable cloud service complete.')

    def _remove_bonding_modules(self, target):
        etc_modules_path = paths.target_path(target, 'etc/modules')
        if not os.path.exists(etc_modules_path):
            LOG.info('Not exits /etc/modules.')
            return
        with chroot.ChrootTarget(target) as inchroot:
            inchroot.execute('sed', '-i', '/bonding/d', '/etc/modules')
        LOG.info('Remove bonding modules.')

    def _get_cloudinit_conf_template_file(self):
        return 'cloud.tpl'

    def _parse_partitions(self, partitions):
        for index, partition in enumerate(partitions):
            old_options = partition.get('options')
            fstype = partition['fs_type']
            new_options = self._fstype_options_handlers(fstype)
            LOG.info('New options:%s', new_options)
            options = utils.merge_list(old_options, new_options)
            partitions[index]['options'] = options
        return partitions

    def _apply_xfs_options(sef) -> tuple:
        return None

    def _apply_swap_options(self) -> tuple:
        return None

    def _fstype_options_handlers(self, fstype) -> tuple:
        if fstype == 'xfs':
            return self._apply_xfs_options()
        elif fstype == 'swap':
            return self._apply_swap_options()
        return None
