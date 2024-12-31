import math
import re

import six

UNIT_SYSTEM_INFO = {
    'IEC': (1024, re.compile(r'(^[-+]?\d*\.?\d+)([KMGT]i?)?(b|bit|B)$')),
    'SI': (1000, re.compile(r'(^[-+]?\d*\.?\d+)([kMGT])?(b|bit|B)$')),
    'mixed': (None, re.compile(r'(^[-+]?\d*\.?\d+)([kKMGT]i?)?(b|bit|B)$')),
}

UNIT_PREFIX_EXPONENT = {
    'k': 1,
    'K': 1,
    'Ki': 1,
    'M': 2,
    'Mi': 2,
    'G': 3,
    'Gi': 3,
    'T': 4,
    'Ti': 4,
}

TRUE_STRINGS = ('1', 't', 'true', 'on', 'y', 'yes')
FALSE_STRINGS = ('0', 'f', 'false', 'off', 'n', 'no')


def string_to_bytes(text, unit_system='IEC', return_int=False):
    """Converts a string into an float representation of bytes.

    The units supported for IEC / mixed::

        Kb(it), Kib(it), Mb(it), Mib(it), Gb(it), Gib(it), Tb(it), Tib(it)
        KB, KiB, MB, MiB, GB, GiB, TB, TiB

    The units supported for SI ::

        kb(it), Mb(it), Gb(it), Tb(it)
        kB, MB, GB, TB

    SI units are interpreted as power-of-ten (e.g. 1kb = 1000b).  Note
    that the SI unit system does not support capital letter 'K'

    IEC units are interpreted as power-of-two (e.g. 1MiB = 1MB =
    1024b)

    Mixed units interpret the "i" to mean IEC, and no "i" to mean SI
    (e.g. 1kb = 1000b, 1kib == 1024b).  Additionaly, mixed units
    interpret 'K' as power-of-ten.  This mode is not particuarly
    useful for new code, but can help with compatability for parsers
    such as GNU parted.

    :param text: String input for bytes size conversion.
    :param unit_system: Unit system for byte size conversion.
    :param return_int: If True, returns integer representation of text
                       in bytes. (default: decimal)
    :returns: Numerical representation of text in bytes.
    :raises ValueError: If text has an invalid value.

    """
    try:
        base, reg_ex = UNIT_SYSTEM_INFO[unit_system]
    except KeyError:
        msg = 'Invalid unit system: "%s"' % unit_system
        raise ValueError(msg)
    match = reg_ex.match(text)
    if match:
        magnitude = float(match.group(1))
        unit_prefix = match.group(2)
        if match.group(3) in ['b', 'bit']:
            magnitude /= 8

        # In the mixed matcher, IEC units (with a trailing 'i') are
        # interpreted as power-of-two, others as power-of-ten
        if unit_system == 'mixed':
            if unit_prefix and not unit_prefix.endswith('i'):
                # For maximum compatability in mixed mode, we understand
                # "K" (which is not strict SI) as "k"
                if unit_prefix.startswith == 'K':
                    unit_prefix = 'k'
                base = 1000
            else:
                base = 1024
    else:
        msg = 'Invalid string format: %s' % text
        raise ValueError(msg)

    if not unit_prefix:
        res = magnitude
    else:
        res = magnitude * pow(base, UNIT_PREFIX_EXPONENT[unit_prefix])
    if return_int:
        return int(math.ceil(res))
    return res


def bool_from_string(subject, strict=False, default=False):
    """Interpret a subject as a boolean.

    A subject can be a boolean, a string or an integer. Boolean type value
    will be returned directly, otherwise the subject will be converted to
    a string. A case-insensitive match is performed such that strings
    matching 't','true', 'on', 'y', 'yes', or '1' are considered True and,
    when `strict=False`, anything else returns the value specified by
    'default'.

    Useful for JSON-decoded stuff and config file parsing.

    If `strict=True`, unrecognized values, including None, will raise a
    ValueError which is useful when parsing values passed in from an API call.
    Strings yielding False are 'f', 'false', 'off', 'n', 'no', or '0'.
    """
    if isinstance(subject, bool):
        return subject
    if not isinstance(subject, six.string_types):
        subject = six.text_type(subject)

    lowered = subject.strip().lower()

    if lowered in TRUE_STRINGS:
        return True
    elif lowered in FALSE_STRINGS:
        return False
    elif strict:
        acceptable = ', '.join("'%s'" % s
                               for s in sorted(TRUE_STRINGS + FALSE_STRINGS))
        msg = _("Unrecognized value '%(val)s', acceptable values are:"
                " %(acceptable)s") % {
                    'val': subject,
                    'acceptable': acceptable
        }
        raise ValueError(msg)
    else:
        return default


