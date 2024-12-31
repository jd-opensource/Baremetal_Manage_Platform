import copy
import logging
# python3 code
from urllib.request import pathname2url

from bmpa.errors import DeviceNotFound
from bmpa.utils import hardware
from bmpa.utils import specs_matcher
from bmpa.utils import strutils

LOG = logging.getLogger(__name__)

# A dictionary in the form {hint name: hint type}
VALID_DEVICE_HINTS = {
    'size': int,
    'model': str,
    'wwn': str,
    'serial': str,
    'vendor': str,
    'wwn_with_extension': str,
    'wwn_vendor_extension': str,
    'name': str,
    'rotational': bool,
    'hctl': str,
    'by_path': str,
}

NOT_VALID_DEVICE_HINTS = [
    'rotational', 'hctl', 'name', 'wwn_with_extension', 'wwn_vendor_extension'
]

DEVICE_HINTS_GRAMMAR = specs_matcher.make_grammar()


def get_device_by_hints(device_hints: dict, validate_all_default_hints: bool = False) -> hardware.BlockDevice:
    block_devices = hardware.list_all_block_devices()
    if (not device_hints) or len(device_hints) == 0:
        raise DeviceNotFound("device_hints shoud not be None.")
    else:
        serialized_devs = [dev.serialize() for dev in block_devices]
        try:
            device = match_device_hints(
                serialized_devs,
                device_hints,
                validate_all_default_hints=validate_all_default_hints)
        except ValueError as e:
            raise DeviceNotFound(
                'No devices could be found using the device hints '
                '%(hints)s because they failed to validate. Error: '
                '%(error)s' % {
                    'hints': device_hints,
                    'error': e
                })

        if not device:
            raise DeviceNotFound(
                "No suitable device was found for "
                "deployment using these hints %s" % device_hints)

    dev_name = device['name']

    LOG.info(
        'Picked device %(dev)s based on '
        'device hints %(hints)s', {
            'dev': dev_name,
            'hints': device_hints
        })
    return device


def match_device_hints(devices,
                       device_hints,
                       validate_all_default_hints=True) -> hardware.BlockDevice:
    """Try to find a device that matches the device hints.

    Try to find a device that matches the device hints. In order
    for a device to be matched it needs to satisfy all the given hints.

    :param devices: A list of dictionaries representing the devices
                    containing one or more of the following keys:

        :name: (String) The device name, e.g /dev/sda
        :size: (Integer) Size of the device in *bytes*
        :model: (String) Device model
        :vendor: (String) Device vendor name
        :serial: (String) Device serial number
        :wwn: (String) Unique storage identifier
        :wwn_with_extension: (String): Unique storage identifier with
                             the vendor extension appended
        :wwn_vendor_extension: (String): United vendor storage identifier
        :rotational: (Boolean) Whether it's a rotational device or
                     not. Useful to distinguish HDDs (rotational) and SSDs
                     (not rotational).
        :hctl: (String): The SCSI address: Host, channel, target and lun.
                         For example: '1:0:0:0'.
        :by_path: (String): The alternative device name,
                  e.g. /dev/disk/by-path/pci-0000:00

    :param device_hints: A dictionary with the device hints.
    :raises: ValueError, if some information is invalid.
    :returns: The first device to match all the hints or None.
    """
    try:
        dev = next(
            find_devices_by_hints(devices, device_hints,
                                  validate_all_default_hints))
    except StopIteration:
        LOG.warning('No device found that matches the device hints %s',
                    device_hints)
    else:
        LOG.info(
            'Device found! The device "%s" matches the '
            'device hints %s', dev, device_hints)
        return dev


