package rabbit_mq

import (
	"fmt"
	"time"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	"github.com/streadway/amqp"
)

//发送给agent/driver
func SendToAgent(routing_key string, message interface{}) error {
	queueDeclare(ironicChannel, routing_key)
	queueBind(ironicChannel, routing_key, routing_key, IRONIC_EXCHANGE)
	return doPublish(ironicChannel, IRONIC_EXCHANGE, routing_key, message)
}

//发送给dhcp-agent
func SendToAgent4Topic(routing_key string, message interface{}) error {
	queueDeclare(ironic4TopicChannel, routing_key)
	queueBind(ironic4TopicChannel, routing_key, routing_key, IRONIC_4TOPIC_EXCHANGE)
	return doPublish(ironic4TopicChannel, IRONIC_4TOPIC_EXCHANGE, routing_key, message)
}

func doPublish(c *amqp.Channel, exchangeName string, routingKey string, val interface{}) (err error) {
	msg, err := commonUtil.Convert2String(val)
	if err != nil {
		return err
	}
	err = c.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(msg),
		})

	fmt.Println(time.Now(), "scheduler publish to mq success, msg info:", "exchangeName:", exchangeName, "routingKey:", routingKey, "msg:", msg)
	return
}
