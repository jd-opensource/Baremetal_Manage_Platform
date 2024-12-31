import os
import re

import netaddr
import six

_IS_IPV6_ENABLED = None


def is_valid_ipv4(address):
    """Verify that address represents a valid IPv4 address.

    :param address: Value to verify
    :type address: string
    :returns: bool

    .. versionadded:: 1.1
    """
    try:
        return netaddr.valid_ipv4(address)
    except netaddr.AddrFormatError:
        return False


def is_valid_ipv6(address):
    """Verify that address represents a valid IPv6 address.

    :param address: Value to verify
    :type address: string
    :returns: bool

    .. versionadded:: 1.1
    """
    if not address:
        return False

    parts = address.rsplit("%", 1)
    address = parts[0]
    scope = parts[1] if len(parts) > 1 else None
    if scope is not None and (len(scope) < 1 or len(scope) > 15):
        return False

    try:
        return netaddr.valid_ipv6(address, netaddr.core.INET_PTON)
    except netaddr.AddrFormatError:
        return False


def is_valid_cidr(address):
    """Verify that address represents a valid CIDR address.

    :param address: Value to verify
    :type address: string
    :returns: bool

    .. versionadded:: 3.8
    """
    try:
        # Validate the correct CIDR Address
        netaddr.IPNetwork(address)
    except (TypeError, netaddr.AddrFormatError):
        return False

    # Prior validation partially verify /xx part
    # Verify it here
    ip_segment = address.split('/')

    if (len(ip_segment) <= 1 or ip_segment[1] == ''):
        return False

    return True


def is_valid_ip(address):
    """Verify that address represents a valid IP address.

    :param address: Value to verify
    :type address: string
    :returns: bool

    .. versionadded:: 1.1
    """
    return is_valid_ipv4(address) or is_valid_ipv6(address)


def is_valid_mac(address):
    """Verify the format of a MAC address.

    Check if a MAC address is valid and contains six octets. Accepts
    colon-separated format only.

    :param address: MAC address to be validated.
    :returns: True if valid. False if not.

    .. versionadded:: 3.17
    """
    m = "[0-9a-f]{2}(:[0-9a-f]{2}){5}$"
    return (isinstance(address, six.string_types)
            and re.match(m, address.lower()))


def _is_int_in_range(value, start, end):
    """Try to convert value to int and check if it lies within

    range 'start' to 'end'.

    :param value: value to verify
    :param start: start number of range
    :param end: end number of range
    :returns: bool
    """
    try:
        val = int(value)
    except (ValueError, TypeError):
        return False
    return (start <= val <= end)


def is_valid_port(port):
    """Verify that port represents a valid port number.

    Port can be valid integer having a value of 0 up to and
    including 65535.

    .. versionadded:: 1.1.1
    """
    return _is_int_in_range(port, 0, 65535)


def is_ipv6_enabled():
    """Check if IPv6 support is enabled on the platform.

    This api will look into the proc entries of the platform to figure
    out the status of IPv6 support on the platform.

    :returns: True if the platform has IPv6 support, False otherwise.

    .. versionadded:: 1.4
    """

    global _IS_IPV6_ENABLED

    if _IS_IPV6_ENABLED is None:
        disabled_ipv6_path = "/proc/sys/net/ipv6/conf/default/disable_ipv6"
        if os.path.exists(disabled_ipv6_path):
            with open(disabled_ipv6_path, 'r') as f:
                disabled = f.read().strip()
            _IS_IPV6_ENABLED = disabled == "0"
        else:
            _IS_IPV6_ENABLED = False
    return _IS_IPV6_ENABLED


def netmask_to_cidr(netmask):
    """netmast to cidr

    :param netmask: netmask ip addr (eg: 255.255.255.0)
    :return: equivalent cidr number to given netmask ip (eg: 24)
    """
    return sum([bin(int(x)).count('1') for x in netmask.split('.')])


def ip_to_cidr(ip_addr, netmask):
    """ip to cidr

    :param ip_addr: ip addr (eg: 10.208.14.165)
    :param netmask: netmask ip addr (eg: 255.255.255.224)
    :return: ip/cidr  (eg: 10.208.14.165/24)
    """
    cidr = netmask_to_cidr(netmask)
    return '%s/%s' % (ip_addr, cidr)
