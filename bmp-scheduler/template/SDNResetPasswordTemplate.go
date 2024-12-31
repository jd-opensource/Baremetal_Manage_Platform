package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type SDNResetPasswordTemplate struct{}

func (b SDNResetPasswordTemplate) accept(network_type string) bool {
	return NetworkType.VPC == network_type
}

func (b SDNResetPasswordTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if id, err = CreateCommand(tx, CommandAction.SDNSubnetUnBinding.Name, CommandType.SDN, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.InitRootDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SDNSubnetBinding.Name, CommandType.SDN, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.PowerOff.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("SDNResetPasswordTemplate_status", "finish")

ROLLBACK:
	tx.Rollback()
	logger.Point("SDNResetPasswordTemplate_status", "failed")

}
