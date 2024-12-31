package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type CollectHardwareInfoTemplate struct{}

func (c CollectHardwareInfoTemplate) InitCommand(logger *log.Logger, request_id, sn, network_type string, allowOverride bool, task string) {

	logger.Info("CollectHardwareInfoTemplate start, ", request_id, sn, network_type)
	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	// 前几个任务是启动agent，后几个任务是收集信息。

	//新增DHCPConfigAddSubnet，CheckInitConfig
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddSubnet.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}
	//采集时不执行CheckInitConfig 20240530 minping
	// if id, err = CreateCommand(tx, CommandAction.CheckInitConfig.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }

	if network_type == NetworkType.RETAIL {
		if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
			goto ROLLBACK
		}
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}

	if network_type == NetworkType.RETAIL {
		if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, "", sn, id, task); err != nil {
			goto ROLLBACK
		}
	}
	// 清除raid
	if allowOverride {
		if id, err = CreateCommand(tx, CommandAction.CleanRaid.Name, CommandType.AGENT, request_id, "", sn, id, task); err != nil {
			goto ROLLBACK
		}
	}

	if id, err = CreateCommand(tx, CommandAction.CollectDiskLocations.Name, CommandType.AGENT, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.CollectHardwareInfo.Name, CommandType.AGENT, request_id, "", sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, "", sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if network_type == NetworkType.RETAIL {
		if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.DRIVER, request_id, "", sn, id, task); err != nil {
			goto ROLLBACK
		}
	}
	CreateCommand(tx, CommandAction.PowerOff.Name, CommandType.DRIVER, request_id, "", sn, id, task)
	tx.Commit()
	logger.Infof("CollectHardwareInfoTemplate success, sn: %s", sn)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("CollectHardwareInfoTemplate failed, sn: %s, err:%s", sn, err.Error())
}
