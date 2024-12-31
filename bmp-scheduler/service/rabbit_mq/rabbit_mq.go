package rabbit_mq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	//********************ironic********************
	IRONIC_EXCHANGE       = "CPS_IRONIC_SCHEDULER"
	IRONIC_EXCHANGE_TYPE  = "direct"
	SCHEDULER_QUEUE_NAME  = "CPS_IRONIC_SCHEDULER"
	SCHEDULER_ROUTING_KEY = "CPS_IRONIC_SCHEDULER"

	IRONIC_4TOPIC_EXCHANGE      = "BMP_SCHEDULER_TOPIC"
	IRONIC_4TOPIC_EXCHANGE_TYPE = "topic"

	API_TO_SCHEDULER_QUEUE_NAME  = "CPS_IRONIC_API_TO_SCHEDULER"
	API_TO_SCHEDULER_ROUTING_KEY = "CPS_IRONIC_INNER_ROUTING_KEY"

	//********************openapi********************
	OPENAPI_EXCHANGE      = "CPS_OPENAPI_EXCHANGE"
	OPENAPI_EXCHANGE_TYPE = "direct"

	OPENAPI_CALLBACK_QUEUE_NAME  = "CPS_OPENAPI_CALLBACK_QUEUE"
	OPENAPI_CALLBACK_ROUTING_KEY = "CPS_OPENAPI_CALLBACK_ROUTING_KEY"
)

var (
	openapiUrl string
	ironicUrl  string

	ironicConn    *amqp.Connection
	ironicChannel *amqp.Channel

	ironic4TopicConn    *amqp.Connection
	ironic4TopicChannel *amqp.Channel

	openapiConn    *amqp.Connection
	openapiChannel *amqp.Channel
)

func DestroyTemplate() {
	ironicChannel.Close()
	ironicConn.Close()
	ironic4TopicChannel.Close()
	ironic4TopicConn.Close()
	openapiChannel.Close()
	openapiConn.Close()
}

