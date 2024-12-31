package processor

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/messageDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/mailLogic"
	messagelogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/messageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	messageType "git.jd.com/cps-golang/ironic-common/ironic/enums/MessageType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/TimeoutPolicy"
	ironicCommon "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

const (
	OK                        = "OK"
	ERROR                     = "ERROR"
	COMMAND_MAX_EXECUTE_COUNT = 3
)

type BaseProcessor struct {
	doProcess    func(*log.Logger, *commandDao.Command)
	afterProcess func(*log.Logger, *commandDao.Command)
}

type Processer interface {
	Process(*log.Logger, *commandDao.Command)
	doProcess(*log.Logger, *commandDao.Command)
	afterProcess(*log.Logger, *commandDao.Command)
	Callback(*log.Logger, *commandDao.Command, ironicCommon.CallbackCommandMessage)
}

type Exception struct {
	Msg string
}

type ProcessAbortException struct {
	Exception
	Msg string
}

var IloUsername string
var IloPassword string
var ProcessorMap map[string]Processer

func init() {
	ProcessorMap = map[string]Processer{}
	ProcessorMap["AddArpStatic"] = NewAddArpStaticProcessor()
	ProcessorMap["DHCPConfigAddSubnet"] = NewDHCPConfigAddSubnetProcessor()
	ProcessorMap["CheckInitConfig"] = NewCheckInitConfigProcessor()
	ProcessorMap["DHCPConfigAddHost"] = NewAddDHCPConfigHostProcessor()
	ProcessorMap["DHCPConfigDelHost"] = NewDelDHCPConfigHostProcessor()
	ProcessorMap["BindingIPv6"] = NewBindingIPv6Processor()
	ProcessorMap["BindingVSI"] = NewBindingVSIProcessor()
	ProcessorMap["CleanArpStatic"] = NewCleanArpStaticProcessor()
	ProcessorMap["CleanBandwidth"] = NewCleanBandwidthProcessor()
	ProcessorMap["CleanBlockDevice"] = NewCleanBlockDeviceProcessor()
	ProcessorMap["CleanRaid"] = NewCleanRaidProcessor()
	ProcessorMap["CollectHardwareInfo"] = NewCollectHardwareInfoProcessor()
	ProcessorMap["CollectDiskLocations"] = NewCollectDiskLocationsProcessor()
	ProcessorMap["CreateVNI"] = NewCreateVNIProcessor()
	ProcessorMap["CreateVRFBalance"] = NewCreateVRFBalanceProcessor()
	ProcessorMap["CreateVRF"] = NewCreateVRFProcessor()
	ProcessorMap["CreateVSIInterface"] = NewCreateVSIInterfaceProcessor()
	ProcessorMap["CreateVSI"] = NewCreateVSIProcessor()
	ProcessorMap["DelDHCPConfigHost"] = NewDelDHCPConfigHostProcessor()
	ProcessorMap["DHCPRestart"] = NewDHCPRestartProcessor()
	ProcessorMap["InitRootDevice"] = NewInitRootDeviceProcessor()
	ProcessorMap["InitRootDeviceV2"] = NewInitRootDeviceProcessorV2()
	ProcessorMap["MakePartitions"] = NewMakePartitionsProcessor()
	ProcessorMap["MakeRaid"] = NewMakeRaidProcessor()
	ProcessorMap["Ping"] = NewPingProcessor()
	ProcessorMap["PowerCycle"] = NewPowerCycleProcessor()
	ProcessorMap["PowerOff"] = NewPowerOffProcessor()
	ProcessorMap["PowerOn"] = NewPowerOnProcessor()
	ProcessorMap["PowerReset"] = NewPowerResetProcessor()
	ProcessorMap["Qcow2MakePartitions"] = NewQcow2MakePartitionsProcessor()
	ProcessorMap["SaveConfigToFtpServer"] = NewSaveConfigToFtpServerProcessor()
	ProcessorMap["SaveSwitchConfig"] = NewSaveSwitchConfigProcessor()
	ProcessorMap["SDNAddAliasIP"] = NewSDNAddAliasIPProcessor()
	ProcessorMap["SDNClean"] = NewSDNCleanProcessor()
	ProcessorMap["SDNDeleteAliasIP"] = NewSDNDeleteAliasIPProcessor()
	ProcessorMap["SDNEipBinding"] = NewSDNEipBindingProcessor()
	ProcessorMap["SDNEipUnBinding"] = NewSDNEipUnBindingProcessor()
	ProcessorMap["SDNRegister"] = NewSDNRegisterProcessor()
	ProcessorMap["SDNSetNetwork"] = NewSDNSetNetworkProcessor()
	ProcessorMap["SDNSubnetBinding"] = NewSDNSubnetBindingProcessor()
	ProcessorMap["SDNSubnetUnBinding"] = NewSDNSubnetUnBindingProcessor()
	ProcessorMap["SetBandwidth"] = NewSetBandwidthProcessor()
	ProcessorMap["SetBond"] = NewSetBondProcessor()
	ProcessorMap["SetDISKBoot"] = NewSetDISKBootProcessor()
	ProcessorMap["SetHostname"] = NewSetHostnameProcessor()
	ProcessorMap["SetKeypairs"] = NewSetKeypairsProcessor()
	// ProcessorMap["SetNetwork"] = NewSetNetworkProcessor()//废弃
	ProcessorMap["SetPassword"] = NewSetPasswordProcessor()
	ProcessorMap["SetPXEBoot"] = NewSetPXEBootProcessor()
	ProcessorMap["SetRetailBond"] = NewSetRetailBondProcessor()
	ProcessorMap["SetUserCmd"] = NewSetUserCmdProcessor()
	ProcessorMap["UnBindingVSI"] = NewUnBindingVSIProcessor()
	ProcessorMap["UploadSystemLog"] = NewUploadSystemLogProcessor()
	ProcessorMap["WriteImage"] = NewWriteImageProcessor()
	ProcessorMap["WriteImageTar"] = NewWriteImageTarProcessor()

	//以下用于cloudinit方式
	ProcessorMap["InitNode"] = NewInitNodeProcessor()
	ProcessorMap["FormatPartitions"] = NewFormatPartitionsProcessor()
	ProcessorMap["MountPartitions"] = NewMountPartitionsProcessor()
	ProcessorMap["SetCloudinitConf"] = NewSetCloudinitConfProcessor()
	// ProcessorMap["SetNocloudUserData"] = NewSetNocloudUserDataProcessor()
	// ProcessorMap["SetNocloudMetaData"] = NewSetNocloudMetaDataProcessor()

	ProcessorMap["SetUserData"] = NewSetNocloudUserDataProcessor()
	ProcessorMap["SetMetaData"] = NewSetNocloudMetaDataProcessor()

	ProcessorMap["SetNetwork"] = NewSetNocloudNetworkProcessor() //linux设置network
	ProcessorMap["InitNode"] = NewInitNodeProcessor()
	ProcessorMap["FormatPartitions"] = NewFormatPartitionsProcessor()
	ProcessorMap["MountPartitions"] = NewMountPartitionsProcessor()

	//以下用于windows方式
	ProcessorMap["SetNetworkWindows"] = NewSetNetworkDataProcessor()
}