def check_string_length(value, name=None, min_length=0, max_length=None):
    """Check the length of specified string.

    :param value: the value of the string
    :param name: the name of the string
    :param min_length: the min_length of the string
    :param max_length: the max_length of the string
    :raises TypeError, ValueError: For any invalid input.

    .. versionadded:: 3.7
    """
    if name is None:
        name = value

    if not isinstance(value, six.string_types):
        msg = _("%s is not a string or unicode") % name
        raise TypeError(msg)

    length = len(value)
    if length < min_length:
        msg = _("%(name)s has %(length)s characters, less than "
                "%(min_length)s.") % {
                    'name': name,
                    'length': length,
                    'min_length': min_length
        }
        raise ValueError(msg)

    if max_length and length > max_length:
        msg = _("%(name)s has %(length)s characters, more than "
                "%(max_length)s.") % {
                    'name': name,
                    'length': length,
                    'max_length': max_length
        }
        raise ValueError(msg)


def validate_integer(value, name, min_value=None, max_value=None):
    """Make sure that value is a valid integer, potentially within range.

    :param value: value of the integer
    :param name: name of the integer
    :param min_value: min_value of the integer
    :param max_value: max_value of the integer
    :returns: integer
    :raises: ValueError if value is an invalid integer

    .. versionadded:: 3.33
    """
    try:
        value = int(str(value))
    except (ValueError, UnicodeEncodeError):
        msg = _('%(value_name)s must be an integer') % {'value_name': name}
        raise ValueError(msg)

    if min_value is not None and value < min_value:
        msg = _('%(value_name)s must be >= %(min_value)d') % {
            'value_name': name,
            'min_value': min_value
        }
        raise ValueError(msg)

    if max_value is not None and value > max_value:
        msg = _('%(value_name)s must be <= %(max_value)d') % {
            'value_name': name,
            'max_value': max_value
        }
        raise ValueError(msg)

    return value


def get_re_items(content, pattern):
    pattern_object = re.compile(pattern, re.MULTILINE)
    items = []
    for m in pattern_object.finditer(content):
        item = m.groups()
        items.append(item[0])
    return items


_SANITIZE_KEYS = ['password', 'admin_pass', 'public_key', 'public_keys']

# Let's build a list of regex objects using the list of
# _SANITIZE_KEYS we already have. This way, we only have to add the new key
# to the list of _SANITIZE_KEYS and we can generate regular expressions
# for XML and JSON automatically.
_SANITIZE_PATTERNS_2 = {}
_SANITIZE_PATTERNS_1 = {}
_SANITIZE_PATTERNS_WILDCARD = {}
_SANITIZE_PATTERNS_ARR = {}

# Some regular expressions have only one parameter, some
# have two parameters. Use different lists of patterns here.
_FORMAT_PATTERNS_1 = [r'(%(key)s[0-9]*\s*[=]\s*)[^\s^\'^\"]+']
_FORMAT_PATTERNS_2 = [
    r'(%(key)s[0-9]*\s*[=]\s*[\"\'])[^\"\']*([\"\'])',
    r'(%(key)s[0-9]*\s*[=]\s*[\"])[^\"]*([\"])',
    r'(%(key)s[0-9]*\s*[=]\s*[\'])[^\']*([\'])',
    r'(%(key)s[0-9]*\s+[\"\'])[^\"\']*([\"\'])',
    r'([-]{2}%(key)s[0-9]*\s+)[^\'^\"^=^\s]+([\s]*)',
    r'(<%(key)s[0-9]*>)[^<]*(</%(key)s[0-9]*>)',
    r'([\"\']%(key)s[0-9]*[\"\']\s*:\s*[\"\'])[^\"\']*'
    r'([\"\'])', r'([\'"][^"\']*%(key)s[0-9]*[\'"]\s*:\s*u?[\'"])[^\"\']*'
    r'([\'"])', r'([\'"][^\'"]*%(key)s[0-9]*[\'"]\s*,\s*\'--?[A-z]+'
    r'\'\s*,\s*u?[\'"])[^\"\']*([\'"])',
    r'(%(key)s[0-9]*\s*--?[A-z]+\s*)\S+(\s*)'
]

