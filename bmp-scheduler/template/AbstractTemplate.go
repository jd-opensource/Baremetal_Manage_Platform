package template

import (
	"fmt"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	"github.com/jinzhu/gorm"
)

func CreateCommand(tx *gorm.DB, action, command_type, request_id, instance_id, sn string, parent_command_id int64, task string) (int64, error) {
	entity := &commandDao.Command{
		RequestID:       request_id,
		Sn:              sn,
		Action:          action,
		Type:            command_type,
		Task:            task,
		InstanceID:      instance_id,
		Status:          CommandStatus.WAIT,
		ExecuteCount:    0,
		TimeoutTime:     time.Now(),
		ParentCommandID: parent_command_id,
		CreatedTime:     int(time.Now().Unix()),
		UpdatedTime:     int(time.Now().Unix()),
	}

	err := tx.Create(entity).Error
	if err != nil {
		fmt.Println("插入command error", err.Error(), int64(entity.ID))
	}

	return int64(entity.ID), err
}
