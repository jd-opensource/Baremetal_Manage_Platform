import json
import logging
import time

from bmpa.errors import BmpError, IncompatibleHardwareMethodError
from bmpa.model.disk import Disk
from bmpa.raid import BaseRaid, RAID_INDEX
from bmpa.utils import executor
from bmpa.utils.executor import ProcessExecutionError


LOG = logging.getLogger(__name__)

EXECUTABLE_FILE = '/usr/sbin/perccli64'


class Perccli64(BaseRaid):

    def list_adapter(self, param='J'):
        cmd = (EXECUTABLE_FILE, '/call show %(param)s' % {'param': param})
        stdout, stderr = executor.execute(*cmd)
        if stderr:
            raise BmpError('Excute list adapter error:%s', stderr)
        LOG.info('List adapter operation : %s.', stdout)
        return stdout

    def get_adapter_ids(self):
        stdout = self.list_adapter()
        adapter_ids = []
        if stdout:
            adapter_ids = []
            json_out = json.loads(stdout)
            for item in json_out['Controllers']:
                adapter_ids.append(item['Command Status']['Controller'])
        if len(adapter_ids) == 0:
            raise BmpError('Get addapter ids error.')

        return adapter_ids

    def display(self, adapter_id=0, param='J'):
        cmd = (EXECUTABLE_FILE, '/c%(adapter_id)s show %(param)s' % {
            'adapter_id': adapter_id,
            'param': param
        })
        stdout, stderr = executor.execute(*cmd)
        if stderr:
            raise BmpError('Excute display error:%s', stderr)
        LOG.info('Display raid adapter %s ', adapter_id)
        return stdout

    def get_disks(self, adapter_id: str = '0'):
        stdout = self.display(adapter_id)
        if stdout:
            LOG.debug('Display out:%s.', stdout)
            return self._parse_disk(stdout)
        return BmpError('Get disks error.')

    def noraid(self, config):
        self._close_jbod_mode()
        raid_level = config.raid_level
        raid_index = RAID_INDEX.get('RAID0')
        if not raid_index:
            raise IncompatibleHardwareMethodError('Not support raid level %s.' %
                                                  raid_level)
        for disk_location in config.physical_disks:
            disk_location = self._parse_physical_disks([disk_location])
            cmd = (EXECUTABLE_FILE, '/c%s' % config.adapter_id, 'add', 'vd',
                   'r%s' % raid_index, 'drives=%s' % disk_location)
            stdout, stderr = executor.execute(*cmd)
            LOG.info('Create %s raid : %s : %s.', raid_level, stdout,
                     stderr)

    def raid0(self, config):
        self._execute_raid(config)

    def raid1(self, config):
        self._execute_raid(config)

    def raid5(self, config):
        self._execute_raid(config)

    def raid10(self, config):
        pdperarray = self._pdperarray(config.physical_disks)
        self._execute_raid(config, pdperarray)

    def clean_raid(self, adapter_id: str = '0', delay_time=0):
        cmd = (EXECUTABLE_FILE, '/c%s/vall' % adapter_id, 'delete',
               'force')
        try:
            stdout, stderr = executor.execute(*cmd)
        except ProcessExecutionError as e:
            if 'No VD' not in e.stdout:
                error_msg = ('clean raid config'
                             'failed with %(err)s' % {
                                 'err': str(e)
                             })
                raise BmpError('Clean raid error:%s.' % error_msg)
        else:
            if 'syntax error' in stdout:
                raise BmpError('Clean raid error:%s.' % stdout)
            LOG.info('Delete %s operation : %s : %s', adapter_id,
                     stdout, stderr)
        time.sleep(delay_time)

    def _parse_disk(self, stdout):
        disks = []
        json_out = json.loads(stdout)
        for item in json_out['Controllers'][0]['Response Data']['PD LIST']:
            slot = int(item['EID:Slt'].split(':')[1])
            enclosure = int(item['EID:Slt'].split(':')[0])
            size = float(item['Size'].split(' ')[0])
            size_unit = item['Size'].split(' ')[1]
            pd_type = item['Intf']
            media_type = item['Med']
            disks.append(
                Disk(enclosure=enclosure,
                     slot=slot,
                     pd_type=pd_type,
                     media_type=media_type,
                     size=size,
                     size_unit=size_unit).__dict__)
        return disks

    def _pdperarray(self, physical_disks):
        if len(physical_disks) % 2 != 0:
            raise BmpError('Raid10 disk location must be 2X.')
        pdperarray_num = 2
        if len(physical_disks) % 4 == 0 and len(physical_disks) > 8:
            pdperarray_num = 4
        return pdperarray_num

    def _execute_raid(self, config, pdperarray=None):
        raid_level = config.raid_level
        raid_index = RAID_INDEX.get(raid_level)
        if not raid_index:
            raise IncompatibleHardwareMethodError('Not support raid level %s.' %
                                                  raid_level)
        disk_location = self._parse_physical_disks(config.physical_disks)
        cmd = (EXECUTABLE_FILE, '-CfgLdAdd', '-r%s' % raid_index,
               disk_location, 'WB', 'Direct', 'CachedBadBBU', '-a',
               config.adapter_id)
        pd = ''
        if pdperarray is not None:
            pd = 'pdperarray=%s' % pdperarray
        cmd = (EXECUTABLE_FILE, '/c%s' % config.adapter_id, 'add', 'vd',
               'r%s' % raid_index, 'drives=%s' % disk_location, pd)
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Create %s raid : %s : %s', raid_level, stdout, stderr)

    def _parse_physical_disks(self, physical_disks):
        disk_strings = [
            f"{disk['enclosure']}:{disk['slot']}" for disk in physical_disks]
        return '%s' % ', '.join(disk_strings)