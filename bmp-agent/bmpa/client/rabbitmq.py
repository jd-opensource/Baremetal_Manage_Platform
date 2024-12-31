import logging
from queue import Empty
from queue import Queue
import time

import pika

from bmpa import config
from bmpa.utils import strutils

LOG = logging.getLogger(__name__)


class Consumer(object):

    def __init__(self):

        self.queue = Queue()
        # type: pika.connection.Connection
        self.connection = None
        # type: pika.channel.Channel
        self.channel = None
        self.consumer_tag = None

        self.username = None
        self.password = None
        self.host = None
        self.port = None
        self.vhost = None
        self.exchange = None
        self.exchange_type = None
        self.queue_name = None
        self.routing_key = None

    def init(self, queue_name: str):
        self.queue_name = queue_name
        self.routing_key = self.queue_name
        self.username = config.mq_username
        self.password = config.mq_password
        self.host = config.mq_host
        self.port = config.mq_port
        self.vhost = config.mq_vhost
        self.exchange = config.mq_exchange
        self.exchange_type = config.mq_exchange_type

    def start(self):
        wait = base = 2
        while True:
            try:
                self._build_connection()
                LOG.info('Begin to comsume message.')
                self.consumer_tag = self.channel.basic_consume(
                    queue=self.queue_name,
                    on_message_callback=self._on_message)
                self.channel.start_consuming()
                wait = base
            except Exception as e:
                wait = min(1024, wait * 2)
                time.sleep(wait)
                LOG.error("Comsume message failed. %s", e, exc_info=1)
            finally:
                try:
                    self._stop()
                except Exception as e:
                    LOG.error(
                        "Close consumer connection failed. "
                        "retry to reconnect, %s",
                        e,
                        exc_info=1)

    def get_message(self):
        try:
            message = self.queue.get()
            # LOG.info("Get message :%s.", message)
            LOG.info("Get message :%s.", strutils.mask_password(message.decode('utf-8')))
            return message
        except Exception as e:
            LOG.warn("Get message error:%s.", str(e))
        return None

    def stop(self):
        self._stop()

    def _connection_parameters(self):
        credentials = pika.PlainCredentials(self.username, self.password)
        return pika.ConnectionParameters(host=self.host,
                                         port=self.port,
                                         virtual_host=self.vhost,
                                         credentials=credentials,
                                         connection_attempts=1,
                                         retry_delay=5,
                                         socket_timeout=5,
                                         heartbeat=5)

    def _build_connection(self):
        """This method connects to RabbitMQ, returning the connection handle.

        When the connection is established, the on_connection_open method
        will be invoked by pika.

        :rtype: pika.SelectConnection
        """
        routing_key = self.queue_name
        parameters = self._connection_parameters()
        LOG.info('Connecting to %s', parameters)
        self._connection = pika.BlockingConnection(parameters=parameters)
        self.channel = self._connection.channel()
        self._setup_exchange(self.channel, self.exchange, self.exchange_type)
        self._setup_queue(self.channel, self.queue_name)
        self._bind_queue_to_exchange(self.channel, self.exchange,
                                     self.queue_name, routing_key)

    def _setup_exchange(self, channel, exchange_name, exchange_type):
        """Setup the exchange on RabbitMQ by invoking the Exchange.Declare RPC

        command. When it is complete, the on_exchange_declareok method will
        be invoked by pika.

        :param str|unicode exchange_name: The name of the exchange to declare
        """
        LOG.info('Declaring exchange : %s', exchange_name)
        # Note: using functools.partial is not required, it is demonstrating
        # how arbitrary data can be passed to the callback when it is called
        channel.exchange_declare(exchange=exchange_name,
                                 exchange_type=exchange_type,
                                 durable=True)

    def _setup_queue(self, channel, queue_name, auto_delete=False):
        """Setup the queue on RabbitMQ by invoking the Queue.Declare RPC

        command. When it is complete, the on_queue_declareok method will
        be invoked by pika.

        :param str|unicode queue_name: The name of the queue to declare.
        """
        LOG.info('Declaring queue %s', queue_name)
        channel.queue_declare(queue=queue_name,
                              durable=True,
                              auto_delete=auto_delete)

    def _bind_queue_to_exchange(self, channel, exchange, queue_name,
                                routing_key):
        """Method invoked by pika when the Queue.Declare RPC call made in

        setup_queue has completed. In this method we will bind the queue
        and exchange together with the routing key by issuing the Queue.Bind
        RPC command. When this command is complete, the on_bindok method will
        be invoked by pika.

        :param pika.frame.Method method_frame: The Queue.DeclareOk frame
        :param str|unicode queue_name: Extra user data (queue name)
        """
        LOG.info('Binding %s to %s with %s', exchange, queue_name,
                 routing_key)
        channel.queue_bind(queue=queue_name,
                           exchange=exchange,
                           routing_key=routing_key)

    def start_consuming(self):
        """This method sets up the consumer by first calling

        add_on_cancel_callback so that the object is notified if RabbitMQ
        cancels the consumer. It then issues the Basic.Consume RPC command
        which returns the consumer tag that is used to uniquely identify the
        consumer with RabbitMQ. We keep the value to use it when we want to
        cancel consuming. The on_message method is passed in as a callback pika
        will invoke when a message is fully received.
        """
        try:
            LOG.info('Issuing consumer related RPC commands')
            self.consumer_tag = self.channel.basic_consume(
                queue=self.receive_queue_name,
                on_message_callback=self.on_message)
            self.channel.start_consuming()
        except Exception as e:
            try:
                self._stop()
                self._build_connection()
            finally:
                message = ("Connect failed. : %s") % str(e)
                LOG.error("Connect failed. %s", message, exc_info=1)

    def _on_message(self, unused_channel, basic_deliver, properties, body):
        """Invoked by pika when a message is delivered from RabbitMQ. The

        channel is passed for your convenience. The basic_deliver object that
        is passed in carries the exchange, routing key, delivery tag and
        a redelivered flag for the message. The properties passed in is an
        instance of BasicProperties with the message properties and the body
        is the message that was sent.

        :param pika.channel.Channel unused_channel: The channel object
        :param pika.Spec.Basic.Deliver: basic_deliver method
        :param pika.Spec.BasicProperties: properties
        :param str|unicode body: The message body
        """
        LOG.info('Received message # %s from %s: %s',
                 basic_deliver.delivery_tag, properties.app_id, strutils.mask_password(body.decode('utf-8')))
        self.queue.put(body)
        self._acknowledge_message(basic_deliver.delivery_tag)

    def _acknowledge_message(self, delivery_tag):
        """Acknowledge the message delivery from RabbitMQ by sending a

        Basic.Ack RPC method for the delivery tag.

        :param int delivery_tag: The delivery tag from the Basic.Deliver frame
        """
        LOG.info('Acknowledging message %s', delivery_tag)
        self.channel.basic_ack(delivery_tag)

    def _stop(self):
        """Cleanly shutdown the connection to RabbitMQ by stopping the consumer

        with RabbitMQ. When RabbitMQ confirms the cancellation, on_cancelok
        will be invoked by pika, which will then closing the channel and
        connection. The IOLoop is started again because this method is invoked
        when CTRL-C is pressed raising a KeyboardInterrupt exception. This
        exception stops the IOLoop which needs to be running for pika to
        communicate with RabbitMQ. All of the commands issued prior to starting
        the IOLoop will be buffered but not processed.
        """

        LOG.info('Stopping')
        self._stop_consuming()
        self._close_channel(self.channel)
        self._close_connection(self.connection)
        LOG.info('Stopped')

    def _stop_consuming(self):
        """Tell RabbitMQ that you would like to stop consuming by sending the

        Basic.Cancel RPC command.
        """
        if self.channel and self.channel.is_open:
            LOG.info('Sending a Basic.Cancel RPC command to RabbitMQ')
            self.channel.basic_cancel(consumer_tag=self.consumer_tag)
        LOG.info("Stop consuming")

    def _close_channel(self, channel):
        """Call to close the channel with RabbitMQ cleanly by issuing the

        Channel.Close RPC command.
        """
        LOG.info('Closing the channel')
        if channel and channel.is_open:
            channel.close()

    def _close_connection(self, connection):
        """This method closes the connection to RabbitMQ."""
        LOG.info('Closing connection')
        if connection and connection.is_open:
            connection.close()

    def __enter__(self):
        return self

    def __exit__(self, exception_type, exception_value, traceback):
        if exception_type is not None:
            LOG.error("%s, %s, %s", exception_type, exception_value,
                      traceback)


