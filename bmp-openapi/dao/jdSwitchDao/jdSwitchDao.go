package jdSwitchDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// JdSwitch 零售中台网络交换机表
type JdSwitch struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`  // 自增id
	Region     string    `gorm:"unique;column:region;type:varchar(32);not null" json:"region"`       // region
	SwitchIP   string    `gorm:"unique;column:switch_ip;type:varchar(32);not null" json:"switch_ip"` // 交换机ip
	Cidr       string    `gorm:"column:cidr;type:varchar(32)" json:"cidr"`                           // 配置的内网cidr
	Mask       string    `gorm:"column:mask;type:varchar(32)" json:"mask"`                           // 配置的内网掩码
	Gateway    string    `gorm:"column:gateway;type:varchar(32)" json:"gateway"`                     // 配置的网关
	PodID      int       `gorm:"column:pod_id;type:int(11)" json:"pod_id"`                           // pod网络id
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`       // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`       // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(2);not null" json:"is_del"`               // 是否删除(0-未删, 1-已删)
}

func (t *JdSwitch) TableName() string {
	return "jd_switch"
}

// AddJdSwitch insert a new JdSwitch into database and returns
// last inserted Id on success.
func AddJdSwitch(logger *log.Logger, m *JdSwitch) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetJdSwitchById retrieves JdSwitch by Id. Returns error if
// Id doesn't exist
func GetJdSwitchById(logger *log.Logger, id int64) (v *JdSwitch, err error) {
	v = &JdSwitch{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneJdSwitch(logger *log.Logger, query map[string]interface{}) (l *JdSwitch, err error) {
	l = &JdSwitch{}
	var db = dao.Model(logger, dao.IronicRdb, JdSwitch{})
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

func GetSwitchCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, JdSwitch{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return

}

// GetAllJdSwitch retrieves all JdSwitch matches certain condition. Returns empty list if
// no records exist
func GetAllJdSwitch(logger *log.Logger, query map[string]interface{}) (ml []*JdSwitch, err error) {

	var db = dao.Model(logger, dao.IronicRdb, JdSwitch{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiJdSwitch retrieves pages JdSwitch matches certain condition. Returns empty list if
// no records exist
func GetMultiJdSwitch(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*JdSwitch, err error) {

	var db = dao.Model(logger, dao.IronicRdb, JdSwitch{})
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

// UpdateJdSwitch updates JdSwitch by Id and returns error if
// the record to be updated doesn't exist
func UpdateJdSwitchById(logger *log.Logger, m *JdSwitch) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, JdSwitch{}).Where("id = ?", m.ID).Updates(m).Error
}

func GetBySwitchIp(logger *log.Logger, switch_ip string) (v *JdSwitch, err error) {
	v = &JdSwitch{}
	err = dao.Where(logger, dao.IronicRdb, "switch_ip = ? and is_del = 0", switch_ip).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}
