package template

import (
	"strings"

	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/NetworkType"
	log "git.jd.com/cps-golang/log"
)

type BasicReinstallInstanceTemplate struct{}

func (b BasicReinstallInstanceTemplate) accept(network_type, interface_mode string) bool {
	return NetworkType.BASIC == network_type
}

func (b BasicReinstallInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {

	var err error
	var id int64

	instance, err := instanceDao.GetInstanceByUuid(logger, instance_id)
	if err != nil {
		logger.Warn("create instance err,get instance err:", instance_id+err.Error())
		return
	}

	osEntity, err := osLogic.GetOsByInstanceId(logger, instance_id)
	if err != nil {
		logger.Warnf("BasicReinstallInstanceTemplate.GetOsByInstanceId error, instance_id:%s, err:%s", instance_id, err.Error())
	}
	tx := dao.GetGormTx(logger)
	tx.Begin()

	// if id, err = CreateCommand(tx, CommandAction.UnBindingVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, 0);err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if !keep_data {

		if id, err = CreateCommand(tx, CommandAction.CleanBlockDevice.Name, CommandType.AGENT,
			request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
		if instance.SystemVolumeRaidID != "" {
			if id, err = CreateCommand(tx, CommandAction.MakeRaid.Name, CommandType.AGENT,
				request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
		}

	}

	if id, err = CreateCommand(tx, "InitNode", CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK

	}

	if strings.ToLower(image_format) == Constants.TAR_IMAGE {
		if make_partitions {
			if id, err = CreateCommand(tx, CommandAction.MakePartitions.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
			if id, err = CreateCommand(tx, "FormatPartitions", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
			if id, err = CreateCommand(tx, "MountPartitions", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
		}

		if id, err = CreateCommand(tx, CommandAction.WriteImageTar.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	} else {
		if id, err = CreateCommand(tx, CommandAction.WriteImage.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	}

	// if id, err = CreateCommand(tx, CommandAction.SetHostname.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.SetNetwork.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.SetUserCmd.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.SetKeypairs.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	// if id, err = CreateCommand(tx, CommandAction.BindingVSI.Name, CommandType.NETWORK, request_id, instance_id, sn, id); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, "SetCloudinitConf", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, "SetMetaData", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	//网络设置
	if osEntity.OsType != "Windows" {
		if id, err = CreateCommand(tx, "SetNetwork", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	} else {
		if id, err = CreateCommand(tx, "SetNetworkWindows", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	}
	if osEntity.OsType != "Windows" {
		if id, err = CreateCommand(tx, "SetUserData", CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	}

	if id, err = CreateCommand(tx, CommandAction.SetDISKBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.PowerCycle.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("BasicReinstallInstanceTemplate_status", "finish")
	return
ROLLBACK:
	tx.Rollback()
	logger.Point("BasicReinstallInstanceTemplate_status", "failed")
}
