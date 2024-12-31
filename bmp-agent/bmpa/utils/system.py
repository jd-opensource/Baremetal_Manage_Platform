import errno
import logging
import os

from jinja2 import DictLoader
from jinja2 import Environment
from jinja2 import FileSystemLoader

from bmpa.errors import BmpError


LOG = logging.getLogger(__name__)


def decode_binary(blob, encoding='utf-8', errors='replace'):
    # Converts a binary type into a text type using given encoding.
    return blob.decode(encoding, errors=errors)


def load_file(path, read_len=None, offset=0, decode=True):
    with open(path, "rb") as fp:
        if offset:
            fp.seek(offset)
        contents = fp.read(read_len) if read_len else fp.read()

    if decode:
        return decode_binary(contents)
    else:
        return contents


def ensure_dir(path, mode=None):
    if path == "":
        path = "."
    try:
        os.makedirs(path)
    except OSError as e:
        if e.errno != errno.EEXIST:
            raise

    if mode is not None:
        os.chmod(path, mode)


def get_kernel_version_by_dir(target):
    LOG.info("Get kernel versioon by dir.")
    kernel_path_candidates = [
        os.path.join(target, "usr/src/kernels"),
        os.path.join(target, "usr/lib/modules")
    ]
    for path in kernel_path_candidates:
        if os.path.exists(path):
            dirs = os.listdir(path)
            if dirs:
                return dirs[0]
    raise BmpError("Get kernels path error.")


def get_relative_path(script):
    """Get the relative path of a script which ships with bmpa.

    :param script: The script name as a string.
    :returns: The relative path of the script.
    """
    cwd = os.path.dirname(os.path.realpath(__file__))
    return os.path.join(cwd, '..', script)


def render_template(template, params=None, is_file=True):
    """Renders Jinja2 template file with given parameters.

    :param template: full path to the Jinja2 template file
    :param params: dictionary with parameters to use when rende ring
    :param is_file: whether template is file or string with template itself
    :returns: the rendered template as a string
    """
    if is_file:
        tmpl_path, tmpl_name = os.path.split(template)
        loader = FileSystemLoader(tmpl_path)
    else:
        tmpl_name = 'template'
        loader = DictLoader({tmpl_name: template})
    env = Environment(loader=loader)
    tmpl = env.get_template(tmpl_name)

    if params is not None:
        return tmpl.render(params, enumerate=enumerate)
    else:
        return tmpl.render()
