package instancePartitionDao

import (
	log "git.jd.com/cps-golang/log"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
)

// InstancePartition 实例分区表
type InstancePartition struct {
	ID                  uint64 `gorm:"primaryKey;column:id" json:"-"`                           // ID
	InstanceID          string `gorm:"column:instance_id" json:"instanceId"`                    // 实例ID
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
	InstancePartitionID string `gorm:"column:instance_partition_id" json:"instancePartitionId"` // 实例分区uuID
	Number              int    `gorm:"column:number" json:"number"`                             // 每个实例的分区序号
}

func (t *InstancePartition) TableName() string {
	return "instance_partition"
}

// AddInstancePartition insert a new InstancePartition into database and returns
// last inserted Id on success.
func AddInstancePartition(logger *log.Logger, m *InstancePartition) (id int64, err error) {
	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetInstancePartitionById retrieves InstancePartition by Id. Returns error if
// Id doesn't exist
func GetInstancePartitionById(logger *log.Logger, id int64) (v *InstancePartition, err error) {
	v = &InstancePartition{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllInstancePartition retrieves all InstancePartition matches certain condition. Returns empty list if
// no records exist
func GetAllInstancePartition(logger *log.Logger, query map[string]interface{}) (ml []*InstancePartition, err error) {

	var db = dao.Model(logger, dao.IronicRdb, InstancePartition{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetOneInstancePartition(logger *log.Logger, query map[string]interface{}) (l *InstancePartition, err error) {
	l = &InstancePartition{}
	var db = dao.Model(logger, dao.IronicRdb, InstancePartition{})
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

// UpdateInterface updates Interface by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstancePartitionById(logger *log.Logger, m *InstancePartition) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, InstancePartition{}).Where("id = ?", m.ID).Updates(m).Error
}

// func GetInstancePartitionsBySn(logger *log.Logger, sn string) (ml []*InstancePartition, err error) {
// 	ml = []*InstancePartition{}
// 	err = dao.Where(logger, dao.IronicRdb, "sn = ? and is_del = 0", sn).Find(&ml).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ml, nil

// }

func GetInstancePartitionsByInstanceId(logger *log.Logger, instanceId string) (ml []*InstancePartition, err error) {
	ml = []*InstancePartition{}
	err = dao.Where(logger, dao.IronicRdb, "instance_id = ? and is_del = 0", instanceId).Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil

}
