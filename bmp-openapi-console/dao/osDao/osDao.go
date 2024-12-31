package osDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Os os
type Os struct {
	ID     uint64 `gorm:"primaryKey;column:id" json:"-"` // ID
	OsID   string `gorm:"column:os_id" json:"osId"`      // 操作系统uuid
	OsName string `gorm:"column:os_name" json:"OsName"`  // 操作系统名称
	OsType string `gorm:"column:os_type" json:"osType"`  // 操作系统分类:linux/windows
	//Platform     string `gorm:"column:platform" json:"platform"`         // suse/centos/ubuntu
	Architecture string `gorm:"column:architecture" json:"architecture"` // 架构:x86/x64/i386/
	Bits         int    `gorm:"column:bits" json:"bits"`                 // 指令宽度:64/32位
	OsVersion    string `gorm:"column:os_version" json:"osVersion"`      // 操作系统版本
	SysUser      string `gorm:"column:sys_user" json:"sysUser"`          // 管理员账户
	CreatedBy    string `gorm:"column:created_by" json:"createdBy"`      // 创建者
	UpdatedBy    string `gorm:"column:updated_by" json:"updatedBy"`      // 更新者
	CreatedTime  int    `gorm:"column:created_time" json:"createdTime"`  // 创建时间戳
	UpdatedTime  int    `gorm:"column:updated_time" json:"updatedTime"`  // 更新时间戳
	DeletedTime  int    `gorm:"column:deleted_time" json:"deletedTime"`  // 删除时间戳
	IsDel        int8   `gorm:"column:is_del" json:"isDel"`              // 是否删除0未删除 1已删除
}

func (t *Os) TableName() string {
	return "os"
}

// AddOs insert a new Os into database and returns
// last inserted Id on success.
func AddOs(logger *log.Logger, m *Os) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetOsById retrieves Os by Id. Returns error if
// Id doesn't exist
func GetOsById(logger *log.Logger, id int64) (v *Os, err error) {
	v = &Os{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetOsByUuid retrieves Os by Uuid. Returns error if
// Id doesn't exist
func GetOsByUuid(logger *log.Logger, uuid string) (v *Os, err error) {
	v = &Os{}
	err = dao.Where(logger, dao.IronicRdb, "os_id = ? and is_del = 0", uuid).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllOs retrieves all os matches certain condition. Returns empty list if
// no records exist
func GetAllOs(logger *log.Logger, query map[string]interface{}) (ml []*Os, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Os{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiOs retrieves all Os matches certain condition. Returns empty list if
// no records exist
func GetMultiOs(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Os, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Os{})
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

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateOsById(logger *log.Logger, m *Os) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Os{}).Where("os_id = ?", m.OsID).Updates(m).Error
}
