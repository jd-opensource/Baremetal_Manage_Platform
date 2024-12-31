package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
)

type ResetSwitchConfigTemplate struct{}

func (c ResetSwitchConfigTemplate) InitCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if id, err = CreateCommand(tx, CommandAction.CreateVRF.Name, CommandType.NETWORK, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CreateVSIInterface.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CreateVNI.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CreateVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.AddArpStatic.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetBandwidth.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.BindingVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SaveSwitchConfig.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.SaveConfigToFtpServer.Name, CommandType.NETWORK, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("ResetSwitchConfigTemplate_status", "finish")

ROLLBACK:
	tx.Rollback()
	logger.Point("ResetSwitchConfigTemplate_status", "failed")
}
