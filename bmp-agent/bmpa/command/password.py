import logging

from bmpa.command.common import Command
import bmpa.service.password as service_password
from bmpa.utils import validates


LOG = logging.getLogger(__name__)


class SetPassword(Command):

    def __init__(self, password, username='root', **kwargs):
        super(SetPassword, self).__init__(**kwargs)
        self.password = password
        self.username = username

    def validate(self):
        validates.is_required(self.password, 'password')

    def execute(self):

        service_password.set_password(self.password, self.username)
