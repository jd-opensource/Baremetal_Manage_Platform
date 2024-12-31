import logging
import os
import signal
import sys
import threading
import time
from typing import Optional

from filelock import FileLock

from bmpda.command import command_service
from bmpda.command.common import Command
from bmpda import config
from bmpda import constants
from bmpda.service import node
from bmpda.utils import dhcp
from bmpda.utils import net


LOG = logging.getLogger(__name__)


class Agent:

    def __init__(self):
        super(Agent, self).__init__()
        self.__init_protocol()
        self.__init_dhcp()
        self._stopped: Optional[threading.Event] = None

    def start(self):
        self.__init_node()
        self.__init_thread()
        self.__bind_signal()

    def stop(self):
        LOG.info('stopped...')
        self._stopped.set()
        self.__unbind_signal()
        LOG.info('stopped.')

    def __init_node(self):
        node_id = "%(prefix)s-%(host_ip)s" % {
            "prefix": config.mq_queue_name_pefix_consumer,
            "host_ip": net.get_host_ip()
        }
        node.set_node_id(node_id)

    def __init_thread(self):
        self._stopped = threading.Event()
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

        __dhcp_ctl_thread = threading.Thread(name='DHCPCtlThread',
                                                  target=self.__dhcp_ctl,
                                                  daemon=True)
        __protocol_start_thread.start()
        __heartbeat_thread.start()
        __command_dispatch_thread.start()
        __command_query_thread.start()
        __dhcp_ctl_thread.start()

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
        LOG.info("Receive a kill signal %s=%s.", signum, frame)
        self.stop()
        sys.exit()

    def __heartbeat(self):
        LOG.info("Start heartbeat.")
        wait = base = config.heart_period
        while not self._stopped.is_set():
            try:
                __protocol.heartbeat()
                wait = base  # reset to base wait time on success
                LOG.info("Trigger a heart period %d.", wait)
            except Exception as exc:
                LOG.error(str(exc))
                wait = min(20, wait * 2 or 1)

            time.sleep(wait)

    def __command_dispatch(self):
        LOG.info("Start command dispatch.")
        command_service.dispatch()

    def __command_query(self):
        LOG.info("Start command query.")
        while not self._stopped.is_set():
            __protocol.receive_command()

    def __protocol_start(self):
        __protocol.start()

    def __init_protocol(self):
        global __protocol
        if config.protocol == constants.PROTOCAL_RABBITMQ:
            from bmpda.protocol.rabbitmq import RabbitmqProtocol
            __protocol = RabbitmqProtocol()
        LOG.info("Init protocol %s.", config.protocol)

    def send_command(command: Command):
        __protocol.send_command(command)

    def __init_dhcp(self):
        dhcp_config_path = os.path.join(config.dhcp_config_dir,
                                        "dhcpd.conf")
        dhcp_config_lock_path = "%s.lock" % dhcp_config_path
        node.set_dhcp_config_path(dhcp_config_path)
        node.set_dhcp_config_lock_path(dhcp_config_lock_path)
        os.chmod(config.dhcp_control_bin, 0o755)
        dhcp.dhcp_ctl('start')
        LOG.info("Init dhcp")

    def __dhcp_ctl(self):
        dhcp_period = 6
        while not self._stopped.is_set():
            try:
                # Only set need_restart_dhcp=False here,
                # so no need to lock when obtaining the value.
                if node.get_need_restart_dhcp():
                    success = dhcp.dhcp_ctl('restart')
                    if success:
                        lock_dhcp = FileLock(
                            node.get_dhcp_config_lock_path(), timeout=3)
                        with lock_dhcp:
                            node.set_need_restart_dhcp(False)
                    else:
                        LOG.error(
                            "DHCP restart failed. retry after %d seconds.",
                            dhcp_period,
                            exc_info=1)
                LOG.info("Trigger a dhcp heart period %d.", dhcp_period)
                time.sleep(dhcp_period)
            except Exception as e:
                LOG.error("DHCP restart failed:%s.", e)
                time.sleep(dhcp_period)
