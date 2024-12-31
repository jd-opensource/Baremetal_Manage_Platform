package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type DHCPRestartProcessor struct {
	BaseProcessor
}

func NewDHCPRestartProcessor() DHCPRestartProcessor {
	b := DHCPRestartProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b DHCPRestartProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b DHCPRestartProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b DHCPRestartProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	az, err := redis.GetObjectFromRedis(Constants.DHCP_RESTART_KEY)
	if err != nil {
		logger.Warn("DHCPRestartProcessor GetObjectFromRedis error:", err.Error())
	}
	if az == "" {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process DHCPRestart error, az %s is null", az)})
	}

	dhcp := driverTypes.DHCPRestart{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(az, dhcp); err != nil {
		logger.Warn("DHCPRestartProcessor SendToAgent error:", command.Sn, err.Error())
		return
	}
	logger.Infof("DHCPRestartProcessor_SendToAgent_Msg: %s", util.ObjToJson(dhcp))
}

func (b DHCPRestartProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
