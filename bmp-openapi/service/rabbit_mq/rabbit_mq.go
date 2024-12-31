package rabbit_mq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	//********************ironic********************
	IRONIC_EXCHANGE = "CPS_IRONIC_SCHEDULER"
	IRONIC_TOPIC    = "direct"

	API_TO_SCHEDULER_QUEUE_NAME  = "CPS_IRONIC_API_TO_SCHEDULER"
	API_TO_SCHEDULER_ROUTING_KEY = "CPS_IRONIC_INNER_ROUTING_KEY"
	SCHEDULER_QUEUE_NAME         = "CPS_IRONIC_SCHEDULER"
	SCHEDULER_ROUTING_KEY        = "CPS_IRONIC_SCHEDULER"

	//********************openapi********************
	OPENAPI_EXCHANGE = "CPS_OPENAPI_EXCHANGE"
	OPENAPI_TOPIC    = "direct"

	OPENAPI_CALLBACK_QUEUE_NAME  = "CPS_OPENAPI_CALLBACK_QUEUE"
	OPENAPI_CALLBACK_ROUTING_KEY = "CPS_OPENAPI_CALLBACK_ROUTING_KEY"

	IRONIC_4TOPIC_EXCHANGE = "BMP_SCHEDULER_TOPIC"
)

var (
	openapiUrl     string
	ironicUrl      string
	ironicConn     *amqp.Connection
	ironicChannel  *amqp.Channel
	openapiConn    *amqp.Connection
	openapiChannel *amqp.Channel
)

func DestroyTemplate() {
	ironicChannel.Close()
	ironicConn.Close()
	//openapiChannel.Close()
	//openapiConn.Close()
}

//连接保活，断开时重连
func keepAliveConn() {

	var ironicCloseChan chan *amqp.Error
	var ironicNotifyClose chan *amqp.Error

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

		ironicCloseChan = make(chan *amqp.Error, 1)
		//openapiCloseChan := make(chan *amqp.Error, 1)
		ironicNotifyClose = ironicChannel.NotifyClose(ironicCloseChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
		//openapiNotifyClose := openapiChannel.NotifyClose(openapiCloseChan)
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
			/*case <-openapiNotifyClose:

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
			}*/
		}
	END:
		fmt.Printf(time.Now().String() + "keepAliveConn once done\n")
		time.Sleep(10 * time.Second)

	}

}

//exchange receive_routing_key 用于ironicmq
func InitMqTemplate(ironic string) (err error) {
	ironicUrl = ironic
	if err := InitIronicMq(); err != nil {
		fmt.Printf(time.Now().String()+" InitMqTemplate error:%s\n", err.Error())
		// return err
	}
	go keepAliveConn()
	//*******************************************
	return nil
}

func InitIronicMq() error {

	var err error
	ironicConn, err = amqp.Dial(ironicUrl)
	if err != nil {
		return err
	}
	ironicChannel, err = ironicConn.Channel()
	if err != nil {
		return err
	}
	err = exchangeDeclare(ironicChannel, IRONIC_EXCHANGE, IRONIC_TOPIC)
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
	err = exchangeDeclare(ironicChannel, IRONIC_EXCHANGE, IRONIC_TOPIC)
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
	err = exchangeDeclare(openapiChannel, OPENAPI_EXCHANGE, OPENAPI_TOPIC)
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
	err = exchangeDeclare(openapiChannel, OPENAPI_EXCHANGE, OPENAPI_TOPIC)
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
