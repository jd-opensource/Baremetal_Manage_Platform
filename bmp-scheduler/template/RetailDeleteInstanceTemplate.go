package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type RetailDeleteInstanceTemplate struct{}

func (b RetailDeleteInstanceTemplate) accept(network_type string) bool {
	return NetworkType.RETAIL == network_type
}

func (b RetailDeleteInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	logger.Info("RetailDeleteInstanceTemplate start", request_id, instance_id, sn)

	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER,
		request_id, instance_id, sn, 0, task); err != nil {
		logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.DHCPConfigAddHost.Name, err.Error())
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.SetPXEBoot.Name, err.Error())
		goto ROLLBACK
	}
	// if network_type == NetworkType.RETAIL {
	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.Ping.Name, err.Error())
		goto ROLLBACK
	}
	// }
	if id, err = CreateCommand(tx, CommandAction.CleanBlockDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.CleanBlockDevice.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CleanRaid.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.CleanRaid.Name, err.Error())
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	logger.Warn("RetailDeleteInstanceTemplate.CreateCommand.error:", CommandAction.UploadSystemLog.Name, err.Error())
	// 	goto ROLLBACK
	// }

	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.DRIVER,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.DHCPConfigDelHost.Name, err.Error())
		goto ROLLBACK
	}

	if _, err = CreateCommand(tx, CommandAction.PowerOff.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.PowerOff.Name, err.Error())
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("RetailDeleteInstanceTemplate success, instance_id: %s", instance_id)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("RetailDeleteInstanceTemplate failed, instance_id: %s", instance_id)
}
