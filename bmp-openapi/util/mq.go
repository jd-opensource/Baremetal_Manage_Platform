package util

import (
	"fmt"
	"log"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/streadway/amqp"
)

// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost，下面是两个/!!!
var (
	host, port, username, password, exchange, vhost, receive_routing_key string
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	host, _ = beego.AppConfig.String("bmp_mq_host")
	port, _ = beego.AppConfig.String("bmp_mq_port")
	username, _ = beego.AppConfig.String("bmp_mq_user")
	password, _ = beego.AppConfig.String("bmp_mq_password")
	vhost, _ = beego.AppConfig.String("bmp_mq_vhost")

	mqurl := "amqp://" + username + ":" + password + "@" + host + ":" + port + "/" + vhost
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     mqurl,
	}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！")

	return rabbitmq
}

// Destory 断开channel和connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQSimple
// 简单模式Step 1.创建简单模式下的RabbitMq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	exchange, _ = beego.AppConfig.String("bmp_mq_exchange")
	receive_routing_key, _ = beego.AppConfig.String("bmp_mq_receive_routing_key")
	return NewRabbitMQ(queueName, exchange, receive_routing_key) //exchange，key要在web界面binds中按照配置进行相应的设置
}

// 简单模式Step 2:简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		true,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(time.Now(), err.Error())
	}
	queueBind(r.channel, r.QueueName, r.QueueName, r.Exchange) //手动绑定 队列名=》exchange
	// 2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	fmt.Println("send success,exchangeName", r.Exchange, "queue:", r.QueueName, "msg:", message)
}
func queueBind(ch *amqp.Channel, queueName, routing_key, exchange string) error {
	fmt.Println("auto bind success,queue:", queueName, "routing_key", routing_key)
	return ch.QueueBind(
		queueName,   // queue name
		routing_key, // routing key
		exchange,    // exchange
		false,
		nil,
	)
}

// ConsumeSimple 使用 goroutine 消费消息
func (r *RabbitMQ) ConsumeSimple() {

	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		true,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(time.Now(), err.Error())
	}
	queueBind(r.channel, r.QueueName, r.QueueName, r.Exchange)
	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(time.Now(), err.Error())
	}

	forever := make(chan bool)
	// 启用协和处理消息
	go func() {
		for d := range msgs {
			// 实现我们要实现的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}
