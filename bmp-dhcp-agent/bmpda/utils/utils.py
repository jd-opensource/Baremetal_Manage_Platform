import os


def _read_params_from_file(filepath: str):
    """Extract key=value pairs from a file.

    :param filepath: path to a file containing key=value pairs separated by
                     whitespace or newlines.
    :returns: a dictionary representing the content of the file
    """
    with open(filepath) as f:
        cmdline = f.read()

    options = cmdline.split()
    params = {}
    for option in options:
        if '=' not in option:
            continue
        k, v = option.split('=', 1)
        params[k] = v

    return params


def get_agent_params(cmdline_path: str = '/proc/cmdline'):
    """Gets parameters passed to the agent via kernel cmdline."""
    return _read_params_from_file(cmdline_path)


def get_relative_path(script):
    """Get the relative path of a script which ships with bmpda

    :param script: The script name as a string.
    :returns: The relative path of the script.
    """
    cwd = os.path.dirname(os.path.realpath(__file__))
    return os.path.join(cwd, '..', script)
