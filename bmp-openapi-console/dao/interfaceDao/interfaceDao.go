package interfaceDao

import (
	"fmt"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
)

type Interface struct {
	ID            int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`   // ID
	InterfaceName string    `gorm:"column:interface_name;type:varchar(16)" json:"interface_name"`        // 网卡名称
	InterfaceType string    `gorm:"column:interface_type;type:varchar(16)" json:"interface_type"`        // 网卡类型：lan/wan
	Sn            string    `gorm:"index:i_interface_sn;column:sn;type:varchar(64);not null" json:"sn"`  // 设备SN
	Mac           string    `gorm:"column:mac;type:varchar(32);not null" json:"mac"`                     // MAC
	SwitchIP      string    `gorm:"column:switch_ip;type:varchar(32);not null" json:"switch_ip"`         // 内网交换机IP
	SwitchPort    string    `gorm:"column:switch_port;type:varchar(32);default:null" json:"switch_port"` // 内网交换机Port
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`        // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                 // 更新时间
	IsDel         int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                // 是否删除(0-未删。1-已删)
}

func (t *Interface) TableName() string {
	return "interface"
}

// AddInterface insert a new Interface into database and returns
// last inserted Id on success.
func AddInterface(logger *log.Logger, m *Interface) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetInterfaceById retrieves Interface by Id. Returns error if
// Id doesn't exist
func GetInterfaceById(logger *log.Logger, id int64) (v *Interface, err error) {
	v = &Interface{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetAllInterface(logger *log.Logger, query map[string]interface{}) (ml []*Interface, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Interface{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetOneInterface(logger *log.Logger, query map[string]interface{}) (l *Interface, err error) {
	l = &Interface{}
	var db = dao.Model(logger, dao.IronicRdb, Interface{})
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

// GetAllInterface retrieves all Interface matches certain condition. Returns empty list if
// no records exist
func GetMultiInterface(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Interface, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Interface{})
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

// UpdateInterface updates Interface by Id and returns error if
// the record to be updated doesn't exist
func UpdateInterfaceById(logger *log.Logger, m *Interface) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Interface{}).Where(logger, dao.IronicRdb, "id = ?", m.ID).Updates(m).Error
}
func DeleteInterfaceBySn(logger *log.Logger, m *Interface) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Interface{}).Where("sn = ?", m.Sn).Delete(m).Error
}
func UpdateMultiInterfaces(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["update_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, Interface{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

type DInterface struct {
	Sn            string    `json:"sn"`
	InterfaceName string    `json:"interface_name"`
	InterfaceType string    `json:"interface_type"`
	Mac           string    `json:"mac"`
	SwitchIp      string    `json:"switch_ip"`
	SwitchPort    string    `json:"switch_port"`
	PodName       string    `json:"pod_name"`
	PodRoom       string    `json:"pod_room"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

func GetBySn(logger *log.Logger, sn string) (ml []*Interface, err error) {

	ml = []*Interface{}
	err = dao.Where(logger, dao.IronicRdb, "sn = ? and is_del = 0", sn).Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil

}

func GetBySns(logger *log.Logger, sns []string) (ml []*DInterface, err error) {

	sql := `select i.sn,i.interface_name,i.interface_type,i.mac,i.switch_ip,i.switch_port,i.create_time,i.update_time,n.name as pod_name,n.room as pod_room FROM interface i LEFT JOIN jd_switch j ON i.switch_ip = j.switch_ip LEFT JOIN net_pod n ON j.pod_id=n.id where 1=1`
	for idx, v := range sns {
		sns[idx] = fmt.Sprintf("'%s'", v)
	}
	sql = sql + fmt.Sprintf(" and i.sn in (%s) and i.is_del=0 and j.is_del=0 and n.is_del=0", strings.Join(sns, ","))
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}
