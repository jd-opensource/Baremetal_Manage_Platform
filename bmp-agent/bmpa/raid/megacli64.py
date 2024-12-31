import logging
import re
import time

from bmpa.errors import BmpError, IncompatibleHardwareMethodError
from bmpa.raid import BaseRaid, RAID_INDEX
from bmpa.model.disk import Disk
from bmpa.utils import executor
from bmpa.utils import strutils

LOG = logging.getLogger(__name__)

EXECUTABLE_FILE = '/usr/sbin/MegaCli64'


class MegaCli64(BaseRaid):

    def list_adapter(self):
        cmd = (EXECUTABLE_FILE, '-AdpGetPciInfo', '-aAll')
        stdout, stderr = executor.execute(*cmd)
        if stderr:
            raise BmpError('Excute list adapter error:%s', stderr)
        LOG.info('List adapter operation : %s.', stdout)
        return stdout

    def get_adapter_ids(self):
        stdout = self.list_adapter()
        adapter_ids = []
        if stdout:
            pattern = re.compile(
                r'PCI\sinformation\sfor\sController\s+(?P<adapter_id>\d+)\s?')
            for match in re.finditer(pattern, stdout):
                adapter_ids.append(match.group('adapter_id'))
        if len(adapter_ids) == 0:
            raise BmpError('Get addapter ids error.')

        return adapter_ids

    def display(self, adapter_id: str = '0'):
        cmd = (EXECUTABLE_FILE, '-PDList', '-a%s' % adapter_id)
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
        adapter_id = config.adapter_id
        raid_index = self._get_raid_index('RAID0')
        for disk_location in config.physical_disks:
            disk_location = self._parse_physical_disks([disk_location])
            self._execute_raid(adapter_id, raid_index, disk_location)

    def raid0(self, config):
        self._close_jbod_mode()
        adapter_id = config.adapter_id
        raid_index = self._get_raid_index(config.raid_level)
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid1(self, config):
        self._close_jbod_mode()
        adapter_id = config.adapter_id
        raid_index = self._get_raid_index(config.raid_level)
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid5(self, config):
        self._close_jbod_mode()
        adapter_id = config.adapter_id
        raid_index = self._get_raid_index(config.raid_level)
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid10(self, config):
        self._close_jbod_mode()
        adapter_id = config.adapter_id
        raid_index = self._get_raid_index(config.raid_level)
        disk_location = self._parse_raid10_physical_disks(
            config.physical_disks)
        cmd = (EXECUTABLE_FILE, '-CfgSpanAdd', '-r10',
               disk_location, 'WB', 'Direct', 'CachedBadBBU', '-a',
               adapter_id)
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Create %s raid%s : %s', raid_index, stdout, stderr)

    def clean_raid(self, adapter_id: str = '0', delay_time=0):
        self._close_jbod_mode()
        self._delete_raid_mode(adapter_id)
        time.sleep(delay_time)

    def _get_raid_index(self, raid_level):
        raid_index = RAID_INDEX.get(raid_level)
        if not raid_index:
            raise IncompatibleHardwareMethodError('Not support raid level %s.' %
                                                  raid_level)
        return raid_index

    def _execute_raid(self, adapter_id, raid_index, disk_location, other_param=''):
        cmd = (EXECUTABLE_FILE, '-CfgLdAdd', '-r%s' % raid_index,
               disk_location, 'WB', 'Direct', 'CachedBadBBU', '-a',
               adapter_id, other_param)
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Create %s raid%s : %s', raid_index, stdout, stderr)

    def _parse_physical_disks(self, physical_disks):
        disk_strings = [
            f"{disk['enclosure']}:{disk['slot']}" for disk in physical_disks]
        return '[%s]' % ', '.join(disk_strings)

    def _parse_raid10_physical_disks(self, physical_disks):
        """Parse raid10 param

        return like "-Array0 [65:2,65:3,65:4,65:5] -Array1 [65:6,65:7,65:8,65:9]"
        """
        if len(physical_disks) % 2 != 0:
            raise BmpError('Raid10 disk location must be 2X.')

        disk_location_groups = []
        # For arrays with 8 disks or fewer, the default value of pdperarray=2 can achieve maximum performance.
        # For arrays with more than 8 disks, setting pdperarray=4 can achieve maximum performance.
        if len(physical_disks) % 4 == 0 and len(physical_disks) > 8:
            pdperarray_num = 4
        else:
            pdperarray_num = 2
        pdperarray = []
        for index, disk in enumerate(physical_disks, 1):
            pdperarray.append(f"{disk['enclosure']}:{disk['slot']}")
            if index % pdperarray_num == 0:
                disk_location_groups.append(pdperarray)
                pdperarray = []

        disk_location_arrays = []
        for index, dl in enumerate(disk_location_groups):
            disk_location_arrays.append('-Array%d' % index)
            disk_location_arrays.append('[%s]' % (','.join(dl)))

        return ' '.join(disk_location_arrays)

    def _parse_disk(self, stdout):
        disks = []
        pattern = re.compile(
            r'Enclosure\sDevice\sID:\s+(?P<enclosure>\d+).*?'
            r'Slot Number:\s+(?P<slot>\d+).*?'
            r'PD Type:\s+(?P<pd_type>\S+).*?'
            r'Raw\sSize:\s+(?P<raw_size>\d+\.\d+)\s+(?P<raw_size_unit>.*?[M|T|G]B).*?'
            r'Coerced\sSize:\s+(?P<core_size>\d+\.\d+)\s+(?P<core_size_unit>.*?[M|T|G]B).*?'
            r'Device Firmware\sLevel:\s*(?P<device_firware_level>\S+).*?'
            r'Inquiry Data:\s*(?P<inquiry_data>[^\n]+).*?'
            r'Media Type:\s+(?P<media_type>[^\n]+).*?', re.DOTALL)
        for match in re.finditer(pattern, stdout):
            enclosure = int(match.group('enclosure'))
            slot = int(match.group('slot'))
            pd_type = match.group('pd_type').strip()
            raw_size = float(match.group('raw_size'))
            raw_size_unit = match.group('raw_size_unit')
            size = float(match.group('core_size'))
            size_unit = match.group('core_size_unit')
            inquiry_data = match.group('inquiry_data').strip()
            device_firware_level = match.group('device_firware_level').strip()
            media_type = strutils.get_initials(
                match.group('media_type').strip())
            disks.append(
                Disk(enclosure=enclosure,
                     slot=slot,
                     pd_type=pd_type,
                     raw_size=raw_size,
                     raw_size_unit=raw_size_unit,
                     size=size,
                     size_unit=size_unit,
                     inquiry_data=inquiry_data,
                     media_type=media_type,
                     device_firware_level=device_firware_level).__dict__)
        return disks

    def _close_jbod_mode(self):
        try:
            # close JBOD mode
            cmd = (EXECUTABLE_FILE, '-AdpSetProp', '-EnableJBOD', '-0',
                   '-aALL')
            stdout, stderr = executor.execute(*cmd)
            LOG.info('Close raid adapter JBOD mode : %s : %s.', stdout,
                     stderr)
        except Exception as e:
            LOG.warn('Close raid adapter JBOD mode error: %s.', str(e))

    def _delete_raid_mode(self, adapter_id: str):
        # clean old raid info
        cmd = (EXECUTABLE_FILE, '-CfgLdDel', '-Lall', 'force', '-aALL')
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Delete raid adapter old raid mode : %s : %s.', stdout,
                 stderr)

        cmd = (EXECUTABLE_FILE, '-CfgForeign', '-Clear', '-a%s' % adapter_id)
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Clean raid adapter old raid mode : %s : %s.', stdout,
                 stderr)
