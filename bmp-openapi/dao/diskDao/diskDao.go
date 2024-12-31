package diskDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Disk Disk
type Disk struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`     // ID
	DiskId    string `gorm:"column:disk_id" json:"diskId"`      //DISK UUID
	Name      string `gorm:"column:name" json:"name"`           //DISK name
	DeviceID  string `gorm:"column:device_id" json:"deviceId"`  // 设备uuid
	Enclosure string `gorm:"column:enclosure" json:"enclosure"` // enclosure
	Slot      int    `gorm:"column:slot" json:"slot"`           // 卡槽槽位
	// 磁盘类型：system,data
	DiskType string `gorm:"column:disk_type" json:"diskType"`
	// control/mvme/panfu
	Types        string `gorm:"column:types" json:"types"`
	Size         string `gorm:"column:size" json:"size"`                  // 硬盘大小，不确定精度(非nvme盘)
	SizeUnit     string `gorm:"column:size_unit" json:"sizeUnit"`         // 硬盘大小单位 MB GB TB ,1024进制
	PdType       string `gorm:"column:pd_type" json:"pdType"`             // 硬盘类型
	AdapterID    int    `gorm:"column:adapter_id" json:"adapterId"`       // 适配ID
	MediaType    string `gorm:"column:media_type" json:"mediaType"`       //
	DevicePath   string `gorm:"column:device_path" json:"devicePath"`     //
	Index        int    `gorm:"column:index" json:"index"`                //
	SerialNumber string `gorm:"column:serial_number" json:"serialNumber"` //
	CreatedBy    string `gorm:"column:created_by" json:"createdBy"`       // 创建者
	UpdatedBy    string `gorm:"column:updated_by" json:"updatedBy"`       // 更新者
	CreatedTime  int    `gorm:"column:created_time" json:"createdTime"`   // 创建时间戳
	UpdatedTime  int    `gorm:"column:updated_time" json:"updatedTime"`   // 更新时间戳
	DeletedTime  int    `gorm:"column:deleted_time" json:"deletedTime"`   // 删除时间戳
	IsDel        int8   `gorm:"column:is_del" json:"isDel"`               // 是否删除0未删除 1已删除
}

func (t *Disk) TableName() string {
	return "disk"
}

// AddDisk insert a new Disk into database and returns
// last inserted Id on success.
func AddDisk(logger *log.Logger, m *Disk) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetDiskById retrieves Disk by Id. Returns error if
// Id doesn't exist
func GetDiskById(logger *log.Logger, id int64) (v *Disk, err error) {
	v = &Disk{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetDiskByUuid retrieves Disk by Uuid. Returns error if
// Id doesn't exist
func GetDiskByUuid(logger *log.Logger, uuid string) (v *Disk, err error) {
	v = &Disk{}
	err = dao.Where(logger, dao.IronicRdb, "Disk_id = ? and is_del = 0", uuid).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllDisk retrieves all Disk matches certain condition. Returns empty list if
// no records exist
func GetAllDisk(logger *log.Logger, query map[string]interface{}) (ml []*Disk, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Disk{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiDisk retrieves all Disk matches certain condition. Returns empty list if
// no records exist
func GetMultiDisk(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Disk, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Disk{})
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

// UpdateDisk updates Disk by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiskById(logger *log.Logger, m *Disk) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Disk{}).Where("device_id = ?", m.DeviceID).Updates(m).Error
}

func GetDiskWithDeviceIdVolumeId(logger *log.Logger, deviceId string, volumeId string) (ml []*Disk, err error) {

	sql := `select a.* from disk a inner join r_device_volume_disks b on a.disk_id = b.disk_id where a.is_del = 0 and b.is_del = 0 `

	sql = sql + fmt.Sprintf(" and b.volume_id = '%s' and b.device_id = '%s'", volumeId, deviceId)

	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
