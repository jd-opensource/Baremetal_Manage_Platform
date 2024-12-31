package rabbit_mq

import (
	"fmt"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	"github.com/streadway/amqp"
	"time"
)

func SendToIronicScheduler(message interface{}) error {
	fmt.Println(time.Now(), "处理结果回传给scheduler", message, "exchange", Exchange, "routing_key", Exchange)
	return doPublish(ironicChannel, Exchange, Exchange, message)
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

	return
}
