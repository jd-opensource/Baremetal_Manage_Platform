import logging
import threading

from bmpa.client.rabbitmq import Consumer
from bmpa.client.rabbitmq import Publisher
from bmpa.command import command_service
from bmpa.command.common import Command
from bmpa.command.common import Heart
from bmpa.protocol import Protocol
from bmpa.serialize import DefaultJsonEncoder
from bmpa.service import node
from bmpa.utils import validates

LOG = logging.getLogger(__name__)


class RabbitmqProtocol(Protocol):

    def __init__(self):
        self.consumer = Consumer()
        self.publisher = Publisher()
        self.encoder = DefaultJsonEncoder()

    def start(self):
        __consumer_thread = threading.Thread(name='ConsumerThread',
                                             target=self._start_consumer,
                                             daemon=True)

        __consumer_thread.start()
        self._start_publisher()

    def stop(self):
        self.consumer.stop()

    def send_command(self, command: Command):
        message = self.encoder.encode(command.serialize())
        LOG.info("Send command: %s", message)
        self.publisher.publish_message(message)

    def receive_command(self):
        LOG.info('Receive command.')
        message = self.consumer.get_message()
        if message is not None:
            command_service.receive_command(message)

    def _is_connect(self):
        return self.publisher.is_connect()

    def heartbeat(self):
        if self._is_connect():
            heart = Heart(sn=node.get_serial_number())
            self.send_command(heart)
        else:
            LOG.warning("RabbitMQ is not connect, not send heartbeat.")

    def _start_consumer(self):
        queue_name = node.get_serial_number()
        validates.is_required(queue_name, 'queu name')
        LOG.info("Comsumer qeue name %s.", queue_name)
        self.consumer.init(queue_name)
        self.consumer.start()

    def _start_publisher(self):
        self.publisher.init()
        self.publisher.start()
