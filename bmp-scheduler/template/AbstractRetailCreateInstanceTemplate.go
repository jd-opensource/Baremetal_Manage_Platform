package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type AbstractRetailCreateInstanceTemplate struct {
	createSetNetworkCommand func(*gorm.DB, string, string, string, int64, string) (int64, error)
}

func (b AbstractRetailCreateInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, make_partitions bool, image_format string, task string) {

	logger.Info("AbstractRetailCreateInstanceTemplate start,", request_id, instance_id, sn, make_partitions, image_format)
	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()
	logger.Info("start insert command ", instance_id)
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER,
		request_id, instance_id, sn, 0, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.DHCPConfigAddHost.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetPXEBoot.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.Ping.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.MakeRaid.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.MakeRaid.Name, err.Error())
		goto ROLLBACK
	}
	if Constants.TAR_IMAGE == image_format {
		if make_partitions {
			if id, err = CreateCommand(tx, CommandAction.MakePartitions.Name, CommandType.AGENT,
				request_id, instance_id, sn, id, task); err != nil {
				logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.MakePartitions.Name, err.Error())
				goto ROLLBACK
			}
		}
		if id, err = CreateCommand(tx, CommandAction.WriteImageTar.Name, CommandType.AGENT,
			request_id, instance_id, sn, id, task); err != nil {
			logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.WriteImageTar.Name, err.Error())
			goto ROLLBACK
		}
	} else {
		if id, err = CreateCommand(tx, CommandAction.WriteImage.Name, CommandType.AGENT,
			request_id, instance_id, sn, id, task); err != nil {
			logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.WriteImage.Name, err.Error())
			goto ROLLBACK
		}
		if make_partitions {
			if id, err = CreateCommand(tx, CommandAction.Qcow2MakePartitions.Name, CommandType.AGENT,
				request_id, instance_id, sn, id, task); err != nil {
				logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.Qcow2MakePartitions.Name, err.Error())
				goto ROLLBACK
			}
		}
	}
	if id, err = CreateCommand(tx, CommandAction.SetHostname.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetHostname.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetKeypairs.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetKeypairs.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetPassword.Name, err.Error())
		goto ROLLBACK
	}
	// TODO 暂时注释
	if id, err = b.createSetNetworkCommand(tx, request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", "createSetNetworkCommand", err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetUserCmd.Name, CommandType.AGENT,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetUserCmd.Name, err.Error())
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT,
	// 	request_id, instance_id, sn, id, task); err != nil {
	// 	logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.UploadSystemLog.Name, err.Error())
	// 	goto ROLLBACK
	// }
	//        id, err = CreateCommand(tx, CommandAction.SaveSwitchConfig.Name, CommandType.NETWORK,
	//                request_id, instance_id, sn, id);
	//
	//        id, err = CreateCommand(tx, CommandAction.SaveConfigToFtpServer.Name, CommandType.NETWORK,
	//                request_id, instance_id, sn, id);
	if id, err = CreateCommand(tx, CommandAction.SetDISKBoot.Name, CommandType.DRIVER,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.SetDISKBoot.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.DRIVER,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.DHCPConfigDelHost.Name, err.Error())
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.PowerCycle.Name, CommandType.DRIVER,
		request_id, instance_id, sn, id, task); err != nil {
		logger.Warn("AbstractRetailCreateInstanceTemplate.CreateCommand.error:", CommandAction.PowerCycle.Name, err.Error())
		goto ROLLBACK
	}
	tx.Commit()
	logger.Infof("AbstractRetailCreateInstanceTemplate success, instanceId:%s", instance_id)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("AbstractRetailCreateInstanceTemplate failed, instanceId:%s", instance_id)
}
