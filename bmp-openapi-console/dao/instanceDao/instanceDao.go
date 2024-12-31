package instanceDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
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
	//是否安装了agent，0->未安装，1->已安装
	IsInstalledAgent string `gorm:"column:is_installed_agent" json:"isInstalledAgent"`
}

func (t *Instance) TableName() string {
	return "instance"
}

// AddInstance insert a new Instance into database and returns
// last inserted Id on success.
func AddInstance(logger *log.Logger, m *Instance) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceByName(logger *log.Logger, name string) (v *Instance, err error) {
	v = &Instance{}
	err = dao.Where(logger, dao.IronicRdb, "instance_name = ? and is_del = 0", name).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceById(logger *log.Logger, InstanceId string) (v *Instance, err error) {
	v = &Instance{}
	err = dao.Where(logger, dao.IronicRdb, "instance_id = ? and is_del = 0", InstanceId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetInstancesByIds retrieves Instance by Ids. Returns error if
// Id doesn't exist
func GetInstancesByIds(logger *log.Logger, InstanceIds []string) (v []*Instance, err error) {
	q := map[string]interface{}{
		"instance_id.in": InstanceIds,
		"is_del":         0,
	}
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, q)
	if err != nil {
		return nil, err
	}
	err = db.Find(&v).Error
	return
}

// GetInstanceById retrieves Instance by Id. Returns error if
// Id doesn't exist
func GetInstanceByDeviceId(logger *log.Logger, DeviceId string) (v *Instance, err error) {
	v = &Instance{}
	err = dao.Where(logger, dao.IronicRdb, "device_id = ? and is_del = 0", DeviceId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetInstanceByUuid retrieves Instance by Uuid. Returns error if
// Id doesn't exist
func GetInstanceCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllInstance retrieves all Instance matches certain condition. Returns empty list if
// no records exist
func GetAllInstance(logger *log.Logger, query map[string]interface{}) (ml []*Instance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiInstance retrieves all Instance matches certain condition. Returns empty list if
// no records exist
func GetMultiInstance(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Instance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Instance{})
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

// UpdateInstance updates Instance by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstanceById(logger *log.Logger, m *Instance) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Instance{}).Where("instance_id = ?", m.InstanceID).Updates(m).Error
}

// description为空也有可能
func UpdateInstanceAnywhere(logger *log.Logger, m *Instance) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Instance{}).Save(m).Error
}

func UpdateInstances(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, Instance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
