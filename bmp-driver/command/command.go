package command

import (
	"encoding/json"
	"fmt"
	"reflect"

	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	beego "github.com/beego/beego/v2/server/web"

	mq "coding.jd.com/aidc-bmp/bmp-driver/service/rabbit_mq"
	log "git.jd.com/cps-golang/log"

	"coding.jd.com/aidc-bmp/bmp-driver/errors"
	"coding.jd.com/aidc-bmp/bmp-driver/util"
)

const (
	OK    = "OK"
	ERROR = "ERROR"
)

var commands []Commandor = []Commandor{}

type Commandor interface {
	ExecuteBefore()
	Execute() error
	ExecuteAfter()
	Accept(string) (Commandor, bool)
}

type BaseCommand struct {
	Logger *log.Logger `json:"-"`
}

func getCommandor(action string) Commandor {
	for _, v := range commands {
		if c, selected := v.Accept(action); selected {
			return c
		}
	}
	return nil
}

func (b *BaseCommand) Prepare(m map[string]interface{}) {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-driver.log")
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	for k, v := range m {
		logger.Point(k, v)
	}
	b.Logger = logger
}

func (b *BaseCommand) Finish() {
	b.Logger.TimeEnd("self_t")
	b.Logger.TimeEnd("all_t")
	b.Logger.Flush()
}

func Run(msg []byte) {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-driver.log")
	logger.Info("ReceiveFromIronicScheduler msg: ", string(msg))
	fmt.Println("ReceiveFromIronicScheduler msg: ", string(msg))
	tmp := map[string]interface{}{}
	if err := json.Unmarshal(msg, &tmp); err != nil {
		logger.Warnf("unmarshal ironic-scheduler message error, msg:%s, error:%s", string(msg), err.Error())
		return
	}
	iAction, ok := tmp["action"]
	if !ok || reflect.TypeOf(iAction).Kind() != reflect.String {
		logger.Warn("unknow command action :", string(msg))
		return
	}
	action := iAction.(string)
	iSn, ok := tmp["sn"]
	if !ok || reflect.TypeOf(iSn).Kind() != reflect.String {
		logger.Warn("unknow command Sn :", string(msg))
		return
	}
	sn := iSn.(string)
	var c Commandor = getCommandor(action)
	if c == nil {
		logger.Warn("unknow command action :", string(msg))
		return
	}
	m := reflect.ValueOf(c).Elem().FieldByName("BaseCommand").MethodByName("Prepare")
	preData := map[string]interface{}{
		"msg":   string(msg),
		"logid": commonUtil.GenerateRandUuid(),
	}
	m.Call([]reflect.Value{reflect.ValueOf(preData)})
	if err := json.Unmarshal(msg, c); err != nil {
		logger.Warn("command unmarshal error :", err.Error())
		return
	}

	var notice map[string]interface{}
	defer func() {
		if err := mq.SendToIronicScheduler(notice); err != nil {
			logger.Warn("send message to ironicScheduler error, msg:%s, error:%s", util.ObjToJson(notice), err.Error())
		} else {
			logger.Infof("send message to ironicScheduler success, msg:%s", util.ObjToJson(notice))
		}
	}()
	//TODO 后面放开
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		t := make([]byte, 1<<16)
	// 		runtime.Stack(t, true)
	// 		notice = map[string]interface{}{
	// 			"sn":      mq.C.Mq.Connection.ReceiveRoutingKey,
	// 			"action":  action,
	// 			"status":  ERROR,
	// 			"message": fmt.Sprintf("run %s failed : %s", action, string(t)),
	// 		}
	// 	}

	// }()

	defer func() {
		m := reflect.ValueOf(c).Elem().FieldByName("BaseCommand").MethodByName("Finish")
		m.Call(nil)
	}()

	c.ExecuteBefore()
	err := c.Execute()
	c.ExecuteAfter()
	var status, message string
	if err != nil {
		status = ERROR
		message = fmt.Sprintf("run %s error:%s.", action, err.Error())
		if errError, ok := err.(errors.Error); ok {
			status = errError.ErrorCode()
			message = fmt.Sprintf("run %s error:%s.", action, errError.Error())
		}
	} else {
		status = OK
		message = fmt.Sprintf("run %s success.", action)
	}
	notice = map[string]interface{}{
		"sn":      sn,
		"action":  action,
		"status":  status,
		"message": message,
	}
	return
}
