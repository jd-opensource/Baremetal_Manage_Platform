package volumeDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Volume 卷管理表
type Volume struct {
	ID int `gorm:"primaryKey;column:id" json:"-"` // 主键
	// 卷uuid
	VolumeID string `gorm:"column:volume_id" json:"volumeId"`
	// 设备类型uuid
	DeviceTypeID string `gorm:"column:device_type_id" json:"deviceTypeId"`
	// 卷名称
	VolumeName string `gorm:"column:volume_name" json:"volumeName"`
	// 卷类型：系统卷，数据卷
	VolumeType string `gorm:"column:volume_type" json:"volumeType"`
	// 硬盘类型（SSD,HDD）
	DiskType string `gorm:"column:disk_type" json:"diskType"`
	// 接口类型（SATA,SAS,NVME,不限制）
	InterfaceType string `gorm:"column:interface_type" json:"interfaceType"`
	// 单盘大小（最小容量）
	VolumeSize string `gorm:"column:volume_size" json:"volumeSize"`
	// 硬盘单位（GB,TB）
	VolumeUnit string `gorm:"column:volume_unit" json:"volumeUnit"`
	// 硬盘数量（最低块数）
	VolumeAmount int `gorm:"column:volume_amount" json:"volumeAmount"`
	//RaidCan       string `gorm:"column:raid_can" json:"raidCan"`             // RAID配置： (RAID,NO RAID)
	//Raid          string `gorm:"column:raid" json:"raid"`                    // RAID模式：RAID1,RIAD10等
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *Volume) TableName() string {
	return "volume"
}

// AddVolume insert a new Volume into database and returns
// last inserted Id on success.
func AddVolume(logger *log.Logger, m *Volume) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetVolumeById retrieves Volume by Id. Returns error if
// Id doesn't exist
func GetVolumeById(logger *log.Logger, VolumeId string) (v *Volume, err error) {
	v = &Volume{}
	err = dao.Where(logger, dao.IronicRdb, "volume_id = ? and is_del = 0", VolumeId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetVolumeByName(logger *log.Logger, VolumeName string) (v *Volume, err error) {
	v = &Volume{}
	err = dao.Where(logger, dao.IronicRdb, "volume_name = ? and is_del = 0", VolumeName).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetVolumeByUuid retrieves Volume by Uuid. Returns error if
// Id doesn't exist
func GetVolumeCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Volume{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllVolume retrieves all Volume matches certain condition. Returns empty list if
// no records exist
func GetAllVolume(logger *log.Logger, query map[string]interface{}) (ml []*Volume, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Volume{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiVolume retrieves all Volume matches certain condition. Returns empty list if
// no records exist
func GetMultiVolume(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Volume, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Volume{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	orderConditions := []string{}
	for i := 0; i < len(sortby); i++ {
		orderConditions = append(orderConditions, fmt.Sprintf("%s %s", sortby[i], order[i]))
	}
	db = db.Order(strings.Join(orderConditions, ","))
	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateVolume updates Volume by Id and returns error if
// the record to be updated doesn't exist
func UpdateVolumeById(logger *log.Logger, m *Volume) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Volume{}).Where("id = ?", m.ID).Save(m).Error
}

func UpdateVolumeByWhere(logger *log.Logger, query map[string]interface{}, updates *Volume) (err error) {
	var db = dao.Model(logger, dao.IronicWdb, Volume{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

// GetrInstanceSshkeyById retrieves rInstanceSshkey by Id. Returns error if
// Id doesn't exist
func GetSystemVolumeByDeviceTypeId(logger *log.Logger, device_type_id string) (v *Volume, err error) {
	v = &Volume{}
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and is_del = 0 and volume_type = 'system'", device_type_id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetAllVolumeByDeviceTypeId(logger *log.Logger, device_type_id string) (ml []*Volume, err error) {
	ml = []*Volume{}
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and is_del = 0", device_type_id).Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil
}
