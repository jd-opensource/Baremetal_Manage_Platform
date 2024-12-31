package command

import (
	mq "coding.jd.com/aidc-bmp/bmp-driver/service/rabbit_mq"
	"coding.jd.com/aidc-bmp/bmp-driver/util"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

type Heart struct {
	*BaseCommand
	Sn string `json:"sn,omitempty"`
}

func init() {

	commands = append(commands, &Heart{})
}

func (d *Heart) Accept(action string) (Commandor, bool) {
	if action == "Heart" {
		return &Heart{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *Heart) Execute() error {
	//TODO
	return nil
}

func (d *Heart) ExecuteBefore() {
	//TODO
}

func (d *Heart) ExecuteAfter() {
	//TODO
}

func Heartbeat() {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-driver.log")
	msg := map[string]interface{}{
		"sn":           mq.RoutingKey,
		"config":       mq.C,
		"heart_period": 300,
	} //心跳格式待再次确认
	if err := mq.SendToIronicScheduler(msg); err != nil {
		logger.Warnf("Heartbeat.SendToIronicScheduler error:%s", err.Error())
	} else {
		logger.Infof("send heart message to ironicScheduler success, msg:%s", util.ObjToJson(msg))
	}
}
