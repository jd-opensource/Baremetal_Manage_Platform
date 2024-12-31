package AdminDao

import (
	"time"

	log "git.jd.com/cps-golang/log"

	"coding.jd.com/aidc-bmp/bmp-operation-api/dao"
)

type CpsUser struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Username   string    `gorm:"unique;column:username;type:varchar(32);not null" json:"username"`
	Password   string    `gorm:"column:password;type:varchar(64)" json:"password"`
	Fullname   string    `gorm:"column:fullname;type:varchar(64)" json:"fullname"`
	Email      string    `gorm:"column:email;type:varchar(128)" json:"email"`
	Mobile     string    `gorm:"column:mobile;type:varchar(32)" json:"mobile"`
	ErpID      string    `gorm:"column:erp_id;type:varchar(32)" json:"erp_id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	IsDel      int8      `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`
}

func (t *CpsUser) TableName() string {
	return "cps_user"
}

// AddImage insert a new Image into database and returns
// last inserted Id on success.
func AddUser(logger *log.Logger, m *CpsUser) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetImageById retrieves Image by Id. Returns error if
// Id doesn't exist
func GetUserById(logger *log.Logger, id int64) (v *CpsUser, err error) {
	v = &CpsUser{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	return
}

func GetUserByName(logger *log.Logger, username string) (v *CpsUser, err error) {
	v = &CpsUser{}
	err = dao.Where(logger, dao.IronicRdb, "username = ? and is_del = 0", username).Take(v).Error
	return
}

// GetMultiImage retrieves all Image matches certain condition. Returns empty list if
// no records exist

func GetAllUser(logger *log.Logger, query map[string]interface{}) (ml []*CpsUser, err error) {

	var db = dao.Model(logger, dao.IronicRdb, CpsUser{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateImage updates Image by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(logger *log.Logger, m *CpsUser) (err error) {
	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, CpsUser{}).Where("id = ?", m.ID).Updates(m).Error
}
