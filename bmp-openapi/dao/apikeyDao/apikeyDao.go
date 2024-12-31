package apikeyDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/jinzhu/gorm"
)

// Apikey Apikey
type Apikey struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	ApikeyID    string `gorm:"column:apikey_id" json:"apikeyId"`       // 秘钥对UUID
	Name        string `gorm:"column:name" json:"name"`                // 秘钥对名称
	ReadOnly    int8   `gorm:"column:read_only" json:"readOnly"`       // 是否支持只读，read_only =1 的时候说明这个key是只读key，不能访问写方法。
	Token       string `gorm:"column:token" json:"token"`              // 32位字符令牌
	Type        string `gorm:"column:type" json:"type"`                // Token类型system/user
	UserID      string `gorm:"column:user_id" json:"userId"`           // 所属用户
	Source      string `gorm:"column:source" json:"source"`            // apikey是通过console创建的还是通过operation创建的
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *Apikey) TableName() string {
	return "apikey"
}

// AddApikey insert a new Apikey into database and returns
// last inserted Id on success.
func AddApikey(logger *log.Logger, m *Apikey) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetApikeyById retrieves Apikey by Id. Returns error if
// Id doesn't exist
func GetApikeyById(logger *log.Logger, ApikeyId string) (v *Apikey, err error) {
	v = &Apikey{}
	err = dao.Where(logger, dao.IronicRdb, "apikey_id = ? and is_del = 0", ApikeyId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetApikeyByToken(logger *log.Logger, token string) (v *Apikey, err error) {
	v = &Apikey{}
	err = dao.Where(logger, dao.IronicRdb, "token = ? and is_del = 0", token).Take(v).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetApikeyByUuid retrieves Apikey by Uuid. Returns error if
// Id doesn't exist
func GetApikeyCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Apikey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllApikey retrieves all Apikey matches certain condition. Returns empty list if
// no records exist
func GetAllApikey(logger *log.Logger, query map[string]interface{}) (ml []*Apikey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Apikey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiApikey retrieves all Apikey matches certain condition. Returns empty list if
// no records exist
func GetMultiApikey(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Apikey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Apikey{})
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

// UpdateApikey updates Apikey by Id and returns error if
// the record to be updated doesn't exist
func UpdateApikeyById(logger *log.Logger, m *Apikey) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Apikey{}).Where("apikey_id = ?", m.ApikeyID).Updates(m).Error
}