def parse_device_hints(device, validate_all_default_hints=True):
    """Parse the device property of a node.

    Parses and validates the device property of a node. These are
    hints for how a node's device is created. The 'size' hint
    should be a positive integer. The 'rotational' hint should be a
    Boolean value.

    :param device: the device dictionary from the node's property.
    :returns: a dictionary with the device hints parsed or
              None if there are no hints.
    :raises: ValueError, if some information is invalid.

    """
    if not device:
        return

    device = copy.deepcopy(device)
    if validate_all_default_hints:
        invalid_hints = set(device) - set(VALID_DEVICE_HINTS)
        if invalid_hints:
            raise ValueError(
                ('The hints "%(invalid_hints)s" are invalid. '
                 'Valid hints are: "%(valid_hints)s"') % {
                     'invalid_hints': ', '.join(invalid_hints),
                     'valid_hints': ', '.join(VALID_DEVICE_HINTS)
                })

    for name, expression in device.items():
        hint_type = VALID_DEVICE_HINTS[name]
        if hint_type is str:
            if not isinstance(expression, str):
                raise ValueError(
                    ('Device hint "%(name)s" is not a string value. '
                     'Hint expression: %(expression)s') % {
                         'name': name,
                         'expression': expression
                    })
            device[name] = _normalize_hint_expression(expression, name)

        elif hint_type is int:
            for v in _extract_hint_operator_and_values(expression,
                                                       name)['values']:
                try:
                    integer = int(v)
                except ValueError:
                    raise ValueError(
                        ('Device hint "%(name)s" is not an integer '
                         'value. Current value: %(expression)s') % {
                             'name': name,
                             'expression': expression
                        })

                if integer <= 0:
                    raise ValueError(
                        ('Device hint "%(name)s" should be a positive '
                         'integer. Current value: %(expression)s') % {
                             'name': name,
                             'expression': expression
                        })

        elif hint_type is bool:
            try:
                device[name] = strutils.bool_from_string(
                    expression, strict=True)
            except ValueError:
                raise ValueError(
                    ('Device hint "%(name)s" is not a Boolean value. '
                     'Current value: %(expression)s') % {
                         'name': name,
                         'expression': expression
                    })

    return _append_operator_to_hints(device)


def find_devices_by_hints(devices,
                          device_hints,
                          validate_all_default_hints=True):
    """Find all devices that match the device hints.

    Try to find devices that match the device hints. In order
    for a device to be matched it needs to satisfy all the given hints.

    :param devices: A list of dictionaries representing the devices
                    containing one or more of the following keys:

        :name: (String) The device name, e.g /dev/sda
        :size: (Integer) Size of the device in *bytes*
        :model: (String) Device model
        :vendor: (String) Device vendor name
        :serial: (String) Device serial number
        :wwn: (String) Unique storage identifier
        :wwn_with_extension: (String): Unique storage identifier with
                             the vendor extension appended
        :wwn_vendor_extension: (String): United vendor storage identifier
        :rotational: (Boolean) Whether it's a rotational device or
                     not. Useful to distinguish HDDs (rotational) and SSDs
                     (not rotational).
        :hctl: (String): The SCSI address: Host, channel, target and lun.
                         For example: '1:0:0:0'.
        :by_path: (String): The alternative device name,
                  e.g. /dev/disk/by-path/pci-0000:00

    :param device_hints: A dictionary with the device hints.
    :raises: ValueError, if some information is invalid.
    :returns: A generator with all matching devices as dictionaries.
    """
    LOG.info(
        'Trying to find devices from "%(devs)s" that match the '
        'device hints "%(hints)s"', {
            'devs': ', '.join([d.get('name') for d in devices]),
            'hints': device_hints
        })
    parsed_hints = parse_device_hints(device_hints)
    LOG.info("parsed_hints:%s", parsed_hints)
    for dev in devices:
        device_name = dev.get('name')

        for hint in parsed_hints:
            hint_type = VALID_DEVICE_HINTS[hint]
            device_value = dev.get(hint)
            hint_value = parsed_hints[hint]

            if hint_type is str:
                try:
                    device_value = _normalize_hint_expression(
                        device_value, hint)
                except ValueError:
                    LOG.warning(
                        'The attribute "%(attr)s" of the device "%(dev)s" '
                        'has an empty value. Skipping device.', {
                            'attr': hint,
                            'dev': device_name
                        })
                    break

            if validate_all_default_hints:
                if hint in NOT_VALID_DEVICE_HINTS:
                    continue

            LOG.info(
                'Trying to match the device hint "%(hint)s" '
                'with a value of "%(hint_value)s" against the same '
                'device\'s (%(dev)s) attribute with a value of '
                '"%(dev_value)s"', {
                    'hint': hint,
                    'dev': device_name,
                    'hint_value': hint_value,
                    'dev_value': device_value
                })

            # Boolean hints are not supported by
            # specs_matcher.match(), so we need to do the comparison
            # ourselves
            if hint_type is bool:
                try:
                    device_value = strutils.bool_from_string(
                        device_value, strict=True)
                except ValueError:
                    LOG.warning(
                        'The attribute "%(attr)s" (with value '
                        '"%(value)s") of device "%(dev)s" is not '
                        'a valid Boolean. Skipping device.', {
                            'attr': hint,
                            'value': device_value,
                            'dev': device_name
                        })
                    break
                if device_value == hint_value:
                    continue

            elif specs_matcher.match(device_value, hint_value):
                continue

            LOG.warn(
                'The attribute "%(attr)s" (with value "%(value)s") '
                'of device "%(dev)s" does not match the hint %(hint)s', {
                    'attr': hint,
                    'value': device_value,
                    'dev': device_name,
                    'hint': hint_value
                })
            break
        else:
            LOG.info("yield dev")
            yield dev


