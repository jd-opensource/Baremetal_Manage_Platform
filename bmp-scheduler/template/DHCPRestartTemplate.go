package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
)

type DHCPRestartTemplate struct{}

func (c DHCPRestartTemplate) InitCommand(logger *log.Logger, request_id, sn string, task string) {

	var err error
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if _, err = CreateCommand(tx, CommandAction.DHCPRestart.Name, CommandType.DRIVER, request_id, "", sn, 0, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("DHCPRestartTemplate success, request_id:%s, sn:%s", request_id, sn)

ROLLBACK:
	tx.Rollback()
	logger.Warnf("DHCPRestartTemplate failed, request_id:%s, sn:%s", request_id, sn)
}
