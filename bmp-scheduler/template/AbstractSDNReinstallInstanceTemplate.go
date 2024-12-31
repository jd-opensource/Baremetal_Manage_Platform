package template

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandAction"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandType"
	log "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

type AbstractSDNReinstallInstanceTemplate struct {
	createSetNetworkCommand func(*gorm.DB, string, string, string, int64, string) (int64, error)
}

func (b AbstractSDNReinstallInstanceTemplate) initCommand(logger *log.Logger, request_id, instance_id, sn string, keep_data, make_partitions bool, image_format string, task string) {

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
	if id, err = CreateCommand(tx, CommandAction.SetPassword.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = b.createSetNetworkCommand(tx, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetUserCmd.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetKeypairs.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	// if id, err = CreateCommand(tx, CommandAction.UploadSystemLog.Name, CommandType.AGENT, request_id, instance_id, sn, id, task); err != nil {
	// 	goto ROLLBACK
	// }
	if id, err = CreateCommand(tx, CommandAction.SDNSubnetBinding.Name, CommandType.SDN, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if id, err = CreateCommand(tx, CommandAction.SetDISKBoot.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}
	if _, err = CreateCommand(tx, CommandAction.PowerCycle.Name, CommandType.DRIVER, request_id, instance_id, sn, id, task); err != nil {
		goto ROLLBACK
	}

	tx.Commit()
	logger.Point("AbstractSDNReinstallInstanceTemplate_status", "finish")

ROLLBACK:
	tx.Rollback()
	logger.Point("AbstractSDNReinstallInstanceTemplate_status", "failed")
}
