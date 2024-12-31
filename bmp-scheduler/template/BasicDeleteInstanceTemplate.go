package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/raidDao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type BasicDeleteInstanceTemplate struct{}

func (b BasicDeleteInstanceTemplate) accept(network_type string) bool {
	return NetworkType.BASIC == network_type
}

func (b BasicDeleteInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, task string) {

	logger.Info("BasicDeleteInstanceTemplate start", request_id, instance_id, sn)

	var err error
	var id int64
	var raid *raidDao.Raid
	tx := dao.GetGormTx(logger)
	tx.Begin()

	instance, err := instanceDao.GetInstanceByUuid(logger, instance_id)
	if err != nil {
		logger.Warn("create instance err,get instance err:", instance_id+err.Error())
		return
	}

	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}

	// if id, err = CreateCommand(tx, CommandAction.UnBindingVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, 0); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CleanArpStatic.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CleanBandwidth.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.SaveSwitchConfig.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.CleanBlockDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if instance.SystemVolumeRaidID != "" {
		raid, err = raidDao.GetRaidByUuid(logger, instance.SystemVolumeRaidID)
		if err != nil {
			logger.Warn("create instance err,get raid by systemVolumeRaidId err:", instance.SystemVolumeRaidID+err.Error())
			return
		}
		if raid.Name != "NORAID" { //如果是NORAID，不执行raid动作//TODO,待确认销毁的时候要不要cleanraid
			id, err = CreateCommand(tx, CommandAction.CleanRaid.Name, CommandType.AGENT, request_id, instance_id, sn, id, task)
			if err != nil {
				goto ROLLBACK
			}
		}
	}

	//if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id); err != nil {
	//	goto ROLLBACK
	//}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.PowerOff.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("BasicDeleteInstanceTemplate_status", "finish")
	return
ROLLBACK:
	tx.Rollback()
	logger.Point("BasicDeleteInstanceTemplate_status", "failed")
}