// getCurrentGoroutineStack 获取当前Goroutine的调用栈，便于排查panic异常
func getCurrentGoroutineStack() string {

	const defaultStackSize = 4096

	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

//Process 处理单个子command的入口
func (b BaseProcessor) Process(logger *log.Logger, command *commandDao.Command) {

	logger.Info("[开始处理command]:", command.Action)
	action := CommandAction.CommandActionDict[command.Action]
	//防止有些指令没有注册策略和过期值，给默认值
	if action.Policy == "" {
		action.Policy = TimeoutPolicy.WARNING
	}
	if action.Timeout == 0 {
		action.Timeout = 180
	}
	command.Status = CommandStatus.RUNNING
	command.ExecuteCount += 1
	command.UpdatedTime = int(time.Now().Unix())
	command.TimeoutTime = time.Now().Add(time.Duration(action.Timeout) * time.Second)
	command.TimeoutPolicy = action.Policy
	if err := commandDao.UpdateCommandById(logger, command); err != nil {
		logger.Warn("Process UpdateCommandById sql error:", err.Error())
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Warnf("processor.Run panic, error:%v, stack:%s", r, getCurrentGoroutineStack())
			// if reflect.TypeOf(r).String() == "processor.ProcessAbortException" {
			// 	abortProcess(logger, command, t)
			// } else {
			// 	sendErrorCallback(logger, command, t)
			// }
		}
	}()
	p, ok := ProcessorMap[command.Action]
	if !ok {
		logger.Warnf("ProcessorMap not found key %s", command.Action)
		return
	}
	p.doProcess(logger, command)
	p.afterProcess(logger, command)
}

