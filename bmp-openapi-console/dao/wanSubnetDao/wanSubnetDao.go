package wanSubnetDao

import (
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// WanSubnet 外网网段
type WanSubnet struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"` // ID
	UUID       string    `gorm:"column:uuid;type:varchar(36);not null" json:"uuid"`                 // 外网网段ID
	Az         string    `gorm:"column:az;type:varchar(32);not null" json:"az"`                     // az
	Cidr       string    `gorm:"column:cidr;type:varchar(32);not null" json:"cidr"`                 // 外网网段
	Gateway    string    `gorm:"column:gateway;type:varchar(32);not null" json:"gateway"`           // 外网网关
	Mask       string    `gorm:"column:mask;type:varchar(32);not null" json:"mask"`                 // 外网掩码
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`      // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`               // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`              // 是否删除(0-未删, 1-已删)
}

func (t *WanSubnet) TableName() string {
	return "wan_subnet"
}

// AddWanSubnet insert a new WanSubnet into database and returns
// last inserted Id on success.
func AddWanSubnet(logger *log.Logger, m *WanSubnet) (id int64, err error) {

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicRdb, m)
}

// GetWanSubnetById retrieves WanSubnet by Id. Returns error if
// Id doesn't exist
func GetWanSubnetById(logger *log.Logger, id int64) (v *WanSubnet, err error) {

	v = &WanSubnet{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllWanSubnet retrieves all WanSubnet matches certain condition. Returns empty list if
// no records exist
func GetMultiWanSubnet(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*WanSubnet, err error) {
	var db = dao.Model(logger, dao.IronicRdb, WanSubnet{})
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
	for _, v := range orderConditions {
		db = db.Order(v)
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateWanSubnet updates WanSubnet by Id and returns error if
// the record to be updated doesn't exist
func UpdateWanSubnetById(logger *log.Logger, m *WanSubnet) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, WanSubnet{}).Where("id = ?", m.ID).Updates(m).Error
}