class Publisher(object):

    def __init__(self):

        self.queue = Queue()
        self.queue_name = None
        self.routing_key = None
        # type: pika.connection.Connection
        self.connection = None
        # type: pika.channel.Channel
        self.channel = None
        self.consumer_tag = None

        self.username = None
        self.password = None
        self.host = None
        self.port = None
        self.vhost = None
        self.exchange = None
        self.exchange_type = None

    def init(self):
        self.queue_name = config.mq_exchange
        self.routing_key = config.mq_exchange

        self.username = config.mq_username
        self.password = config.mq_password
        self.host = config.mq_host
        self.port = config.mq_port
        self.vhost = config.mq_vhost
        self.exchange = config.mq_exchange
        self.exchange_type = config.mq_exchange_type

    def publish_message(self, message):
        self.queue.put(message)

    def start(self):
        self._build_connection()
        while True:
            try:
                message = self.queue.get(timeout=0.2)
                if not message:
                    LOG.warn("Get a None message from publish queue.")
                    continue
                try:
                    # 未收到ack确认信息时会抛出异常
                    self.channel.basic_publish(
                        exchange=self.exchange,
                        routing_key=self.routing_key,
                        body=message,
                        properties=pika.BasicProperties(
                            content_type='application/json',
                            app_id=self.queue_name,
                            delivery_mode=2))

                    LOG.debug("Deliver message to %s %s %s.", self.exchange,
                              self.routing_key, strutils.mask_password(message))
                except Exception as e:
                    LOG.error(
                        "Publish message:%s failed, put message into queue. %s",
                        message,
                        e,
                        exc_info=1)
                    self.queue.put(message)
                    self._close_channel(self.channel)
                    self._close_connection(self.connection)
                    time.sleep(1)
                    self._build_connection()
                    time.sleep(1)
            except Empty:
                try:
                    if not self.is_connect():
                        self._build_connection()
                        LOG.info('Connection is connect.')
                        continue
                except Exception as e:
                    time.sleep(1)
                    LOG.info('Error:%s', e)
                    continue
                LOG.debug("Publish queue is empty.")
                time.sleep(1)
                continue
            except Exception as e:
                LOG.error("Publish message failed. %s", e, exc_info=1)
                time.sleep(1)
                continue

    def stop(self):
        self._stop()

    def is_connect(self):
        if self.channel and self.channel.is_open:
            LOG.debug("RabbitMQ connection is open.")
            return True
        else:
            LOG.debug("RabbitMQ connection is closed.")
            return False

    def _connection_parameters(self):
        credentials = pika.PlainCredentials(self.username, self.password)
        return pika.ConnectionParameters(host=self.host,
                                         port=self.port,
                                         virtual_host=self.vhost,
                                         credentials=credentials,
                                         connection_attempts=1,
                                         retry_delay=5,
                                         socket_timeout=5,
                                         heartbeat=5)

    def _build_connection(self):
        """This method connects to RabbitMQ, returning the connection handle.

        When the connection is established, the on_connection_open method
        will be invoked by pika.

        :rtype: pika.SelectConnection
        """
        parameters = self._connection_parameters()
        LOG.info('Connecting to %s', parameters)
        self._connection = pika.BlockingConnection(parameters=parameters)
        self.channel = self._connection.channel()
        self.channel.confirm_delivery()
        self._setup_exchange(self.channel, self.exchange, self.exchange_type)
        self._setup_queue(self.channel, self.queue_name)
        self._bind_queue_to_exchange(self.channel, self.exchange,
                                     self.queue_name, self.routing_key)

    def _setup_exchange(self, channel, exchange_name, exchange_type):
        """Setup the exchange on RabbitMQ by invoking the Exchange.Declare RPC

        command. When it is complete, the on_exchange_declareok method will
        be invoked by pika.

        :param str|unicode exchange_name: The name of the exchange to declare
        """
        LOG.info('Declaring exchange : %s', exchange_name)
        # Note: using functools.partial is not required, it is demonstrating
        # how arbitrary data can be passed to the callback when it is called
        channel.exchange_declare(exchange=exchange_name,
                                 exchange_type=exchange_type,
                                 durable=True)

    def _setup_queue(self, channel, queue_name, auto_delete=False):
        """Setup the queue on RabbitMQ by invoking the Queue.Declare RPC

        command. When it is complete, the on_queue_declareok method will
        be invoked by pika.

        :param str|unicode queue_name: The name of the queue to declare.
        """
        LOG.info('Declaring queue %s', queue_name)
        channel.queue_declare(queue=queue_name,
                              durable=True,
                              auto_delete=auto_delete)

    def _bind_queue_to_exchange(self, channel, exchange, queue_name,
                                routing_key):
        """Method invoked by pika when the Queue.Declare RPC call made in

        setup_queue has completed. In this method we will bind the queue
        and exchange together with the routing key by issuing the Queue.Bind
        RPC command. When this command is complete, the on_bindok method will
        be invoked by pika.

        :param pika.frame.Method method_frame: The Queue.DeclareOk frame
        :param str|unicode queue_name: Extra user data (queue name)
        """
        LOG.info('Binding %s to %s with %s', exchange, queue_name,
                 routing_key)
        channel.queue_bind(queue=queue_name,
                           exchange=exchange,
                           routing_key=routing_key)

    def _stop(self):
        """Cleanly shutdown the connection to RabbitMQ by stopping the consumer

        with RabbitMQ. When RabbitMQ confirms the cancellation, on_cancelok
        will be invoked by pika, which will then closing the channel and
        connection. The IOLoop is started again because this method is invoked
        when CTRL-C is pressed raising a KeyboardInterrupt exception. This
        exception stops the IOLoop which needs to be running for pika to
        communicate with RabbitMQ. All of the commands issued prior to starting
        the IOLoop will be buffered but not processed.
        """

        LOG.info('Stopping')
        self._close_channel(self.channel)
        self._close_connection(self.connection)
        LOG.info('Stopped')

    def _close_channel(self, channel):
        """Call to close the channel with RabbitMQ cleanly by issuing the

        Channel.Close RPC command.
        """
        LOG.info('Closing the channel')
        if channel and channel.is_open:
            channel.close()

    def _close_connection(self, connection):
        """This method closes the connection to RabbitMQ."""
        LOG.info('Closing connection')
        if connection and connection.is_open:
            connection.close()

    def __enter__(self):
        return self

    def __exit__(self, exception_type, exception_value, traceback):
        if exception_type is not None:
            LOG.error("%s, %s, %s", exception_type, exception_value,
                      traceback)
