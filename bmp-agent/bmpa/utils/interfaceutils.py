#!/usr/bin/env python
# -*- coding: utf-8 -*-

import binascii
import ctypes
import fcntl
import logging
import os
import select
import socket
import struct
import sys
import time

import netifaces
import pyudev

from bmpa.errors import BmpError, IncompatibleHardwareMethodError
from bmpa.serialize import SerializableComparable
from bmpa.utils import executor
from bmpa.utils import hardware
from bmpa.utils import netutils


LOG = logging.getLogger(__name__)

DEFAULT_DHCP_WAIT_TIMEOUT = 600
DHCP_RETRY_INTERVAL = 2

LLDP_TIMEOUT = 60.0

LLDP_ETHERTYPE = 0x88cc
IFF_PROMISC = 0x100
SIOCGIFFLAGS = 0x8913
SIOCSIFFLAGS = 0x8914
DMAC_TYPE_NEAREST_BRIDGE = 0xe


class ifreq(ctypes.Structure):
    """Class for setting flags on a socket."""
    _fields_ = [("ifr_ifrn", ctypes.c_char * 16),
                ("ifr_flags", ctypes.c_short)]


class RawPromiscuousSockets(object):
    def __init__(self, interface_names, protocol):
        """Initialize context manager.

        :param interface_names: a list of interface names to bind to
        :param protocol: the protocol to listen for
        :returns: A list of tuple of (interface_name, bound_socket), or [] if
                  there is an exception binding or putting the sockets in
                  promiscuous mode
        """
        if not interface_names:
            raise ValueError('interface_names must be a non-empty list of '
                             'network interface names to bind to.')
        self.protocol = protocol
        # A 3-tuple of (interface_name, socket, ifreq object)
        self.interfaces = [(name, self._get_socket(), ifreq())
                           for name in interface_names]

    def __enter__(self):
        for interface_name, sock, ifr in self.interfaces:
            LOG.info(
                'Interface %s entering promiscuous mode to capture ', interface_name)
            try:
                ifr.ifr_ifrn = interface_name.encode()
                # Get current flags
                fcntl.ioctl(sock.fileno(), SIOCGIFFLAGS, ifr)  # G for Get
                # bitwise or the flags with promiscuous mode, set the new flags
                ifr.ifr_flags |= IFF_PROMISC
                fcntl.ioctl(sock.fileno(), SIOCSIFFLAGS, ifr)  # S for Set
                # Bind the socket so it can be used
                LOG.debug(
                    'Binding interface %(interface)s for protocol '
                    '%(proto)s', {
                        'interface': interface_name,
                        'proto': self.protocol
                    })
                sock.bind((interface_name, self.protocol))
            except Exception:
                LOG.error('Failed to open all RawPromiscuousSockets, '
                          'attempting to close any opened sockets.')
                self.__exit__(*sys.exc_info())
                raise

        # No need to return each interfaces ifreq.
        return [(sock[0], sock[1]) for sock in self.interfaces]

    def __exit__(self, exception_type, exception_val, trace):
        for name, sock, ifr in self.interfaces:
            # bitwise or with the opposite of promiscuous mode to remove
            ifr.ifr_flags &= ~IFF_PROMISC
            try:
                fcntl.ioctl(sock.fileno(), SIOCSIFFLAGS, ifr)
                sock.close()
            except Exception:
                LOG.exception('Failed to close raw socket for interface %s',
                              name)

    def _get_socket(self):
        return socket.socket(socket.AF_PACKET, socket.SOCK_RAW, self.protocol)


class NetworkInterface(SerializableComparable):
    serializable_fields = ('name', 'mac_address', 'ipv4_address',
                           'has_carrier', 'lldp', 'vendor', 'product',
                           'client_id', 'biosdevname', 'switch_index',
                           'switch_manage_ip', 'bus_info')

    def __init__(self,
                 name,
                 mac_addr,
                 ipv4_address=None,
                 has_carrier=True,
                 lldp=None,
                 vendor=None,
                 product=None,
                 client_id=None,
                 biosdevname=None,
                 switch_index=None,
                 switch_manage_ip=None,
                 bus_info=None):
        self.name = name
        self.mac_address = mac_addr
        self.ipv4_address = ipv4_address
        self.has_carrier = has_carrier
        self.lldp = lldp
        self.vendor = vendor
        self.product = product
        self.biosdevname = biosdevname
        # client_id is used for InfiniBand only. we calculate the DHCP
        # client identifier Option to allow DHCP to work over InfiniBand.
        # see https://tools.ietf.org/html/rfc4390
        self.client_id = client_id
        self.switch_index = switch_index
        self.switch_manage_ip = switch_manage_ip
        self.bus_info = bus_info

    def __repr__(self):
        return "NetworkInterface %s" % self.serialize()


