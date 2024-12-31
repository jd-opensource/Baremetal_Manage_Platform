package apiActor

import (
	"encoding/json"
	"runtime"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/processor"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	log "git.jd.com/cps-golang/log"
)

type ApiActor interface {
	Do(*log.Logger, string, string)
}

type BaseActor struct {
}

var ApiActorMap map[string]ApiActor

func init() {
	ApiActorMap = map[string]ApiActor{}
	ApiActorMap["CollectHardwareInfo"] = NewCollectHardwareInfoActor()
	ApiActorMap["CreateInstances"] = NewCreateInstancesActor()
	ApiActorMap["CutDeviceStock"] = NewCutDeviceStockActor()
	ApiActorMap["DeleteInstance"] = NewDeleteInstanceActor()
	ApiActorMap["InstanceResetPassword"] = NewInstanceResetPasswordActor()
	ApiActorMap["ModifyBandwidth"] = NewModifyBandwidthActor()
	ApiActorMap["ReinstallInstance"] = NewReinstallInstanceActor()
	ApiActorMap["ResetSwitchConfig"] = NewResetSwitchConfigActor()
	ApiActorMap["RestartDhcp"] = NewRestartDhcpActor()
	ApiActorMap["RestartInstances"] = NewRestartInstancesActor()
	ApiActorMap["StartInstances"] = NewStartInstancesActor()
	ApiActorMap["StopInstances"] = NewStopInstancesActor()
	ApiActorMap["PutawayDevice"] = NewPutawayDeviceActor()

}

//业务代码可以直接panic终止运行
func CatchException() {
	if r := recover(); r != nil {
		t := make([]byte, 1<<16)
		runtime.Stack(t, true)
		// b.LogPoints.Warn(string(t))
	}
}

func GetAndStartFirstCommand(logger *log.Logger, sn string) {
	// 不允许状态为running的status出现
	command_running_count, err := commandDao.CountBySnAndStatus(logger, sn, CommandStatus.RUNNING)
	if err != nil {
		logger.Warn(sn, "CountBySnAndStatus sql error:", err.Error())
	}
	if command_running_count > 0 {
		logger.Warn(sn, " commands are running. skip start message")
		return
	}
	// 不允许状态为error的status出现
	command_error_count, err := commandDao.CountBySnAndStatus(logger, sn, CommandStatus.ERROR)
	if err != nil {
		logger.Warn(sn, "CountBySnAndStatus sql error:", err.Error())
	}
	if command_error_count > 0 {
		logger.Warn(sn, " commands are error. skip start message")
		return
	}
	param := map[string]interface{}{
		"sn":     sn,
		"status": CommandStatus.WAIT,
		"is_del": 0,
	}
	command, err := commandDao.GetFirstCommand(logger, param)
	if err != nil {
		logger.Warn(sn, "GetFirstCommand sql error:", err.Error())
		return
	}
	b, _ := json.Marshal(command)
	logger.Info("get first cmd: ", string(b))
	processor.ProcessorMap[command.Action].Process(logger, command)

}