// ironic-agent/driver 投递来command异步执行结果的入口
func (b BaseProcessor) CallBack(logger *log.Logger, command *commandDao.Command, msg ironicCommon.CallbackCommandMessage) {
	val, _ := json.Marshal(msg)
	v, _ := json.Marshal(command)
	if strings.ToUpper(msg.Status) == OK {
		logger.Info("CallBack SuccessExecute command from driver/agent: ", string(v))
		b.AfterSuccessExecute(logger, command)
	} else {
		logger.Warn("CallBack FailedExecute command from driver/agent: ", string(v))
		b.AfterFailedExecute(logger, command, string(val))
	}
	if err := redis.SetObjectToRedis(fmt.Sprintf(constants.COMMAND_RESULT, fmt.Sprint(command.ID)), string(val)); err != nil {
		logger.Warn("BaseProcessor CallBack SetObjectToRedis error:", command.ID, err.Error())
	}
}

func (b BaseProcessor) AfterSuccessExecute(logger *log.Logger, command *commandDao.Command) {
	logger.Info("[command指令处理完成]:", command.Action, "准备将command状态更新为finish")
	command.Status = CommandStatus.FINISH
	command.UpdatedTime = int(time.Now().Unix())
	if err := commandDao.UpdateCommandById(logger, command); err != nil {
		logger.Warn("afterSuccessExecute UpdateCommandById sql error:", err.Error())
	}
	instance := &instanceDao.Instance{
		InstanceID: command.InstanceID,
		Reason:     "",
	}
	instanceDao.UpdateSuccessInstanceByInstanceId(logger, instance) //实例失败原因置为空
	instanceDao.UpdateInstanceByInstanceId(logger, instance)
	logger.Info("更新实例创建原因为空，success", command.InstanceID)
	next_command, err := commandDao.GetCommandByParentId(logger, int64(command.ID))
	if err != nil {
		logger.Warn("afterSuccessExecute GetCommandByParentId sql error:", command.ID, err.Error())
	}
	v, _ := json.Marshal(next_command)
	if next_command != nil {
		logger.Info("start success_execute_next_command: ", string(v))
		b.Process(logger, next_command)
	} else {
		//采集任务时，设置device的采集状态为成功
		deviceEntity, err := deviceDao.GetBySn(logger, command.Sn)
		if err != nil {
			logger.Warnf("AfterSuccessExecute.GetBySn error, sn:%s, error:%s", command.Sn, err.Error())
		} else {
			deviceEntity.CollectStatus = "1" //1表示采集成功
			deviceEntity.CollectFailReason = ""
			if err := deviceDao.UpdateDeviceAnywhere(logger, deviceEntity); err != nil {
				logger.Warnf("AfterSuccessExecute.UpdateDeviceAnywhere error, sn:%s, error:%s", command.Sn, err.Error())
			}
		}

		auditLogLogic.SetAuditLogToSuccess(logger, command.Sn, command.InstanceID, command.Task)

		logger.Infof("全部指令执行完毕，开始更新设备%s实例状态", command.Sn)
		instanceLogic.CallBack(logger, command.InstanceID)
	}
}