//连接保活，断开时重连
func keepAliveConn() error {

	var ironicCloseChan chan *amqp.Error
	var ironic4TopicCloseChan chan *amqp.Error
	var openapiCloseChan chan *amqp.Error
	var ironicNotifyClose chan *amqp.Error
	var ironic4TopicNotifyClose chan *amqp.Error
	var openapiNotifyClose chan *amqp.Error

	for {

		if ironicConn == nil || ironicConn.IsClosed() { //连接已断开，需要重连
			fmt.Printf(time.Now().String() + "ironic conn is closed\n")
			//新建连接和信道
			if err := InitIronicMq(); err != nil {
				fmt.Printf(time.Now().String()+"keepAliveConn.InitIronicMq error:%s\n", err.Error())
				goto END
			} else {
				fmt.Printf(time.Now().String() + "keepAliveConn.InitIronicMq success\n")
			}
		}
		if ironic4TopicConn == nil || ironic4TopicConn.IsClosed() { //连接已断开，需要重连
			fmt.Printf(time.Now().String() + "ironic conn is closed\n")
			//新建连接和信道
			if err := InitIronic4TopicMq(); err != nil {
				fmt.Printf(time.Now().String()+"keepAliveConn.InitIronic4TopicMq error:%s\n", err.Error())
				goto END
			} else {
				fmt.Printf(time.Now().String() + "keepAliveConn.InitIronic4TopicMq success\n")
			}
		}
		if openapiConn == nil || openapiConn.IsClosed() { //连接已断开，需要重连
			fmt.Printf(time.Now().String() + "ironic conn is closed\n")
			//新建连接和信道
			if err := InitOpenapiMq(); err != nil {
				fmt.Printf(time.Now().String()+"keepAliveConn.InitOpenapiMq error:%s\n", err.Error())
				goto END
			} else {
				fmt.Printf(time.Now().String() + "keepAliveConn.InitOpenapiMq success\n")
			}
		}

		ironicCloseChan = make(chan *amqp.Error, 1)
		ironic4TopicCloseChan = make(chan *amqp.Error, 1)
		openapiCloseChan = make(chan *amqp.Error, 1)
		ironicNotifyClose = ironicChannel.NotifyClose(ironicCloseChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
		ironic4TopicNotifyClose = ironic4TopicChannel.NotifyClose(ironic4TopicCloseChan)
		openapiNotifyClose = openapiChannel.NotifyClose(openapiCloseChan)
		select {
		case <-ironicNotifyClose:
			fmt.Printf(time.Now().String() + "ironic channel error\n")
			if ironicConn == nil || ironicConn.IsClosed() { //连接已断开，需要重连
				fmt.Printf(time.Now().String() + "ironic conn is closed\n")
				time.Sleep(5 * time.Second)
				//新建连接和信道
				if err := InitIronicMq(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitIronicMq error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitIronicMq success\n")
				}

			} else { //连接还在，只是信道异常，需要获取新信道
				fmt.Printf(time.Now().String() + "ironic conn is open but channel is bad\n")
				time.Sleep(5 * time.Second)
				//获取新信道
				if err := InitIronicMqForNewChannel(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitIronicMqForNewChannel error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitIronicMqForNewChannel success\n")
				}
			}
		case <-ironic4TopicNotifyClose:
			fmt.Printf(time.Now().String() + "ironic4topic channel error\n")
			if ironic4TopicConn == nil || ironic4TopicConn.IsClosed() { //连接已断开，需要重连
				fmt.Printf(time.Now().String() + "ironic4topic conn is closed\n")
				time.Sleep(5 * time.Second)
				//新建连接和信道
				if err := InitIronic4TopicMq(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitIronic4topicMq error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitIronic4topicMq success\n")
				}

			} else { //连接还在，只是信道异常，需要获取新信道
				fmt.Printf(time.Now().String() + "ironic4topic conn is open but channel is bad\n")
				time.Sleep(5 * time.Second)
				//获取新信道
				if err := InitIronic4TopicMqForNewChannel(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitIronic4TopicMqForNewChannel error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitIronic4TopicMqForNewChannel success\n")
				}
			}
		case <-openapiNotifyClose:

			fmt.Printf(time.Now().String() + "openapi channel error\n")
			if openapiConn == nil || openapiConn.IsClosed() { //连接已断开，需要重连
				fmt.Printf(time.Now().String() + "openapi conn is closed\n")
				time.Sleep(5 * time.Second)
				//新建连接和信道
				if err := InitOpenapiMq(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitOpenapiMq error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitOpenapiMq success\n")
				}

			} else { //连接还在，只是信道异常，需要获取新信道
				fmt.Printf(time.Now().String() + "openapi conn is open but channel is bad\n")
				time.Sleep(5 * time.Second)
				//获取新信道
				if err := InitOpenapiMqForNewChannel(); err != nil {
					fmt.Printf(time.Now().String()+"keepAliveConn.InitOpenapiMqForNewChannel error:%s\n", err.Error())
				} else {
					fmt.Printf(time.Now().String() + "keepAliveConn.InitOpenapiMqForNewChannel success\n")
				}
			}
		}
	END:
		fmt.Printf(time.Now().String() + "keepAliveConn once done\n")
		time.Sleep(10 * time.Second)
	}

}

//exchange receive_routing_key 用于ironicmq
func InitMqTemplate(openapi string, ironic string) (err error) {

	openapiUrl = openapi
	ironicUrl = ironic

	if err := InitIronicMq(); err != nil {
		fmt.Printf(time.Now().String()+" InitIronicMq error:%s\n", err.Error())
		//return err
	}
	if err := InitOpenapiMq(); err != nil {
		fmt.Printf(time.Now().String()+" InitOpenapiMq error:%s\n", err.Error())
		//return err
	}
	if err := InitIronic4TopicMq(); err != nil {
		fmt.Printf(time.Now().String()+" InitIronic4TopicMq error:%s\n", err.Error())
		//return err
	}

	go keepAliveConn()

	//********************ironic-topic***********************

	return nil
}

func InitIronic4TopicMq() error {

	var err error

	//********************ironic-direct***********************
	ironic4TopicConn, err = amqp.Dial(ironicUrl)
	if err != nil {
		return err
	}
	ironic4TopicChannel, err = ironic4TopicConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(ironic4TopicChannel, IRONIC_4TOPIC_EXCHANGE, IRONIC_4TOPIC_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

//连接正常，信道异常时，获取新信道
func InitIronic4TopicMqForNewChannel() error {

	var err error
	ironic4TopicChannel, err = ironic4TopicConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(ironic4TopicChannel, IRONIC_4TOPIC_EXCHANGE, IRONIC_4TOPIC_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

func InitIronicMq() error {

	var err error

	//********************ironic-direct***********************
	ironicConn, err = amqp.Dial(ironicUrl)
	if err != nil {
		return err
	}
	ironicChannel, err = ironicConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(ironicChannel, IRONIC_EXCHANGE, IRONIC_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

//连接正常，信道异常时，获取新信道
func InitIronicMqForNewChannel() error {

	var err error
	ironicChannel, err = ironicConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(ironicChannel, IRONIC_EXCHANGE, IRONIC_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

func InitOpenapiMq() error {

	var err error
	openapiConn, err = amqp.Dial(openapiUrl)
	if err != nil {
		return err
	}
	openapiChannel, err = openapiConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(openapiChannel, OPENAPI_EXCHANGE, OPENAPI_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

//连接正常，信道异常时，获取新信道
func InitOpenapiMqForNewChannel() error {

	var err error
	openapiChannel, err = openapiConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(openapiChannel, OPENAPI_EXCHANGE, OPENAPI_EXCHANGE_TYPE)
	if err != nil {
		return err
	}
	return nil
}

func exchangeDeclare(ch *amqp.Channel, exchange, topicType string) error {
	return ch.ExchangeDeclare(
		exchange,  // name
		topicType, // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
}
