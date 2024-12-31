package template

import (
	"fmt"

	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/raidDao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type BasicCreateInstanceTemplateForWindows struct{}

func (b BasicCreateInstanceTemplateForWindows) accept(network_type, interface_mode string) bool {
	return NetworkType.BASIC == network_type
}

func (b BasicCreateInstanceTemplateForWindows) initCommand(logger *log.Logger, request_id, instance_id, sn string, make_partitions bool, image_format string, task string) {

	logger.Info("BasicCreateInstanceTemplate start", request_id, instance_id, sn)
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

	if err != nil {
		logger.Warn("create instance err,get raid by systemVolumeRaidId err:", instance.SystemVolumeRaidID+err.Error())
		return
	}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if instance.SystemVolumeRaidID != "" {
		raid, err = raidDao.GetRaidByUuid(logger, instance.SystemVolumeRaidID)
		if err != nil {
			logger.Warn("create instance err,get raid by systemVolumeRaidId err:", instance.SystemVolumeRaidID+err.Error())
			return
		}
		if raid.Name != "NORAID" { //如果是NORAID，不执行raid动作
			if id, err = CreateCommand(tx, CommandAction.MakeRaid.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
		}
	}

	if id, err = CreateCommand(tx, "InitNode", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	//if id, err = CreateCommand(tx, CommandAction.MakePartitions.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	//	goto ROLLBACK
	//}
	if id, err = CreateCommand(tx, CommandAction.WriteImage.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	//if id, err = CreateCommand(tx, CommandAction.SetHostname.Name, CommandType.AGENT, request_id, instance_id, sn, id); err != nil {
	//	goto ROLLBACK
	//}
	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	//if id, err = CreateCommand(tx, CommandAction.SetNetwork.Name, CommandType.AGENT, request_id, instance_id, sn, id); err != nil {
	//	goto ROLLBACK
	//}
	//if id, err = CreateCommand(tx, CommandAction.SetUserCmd.Name, CommandType.AGENT, request_id, instance_id, sn, id); err != nil {
	//	goto ROLLBACK
	//}
	//if id, err = CreateCommand(tx, "SetCloudinitConf", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	//	goto ROLLBACK
	//}
	//if id, err = CreateCommand(tx, "SetNocloudUserData", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	//	goto ROLLBACK
	//}
	if id, err = CreateCommand(tx, "SetMetaData", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, "SetNetworkWindows", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	//if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id); err != nil {
	//	goto ROLLBACK
	//}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.CreateVRF.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CreateVRFBalance.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CreateVSIInterface.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CreateVNI.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.CreateVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.AddArpStatic.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.SetBandwidth.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.BindingVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.SaveSwitchConfig.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SetDISKBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.PowerCycle.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	fmt.Println("insert command finished")
	logger.Point("BasicCreateInstanceTemplate_status", "finish")
	return
ROLLBACK:
	tx.Rollback()
	logger.Point("BasicCreateInstanceTemplate_status", "failed")
}