func (b BaseProcessor) AfterFailedExecute(logger *log.Logger, command *commandDao.Command, failed_msg string) {
	v, _ := json.Marshal(command)
	// fmt.Println(command.Action, "已经执行的次数", command.ExecuteCount)
	if command.ExecuteCount > COMMAND_MAX_EXECUTE_COUNT {
		return
	} else if command.ExecuteCount == COMMAND_MAX_EXECUTE_COUNT {
		logger.Warn("execute faild max time, command: ", string(v))
		command.ExecuteCount += 1
		command.Status = CommandStatus.ERROR
		command.UpdatedTime = int(time.Now().Unix())
		if err := commandDao.UpdateCommandById(logger, command); err != nil {
			logger.Warn("BaseProcessor afterFailedExecute UpdateCommandById sql error:", err.Error())
		}
		msg := &ironicCommon.CallbackCommandMessage{}
		if err := json.Unmarshal([]byte(failed_msg), msg); err != nil {
			logger.Warn("BaseProcessor unmarshal json error:", err.Error())
		}
		var instanceEntity *instanceDao.Instance
		var err error
		if command.InstanceID != "" {

			instanceEntity, err = instanceDao.GetInstanceByUuid(logger, command.InstanceID)
			if err != nil {
				logger.Warnf("AfterFailedExecute.GetInstanceByUuid error, command_id:%d, instance_id:%s, error:%s", command.ID, command.InstanceID, err.Error())
				return
			}

			instanceEntity.Status = InstanceStatus.Instance_To_Error_Status[command.Task] //实例状态置为特定错误

			if command.Type == "driver" {
				instanceEntity.Reason = msg.Status
			}
			if instanceEntity.Reason == "" {
				instanceEntity.Reason = command.Action + "Error"
			}

			instanceDao.UpdateInstanceByInstanceId(logger, instanceEntity)
			logger.Info("更新实例创建失败原因", command.InstanceID, msg.Status)

		}
		if command.Task == "CollectHardwareInfo" { //采集失败也需要记录失败原因
			deviceEntity, err := deviceDao.GetBySn(logger, command.Sn)
			if err != nil {
				logger.Warnf("AfterFailedExecute.CollectHardwareInfo.GetBySn error, sn:%s, error:%s", command.Sn, err.Error())
			} else {
				deviceEntity.CollectStatus = "4" //4表示采集失败
				deviceEntity.CollectFailReason = command.Action + "Error"
				if err := deviceDao.UpdateDeviceById(logger, deviceEntity); err != nil {
					logger.Warnf("AfterFailedExecute.CollectHardwareInfo.UpdateDeviceById error, sn:%s, error:%s", command.Sn, err.Error())
				}
			}
		}
		webMsg := messageDao.WebMessage{
			MessageID: util.GetUuid("m-"),
			// UserID:         logger.GetPoint("user_id").(string),?????
			SN:             command.Sn,
			MessageType:    messageType.MessageTypeOperation,
			MessageSubType: command.Task, //注意在common库中和task(message)保持一致
			ResourceType:   "instance",
			Result:         "failed",
			FinishTime:     int(time.Now().Unix()),
			Detail:         command.Sn,
		}
		mailSubject := ""
		mailTpl := ""
		if instanceEntity != nil {
			webMsg.UserID = instanceEntity.UserID
			webMsg.InstanceID = instanceEntity.InstanceID
			webMsg.InstanceName = instanceEntity.InstanceName
			webMsg.Detail = fmt.Sprintf("%s/%s", instanceEntity.InstanceName, command.Sn)

			user, err := userDao.GetUserById(logger, instanceEntity.UserID)
			if err != nil {
				logger.Warnf("AfterFailedExecute.GetUserById error, user_id:%s, error:%s", instanceEntity.UserID, err.Error())
			}
			if user != nil {
				webMsg.UserName = user.UserName
			}

			mailSubject = fmt.Sprintf("操作消息-【%s】失败通知", webMsg.MessageSubType)
			mailTpl = mailLogic.OperationAfterInstanceTemplate

		} else {
			mailSubject = fmt.Sprintf("操作消息-【%s】失败通知", webMsg.MessageSubType)
			mailTpl = mailLogic.OperationBeforeInstanceTemplate
		}

		if err := messagelogic.SendWebNotice(logger, webMsg); err != nil {
			logger.Warnf("BaseProcessor.AfterFailedExecute.SendWebNotice error:%s", err.Error())
		}
		//发送邮件通知
		err = mailLogic.SendMailByTpl(logger, mailTpl, mailSubject, WebMessage2MailMessage(webMsg))
		if err != nil {
			logger.Warnf("doLicenseTimeoutCron.SendMail error:%s", err.Error())
		}
		// sendErrorMailMessage(logger, command, []byte(failed_msg))

		auditLogLogic.SetAuditLogToFail(logger, command.Sn, command.InstanceID, command.Task, command.Action+"Error")

	} else {
		logger.Info("execute again, command: ", string(v))
		b.Process(logger, command) //重新执行一次
	}
}

