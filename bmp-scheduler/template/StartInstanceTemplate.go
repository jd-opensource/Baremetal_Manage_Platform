package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
)

type StartInstanceTemplate struct{}

func (c StartInstanceTemplate) InitCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	var err error
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if _, err = CreateCommand(tx, CommandAction.PowerOn.Name, CommandType.DRIVER, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("StartInstanceTemplate success, instanceId:%s", instance_id)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("StartInstanceTemplate failed, instanceId:%s", instance_id)

}