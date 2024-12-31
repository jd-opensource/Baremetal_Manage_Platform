package rDeviceTypeImagePartitionDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// RDeviceTypeImagePartition device_type/image
type RDeviceTypeImagePartition struct {
	ID                  uint64 `gorm:"primaryKey;column:id" json:"-"`                           // ID
	ImageID             string `gorm:"column:image_id" json:"imageId"`                          // 镜像ID
	DeviceTypeID        string `gorm:"column:device_type_id" json:"deviceTypeId"`               // 设备类型
	BootMode            string `gorm:"column:boot_mode" json:"bootMode"`                        // boot类型：bios、uefi
	PartitionType       string `gorm:"column:partition_type" json:"partitionType"`              // 分区类型：root、boot、system、data
	PartitionSize       int    `gorm:"column:partition_size" json:"partitionSize"`              // 分区大小，单位MB
	PartitionFsType     string `gorm:"column:partition_fs_type" json:"partitionFsType"`         // 文件系统类型：xfs
	PartitionMountPoint string `gorm:"column:partition_mount_point" json:"partitionMountPoint"` // 分区目录
	PartitionLabel      string `gorm:"column:partition_label" json:"partitionLabel"`            // 分区标签：l_分区目录
	SystemDiskLabel     string `gorm:"column:system_disk_label" json:"systemDiskLabel"`         // 系统盘分区格式：gpt、msdos（做完RAID系统盘大于4T必用gpt）
	DataDiskLabel       string `gorm:"column:data_disk_label" json:"dataDiskLabel"`             // 数据盘分区格式：gpt、msdos
	CreatedBy           string `gorm:"column:created_by" json:"createdBy"`                      // 创建者
	UpdatedBy           string `gorm:"column:updated_by" json:"updatedBy"`                      // 更新者
	CreatedTime         int    `gorm:"column:created_time" json:"createdTime"`                  // 创建时间戳
	UpdatedTime         int    `gorm:"column:updated_time" json:"updatedTime"`                  // 更新时间戳
	DeletedTime         int    `gorm:"column:deleted_time" json:"deletedTime"`                  // 删除时间戳
	IsDel               int8   `gorm:"column:is_del" json:"isDel"`                              // 是否删除(0-未删, 1-已删)
}

func (t *RDeviceTypeImagePartition) TableName() string {
	return "r_device_type_image_partition"
}

// AddRDeviceTypeImage insert a new RDeviceTypeImage into database and returns
// last inserted Id on success.
func AddRDeviceTypeImagePartition(logger *log.Logger, m *RDeviceTypeImagePartition) (id int64, err error) {
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func AddMultiRDeviceTypeImagePartition(logger *log.Logger, m []*RDeviceTypeImagePartition) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, r := range m {
		if err := tx.Create(r).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetRDeviceTypeImageById retrieves RDeviceTypeImage by Id. Returns error if
// Id doesn't exist
func GetRDeviceTypeImagePartitionById(logger *log.Logger, id int64) (v *RDeviceTypeImagePartition, err error) {
	v = &RDeviceTypeImagePartition{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetAllRDeviceTypeImagePartition(logger *log.Logger, query map[string]interface{}) (ml []*RDeviceTypeImagePartition, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImagePartition{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiRDeviceTypeImage retrieves all RDeviceTypeImage matches certain condition. Returns empty list if
// no records exist
func GetMultiRDeviceTypeImagePartition(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*RDeviceTypeImagePartition, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImagePartition{})
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

// UpdateRDeviceTypeImage updates RDeviceTypeImage by Id and returns error if
// the record to be updated doesn't exist
func UpdateRDeviceTypeImagePartitionById(logger *log.Logger, m *RDeviceTypeImagePartition) (err error) {
	return dao.Model(logger, dao.IronicWdb, RDeviceTypeImagePartition{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateMultiRDeviceTypeImagePartition(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	//updates["updated_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, RDeviceTypeImagePartition{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func GetOneRDeviceTypeImagePartition(logger *log.Logger, query map[string]interface{}) (l *RDeviceTypeImagePartition, err error) {
	l = &RDeviceTypeImagePartition{}
	var db = dao.Model(logger, dao.IronicRdb, RDeviceTypeImagePartition{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func GetByDeviceTypeAndImageId(logger *log.Logger, deviceTypeId, imageId string) (ml []*RDeviceTypeImagePartition, err error) {
	sql := fmt.Sprintf("select * from r_device_type_image_partition r where device_type_id='%s' and image_id='%s' and is_del=0", deviceTypeId, imageId) //or device_type='common'
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
