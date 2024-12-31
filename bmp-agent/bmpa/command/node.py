import logging

from bmpa.command.common import Command
from bmpa.service import initnode
import bmpa.utils.validates as validates

LOG = logging.getLogger(__name__)


class InitNode(Command):

    def __init__(self, node_data, **kwargs):
        super(InitNode, self).__init__(**kwargs)
        self.node_data = node_data
        validates.is_required(node_data, 'node_data')
        validates.check_field(node_data, 'os_type')
        validates.check_field(node_data, 'os_version')
        self.res_node = None

    def execute(self):
        self.res_node = initnode.init_node(self.node_data)

    def execute_after(self):
        if isinstance(self.res_node, dict):
            for key, value in self.res_node.items():
                self.data[key] = value