def get_lldp_info(interface_names):
    """Get LLDP info from the switch(es) the agent is connected to.

    Listens on either a single or all interfaces for LLDP packets, then
    parses them. If no LLDP packets are received before lldp_timeout,
    returns a dictionary in the form {'interface': [],...}.

    :param interface_names: The interface to listen for packets on. If
                           None, will listen on each interface.
    :return: A dictionary in the form
             {'interface': [(lldp_type, lldp_data)],...}
    """
    with RawPromiscuousSockets(interface_names, LLDP_ETHERTYPE) as interfaces:
        try:
            return _get_lldp_info(interfaces)
        except Exception as e:
            LOG.exception('Error while getting LLDP info: %s', str(e))
            raise


def _get_lldp_info(interfaces):
    """Wait for packets on each socket, parse the received LLDP packets.

    """
    LOG.debug('Getting LLDP info for interfaces %s', interfaces)

    lldp_info = {}
    if not interfaces:
        return {}

    while interfaces:
        LOG.info(
            'Waiting on LLDP info for interfaces: %(interfaces)s, '
            'timeout: %(timeout)s', {
                'interfaces': interfaces,
                'timeout': LLDP_TIMEOUT
            })

        socks = [interface[1] for interface in interfaces]
        # rlist is a list of sockets ready for reading
        rlist, _, _ = select.select(socks, [], [], LLDP_TIMEOUT)
        if not rlist:
            # Empty read list means timeout on all interfaces
            LOG.warning('LLDP timed out, remaining interfaces: %s',
                        interfaces)
            break

        for s in rlist:
            # Find interface name matching socket ready for read
            # Create a copy of interfaces to avoid deleting while iterating.
            for index, interface in enumerate(list(interfaces)):
                if s == interface[1]:
                    try:
                        dmac_type, lldp_info[
                            interface[0]] = _receive_lldp_packets(s)
                    except socket.error:
                        LOG.exception(
                            'Socket for network interface %s said '
                            'that it was ready to read we were '
                            'unable to read from the socket while '
                            'trying to get LLDP packet. Skipping '
                            'this network interface.', interface[0])
                    else:
                        LOG.info('Found LLDP info for interface: %s',
                                 interface[0])
                    # Remove interface from the list, only need one packet
                    if dmac_type == DMAC_TYPE_NEAREST_BRIDGE:
                        del interfaces[index]

    # Add any interfaces that didn't get a packet as empty lists
    for name, _sock in interfaces:
        lldp_info[name] = []

    return lldp_info


def collect_lldp_data(interfaces_names):
    """Collect and convert LLDP info from the node.

    In order to process the LLDP information later, the raw data needs to
    be converted for serialization purposes.

    :param interface_names: list of names of node's interfaces.
    :return: a dict, containing the lldp data from every interface.
    """
    interface_names = [name for name in interfaces_names if name != 'lo']
    lldp_data = {}
    try:
        raw_lldp_data = get_lldp_info(interface_names)
    except Exception:
        # NOTE(sambetts) The get_lldp_info function will log this exception
        # and we don't invalidate any existing data in the cache if we fail
        # to get data to replace it so just return.
        return lldp_data
    for ifname, tlvs in raw_lldp_data.items():
        # NOTE(sambetts) Convert each type-length-value (TLV) value to hex
        # so that it can be serialised safely
        processed_tlvs = []
        for typ, data in tlvs:
            try:
                #    binascii.hexlify(data).decode()))
                processed_tlvs.append((typ, data))
            except (binascii.Error, binascii.Incomplete) as e:
                LOG.warning(
                    'An error occurred while processing TLV type '
                    '%s for interface %s: %s', typ, ifname, e)
        lldp_data[ifname] = processed_tlvs
    return lldp_data


def _receive_lldp_packets(sock):
    """Receive LLDP packets and process them.

    :param sock: A bound socket
    :return: A list of tuples in the form (lldp_type, lldp_data)
    """
    pkt = sock.recv(1600)
    # Filter invalid packets
    if not pkt or len(pkt) < 14:
        return []
    dmac_type = struct.unpack('!H', pkt[4:6])[0]
    # Skip header (dst MAC, src MAC, ethertype)
    pkt = pkt[14:]
    return dmac_type, _parse_tlv(pkt)


