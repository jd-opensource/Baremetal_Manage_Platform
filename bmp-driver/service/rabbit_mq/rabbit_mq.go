package rabbit_mq

import (
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/streadway/amqp"
)

const (
	TOPIC     = "direct"
	QueueName = "IRONIC_DRIVER"
)

var C Conf

var (
	Exchange   string
	RoutingKey string

	ironicUrl     string
	ironicConn    *amqp.Connection
	ironicChannel *amqp.Channel
)

type Conf struct {
	Heart struct {
		Period int `yaml:"period"`
	} `yaml:"heart"`
	Mq struct {
		Hosts []struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"hosts"`
		Connection struct {
			Username          string `yaml:"username"`
			Password          string `yaml:"password"`
			Vhost             string `yaml:"vhost"`
			Exchange          string `yaml:"exchange"`
			ReceiveRoutingKey string `yaml:"receive_routing_key"`
		} `yaml:"connection"`
	} `yaml:"mq"`
	Log struct {
		LogPath string `yaml:"logPath"`
	} `yaml:"log"`
}

func DestroyTemplate() {
	ironicChannel.Close()
	ironicConn.Close()
}

//exchange receive_routing_key 用于ironicmq
func InitMqTemplate() (err error) {

	//ironicUrls := []string{}
	//hosts := C.Mq.Hosts
	//for _, v := range hosts {
	//	url := "amqp://" + C.Mq.Connection.Username + ":" + C.Mq.Connection.Password + "@" + v.Host + ":" + strconv.Itoa(v.Port) + "/" + C.Mq.Connection.Vhost
	//	ironicUrls = append(ironicUrls, url)
	//}
	//
	//Exchange = C.Mq.Connection.Exchange
	//RoutingKey = C.Mq.Connection.ReceiveRoutingKey
	//for _, v := range ironicUrls { //多地址挑一个能dial通的
	host, _ := beego.AppConfig.String("bmp_mq_host")
	port, _ := beego.AppConfig.String("bmp_mq_port")
	user, _ := beego.AppConfig.String("bmp_mq_user")
	password, _ := beego.AppConfig.String("bmp_mq_password")
	vhost, _ := beego.AppConfig.String("bmp_mq_vhost")
	Exchange, _ = beego.AppConfig.String("bmp_mq_exchange")
	RoutingKey, _ = beego.AppConfig.String("bmp_mq_receive_routing_key")
	ironicUrl = "amqp://" + user + ":" + password + "@" + host + ":" + port + "/" + vhost
	fmt.Println("ironicUrl mq url:", ironicUrl, Exchange, RoutingKey)
	if err := InitIronicMq(); err != nil {
		fmt.Println("InitMqTemplate.InitIronicMq error:%s", err.Error())
	}

	go keepAliveConn()
	return nil
}

//初始化connection和channel
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
	err = exchangeDeclare(ironicChannel, Exchange, TOPIC)
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
	err = exchangeDeclare(ironicChannel, Exchange, TOPIC)
	if err != nil {
		return err
	}
	return nil
}

//连接保活，断开时重连
func keepAliveConn() error {

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
		ironicNotifyClose = ironicChannel.NotifyClose(ironicCloseChan) //一旦消费者的channel有错误，产生一个amqp.Error，channel监听并捕捉到这个错误
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

		}
	END:
		fmt.Printf(time.Now().String() + "keepAliveConn once done\n")
		time.Sleep(10 * time.Second)
	}
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
