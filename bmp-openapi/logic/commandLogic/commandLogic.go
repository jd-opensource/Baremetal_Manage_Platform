package commandLogic

import (
	"fmt"
	"strings"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-openapi/service/rabbit_mq"
	"coding.jd.com/aidc-bmp/bmp-openapi/service/redis"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	commandStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	rabbitEvent "git.jd.com/cps-golang/ironic-common/ironic/event"
	rabbitIronicMsgCommand "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	"github.com/jinzhu/gorm"
)

func CommandEntity2Command(c *commandDao.Command) *responseTypes.Command {
	return &responseTypes.Command{
		Id:              int64(c.ID),
		RequestId:       c.RequestID,
		Sn:              c.Sn,
		InstanceId:      c.InstanceID,
		Action:          c.Action,
		Type:            c.Type,
		Status:          c.Status,
		ParentCommandId: c.ParentCommandID,
		ExecuteCount:    c.ExecuteCount,
		//CreateTime:      c.CreateTime,
		//UpdateTime:      c.UpdateTime,
	}
}

func Query(logger *log.Logger, query *requestTypes.QueryCommandsRequest, p util.Pageable) ([]*responseTypes.Command, int64, error) {
	offset, limit := p.Pageable2offset()
	param := map[string]interface{}{
		"is_del": "0",
	}
	if query.RequestId != "" {
		param["request_id"] = query.RequestId
	}
	if query.InstanceId != "" {
		param["instance_id"] = query.InstanceId
	}
	if query.Sn != "" {
		param["sn"] = query.Sn
	}

	count, err := commandDao.GetCommandCount(logger, param)
	if err != nil {
		logger.Warn("GetCommandCount sql error:", err.Error())
		return nil, 0, err
	}

	command_entities, err := commandDao.GetMultiCommand(logger, param, nil, nil, nil, offset, limit)
	if err != nil {
		logger.Warn("QueryCommands GetMultiCommand sql error:", err.Error())
		return nil, 0, err
	}
	resCmds := []*responseTypes.Command{}
	for _, entity := range command_entities {
		command := CommandEntity2Command(entity)
		if strings.ToLower(entity.Status) == commandStatus.FINISH || strings.ToLower(entity.Status) == commandStatus.ERROR {
			result, err := redis.GetObjectFromRedis(fmt.Sprintf(constants.COMMAND_RESULT, fmt.Sprint(entity.ID)))
			if err != nil {
				logger.Warnf("QueryCommands GetObjectFromRedis error, id:%d, error:%s", entity.ID, err.Error())
			} else {
				command.Result = result
			}
		}
		resCmds = append(resCmds, command)
	}
	return resCmds, count, nil

}

func CancelCommand(logger *log.Logger, sn string) error {
	param := map[string]interface{}{
		"is_del":    0,
		"sn":        sn,
		"status.in": []string{commandStatus.RUNNING, commandStatus.WAIT},
	}

	commands, err := commandDao.GetAllCommand(logger, param)
	if err != nil {
		logger.Warnf("CancelCommand GetMultiCommand sql error, sn:%s, error:%s", sn, err.Error())
		return err
	}

	command := commands[0]
	if exist := util.InArray(command.Status, []string{commandStatus.RUNNING, commandStatus.WAIT, commandStatus.ERROR}); exist {
		command.Status = commandStatus.CANCEL
		if err := commandDao.UpdateCommandById(logger, command); err != nil {
			logger.Warnf("CancelCommand UpdateCommandById sql error, id:%d, error:%s", command.ID, err.Error())
			return err
		}
	}

	for command.ParentCommandID != 0 {
		command, err = commandDao.GetCommandById(logger, command.ParentCommandID)
		if err != nil {
			logger.Warnf("CancelCommand GetCommandById sql error, id:%d, error:%s", command.ParentCommandID, err.Error())
			continue
		}

		if exist := util.InArray(command.Status, []string{commandStatus.RUNNING, commandStatus.WAIT, commandStatus.ERROR}); exist {
			command.Status = commandStatus.CANCEL
			if err := commandDao.UpdateCommandById(logger, command); err != nil {
				logger.Warnf("CancelCommand UpdateCommandById sql error, id:%d, error:%s", command.ID, err.Error())
				continue
			}
		}
	}
	return nil
}

func RetryCommand(logger *log.Logger, offsetCommandId int64) error {

	logid := logger.GetPoint("logid").(string)
	command, err := commandDao.GetCommandById(logger, offsetCommandId)
	if err != nil {
		logger.Warnf("RetryCommand GetCommandById sql error, id:%d, error:%s", offsetCommandId, err.Error())
		return err
	}
	command.Status = commandStatus.WAIT
	command.TimeoutTime = time.Now()
	if err := commandDao.UpdateCommandById(logger, command); err != nil {
		logger.Warnf("UpdateCommandById sql error, id:%d, error:%s", command.ID, err.Error())
		return err
	}

	var childCmd *commandDao.Command
	childCmd, err = commandDao.GetCommandByParentId(logger, command.ID)
	if err != nil {
		logger.Warnf("retry command GetCommandByParentId error, id:%d, err:%s", command.ID, err.Error())
	}
	for err == nil {
		logger.Infof("retry command id:%d, childcmd id:%d", command.ID, childCmd.ID)
		childCmd.Status = commandStatus.WAIT
		childCmd.TimeoutTime = time.Time{}
		if err := commandDao.UpdateCommandById(logger, childCmd); err != nil {
			logger.Warnf("RetryCommand UpdateCommandById sql error, id:%d, error:%s", childCmd.ID, err.Error())
			return err
		}
		childCmd, err = commandDao.GetCommandByParentId(logger, childCmd.ID)
	}
	if !gorm.IsRecordNotFoundError(err) {
		return err
	}

	msg := rabbitIronicMsgCommand.CallbackCommandMessage{
		CommandMessage: rabbitIronicMsgCommand.CommandMessage{
			Sn:     command.Sn,
			Action: "Start",
		},
		Status: "OK",
	}
	event, err := rabbitEvent.NewEvent(msg, logid, "")
	if err != nil {
		logger.Warnf("RetryCommand convert event msg error, sn:%s, error:%s", command.Sn, err.Error())
		return err
	}
	if err := rabbitMq.SendToApi2Scheduler(event); err != nil {
		logger.Warnf("RetryCommand SendToScheduler error, msg:%s, error:%s", util.ObjToJson(msg), err.Error())
		return err
	}

	logger.Infof("RetryCommand SendToScheduler success, msg: %s", util.ObjToJson(msg))

	return nil
}
