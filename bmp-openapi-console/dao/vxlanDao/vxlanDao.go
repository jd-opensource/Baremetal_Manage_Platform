package vxlanDao

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Vxlan vxlan
type Vxlan struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`          // ID
	VxlanID    int64     `gorm:"column:vxlan_id;type:bigint(20) unsigned" json:"vxlan_id"`                   // vxlan_id
	Tenant     string    `gorm:"index:i_vxlan_tenant;column:tenant;type:varchar(64);not null" json:"tenant"` // 租户
	Az         string    `gorm:"column:az;type:varchar(32)" json:"az"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`          // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`         // 是否删除(0-未删, 1-已删)
}

func (t *Vxlan) TableName() string {
	return "vxlan"
}

// AddVxlan insert a new Vxlan into database and returns
// last inserted Id on success.
func AddVxlan(logger *log.Logger, m *Vxlan) (id int64, err error) {

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetVxlanById retrieves Vxlan by Id. Returns error if
// Id doesn't exist
func GetVxlanById(logger *log.Logger, id int64) (v *Vxlan, err error) {
	v = &Vxlan{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

//GetMaxVxlanID 没找到max方法，自己写
func GetMaxVxlanId(logger *log.Logger, cloumn string) (n int64, err error) {

	sql := fmt.Sprintf("select max(%s) from vxlan", cloumn)
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&n).Error
	return
}

func GetOrCreateVxlan(logger *log.Logger, query map[string]interface{}) (l *Vxlan, err error) {
	l = &Vxlan{}
	var db = dao.Model(logger, dao.IronicRdb, Vxlan{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err == nil {
		return
	}
	//not found
	maxVlanID, _ := GetMaxVxlanId(logger, "vxlan_id")
	if maxVlanID < 0 {
		return nil, errors.New("get max vxlan_id error!!")
	}
	s := Vxlan{
		// Az:      query["az"].(string),
		// Tenant:  query["tenant"].(string),
		VxlanID: maxVlanID + 1,
	}
	n, err := dao.CreateAndGetId(logger, dao.IronicRdb, s)
	if err != nil {
		return nil, err
	}
	s.ID = n
	return &s, nil
}

func GetOneVxlan(logger *log.Logger, query map[string]interface{}) (l *Vxlan, err error) {
	l = &Vxlan{}
	var db = dao.Model(logger, dao.IronicRdb, Vxlan{})
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

// GetMultiVxlan retrieves all Vxlan matches certain condition. Returns empty list if
// no records exist
func GetMultiVxlan(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Vxlan, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Vxlan{})
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

	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateVxlan updates Vxlan by Id and returns error if
// the record to be updated doesn't exist
func UpdateVxlanById(logger *log.Logger, m *Vxlan) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Vxlan{}).Where("id = ?", m.ID).Updates(m).Error
}

func GetByTenant(logger *log.Logger, tanant, az string) (l *Vxlan, err error) {
	l = &Vxlan{}
	// err = dao.Where(logger, dao.IronicRdb, "tenant = ? and az = ? and is_del = 0", tanant, az).Take(l).Error
	err = dao.Where(logger, dao.IronicRdb, "is_del = 0").Take(l).Error

	if err != nil {
		return nil, err
	}
	return l, nil
}
