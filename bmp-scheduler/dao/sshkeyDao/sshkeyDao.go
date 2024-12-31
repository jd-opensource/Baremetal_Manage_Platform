package sshkeyDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// Sshkey Sshkey
type Sshkey struct {
	ID          int64  `gorm:"id" json:"id"`
	SshkeyId    string `gorm:"sshkey_id" json:"SshkeyId"`              // 秘钥uuid
	Name        string `gorm:"name" json:"name"`                       // 秘钥名称
	Key         string `gorm:"key" json:"key"`                         // 公钥，格式：ssh-rsa AAA
	FingerPrint string `gorm:"finger_print" json:"fingerPrint"`        // 公钥指纹
	CreatedBy   string `gorm:"created_by" json:"createdBy"`            // 创建者
	UpdatedBy   string `gorm:"updated_by" json:"updatedBy"`            // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *Sshkey) TableName() string {
	return "ssh_key"
}

// AddSshkey insert a new Sshkey into database and returns
// last inserted Id on success.
func AddSshkey(logger *log.Logger, m *Sshkey) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetSshkeyById retrieves Sshkey by Id. Returns error if
// Id doesn't exist
func GetSshkeyById(logger *log.Logger, SshkeyId string) (v *Sshkey, err error) {
	v = &Sshkey{}
	err = dao.Where(logger, dao.IronicRdb, "sshkey_id = ? and is_del = 0", SshkeyId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetSshkeyByUuid retrieves Sshkey by Uuid. Returns error if
// Id doesn't exist
func GetSshkeyCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Sshkey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllSshkey retrieves all Sshkey matches certain condition. Returns empty list if
// no records exist
func GetAllSshkey(logger *log.Logger, query map[string]interface{}) (ml []*Sshkey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Sshkey{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiSshkey retrieves all Sshkey matches certain condition. Returns empty list if
// no records exist
func GetMultiSshkey(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Sshkey, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Sshkey{})
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

// UpdateSshkey updates Sshkey by Id and returns error if
// the record to be updated doesn't exist
func UpdateSshkeyById(logger *log.Logger, m *Sshkey) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Sshkey{}).Where("Sshkey_id = ?", m.SshkeyId).Updates(m).Error
}
