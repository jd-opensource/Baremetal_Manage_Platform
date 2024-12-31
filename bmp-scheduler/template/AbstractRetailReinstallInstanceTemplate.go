package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type AbstractRetailReinstallInstanceTemplate struct {
	createSetNetworkCommand func(*gorm.DB, string, string, string, int64, string) (int64, error)
}

func (b AbstractRetailReinstallInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {
	logger.Info("AbstractRetailReinstallInstanceTemplate start,", request_id, instance_id, sn, keep_data, make_partitions, image_format)
	var err error
	var id int64
	tx := dao.GetGormTx(logger)
	tx.Begin()

	if id, err = CreateCommand(tx, CommandAction.DHCPConfigAddHost.Name, CommandType.DRIVER, request_id, instance_id, sn, 0, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.SetPXEBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.Ping.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if keep_data {
		if id, err = CreateCommand(tx, CommandAction.InitRootDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	} else {
		if id, err = CreateCommand(tx, CommandAction.CleanBlockDevice.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
		if id, err = CreateCommand(tx, CommandAction.MakeRaid.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
			goto ROLLBACK
		}
	}
	make_partitions = true //不管是自定义分区还是默认分区，都是true
	if Constants.TAR_IMAGE == image_format {

		if make_partitions {
			if id, err = CreateCommand(tx, CommandAction.MakePartitions.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
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
		if make_partitions {
			if id, err = CreateCommand(tx, CommandAction.Qcow2MakePartitions.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
				goto ROLLBACK
			}
		}
	}
	if id, err = CreateCommand(tx, CommandAction.SetHostname.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.SetKeypairs.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	//TODO 暂时注释
	if id, err = b.createSetNetworkCommand(tx, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetUserCmd.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SetDISKBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.DHCPConfigDelHost.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.PowerCycle.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Infof("AbstractRetailReinstallInstanceTemplate success, instanceId: %s", instance_id)
	return
ROLLBACK:
	tx.Rollback()
	logger.Warnf("AbstractRetailReinstallInstanceTemplate failed, instanceId: %s", instance_id)
}