def _parse_tlv(buff):
    """Iterate over a buffer and generate structured TLV data.

    :param buff: An ethernet packet with the header trimmed off (first
                 14 bytes)
    """
    lldp_info = []
    while len(buff) >= 2:
        # TLV structure: type (7 bits), length (9 bits), val (0-511 bytes)
        tlvhdr = struct.unpack('!H', buff[:2])[0]
        tlvtype = (tlvhdr & 0xfe00) >> 9
        tlvlen = (tlvhdr & 0x01ff)
        tlvdata = buff[2:tlvlen + 2]
        buff = buff[tlvlen + 2:]
        lldp_info.append((tlvtype, tlvdata))

    if buff:
        LOG.warning("Trailing byte received in an LLDP package: %r", buff)
    return lldp_info


def get_interface_bus_info(interface_name):
    context = pyudev.Context()
    for device in context.list_devices(subsystem='net'):
        if device.get('INTERFACE') == interface_name:
            path = device.get('ID_PATH')
            if path:
                parts = path.split('-')
                return parts[1] if len(parts) > 1 else path
    return None


def get_interface_info(interface_name):
    mac_addr = get_mac_addr(interface_name)
    if mac_addr is None:
        raise IncompatibleHardwareMethodError()

    return NetworkInterface(
        interface_name,
        mac_addr,
        ipv4_address=get_ipv4_addr(interface_name),
        has_carrier=interface_has_carrier(interface_name),
        vendor=hardware.get_device_info(interface_name, 'net',
                                        'vendor'),
        product=hardware.get_device_info(interface_name, 'net',
                                         'device'),
        biosdevname=hardware.get_bios_given_nic_name(interface_name),
        bus_info=get_interface_bus_info(interface_name))


def list_network_interfaces_meta(include_lldp=True, is_up_eth=True):
    network_interfaces = []
    interface_names = list_interface_name()
    if is_up_eth:
        for if_name in interface_names:
            is_inactive = is_inactive_eth(if_name)
            if is_inactive:
                try:
                    executor.execute(
                        'ifconfig %s up' % if_name, shell=True)
                finally:
                    LOG.info('ifconfig %s up', if_name)
    lldp_data = None
    if include_lldp:
        lldp_data = collect_lldp_data(interface_names)

    for if_name in interface_names:
        result = get_interface_info(if_name)
        if (lldp_data and type(lldp_data) is dict
                and if_name in lldp_data.keys()):
            lldp = lldp_data.get(if_name)
            if len(lldp) <= 0:
                continue
            switch_index = parse_switch_index(lldp)
            if switch_index is None:
                continue
            LOG.info("%s switch index %s", if_name, switch_index)
            switch_manage_ip = parse_switch_management_ip(lldp)
            LOG.info("%s switch management ip %s", if_name,
                     switch_manage_ip)
            result.switch_index = switch_index
            result.switch_manage_ip = switch_manage_ip
        network_interfaces.append(result)

    return network_interfaces


def parse_switch_index(lldp):
    """doc: https://en.wikipedia.org/wiki/Link_Layer_Discovery_Protocol

     tlv type:
     0	End of LLDPDU	Mandatory
     1	Chassis ID	Mandatory
     2	Port ID	Mandatory
     3	Time To Live	Mandatory
     4	Port description	Optional
     5	System name	Optional
     6	System description	Optional
     7	System capabilities	Optional
     8	Management address	Optional
     9–126	Reserved	-
     127	Custom TLVs	Optional

     data struct：
     [(1, '\x04\xdc!\xe2u\xa01'),
     (2, '\x0525GE1/0/25'),
     (3, '\x00x'),
     (5, 'SQV03_Physics-Cloud_T4-L3-161_TOR-HUAWEI-25G'),
     (6, 'Huawei Versatile Routing Platform Software\r\nVRP (R) software, Version 8.180 (CE6865EI V200R005C10SPC800)\r\nCopyright (C) 2012-2018 Huawei Technologies Co., Ltd.\r\nHUAWEI CE6865-48S8CQ-EI\r\n'),
     (7, '\x00\x14\x00\x14'),
     (8, '\x05\x01\n\xd0\x08Y\x02\x00\x00\x00>\x11\x06\x0f+\x06\x01\x04\x01\x8f[\x05\x19)\x01\x02\x01\x01\x01'),
     (127, '\x00\x80\xc2\x01\x00d'), (127, '\x00\x80\xc2\x02\x00\x00\x00'),
     (127, '\x00\x80\xc2\x03\x00d\x07VLAN100'),
     (127, '\x00\x12\x0f\x01\x00\x00\x00\x00\x00'),
     (127, '\x00\x12\x0f\x03\x01\x00\x00\x00\x00'),
     (127, '\x00\x12\x0f\x04$\x00'), (0, '')]
    """

    for tlv in lldp:
        if tlv[0] == 2:
            if isinstance(tlv[1][1:], bytes):
                try:
                    return tlv[1][1:].decode('utf-8')
                except Exception as e:
                    LOG.info("Invalid lldp info : %s, err:%s.", lldp, e)
                    return None

            return tlv[1][1:]
    LOG.info("Invalid lldp info : %s." % lldp)
    return None


