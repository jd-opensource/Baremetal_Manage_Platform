package rInstanceSshkeyDao

import (
	"fmt"
	"strings"
	"time"

	//"coding.jd.com/aidc-bmp/bmp-scheduler/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// rInstanceSshkey device_type/image关系
type RInstanceSshkey struct {
	ID          uint64 `gorm:"primaryKey;column:id" json:"-"`          // ID
	InstanceID  string `gorm:"column:instance_id" json:"instanceId"`   // 实例ID
	SSHkeyID    string `gorm:"column:sshkey_id" json:"sshkeyId"`       // sshkeyid
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *RInstanceSshkey) TableName() string {
	return "r_instance_sshkey"
}

// AddrInstanceSshkey insert a new rInstanceSshkey into database and returns
// last inserted Id on success.
func AddrInstanceSshkey(logger *log.Logger, m *RInstanceSshkey) (id int64, err error) {
	//
	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func AddMultirInstanceSshkey(logger *log.Logger, m []*RInstanceSshkey) (id int64, err error) {

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

// GetrInstanceSshkeyById retrieves rInstanceSshkey by Id. Returns error if
// Id doesn't exist
func GetrInstanceSshkeyById(logger *log.Logger, id int64) (v *RInstanceSshkey, err error) {
	v = &RInstanceSshkey{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetAllrInstanceSshkey(logger *log.Logger, query map[string]interface{}) (ml []*RInstanceSshkey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RInstanceSshkey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultirInstanceSshkey retrieves all rInstanceSshkey matches certain condition. Returns empty list if
// no records exist
func GetMultirInstanceSshkey(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*RInstanceSshkey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RInstanceSshkey{})
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

// UpdaterInstanceSshkey updates rInstanceSshkey by Id and returns error if
// the record to be updated doesn't exist
func UpdaterInstanceSshkeyById(logger *log.Logger, m *RInstanceSshkey, deviceTypeId, imageId string) (err error) {
	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, RInstanceSshkey{}).Where("device_type_id = ? and image_id = ? and is_del = ?", deviceTypeId, imageId, 0).Update(m).Error
}

func UpdateMultirInstanceSshkey(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["update_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, RInstanceSshkey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

//func AssociatedImageMultirInstanceSshkey(logger *log.Logger, items []*RInstanceSshkey) (err error) {
//	//paris := []string{}
//	for _, item := range items {
//		//paris = append(paris, fmt.Sprintf("('%s', '%s')", item.DeviceTypeID, item.ImageID))
//		err = dao.Model(logger, dao.IronicRdb, RInstanceSshkey{}).Where("device_type_id = ? and image_id = ? and is_del = 0", item.DeviceTypeID, item.ImageID).Update(item).Error
//	}
//	//fmt.Sprintf(`(device_type_id, image_id) in (%s)`, strings.Join(paris, ","))
//
//	//updates := map[string]interface{}{
//	//	"is_del":      0,
//	//	"update_time": time.Now(),
//	//}
//	return
//}

func GetOnerInstanceSshkey(logger *log.Logger, query map[string]interface{}) (l *RInstanceSshkey, err error) {
	l = &RInstanceSshkey{}
	var db = dao.Model(logger, dao.IronicRdb, RInstanceSshkey{})
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

func GetByDeviceTypeIDAndImageID(logger *log.Logger, deviceTypeID, imageID string) (l *RInstanceSshkey, err error) {
	l = &RInstanceSshkey{}
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and image_id = ? and is_del = ?", deviceTypeID, imageID, 0).Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}
func GetByDeviceTypeID(logger *log.Logger, deviceTypeID string) (l []*RInstanceSshkey, err error) {
	err = dao.Where(logger, dao.IronicRdb, "device_type_id = ? and is_del = ?", deviceTypeID, 0).Find(&l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func GetByDeviceTypeAndImageIdPatition(logger *log.Logger, device_type, image_id string) (ml []*RInstanceSshkey, err error) {

	sql := fmt.Sprintf("select r from r_device_type_image r where (device_type='%s'  or device_type='common') and image_id='%s' and is_del=0", device_type, image_id)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
func QueryDeviceTypeImagesCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, RInstanceSshkey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}
