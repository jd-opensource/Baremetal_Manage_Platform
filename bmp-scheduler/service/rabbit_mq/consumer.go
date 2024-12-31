package rabbit_mq

import (
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

//消费从openapi投递过来的msg
func ReceiveFromIronicApi2Scheduler() (<-chan amqp.Delivery, error) {

	if ironicChannel == nil {
		return nil, errors.New("ironicChannel is nil")
	}
	q, err := queueDeclare(ironicChannel, API_TO_SCHEDULER_QUEUE_NAME)
	if err != nil {
		return nil, err
	}
	if err := queueBind(ironicChannel, q.Name, API_TO_SCHEDULER_ROUTING_KEY, IRONIC_EXCHANGE); err != nil {
		return nil, err
	}
	msgChan, err := doConsume(ironicChannel, q.Name)
	if err != nil {
		return nil, err
	}
	return msgChan, nil
}

//消费从agent/dhcp-agent投递过来的回执msg
func ReceiveFromIronicAgentDriver() (<-chan amqp.Delivery, error) {

	if ironicChannel == nil {
		return nil, errors.New("ironicChannel is nil")
	}

	q, err := queueDeclare(ironicChannel, SCHEDULER_QUEUE_NAME)
	if err != nil {
		fmt.Println(time.Now(), "queueDeclare error:", err.Error())
		return nil, err
	}
	if err := queueBind(ironicChannel, q.Name, SCHEDULER_ROUTING_KEY, IRONIC_EXCHANGE); err != nil {
		fmt.Println(time.Now(), "queueBind error:", err.Error())
		return nil, err
	}
	msgChan, err := doConsume(ironicChannel, q.Name)
	if err != nil {
		fmt.Println(time.Now(), "doConsume error:", err.Error())
		return nil, err
	}
	return msgChan, nil
}

func queueDeclare(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func queueBind(ch *amqp.Channel, queueName, routing_key, exchange string) error {
	return ch.QueueBind(
		queueName,   // queue name
		routing_key, // routing key
		exchange,    // exchange
		false,
		nil,
	)
}

func doConsume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
}