func WebMessage2MailMessage(p messageDao.WebMessage) mailLogic.MailMessage {
	m := mailLogic.MailMessage{
		MessageID:      p.MessageID,
		MessageType:    p.MessageType,
		MessageSubType: messagelogic.Types2ShowZh[p.MessageSubType],
		ResourceType:   p.ResourceType,
		FinishTime:     time.Unix(int64(p.FinishTime), 0).Format("2006-01-02 15:04:05"),
		Detail:         p.Detail,
		SN:             p.SN,
		UserID:         p.UserID,
		UserName:       p.UserName,
		InstanceID:     p.InstanceID,
		InstanceName:   p.InstanceName,
	}
	if m.MessageSubType == "" {
		m.MessageSubType = p.MessageSubType
	}
	if m.UserName == "" {
		m.UserName = "admin"
	}
	return m
}

func abortProcess(logger *log.Logger, command *commandDao.Command, e []byte) {
	command.Status = CommandStatus.ERROR
	command.UpdatedTime = int(time.Now().Unix())
	if err := commandDao.UpdateCommandById(logger, command); err != nil {
		logger.Warn("abortProcess UpdateCommandById sql error:", err.Error())
	}
	err := redis.SetObjectToRedis(fmt.Sprintf(constants.COMMAND_RESULT, fmt.Sprint(command.ID)), string(e))
	if err != nil {
		logger.Warn("QueryCommands SetObjectToRedis error:", err.Error())
	}
	sendErrorMailMessage(logger, command, e)
}

func sendErrorCallback(logger *log.Logger, command *commandDao.Command, e []byte) {
	//TODO
}

func sendErrorMailMessage(logger *log.Logger, command *commandDao.Command, e []byte) {
	//TODO
}

func loadInstanceExtra(logger *log.Logger, instance_id string) *types.InstanceExtra {
	extra_json, err := redis.GetObjectFromRedis(fmt.Sprintf(constants.INSTANCE_EXTRA_KEY, instance_id))
	if err != nil || extra_json == "" {
		logger.Warn("loadInstanceExtra GetObjectFromRedis error:", instance_id, err.Error())
		return nil
	}
	var r = types.InstanceExtra{}
	if json.Unmarshal([]byte(extra_json), &r); err != nil {
		logger.Warn("loadInstanceExtra Unmarshal error:", extra_json, err.Error())
		return nil
	}
	return &r
}
