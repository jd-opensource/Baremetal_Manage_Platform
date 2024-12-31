
import logging
import threading

from bmpda.client.rabbitmq import Consumer
from bmpda.client.rabbitmq import Publisher
from bmpda.command import command_service
from bmpda.command.common import Command
from bmpda.command.common import Heart
from bmpda.protocol import Protocol
from bmpda.serialize import DefaultJsonEncoder
from bmpda.service import node

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
            heart = Heart(sn=node.get_node_id())
            self.send_command(heart)
        else:
            LOG.warning("RabbitMQ is not connect, not send heartbeat.")

    def _start_consumer(self):
        LOG.info("Comsumer qeue name %s.", node.get_node_id())
        self.consumer.init(node.get_node_id())
        self.consumer.start()

    def _start_publisher(self):
        self.publisher.init(routing_key=node.get_node_id())
        self.publisher.start()
