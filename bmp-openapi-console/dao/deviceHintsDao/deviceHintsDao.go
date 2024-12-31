package deviceHintsDao

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// DeviceHints 根设备所在磁盘
type DeviceHints struct {
	ID int64 `gorm:"unique;column:id;type:bigint(20) unsigned;not null" json:"id"` // 设备ID编号

	Sn string `gorm:"column:sn;type:varchar(64);not null" json:"sn"` // 设备SN
	//volume UUID
	VolumeId           string    `gorm:"column:volume_id;type:varchar(128)" json:"volume_id"`
	VolumeType         string    `gorm:"column:volume_type" json:"volumeType"`                                     // 卷类型：系统卷，数据卷
	Size               int64     `gorm:"column:size;type:bigint(20);not null" json:"size"`                         // 以GiB为单位的设备尺寸
	Rotational         int8      `gorm:"column:rotational;type:tinyint(4);not null" json:"rotational"`             // 转动，可以区分HDDs(旋转-磁盘)和SSDs(非旋转-固态硬盘)
	Wwn                string    `gorm:"column:wwn;type:varchar(128)" json:"wwn"`                                  // 唯一的存储标识符
	Name               string    `gorm:"column:name;type:varchar(64);not null" json:"name"`                        // 设备名称，例如/dev/md0
	WwnVendorExtension string    `gorm:"column:wwn_vendor_extension;type:varchar(64)" json:"wwn_vendor_extension"` // 唯一的供应商存储标识符
	WwnWithExtension   string    `gorm:"column:wwn_with_extension;type:varchar(64)" json:"wwn_with_extension"`     // 带有供应商附加扩展的唯一的存储标识符
	Model              string    `gorm:"column:model;type:varchar(64)" json:"model"`                               // 设备标识符
	Serial             string    `gorm:"column:serial;type:varchar(64)" json:"serial"`                             // 磁盘序列号
	Hctl               string    `gorm:"column:hctl;type:varchar(64)" json:"hctl"`                                 // SCSI地址（主机名、channel通道、Target和Lun）
	ByPath             string    `gorm:"column:by_path;type:varchar(128)" json:"by_path"`                          // 磁盘by_path路径
	Vendor             string    `gorm:"column:vendor;type:varchar(64)" json:"vendor"`                             // 设备供应商
	CreateTime         time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`             // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                      // 更新时间
	IsDel              int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                     // 是否删除(0-未删, 1-已删)
}

func (t *DeviceHints) TableName() string {
	return "device_hints"
}

// AddDeviceHints insert a new DeviceHints into database and returns
// last inserted Id on success.
func AddDeviceHints(logger *log.Logger, m *DeviceHints) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetDeviceHintsById retrieves DeviceHints by Id. Returns error if
// Id doesn't exist
func GetDeviceHintsById(logger *log.Logger, id int) (v *DeviceHints, err error) {
	v = &DeviceHints{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetSystemVolumnDeviceHintsBySn(logger *log.Logger, sn string) (v *DeviceHints, err error) {
	v = &DeviceHints{}
	err = dao.Where(logger, dao.IronicRdb, "sn = ? and volume_type = 'system' and is_del = 0", sn).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetDeviceHintsBySnAndVolumeId(logger *log.Logger, sn, volume_id string) (v *DeviceHints, err error) {
	v = &DeviceHints{}
	err = dao.Where(logger, dao.IronicRdb, "sn = ? and volume_id = ? and is_del = 0", sn, volume_id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllDeviceHints retrieves all DeviceHints matches certain condition. Returns empty list if
// no records exist
func GetAllDeviceHints(logger *log.Logger, query map[string]interface{}) (ml []*DeviceHints, err error) {

	var db = dao.Model(logger, dao.IronicRdb, DeviceHints{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateDeviceHintsById updates DeviceHints by Id and returns error if
// the record to be updated doesn't exist
func UpdateDeviceHintsById(logger *log.Logger, m *DeviceHints) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, DeviceHints{}).Where("id = ?", m.ID).Updates(m).Error
}

func DeleteDeviceHintsBySn(logger *log.Logger, sn string) (err error) {

	var db = dao.Model(logger, dao.IronicWdb, DeviceHints{})
	query := map[string]interface{}{
		"sn":     sn,
		"is_del": 0,
	}
	updates := map[string]interface{}{
		"is_del":      1,
		"update_time": time.Now(),
	}

	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}
