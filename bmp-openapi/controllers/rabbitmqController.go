package controllers

import (
	"time"

	util "coding.jd.com/aidc-bmp/bmp-openapi/util"
)

// DeviceController operations for Device
type RabbitMqController struct {
	BaseController
}

// QueryDevices ...
// @router / [get]
func (c *RabbitMqController) Send() {

	rabbitmq := util.NewRabbitMQSimple("cn-north-1c")
	rabbitmq.PublishSimple("Hello goFrame!" + time.Now().Format("2006-01-02 15:04:05"))

	return
}
func (c *RabbitMqController) Receive() {

	//util.Receive()
	rabbitmq := util.NewRabbitMQSimple("cn-north-1c")

	rabbitmq.ConsumeSimple()

	return
}
