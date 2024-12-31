package apiActor

import (
	"encoding/json"
	"fmt"

	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type RestartDhcpActor struct {
	BaseActor
}

func NewRestartDhcpActor() RestartDhcpActor {
	return RestartDhcpActor{}
}

func (s RestartDhcpActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.RestartDhcpMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("RestartDhcpActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	if err := redis.SetObjectToRedisWithExpire(Constants.DHCP_RESTART_KEY, message.Az, 60*60*24); err != nil {
		logger.Warn("RestartDhcpActor SetObjectToRedis error:", message.Az, err.Error())
	}
	DEFAULT_SN := "DHCP_%s"
	sn := fmt.Sprintf(DEFAULT_SN, message.Az)
	templater := template.DHCPRestartTemplate{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, message.RequestId, sn, task)
	GetAndStartFirstCommand(logger, sn)
}
