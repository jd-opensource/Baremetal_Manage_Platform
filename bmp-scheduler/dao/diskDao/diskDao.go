package diskDao

import (
	"fmt"

	log "git.jd.com/cps-golang/log"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
)

type Disk struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`     // ID
	DiskId    string `gorm:"column:disk_id" json:"diskId"`      //DISK UUID
	Name      string `gorm:"column:name" json:"name"`           //DISK name
	DeviceID  string `gorm:"column:device_id" json:"deviceId"`  // 设备uuid
	Enclosure string `gorm:"column:enclosure" json:"enclosure"` // enclosure
	Slot      int    `gorm:"column:slot" json:"slot"`           // 卡槽槽位
	DiskType  string `gorm:"column:disk_type" json:"diskType"`  // 磁盘类型：system,data
	Size      string `gorm:"column:size" json:"size"`           // 硬盘大小，不确定精度(非nvme盘)
	SizeUnit  string `gorm:"column:size_unit" json:"sizeUnit"`  // 硬盘大小单位 MB GB TB ,1024进制
	//接口类型
	PdType    string `gorm:"column:pd_type" json:"pdType"`       // 硬盘类型
	AdapterID int    `gorm:"column:adapter_id" json:"adapterId"` // 适配ID
	//磁盘类型
	MediaType  string `gorm:"column:media_type" json:"mediaType"`
	DevicePath string `gorm:"column:device_path" json:"devicePath"`
	Serial     string `gorm:"column:serial_number" json:"serial"`

	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
	Types       string `gorm:"column:types" json:"types"`              // controller/nvme/panfu
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

// GetAllDisk retrieves all Disk matches certain condition. Returns empty list if
// no records exist
func GetAllDisk(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string) (ml []*Disk, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Disk{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateDisk updates Disk by Id and returns error if
// the record to be updated doesn't exist
func UpdateDiskById(logger *log.Logger, m *Disk) (err error) {
	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Disk{}).Where("id = ?", m.ID).Updates(m).Error
}

// //TODO 有时间试试gorm v2支持的批量插入
func AddMultiDisk(logger *log.Logger, m []*Disk) (id int64, err error) {

	// tx := dao.GetGormTx(logger)
	// tx.Begin()
	// for _, device := range m {
	// 	//device.CreateTime = time.Now()
	// 	//device.UpdateTime = time.Now()
	// 	if err := tx.Create(device).Error; err != nil {
	// 		tx.Rollback()
	// 		return 0, err
	// 	}
	// }
	// tx.Commit()
	// return int64(len(m)), nil
	for _, item := range m {
		_, err := AddDisk(logger, item)
		if err != nil {
			logger.Warnf("AddMultiDisk error, device_id:%s", item.DeviceID)
		}
	}
	return 0, nil

}

func UpdateMultiDisks(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, Disk{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func GetVolumeDisks(logger *log.Logger, deviceId, volumeId string) (ml []*Disk, err error) {
	ml = []*Disk{}
	sql := `select a.* from disk a inner join r_device_volume_disks b on a.disk_id = b.disk_id and a.is_del = 0 and b.is_del = 0 `

	sql = sql + fmt.Sprintf(" and b.device_id = '%s' and b.volume_id = '%s'", deviceId, volumeId)

	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil
}

func GetDisks(logger *log.Logger, query map[string]interface{}) (ml []*Disk, err error) {
	ml = []*Disk{}
	var db = dao.Model(logger, dao.IronicRdb, Disk{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return
}
