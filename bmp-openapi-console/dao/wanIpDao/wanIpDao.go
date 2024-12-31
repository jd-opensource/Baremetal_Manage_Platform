package wanIpDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// WanIP 外网ip
type WanIp struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`         // ID
	Az          string    `gorm:"unique_index:i_wan_ip_az_ip;column:az;type:varchar(32);not null" json:"az"` // az
	WanSubnetID string    `gorm:"column:wan_subnet_id;type:varchar(36);not null" json:"wan_subnet_id"`       // 外网网段ID
	IP          string    `gorm:"unique_index:i_wan_ip_az_ip;column:ip;type:varchar(32);not null" json:"ip"` // 外网ip
	Status      int8      `gorm:"column:status;type:tinyint(4);not null" json:"status"`                      // 是否已使用(0-未使用, 1-已使用)
	LineType    string    `gorm:"column:line_type;type:varchar(20)" json:"line_type"`                        // 链路类型：bgp、dynamic_bgp
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`              // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                       // 更新时间
	IsDel       int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                      // 是否删除(0-未删, 1-已删)
}

func (t *WanIp) TableName() string {
	return "wan_ip"
}

// AddWanIp insert a new WanIp into database and returns
// last inserted Id on success.
func AddWanIp(logger *log.Logger, m *WanIp) (id int64, err error) {

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// AddMultiWanIp insert a new WanIp into database and returns
// last inserted Id on success.
func AddMultiWanIp(logger *log.Logger, m []*WanIp) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, device := range m {
		device.CreateTime = time.Now()
		device.UpdateTime = time.Now()
		if err := tx.Create(device).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetWanIpById retrieves WanIp by Id. Returns error if
// Id doesn't exist
func GetWanIpById(logger *log.Logger, id int64) (v *WanIp, err error) {

	v = &WanIp{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetAllWanIp(logger *log.Logger, query map[string]interface{}) (ml []*WanIp, err error) {

	var db = dao.Model(logger, dao.IronicRdb, WanIp{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiWanIp retrieves all WanIp matches certain condition. Returns empty list if
// no records exist
func GetMultiWanIp(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*WanIp, err error) {

	var db = dao.Model(logger, dao.IronicRdb, WanIp{})
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

// UpdateWanIp updates WanIp by Id and returns error if
// the record to be updated doesn't exist
func UpdateWanIpById(logger *log.Logger, m *WanIp) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, WanIp{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateMultiWanIps(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {
	updates["update_time"] = time.Now()
	var db = dao.Model(logger, dao.IronicWdb, WanIp{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func QueryUnusedByAz(logger *log.Logger, az string, offset, limit int64) (ml []*WanIp, err error) {

	// err = dao.Where(logger, dao.IronicRdb, "az = ? and status = 0 and is_del = 0", az).Find(&ml).Error
	err = dao.Where(logger, dao.IronicRdb, "status = 0 and is_del = 0").Find(&ml).Error
	return

}

func QueryByAzAndIP(logger *log.Logger, az string, ips []string) (ml []*WanIp, err error) {
	query := map[string]interface{}{
		"ip.in": ips,
		// "az":     az,
		"is_del": "0",
	}
	var db = dao.Model(logger, dao.IronicRdb, WanIp{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}

	err = db.Find(&ml).Error
	return
}
