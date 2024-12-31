package dao

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Instance Instance
type Instance struct {
	ID                 int    `gorm:"primaryKey;column:id" json:"-"`
	InstanceID         string `gorm:"column:instance_id" json:"instanceId"`                   // 实例ID（uuid）
	IDcID              string `gorm:"column:idc_id" json:"idcId"`                             // 机房uuid
	ProjectID          string `gorm:"column:project_id" json:"projectId"`                     // 项目id
	UserID             string `gorm:"column:user_id" json:"userId"`                           // 用户uuid
	InstanceName       string `gorm:"column:instance_name" json:"instanceName"`               // 实例名称
	DeviceID           string `gorm:"column:device_id" json:"deviceId"`                       // 设备uuid
	DeviceTypeID       string `gorm:"column:device_type_id" json:"deviceTypeId"`              // 机型uuid
	Hostname           string `gorm:"column:hostname" json:"hostname"`                        // 主机名
	Status             string `gorm:"column:status" json:"status"`                            // 失败原因
	Reason             string `gorm:"column:reason" json:"reason"`                            // 运行状态
	ImageID            string `gorm:"column:image_id" json:"imageId"`                         // 镜像uuid
	SystemVolumeRaidID string `gorm:"column:system_volume_raid_id" json:"systemVolumeRaidId"` // 系统盘raid
	Locked             string `gorm:"column:locked" json:"locked"`                            // 是否锁定解锁锁定:locked,解锁unlocked
	DataVolumeRaidID   string `gorm:"column:data_volume_raid_id" json:"dataVolumeRaidId"`     // 数据盘raid
	Description        string `gorm:"column:description" json:"description"`                  // 描述
	BootMode           string `gorm:"column:boot_mode" json:"bootMode"`                       // boot类型：UEFI Legacy/BIOS
	CreatedBy          string `gorm:"column:created_by" json:"createdBy"`                     // 创建者
	UpdatedBy          string `gorm:"column:updated_by" json:"updatedBy"`                     // 更新者
	CreatedTime        int    `gorm:"column:created_time" json:"createdTime"`                 // 创建时间戳
	UpdatedTime        int    `gorm:"column:updated_time" json:"updatedTime"`                 // 更新时间戳
	DeletedTime        int    `gorm:"column:deleted_time" json:"deletedTime"`                 // 删除时间戳
	IsDel              int8   `gorm:"column:is_del" json:"isDel"`                             // 是否删除0未删除 1已删除
}

func (t *Instance) TableName() string {
	return "instance"
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceById(logger *log.Logger, InstanceId string) (v *Instance, err error) {
	v = &Instance{}
	err = Where(logger, IronicRdb, "instance_id = ? and is_del = 0", InstanceId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}
