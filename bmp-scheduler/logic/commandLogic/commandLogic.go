package commandLogic

import (
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	log "git.jd.com/cps-golang/log"
)

func Create(logger *log.Logger, param types.CreateCommandRequest) int64 {
	entity := &commandDao.Command{
		RequestID:       param.RequestId,
		Sn:              param.Sn,
		Action:          param.Action,
		Type:            param.CommandType,
		InstanceID:      param.InstanceId,
		Status:          CommandStatus.WAIT,
		ExecuteCount:    0,
		ParentCommandID: param.ParentCommandId,
		CreatedTime:     int(time.Now().Unix()),
		UpdatedTime:     int(time.Now().Unix()),
	}
	id, err := commandDao.AddCommand(logger, entity)
	if err != nil {
		panic(err)
	}
	return id
}

func GetBySnAndActionAndStatus(logger *log.Logger, sn, action, status string) (*commandDao.Command, error) {
	param := map[string]interface{}{
		"sn":     sn,
		"action": action,
		"status": status,
		"is_del": 0,
	}
	return commandDao.GetOneCommand(logger, param)
}

func QueryTimeoutCommands(logger *log.Logger, stime, etime time.Time) ([]*commandDao.Command, error) {
	statuses := []string{CommandStatus.RUNNING, CommandStatus.ERROR}
	param := map[string]interface{}{
		"is_del":          0,
		"timeout_time.gt": stime.String(),
		"timeout_time.lt": etime.String(),
		"status.in":       statuses,
	}
	return commandDao.GetAllCommand(logger, param)
}

func QueryCommandsByRequestIdAndSn(logger *log.Logger, request_id, sn string) ([]*commandDao.Command, error) {
	param := map[string]interface{}{
		"is_del":     0,
		"request_id": request_id,
		"sn":         sn,
	}
	return commandDao.GetAllCommand(logger, param)
}
