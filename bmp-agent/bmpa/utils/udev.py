import logging
import shlex

from bmpa.utils import executor

LOG = logging.getLogger(__name__)

try:
    shlex_quote = shlex.quote
except AttributeError:
    # python2.7 uses pipes.quote
    import pipes
    shlex_quote = pipes.quote


def udevadm_info(path=None):
    """Return a dictionary populated by properties of the device specified

        in the `path` variable via querying udev 'property' database.
    :params: path: path to device, either /dev or /sys
    :returns: dictionary of key=value pairs as exported from the udev database
    :raises: ValueError path is None, ProcessExecutionError on exec error.
    """
    if not path:
        raise ValueError('Invalid path: "%s"' % path)

    info_cmd = ['udevadm', 'info', '--query=property', '--export', path]
    output, _ = executor.execute(*info_cmd)

    # strip for trailing empty line
    info = {}
    for line in output.splitlines():
        if not line:
            continue
        # maxsplit=1 gives us key and remaininng part of line is value
        # py2.7 on Trusty doesn't have keyword, pass as argument
        key, value = line.split('=', 1)
        if not value:
            value = None
        if value:
            # preserve spaces in values to match udev database
            try:
                parsed = shlex.split(value)
            except ValueError:
                # strip the leading/ending single tick from udev output before
                # escaping the value to prevent their inclusion in the result.
                trimmed_value = value[1:-1]
                try:
                    quoted = shlex_quote(trimmed_value)
                    LOG.debug(
                        'udevadm_info: quoting shell-escape chars '
                        'in %s=%s -> %s', key, value, quoted)
                    parsed = shlex.split(quoted)
                except ValueError:
                    escaped_value = (trimmed_value.replace("'", "_").replace(
                        '"', "_"))
                    LOG.debug(
                        'udevadm_info: replacing shell-escape chars '
                        'in %s=%s -> %s', key, value, escaped_value)
                    parsed = shlex.split(escaped_value)
            if ' ' not in value:
                info[key] = parsed[0]
            else:
                # special case some known entries with spaces, e.g. ID_SERIAL
                # and DEVLINKS, see tests/unittests/test_udev.py
                if key == "DEVLINKS":
                    info[key] = shlex.split(parsed[0])
                elif key == 'ID_SERIAL':
                    info[key] = parsed[0]
                else:
                    info[key] = parsed

    return info
