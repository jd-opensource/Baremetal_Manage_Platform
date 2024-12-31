package keypairDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// Keypair 密钥对
type Keypair struct {
	ID          int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"` // ID
	UUID        string    `gorm:"unique;column:uuid;type:varchar(36);not null" json:"uuid"`          // 密钥对ID
	Tenant      string    `gorm:"column:tenant;type:varchar(64);not null" json:"tenant"`             // 租户
	Region      string    `gorm:"column:region;type:varchar(32);not null" json:"region"`             // region
	Az          string    `gorm:"column:az;type:varchar(32)" json:"az"`                              // az
	Name        string    `gorm:"column:name;type:varchar(64);not null" json:"name"`                 // 密钥对名称
	PublicKey   string    `gorm:"column:public_key;type:varchar(2000);not null" json:"public_key"`   // 公钥内容
	FingerPrint string    `gorm:"column:finger_print;type:varchar(64);not null" json:"finger_print"` // 指纹
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`      // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`               // 更新时间
	IsDel       int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`              // 是否删除(0-未删, 1-已删)
}

func (t *Keypair) TableName() string {
	return "keypair"
}

// AddKeypair insert a new Keypair into database and returns
// last inserted Id on success.
func AddKeypair(logger *log.Logger, m *Keypair) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetKeypairById retrieves Keypair by Id. Returns error if
// Id doesn't exist
func GetKeypairById(logger *log.Logger, id int64) (v *Keypair, err error) {
	v = &Keypair{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneKeypair(logger *log.Logger, query map[string]interface{}) (l *Keypair, err error) {
	l = &Keypair{}
	var db = dao.Model(logger, dao.IronicRdb, Keypair{})
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

// GetAllKeypair retrieves all Keypair matches certain condition. Returns empty list if
// no records exist
func GetAllKeypair(logger *log.Logger, query map[string]interface{}) (ml []*Keypair, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Keypair{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiKeypair retrieves all Keypair matches certain condition. Returns empty list if
// no records exist
func GetMultiKeypair(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Keypair, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Keypair{})
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

func GetKeypairCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Keypair{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return

}

// UpdateKeypair updates Keypair by Id and returns error if
// the record to be updated doesn't exist
func UpdateKeypairById(logger *log.Logger, m *Keypair) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Keypair{}).Where("id = ?", m.ID).Updates(m).Error
}

func GetByUuidAndTenant(logger *log.Logger, uuid, tenant string) (v *Keypair, err error) {
	v = &Keypair{}
	err = dao.Where(logger, dao.IronicRdb, "uuid = ? and tenant = ? and is_del = 0", uuid, tenant).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func QueryByParam(logger *log.Logger, query map[string]interface{}, offset, limit int64) (ml []*Keypair, err error) {
	return GetMultiKeypair(logger, query, nil, nil, nil, offset, limit)
}
