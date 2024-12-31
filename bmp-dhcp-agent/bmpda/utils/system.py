import logging
import os

from jinja2 import DictLoader
from jinja2 import Environment
from jinja2 import FileSystemLoader


LOG = logging.getLogger(__name__)


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
