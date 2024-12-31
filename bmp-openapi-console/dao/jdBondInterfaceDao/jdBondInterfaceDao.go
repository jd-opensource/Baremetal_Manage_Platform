package jdBondInterfaceDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// JdBondInterface 零售中台网络端口Bond表
type JdBondInterface struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`            // 自增id
	GroupID    int       `gorm:"column:group_id;type:int(11)" json:"group_id"`                                 // 端口聚合组ID
	Sn         string    `gorm:"index:i_sn;column:sn;type:varchar(64);not null" json:"sn"`                     // 设备sn
	SwitchIP   string    `gorm:"index:i_switchIp;column:switch_ip;type:varchar(32);not null" json:"switch_ip"` // 交换机ip
	PrivateIP  string    `gorm:"column:private_ip;type:varchar(30)" json:"private_ip"`                         // 内网IP
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`                 // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`                 // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(2);not null" json:"is_del"`                         // 是否删除(0-未删, 1-已删)
}

func (t *JdBondInterface) TableName() string {
	return "jd_bond_interface"
}

// AddJdBondInterface insert a new JdBondInterface into database and returns
// last inserted Id on success.
func AddJdBondInterface(logger *log.Logger, m *JdBondInterface) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetJdBondInterfaceById retrieves JdBondInterface by Id. Returns error if
// Id doesn't exist
func GetJdBondInterfaceById(logger *log.Logger, id int64) (v *JdBondInterface, err error) {
	v = &JdBondInterface{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetAllJdBondInterface(logger *log.Logger, query map[string]interface{}) (ml []*JdBondInterface, err error) {

	var db = dao.Model(logger, dao.IronicRdb, JdBondInterface{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiJdBondInterface retrieves all JdBondInterface matches certain condition. Returns empty list if
// no records exist
func GetMultiJdBondInterface(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*JdBondInterface, err error) {

	var db = dao.Model(logger, dao.IronicRdb, JdBondInterface{})
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

// UpdateJdBondInterface updates JdBondInterface by Id and returns error if
// the record to be updated doesn't exist
func UpdateJdBondInterfaceById(logger *log.Logger, m *JdBondInterface) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, JdBondInterface{}).Where(logger, dao.IronicRdb, "id = ?", m.ID).Updates(m).Error //update 有问题，参考下面这个方法
}
func DeleteJdBondInterfaceBySn(logger *log.Logger, m *JdBondInterface) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, JdBondInterface{}).Where("sn = ?", m.Sn).Updates(m).Error
}
