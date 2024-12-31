package rabbit_mq

import (
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

//ironic scheduler consumer
func ReceiveFromIronicScheduler() (<-chan amqp.Delivery, error) {

	if ironicChannel == nil {
		return nil, errors.New("ironicChannel nil")
	}
	q, err := queueDeclare(ironicChannel, QueueName)
	if err != nil {
		fmt.Println(time.Now(), "queueDeclare error:", err.Error())
		return nil, err
	}
	if err := queueBind(ironicChannel, q.Name, RoutingKey, Exchange); err != nil {
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
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func queueBind(ch *amqp.Channel, queueName, binding_key, exchange string) error {
	return ch.QueueBind(
		queueName,   // queue name
		binding_key, // binding key
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