def _extract_hint_operator_and_values(hint_expression, hint_name):
    """Extract the operator and value(s) of a device hint expression.

    A device hint expression could contain one or more values
    depending on the operator. This method extracts the operator and
    value(s) and returns a dictionary containing both.

    :param hint_expression: The hint expression string containing value(s)
                            and operator (optionally).
    :param hint_name: The name of the hint. Used for logging.
    :raises: ValueError if the hint_expression is empty.
    :returns: A dictionary containing:

        :op: The operator. An empty string in case of None.
        :values: A list of values stripped and converted to lowercase.
    """
    expression = str(hint_expression).strip().lower()
    if not expression:
        raise ValueError(
            ('Device hint "%s" expression is empty') % hint_name)

    # parseString() returns a list of tokens which the operator (if
    # present) is always the first element.
    ast = DEVICE_HINTS_GRAMMAR.parseString(expression)
    if len(ast) <= 1:
        # hint_expression had no operator
        return {'op': '', 'values': [expression]}

    op = ast[0]
    return {
        'values': [v.strip() for v in re.split(op, expression) if v],
        'op': op
    }


def _normalize_hint_expression(hint_expression, hint_name):
    """Normalize a string type hint expression.

    A string-type hint expression contains one or more operators and
    one or more values: [<op>] <value> [<op> <value>]*. This normalizes
    the values by url-encoding white spaces and special characters. The
    operators are not normalized. For example: the hint value of "<or>
    foo bar <or> bar" will become "<or> foo%20bar <or> bar".

    :param hint_expression: The hint expression string containing value(s)
                            and operator (optionally).
    :param hint_name: The name of the hint. Used for logging.
    :raises: ValueError if the hint_expression is empty.
    :returns: A normalized string.
    """
    hdict = _extract_hint_operator_and_values(hint_expression, hint_name)
    result = hdict['op'].join(
        [' %s ' % pathname2url(t) for t in hdict['values']])
    return (hdict['op'] + result).strip()


def _append_operator_to_hints(device):
    """Add an equal (s== or ==) operator to the hints.

    For backwards compatibility, for device hints where no operator
    means equal, this method adds the equal operator to the hint. This is
    needed when using oslo.utils.specs_matcher methods.

    :param device: The device hints dictionary.
    """
    for name, expression in device.items():
        # The specs_matcher does not support boolean,
        # so we don't need to append any operator for it.
        if VALID_DEVICE_HINTS[name] is bool:
            continue

        expression = str(expression)
        ast = DEVICE_HINTS_GRAMMAR.parseString(expression)
        if len(ast) > 1:
            continue

        op = 's== %s' if VALID_DEVICE_HINTS[name] is str else '== %s'
        device[name] = op % expression

    return device
