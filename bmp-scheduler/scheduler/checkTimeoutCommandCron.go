package scheduler

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/commandLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/processor"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/TimeoutPolicy"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
)

func doCheckTimeoutCommandCron() error {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-scheduler-checkTimeoutCommandCron.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	logger.Point("task_type", "CheckTimeoutCommandCron")
	//没有requestid，临时生成
	logger.Point("logid", commonUtil.GenerateRandUuid())
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	logger.Info("doCheckTimeoutCommandCron starting......")

	key := "SchedulerLock:cps-ironic:check_timeout_command_cron"
	is_set, err := redis.SetNxObjectToRedis(key, "1", 30)
	//fmt.Println("设置redis key: SchedulerLock:cps-ironic:check_timeout_command_cro", is_set)
	if err != nil {
		logger.Warn("CheckTimeoutCommand SetNxObjectToRedis error:", err.Error())
		return nil
	}
	if is_set {
		commands, err := commandLogic.QueryTimeoutCommands(logger, time.Now().Add(-time.Hour*8), time.Now())
		if err != nil {
			logger.Warn("CheckTimeoutCommand QueryTimeoutCommands sql error:", err.Error())
			return nil
		}
		length := len(commands)
		wg := sync.WaitGroup{}
		wg.Add(length)
		for i := 0; i < length; i++ {
			defer wg.Done()
			go dealTimeOutCmd(logger, commands[i])
		}
		wg.Wait()
	}
	return nil
}

func dealTimeOutCmd(logger *log.Logger, command *commandDao.Command) {
	val, _ := json.Marshal(command)
	logger.Info("get timeout comand :", string(val))
	defer func() {
		if r := recover(); r != nil {
			logger.Warn("dealTimeOutCmd processor.Run panic info:", r)
		}
	}()
	if command.Action == CommandAction.Ping.Name {
		processor.BaseProcessor{}.AfterFailedExecute(logger, command, string(val))
		return
	}

	if command.Task == "CollectHardwareInfo" {
		//采集任务如果失败则直接软删除，不影响后面的装机
		q := map[string]interface{}{
			"request_id": command.RequestID,
			"sn":         command.Sn,
			"is_del":     0,
		}
		u := map[string]interface{}{
			"is_del": 1,
		}
		if err := commandDao.UpdateMultiCommands(logger, q, u); err != nil {
			logger.Warnf("dealTimeOutCmd UpdateMultiCommands for CollectHardwareInfo error, sn:%s, error:%s", command.Sn, err.Error())
			return
		}

		// if err := devicepreDao.DeleteBySn(logger, command.Sn); err != nil {
		// 	logger.Warnf("dealTimeOutCmd DeletedevicePreBySn for CollectHardwareInfo error, sn:%s, error:%s", command.Sn, err.Error())
		// }
		return
	}

	if command.TimeoutPolicy == TimeoutPolicy.SKIP {
		processor.BaseProcessor{}.AfterSuccessExecute(logger, command)
	} else {

		if alreadyWarning(logger, command.RequestID, command.Sn) {
			logger.Warn(fmt.Sprintf("already warning for request_id:%s, sn:%s", command.RequestID, command.Sn))
			return
		}
		//sendCommandTimeoutEmail(logger, command)

		//TODO
		command.Status = CommandStatus.ERROR
		command.UpdatedTime = int(time.Now().Unix())
		if err := commandDao.UpdateCommandById(logger, command); err != nil {
			logger.Warn("BaseProcessor afterFailedExecute UpdateCommandById sql error:", err.Error())
		}

		if command.InstanceID != "" {
			instanceEntity, err := instanceDao.GetInstanceByUuid(logger, command.InstanceID)
			if err != nil {
				logger.Warnf("dealTimeOutCmd GetInstanceByUuid error, InstanceID:%s, error:%s", command.InstanceID, err.Error())
			} else {
				instanceEntity.Status = InstanceStatus.Instance_To_Error_Status[command.Task]
				instanceEntity.Reason = command.Action + "Error"
				if err := instanceDao.UpdateInstanceByInstanceId(logger, instanceEntity); err != nil {
					logger.Warnf("更新实例创建失败原因", command.InstanceID, command.Action)
				}
			}
		}

	}
}

func sendCommandTimeoutEmail(logger *log.Logger, command *commandDao.Command) {
	logger.Info("sendCommandTimeoutEmail start......", util.ObjToJson(command))
	device, err := deviceLogic.GetBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("sendCommandTimeoutEmail GetBySn sql error:", command.Sn, err.Error())
	}
	if device == nil {
		logger.Warn("no device found by sn ", command.Sn)
		return
	}
	commands, err := commandLogic.QueryCommandsByRequestIdAndSn(logger, command.RequestID, command.Sn)
	if err != nil {
		logger.Warn("sendCommandTimeoutEmail QueryCommandsByRequestIdAndSn sql error:", command.RequestID, command.Sn, err.Error())
	}
	if len(commands) == 0 {
		return
	}
	var instance *instanceDao.Instance
	if command.RequestID == "" {
		instance = nil
	} else {
		instance, err = instanceLogic.GetByInstanceId(logger, command.InstanceID)
		if err != nil {
			logger.Warn("sendCommandTimeoutEmail GetByInstanceId sql error:", command.InstanceID, err.Error())
			return
		}
	}
	if instance == nil {
		logger.Warn("sendCommandTimeoutEmail empty instance", command.InstanceID)
		return
	}
	//屏蔽发邮件逻辑
	//mailData := emailTemplate.CheckTimeoutCommandTpl{
	//	CurrentDate: time.Now().String(),
	//	Commands:    commands,
	//	Device:      *device,
	//	Instance:    *instance,
	//}
	//
	//logger.Info("[mailData debug]", util.ObjToJson(mailData))
	//receivers := mailLogic.GetMailReceiver(logger, MailType.COMMAND_WARNING)
	//if len(receivers) == 0 {
	//	logger.Warn("no mail receiver found, skip send command warning email")
	//	return
	//}
	//subject := fmt.Sprintf("[%s]指令监控", "bmp-ironic-scheduler")
	//if err := email.SendMailByTpl(emailTemplate.CheckTimeoutCommandMailContent, subject, mailData, receivers); err != nil {
	//	logger.Warn("CheckTimeoutCommand SendMailByTpl error:", err.Error())
	//} else {
	//	logger.Point("sendCommandTimeoutEmail_status", "succeed")
	//}
	//
	//logger.Info("sendCommandTimeoutEmail end......")

}

func alreadyWarning(logger *log.Logger, request_id, sn string) bool {
	key := fmt.Sprintf(Constants.REDIS_KEY_MONITOR_COMMAND_FORMAT, request_id, sn)
	res, err := redis.Keys(key)
	if err != nil {
		logger.Warn("alreadyWarning Keys error:", request_id, sn, err.Error())
	}
	if len(res) > 0 {
		return true
	}
	if err := redis.SetObjectToRedisWithExpire(key, "1", Constants.TIMEOUT_MONITOR_COMMAND); err != nil {
		logger.Warn("alreadyWarning SetObjectToRedisWithExpire error:", key, err.Error())
	}
	return false
}
