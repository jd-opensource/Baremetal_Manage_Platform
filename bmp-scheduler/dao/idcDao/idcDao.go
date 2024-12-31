package idcDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

type Idc struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                  // ID
	IDcID           string `gorm:"column:idc_id" json:"idcId"`                     // 机房uuid
	Name            string `gorm:"column:name" json:"name"`                        // 机房名称
	Shortname       string `gorm:"column:shortname" json:"shortname"`              // 机房名称缩写
	Level           string `gorm:"column:level" json:"level"`                      // 机房等级
	Address         string `gorm:"column:address" json:"address"`                  // 机房地址
	IloUser         string `gorm:"column:ilo_user" json:"iloUser"`                 // 带外登录账号
	IloPassword     string `gorm:"column:ilo_password" json:"iloPassword"`         // 带外登录密码
	SwitchUser1     string `gorm:"column:switch_user1" json:"switchUser1"`         // 交换机1登录账号
	SwitchPassword1 string `gorm:"column:switch_password1" json:"switchPassword1"` // 交换机1登录密码
	SwitchUser2     string `gorm:"column:switch_user2" json:"switchUser2"`         // 交换机2登录账号
	SwitchPassword2 string `gorm:"column:switch_password2" json:"switchPassword2"` // 交换机2登录密码
	CreatedBy       string `gorm:"column:created_by" json:"createdBy"`             // 创建者
	UpdatedBy       string `gorm:"column:updated_by" json:"updatedBy"`             // 更新者
	CreatedTime     int    `gorm:"column:created_time" json:"createdTime"`         // 创建时间戳
	UpdatedTime     int    `gorm:"column:updated_time" json:"updatedTime"`         // 更新时间戳
	DeletedTime     int    `gorm:"column:deleted_time" json:"deletedTime"`         // 删除时间戳
	IsDel           int8   `gorm:"column:is_del" json:"isDel"`                     // 是否删除0未删除 1已删除
}

func (t *Idc) TableName() string {
	return "idc"
}

// AddIdc insert a new Idc into database and returns
// last inserted Id on success.
func CreateIdc(logger *log.Logger, m *Idc) (id int64, err error) {
	return dao.CreateAndGetId(logger, dao.IronicRdb, m)
}

// GetIdcById retrieves Idc by Id. Returns error if
// Id doesn't exist
func GetIdcById(logger *log.Logger, idcId string) (v *Idc, err error) {
	v = &Idc{}
	err = dao.Where(logger, dao.IronicRdb, "idc_id = ? and is_del = 0", idcId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetOneIdc(logger *log.Logger, query map[string]interface{}) (l *Idc, err error) {
	l = &Idc{}
	var db = dao.Model(logger, dao.IronicRdb, Idc{})
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

// GetAllIdc retrieves all Idc matches certain condition. Returns empty list if
// no records exist
func GetAllIdc(logger *log.Logger, query map[string]interface{}) (ml []*Idc, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Idc{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiIdc retrieves all Idc matches certain condition. Returns empty list if
// no records exist
func GetMultiIdc(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Idc, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Idc{})
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

func GetIdcCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Idc{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return

}

// UpdateIdc updates Idc by Id and returns error if
// the record to be updated doesn't exist
func UpdateIdcById(logger *log.Logger, m *Idc, dcId string) (err error) {
	return dao.Model(logger, dao.IronicWdb, Idc{}).Where("idc_id = ?", dcId).Update(m).Error
}

//func DeleteIdcById(logger *log.Logger, m *Idc) (err error) {
//	m.UpdatedBy = "admin"
//	//m.UpdatedAt = time.Now().Unix()
//	//m.De = time.Now().Unix()
//	return dao.Model(logger, dao.IronicWdb, Idc{}).Where("idc_id = ?", m.DcId).Update(m).Error
//}

func QueryByParam(logger *log.Logger, query map[string]interface{}, offset, limit int64) (ml []*Idc, err error) {
	return GetMultiIdc(logger, query, nil, []string{"id"}, []string{"desc"}, offset, limit)
}
