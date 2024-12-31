package task

import (
	"fmt"
	"time"

	"coding.jd.com/bmp/cps-agent/api"
	"coding.jd.com/bmp/cps-agent/api/models"
	"coding.jd.com/bmp/cps-agent/global"

	"github.com/astaxie/beego/logs"
)

type HeartbeatTask struct {
}

func init() {
	Register(&HeartbeatTask{})
}

func (t *HeartbeatTask) Name() string {
	return "HeartbeatTask"
}

func (t *HeartbeatTask) Spec() string {
	return global.DefaultString(global.AgentConfig.Agent.Heartbeat.Scheduler, "0/60 * * * * *")
}

func (t *HeartbeatTask) TaskFunc() error {

	logs.Debug("Heartbeat")
	url := global.AgentConfig.Proxy.URL + "heartbeat"
	version := fmt.Sprintf("%s.%s", global.VERSION, global.BUILD_TIME)
	// global.SerialNumber = "minping-mock-sn" //TODO 记得注释掉
	status := models.AgentStatus{SN: global.SerialNumber, AgentVersion: version, Timestamp: time.Now().Unix(), Status: models.AGENT_OK}

	if err := api.Heartbeat(url, status); err != nil {
		logs.Info("Post Heartbeat Error:%s", err.Error())
	}
	return nil
}

func (t *HeartbeatTask) Shutdown() error {
	logs.Debug("Heartbeat Shutdown")
	url := global.AgentConfig.Proxy.URL + "heartbeat"
	version := fmt.Sprintf("%s.%s", global.VERSION, global.BUILD_TIME)
	status := models.AgentStatus{SN: global.SerialNumber, AgentVersion: version, Timestamp: time.Now().Unix(), Status: models.AGENT_STOP}

	api.Heartbeat(url, status)
	return nil
}
