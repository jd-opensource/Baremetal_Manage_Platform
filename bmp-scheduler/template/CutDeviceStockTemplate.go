package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
)

type CutDeviceStockTemplate struct{}

func (c CutDeviceStockTemplate) InitCommand(logger *log.Logger, request_id, sn string, task string) {

	var err error
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if _, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.DRIVER, request_id, "", sn, 0, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("CutDeviceStockTemplate success, sn:%s", sn)

ROLLBACK:
	tx.Rollback()
	logger.Warnf("CutDeviceStockTemplate failed, sn:%s", sn)
}