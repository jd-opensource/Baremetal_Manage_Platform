import logging
from queue import Queue, Empty

from bmpa.command.common import COMMAND_MAP
from bmpa.command.common import SystemCommandError
from bmpa import engine
from bmpa.errors import UnkownCommandError
from bmpa import serialize
from bmpa.service import node
from bmpa.utils import classloader
from bmpa.utils import strutils


LOG = logging.getLogger(__name__)


class CommandService:

    def __init__(self) -> None:
        self._commands = Queue()
        self.decoder = serialize.DefaultJsonDecoder()

    def dispatch(self):
        LOG.info('Command dispatch start.')
        while True:
            command_str = None
            action = None
            request_id = None
            try:
                command_str = self._commands.get()
                if not command_str:
                    LOG.warn("Get a None message: %.", command_str)
                    continue
                command_str = command_str.decode('utf8')
                LOG.info('Command: %s', strutils.mask_password(command_str))
                command = self.decoder.decode(command_str)
                request_id = command.get('request_id', None)
                if type(command) is dict and "action" in command.keys():
                    action = command.get('action')
                    if action in COMMAND_MAP:
                        class_path = COMMAND_MAP.get(action)
                        command_clazz = classloader.load_class(class_path)
                        command_object = command_clazz(**command)

                        command_object.run()
                        engine.Engine.send_command(command_object)
                        continue

                LOG.error("Unknow command action : %s.",
                          strutils.mask_password(command_str))
                raise UnkownCommandError(
                    details=strutils.mask_password(command_str))
            except Empty:
                continue
            except Exception as e:
                LOG.error("Process command failed : %s.",
                          str(e),
                          exc_info=1)
                cmd = SystemCommandError(sn=node.get_serial_number(),
                                         action=action or 'SystemCommandError',
                                         code=getattr(e, 'code', 500),
                                         type=e.__class__.__name__,
                                         details=getattr(e, 'details', None),
                                         message=getattr(e, 'message', str(e)),
                                         request_id=request_id)
                engine.Engine.send_command(cmd)

    def receive_command(self, command):
        self._commands.put(command)
