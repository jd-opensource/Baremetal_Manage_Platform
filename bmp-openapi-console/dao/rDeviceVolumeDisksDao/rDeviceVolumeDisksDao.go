package rDeviceVolumeDisksDao

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// RDeviceVolumeDisks 设备-卷-磁盘 关系表
type RDeviceVolumeDisks struct {
	ID          uint32 `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                  // 主键
	DeviceID    string `gorm:"index:i_device_id;column:device_id;type:varchar(255);not null" json:"device_id"` // 设备uuid
	VolumeID    string `gorm:"index:i_volume_id;column:volume_id;type:varchar(255);not null" json:"volume_id"` // volume uuid
	DiskID      string `gorm:"index:i_disk_id;column:disk_id;type:varchar(255);not null" json:"disk_id"`       // disk uuid
	CreatedBy   string `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`                 // 创建者
	UpdatedBy   string `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`                 // 更新者
	CreatedTime int    `gorm:"column:created_time;type:int(255);not null" json:"created_time"`                 // 创建时间
	UpdatedTime int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`                  // 更新时间
	DeletedTime int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`                  // 删除时间
	IsDel       int8   `gorm:"index:i_is_del;column:is_del;type:tinyint(4);not null" json:"is_del"`            // 是否删除0未删除 1已删除
}

func (t *RDeviceVolumeDisks) TableName() string {
	return "r_device_volume_disks"
}

// GetAllRDeviceTypeRaid retrieves all raid matches certain condition. Returns empty list if
// no records exist
func GetAllRDeviceVolumeDisks(logger *log.Logger, query map[string]interface{}) (ml []*RDeviceVolumeDisks, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RDeviceVolumeDisks{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetRDeviceTypeRaidById retrieves RDeviceTypeRaid by Id. Returns error if
// Id doesn't exist
func GetRDeviceVolumeDiskByVolumeIdDiskId(logger *log.Logger, volumeId, diskId string) (v *RDeviceVolumeDisks, err error) {
	v = &RDeviceVolumeDisks{}
	err = dao.Where(logger, dao.IronicRdb, "volume_id = ? and disk_id = ? and is_del = 0", volumeId, diskId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func UpdateMultiRDeviceVolumeDisk(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = int(time.Now().Unix())
	var db = dao.Model(logger, dao.IronicWdb, RDeviceVolumeDisks{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func AddMultiRDeviceVolumeDisk(logger *log.Logger, m []*RDeviceVolumeDisks) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, r := range m {
		//r.CreateTime = time.Now()
		//r.UpdateTime = time.Now()
		if err := tx.Create(r).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}
