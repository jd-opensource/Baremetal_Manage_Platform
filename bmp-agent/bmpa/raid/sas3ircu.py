import logging
import re
import time

from bmpa.errors import BmpError, IncompatibleHardwareMethodError
from bmpa.model.disk import Disk
from bmpa.raid import BaseRaid, RAID_INDEX
from bmpa.utils import executor

LOG = logging.getLogger(__name__)

EXECUTABLE_FILE = '/usr/sbin/sas3ircu'


class Sas3ircu(BaseRaid):

    def list_adapter(self):
        """Avago Technologies SAS3 IR Configuration Utility.

        Version 13.00.00.00 (2016.03.08)
        Copyright (c) 2009-2016 Avago Technologies. All rights reserved.


                Adapter       Vendor  Device                      SubSys  SubSys
        Index    Type           ID      ID    Pci Address         Ven ID  Dev ID
        -----  ------------  ------  ------  -----------------   ------  ------
        0     SAS3008        1000h   97h    00h:01h:00h:00h     1000h   3090h
        SAS3IRCU: Utility Completed Successfully.
        """
        cmd = (EXECUTABLE_FILE, 'list')
        stdout, stderr = executor.execute(*cmd)
        if stderr:
            raise BmpError('Excute list adapter error:%s', stderr)
        LOG.info('List adapter operation : %s.', stdout)
        return stdout

    def get_adapter_ids(self):
        stdout = self.list_adapter()
        if stdout:
            pattern = r'\s+(\d+)\s+SAS3008'
            pattern_object = re.compile(pattern, re.MULTILINE)
            adapter_ids = []
            for m in pattern_object.finditer(stdout):
                adapter_ids.append(m.group(1))
            if len(adapter_ids) > 0:
                return adapter_ids

        raise BmpError('Get addapter ids error.')

    def display(self, adapter_id=0):
        cmd = (EXECUTABLE_FILE, adapter_id, 'display')
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
        # only list adapter
        self.list_adapter()

    def raid0(self, config):
        raid_index = self._get_raid_index(config.raid_level)
        adapter_id = config.adapter_id
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid1(self, config):
        raid_index = self._get_raid_index(config.raid_level)
        adapter_id = config.adapter_id
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid5(self, config):
        raid_index = self._get_raid_index(config.raid_level)
        adapter_id = config.adapter_id
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def raid10(self, config):
        raid_index = self._get_raid_index(config.raid_level)
        adapter_id = config.adapter_id
        disk_location = self._parse_physical_disks(config.physical_disks)
        self._execute_raid(adapter_id, raid_index, disk_location)

    def clean_raid(self, adapter_id: str = '0', delay_time=0):
        time.sleep(30)
        cmd = (EXECUTABLE_FILE, adapter_id, 'delete', 'noprompt')
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Delete %s operation : %s : %s.', adapter_id, stdout,
                 stderr)
        time.sleep(delay_time)

    def _parse_disk(self, stdout):
        disks = []
        pattern = re.compile(
            r'\s+Enclosure\s\#\s+:\s+(?P<enclosure>\d+)\s.*?'
            r'\s+Slot\s\#\s+:\s+(?P<slot>\d+)\s.*?'
            r'\s+Size\s\(in\s+(?P<size_unit>\w+).*?:\s+(?P<size>\S+)/.*?'
            r'\s+Protocol\s+:\s+(?P<pd_type>\S+)\s.*?'
            r'\s+Drive Type\s+:\s+(?P<media_type>\S+)\s.*?', re.DOTALL)
        for match in re.finditer(pattern, stdout):
            enclosure = int(match.group('enclosure').strip())
            slot = int(match.group('slot').strip())
            pd_type = match.group('pd_type').strip()
            media_type = match.group(
                'media_type').strip().rsplit("_", 1)[-1].upper()
            size = float(match.group('size').strip())
            size_unit = match.group('size_unit').strip()
            disks.append(
                Disk(enclosure=enclosure,
                     slot=slot,
                     pd_type=pd_type,
                     media_type=media_type,
                     size=size,
                     size_unit=size_unit).__dict__)
        return disks

    def _get_raid_index(self, raid_level):
        raid_index = RAID_INDEX.get(raid_level)
        if not raid_index:
            raise IncompatibleHardwareMethodError('Not support raid level %s.' %
                                                  raid_level)
        return raid_index

    def _execute_raid(self, adapter_id, raid_index, disk_location, other_param=''):
        cmd = ((
            EXECUTABLE_FILE,
            adapter_id,
            'create',
            'RAID%s' % raid_index,
            'MAX') +
            disk_location + (
            'noprompt',
            other_param)
        )
        stdout, stderr = executor.execute(*cmd)
        LOG.info('Create %s raid%s : %s', raid_index, stdout, stderr)

    def _parse_physical_disks(self, physical_disks):
        disk_strings = [
            f"{disk['enclosure']}:{disk['slot']}" for disk in physical_disks]
        return tuple(disk_strings)
