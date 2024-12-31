import logging
import signal
import sys
import threading
import time

from bmpa.command import command_service
from bmpa.command.common import Command
from bmpa import config
from bmpa import constants
import bmpa.lifecycle
from bmpa.service import node
from bmpa.utils import hardware

LOG = logging.getLogger(__name__)


class Engine(bmpa.lifecycle.LifeCycle):

    def __init__(self):
        super(Engine, self).__init__()
        self.__init_protocol()

    def do_start(self):
        self.__init_node()
        self.__init_thread()
        self.__bind_signal()

    def do_stop(self):
        self.__unbind_signal()
        LOG.info('stop.....')

    def __init_node(self):
        vendor = hardware.get_system_vendor_info()
        node.set_system_vendor_info(vendor)
        node.set_serial_number(vendor.serial_number)

    def __init_thread(self):
        __protocol_start_thread = threading.Thread(
            name='ProtocalStartThread',
            target=self.__protocol_start,
            daemon=True)
        __heartbeat_thread = threading.Thread(name='HeartbeatThread',
                                              target=self.__heartbeat,
                                              daemon=True)
        __command_dispatch_thread = threading.Thread(
            name='CommandDispatchThread',
            target=self.__command_dispatch,
            daemon=True)

        __command_query_thread = threading.Thread(name='CommandQueryThread',
                                                  target=self.__command_query,
                                                  daemon=True)
        __protocol_start_thread.start()
        __heartbeat_thread.start()
        __command_dispatch_thread.start()
        __command_query_thread.start()

    def __bind_signal(self):
        # bind signal
        self.hold_sigint = signal.signal(signal.SIGINT, self.__handler_signal)
        self.hold_sigterm = signal.signal(signal.SIGTERM,
                                          self.__handler_signal)

    def __unbind_signal(self):
        # unbind signal
        signal.signal(signal.SIGINT, self.hold_sigint)
        signal.signal(signal.SIGTERM, self.hold_sigterm)

    def __handler_signal(self, signum, frame):
        LOG.info("receive a kill signal %s=%s.", signum, frame)
        self.stop()
        sys.exit()

    def __heartbeat(self):
        wait = base = config.heart_period
        while not self.should_stop():
            try:
                __protocol.heartbeat()
                wait = base  # reset to base wait time on success
                LOG.info("Trigger a heart period %d.", wait)
            except Exception as exc:
                LOG.error(str(exc))
                wait = min(20, wait * 2 or 1)

            time.sleep(wait)

    def __command_dispatch(self):
        while not self.should_stop():
            command_service.dispatch()

    def __command_query(self):
        while not self.should_stop():
            __protocol.receive_command()

    def __protocol_start(self):
        # while not self.should_stop():
        __protocol.start()

    def __init_protocol(self):
        global __protocol
        if config.protocol == constants.PROTOCAL_RABBITMQ:
            from bmpa.protocol.rabbitmq import RabbitmqProtocol
            __protocol = RabbitmqProtocol()

    def send_command(command: Command):
        __protocol.send_command(command)
