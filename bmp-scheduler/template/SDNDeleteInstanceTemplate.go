package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type SDNDeleteInstanceTemplate struct{}

func (b SDNDeleteInstanceTemplate) accept(network_type string) bool {
	return NetworkType.VPC == network_type
}

func (b SDNDeleteInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if id, err = CreateCommand(tx, CommandAction.SDNEipUnBinding.Name, CommandType.SDN, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SDNDeleteAliasIP.Name, CommandType.SDN, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SDNSubnetUnBinding.Name, CommandType.SDN, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CleanBlockDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CleanRaid.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.PowerOff.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.SDNClean.Name, CommandType.SDN, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("SDNDeleteInstanceTemplate_status", "finish")

ROLLBACK:
	tx.Rollback()
	logger.Point("SDNDeleteInstanceTemplate_status", "failed")

}
