package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
)

type PutawayDeviceTemplate struct{}

func (c PutawayDeviceTemplate) InitCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	var err error
	tx := dao.GetGormTx(logger)
	tx.Begin()
	var id int64

	//DHCPConfigAddSubnet，CheckInitConfig
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddSubnet.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}

	if _, err = CreateCommand(tx, CommandAction.CheckInitConfig.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("PutawayDeviceTemplate success, sn:%s", sn)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("PutawayDeviceTemplate failed, sn:%s", sn)
}