# like this "public_keys": ["ssha 1", "ssha 2"]
_FORMAT_PATTERNS_ARR = [
    r'([\"\']\s*%(key)s[0-9]*\s*[\"\']\s*:\s*\[)([^\]]*)(\])']

# Keep a separate list of patterns by key so we only
# need to apply the substitutions for keys we find using a quick "in"
# test.
for key in _SANITIZE_KEYS:
    _SANITIZE_PATTERNS_1[key] = []
    _SANITIZE_PATTERNS_2[key] = []
    _SANITIZE_PATTERNS_WILDCARD[key] = []
    _SANITIZE_PATTERNS_ARR[key] = []

    for pattern in _FORMAT_PATTERNS_2:
        reg_ex = re.compile(pattern % {'key': key}, re.DOTALL | re.IGNORECASE)
        _SANITIZE_PATTERNS_2[key].append(reg_ex)

    for pattern in _FORMAT_PATTERNS_1:
        reg_ex = re.compile(pattern % {'key': key}, re.DOTALL | re.IGNORECASE)
        _SANITIZE_PATTERNS_1[key].append(reg_ex)

    for pattern in _FORMAT_PATTERNS_ARR:
        reg_ex = re.compile(pattern % {'key': key}, re.DOTALL | re.IGNORECASE)
        _SANITIZE_PATTERNS_ARR[key].append(reg_ex)


# Before submitting a patch to add a new argument to
# this function to allow the caller to pass in "extra" or "additional"
# or "replacement" patterns to be masked out, please note that we have
# discussed that feature many times and always rejected it based on
# the desire to have Oslo functions behave consistently across all
# projects and *especially* to have security features work the same
# way no matter where they are used. If every project adopted its own
# set patterns for secret values, it would be very difficult to audit
# the logging to ensure that everything is properly masked. So, please
# either add your pattern to the module-level variables at the top of
# this file or, even better, pick an existing pattern or key to use in
# your application to ensure that the value is masked by this
# function.
def mask_password(message, secret="***"):  # nosec
    """Replace password with *secret* in message.

    :param message: The string which includes security information.
    :param secret: value with which to replace passwords.
    :returns: The unicode value of message with the password fields masked.

    For example:

    >>> mask_password("'password' : 'aaaaa'")
    "'password' : '***'"
    """

    try:
        message = str(message)
    except UnicodeDecodeError:
        pass

    substitute1 = r'\g<1>' + secret
    substitute2 = r'\g<1>' + secret + r'\g<2>'
    substitute_wildcard = r'\g<1>'

    def substitute_arr(secret):

        def replace(match):
            pattern = r'(?P<quote>["\'])(.*?)(?P=quote)'
            reg_ex = re.compile(pattern, re.DOTALL | re.IGNORECASE)
            substitute = r'\g<quote>' + secret + r'\g<quote>'
            return match.group(1) + re.sub(reg_ex, substitute,
                                           match.group(2)) + match.group(3)

        return replace

    # Check to see if anything in message contains any key
    # specified in _SANITIZE_KEYS, if not then just return the message since
    # we don't have to mask any passwords.
    for key in _SANITIZE_KEYS:
        if key in message.lower():
            for pattern in _SANITIZE_PATTERNS_2[key]:
                message = re.sub(pattern, substitute2, message)
            for pattern in _SANITIZE_PATTERNS_1[key]:
                message = re.sub(pattern, substitute1, message)
            # Those case are poorly handled by previous
            # patterns. They are passwords with quotes or double quotes.
            # They also needs a different way to substitute group this is why
            # they aren't fix in the pattern 1 or 2.
            for pattern in _SANITIZE_PATTERNS_WILDCARD[key]:
                message = re.sub(pattern, substitute_wildcard, message)

            for pattern in _SANITIZE_PATTERNS_ARR[key]:
                message = re.sub(pattern, substitute_arr(secret), message)
    return message


def get_initials(text):
    """Generate a string of initials from the input text.

    This function takes a string that contains multiple words separated by spaces
    and returns a new string that consists of the uppercase first letter of each word.

    Parameters:
    - text (str): A string of text that may contain multiple words.

    Returns:
    - str: A string of uppercase initials of the words in the input text.

    Usage examples:
    - get_initials("Hard Disk Device") returns "HDD"
    """
    words = text.split()
    initials = [word[0].upper() for word in words]
    return ''.join(initials)
