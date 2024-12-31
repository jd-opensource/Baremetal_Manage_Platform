package rabbit_mq

import (
	"fmt"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	"github.com/streadway/amqp"
)

func SendToApi2Scheduler(obj interface{}) error {
	return doPublish(ironicChannel, IRONIC_EXCHANGE, API_TO_SCHEDULER_ROUTING_KEY, obj)
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
	fmt.Println(exchangeName, routingKey, msg)
	return
}
