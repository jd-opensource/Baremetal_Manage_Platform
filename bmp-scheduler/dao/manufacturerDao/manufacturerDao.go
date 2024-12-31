package manufacturerDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// Manufacturer 产品型号信息
type Manufacturer struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`  // ID
	DeviceType   string    `gorm:"column:device_type;type:varchar(100);not null" json:"device_type"`   // 设备类型
	Manufacturer string    `gorm:"column:manufacturer;type:varchar(64);not null" json:"manufacturer"`  // 厂商
	ProductName  string    `gorm:"column:product_name;type:varchar(128);not null" json:"product_name"` // 产品型号名称
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`       // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                // 更新时间
	IsDel        int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`               // 是否删除：1删除，0未删除
}

func (t *Manufacturer) TableName() string {
	return "manufacturer"
}

// AddManufacturer insert a new Manufacturer into database and returns
// last inserted Id on success.
func AddManufacturer(logger *log.Logger, m *Manufacturer) (id int64, err error) {

	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetManufacturerById retrieves Manufacturer by Id. Returns error if
// Id doesn't exist
func GetManufacturerById(logger *log.Logger, id int64) (v *Manufacturer, err error) {
	v = &Manufacturer{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneManufacturer(logger *log.Logger, query map[string]interface{}) (v *Manufacturer, err error) {
	v = &Manufacturer{}
	var db = dao.Model(logger, dao.IronicRdb, Manufacturer{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil

}

// GetAllManufacturer retrieves all Manufacturer matches certain condition. Returns empty list if
// no records exist
func GetAllManufacturer(logger *log.Logger, query map[string]interface{}) (ml []*Manufacturer, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Manufacturer{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetMultiManufacturer(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Manufacturer, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Manufacturer{})
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

// UpdateManufacturer updates Manufacturer by Id and returns error if
// the record to be updated doesn't exist
func UpdateManufacturerById(logger *log.Logger, m *Manufacturer) (err error) {
	m.UpdateTime = time.Now()

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Manufacturer{}).Where("id = ?", m.ID).Updates(m).Error
}