def parse_switch_management_ip(lldp):
    for tlv in lldp:
        if tlv[0] == 8:
            management_ip = tlv[1][2:6]
            return "%d.%d.%d.%d" % (management_ip[0],
                                    management_ip[1], management_ip[2],
                                    management_ip[3])
    raise BmpError("Invalid lldp info : %s." % lldp)


def list_interface_name():
    iface_names = os.listdir('/sys/class/net')
    iface_names = [
        name for name in iface_names if hardware.is_device(name)
    ]
    return iface_names


def get_interface_flags(interface_name):
    """Get interface flags.

    :param  iface_name
    :type   string

    """
    try:
        with open('/sys/class/net/%s/flags' % interface_name, 'r') as f:
            return f.read().strip()
    except IOError:
        LOG.warn("Can't find interface %s flags", interface_name)


def is_inactive_eth(interface_name):
    """Ethernet adapter is not activated.

    :param  iface_name
    :type   string

    """
    status = get_interface_flags(interface_name)
    if status is not None:
        if status == "0x1002":
            return True
    return False


def list_network_interfaces():
    network_interfaces = []
    iface_names = list_interface_name()

    for iface_name in iface_names:
        result = get_interface_info(interface_name=iface_name)
        network_interfaces.append(result)
    return network_interfaces


def get_ipv4_addr(if_name):
    try:
        addrs = netifaces.ifaddresses(if_name)
        return addrs[netifaces.AF_INET][0]['addr']
    except (ValueError, IndexError, KeyError):
        # No default IPv4 address found
        return None


def get_mac_addr(interface_name):
    try:
        addrs = netifaces.ifaddresses(interface_name)
        return addrs[netifaces.AF_LINK][0]['addr']
    except (ValueError, IndexError, KeyError):
        # No mac address found
        return None


def interface_has_carrier(interface_name):
    path = '/sys/class/net/{}/carrier'.format(interface_name)
    try:
        with open(path, 'rt') as fp:
            return fp.read().strip() == '1'
    except EnvironmentError:
        LOG.debug('No carrier information for interface %s', interface_name)
        return False


def wrap_ipv6(ip):
    if netutils.is_valid_ipv6(ip):
        return "[%s]" % ip
    return ip


def get_wildcard_address():
    if netutils.is_ipv6_enabled():
        return "::"
    return "0.0.0.0"


def _normalize_mac(mac):
    """Convert MAC to a well-known format aa:bb:cc:dd:ee:ff."""
    if '-' in mac:
        # pxelinux format is 01-aa-bb-cc-dd-ee-ff
        mac = mac.split('-', 1)[1]
        mac = mac.replace('-', ':')
    return mac.lower()


def wait_for_dhcp():
    """Wait until NIC's get their IP addresses via DHCP or timeout happens.

    Depending on the value of inspection_dhcp_all_interfaces configuration
    option will wait for either all or only PXE booting NIC.

    Note: only supports IPv4 addresses for now.

    :return: True if all NIC's got IP addresses, False if timeout happened.
             Also returns True if waiting is disabled via configuration.
    """
    threshold = time.time() + DEFAULT_DHCP_WAIT_TIMEOUT
    while time.time() <= threshold:
        interfaces = list_network_interfaces_meta(include_lldp=False)
        for iface in interfaces:
            if iface.ipv4_address:
                LOG.info("capture ipv4 address : %s", iface)
                return (True, iface.name)

        LOG.debug('Still waiting for interfaces to get IP addresses')
        time.sleep(DHCP_RETRY_INTERVAL)

    LOG.warning(
        'Not all network interfaces received IP addresses in \
                %(timeout)d seconds', {'timeout': DEFAULT_DHCP_WAIT_TIMEOUT})
    return (False, None)


def dhclient():
    """Run dhclient

    """
    try:
        executor.execute('dhclient', shell=True)
    except Exception as e:
        LOG.error('Run dhclient error:%s.' % e)
