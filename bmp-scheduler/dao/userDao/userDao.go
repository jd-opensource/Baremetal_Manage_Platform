package userDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// user user
type User struct {
	ID               int    `gorm:"primaryKey;column:id" json:"-"`
	UserID           string `gorm:"column:user_id" json:"userId"`                              // 用户uuid
	RoleID           string `gorm:"column:role_id" json:"roleId"`                              // 角色uuid
	UserName         string `gorm:"column:user_name" json:"userName"`                          // 用户名，唯一
	Email            string `gorm:"column:email" json:"email"`                                 // 邮箱
	PhonePrefix      string `gorm:"column:phone_prefix" json:"phonePrefix"`                    // 国家地区码，如86
	PhoneNumber      string `gorm:"column:phone_number" json:"phoneNumber"`                    // 手机号
	DefaultProjectID string `gorm:"column:default_project_id" json:"defaultProjectId"`         // 所属项目uuid
	Language         string `gorm:"column:language" json:"language"`                           // 语言（中文/English）
	Timezone         string `gorm:"column:timezone;type:varchar(64);not null" json:"timezone"` // timezone
	Password         string `gorm:"column:password" json:"password"`                           // 密码，sha256非对称加密后存储
	Description      string `gorm:"column:description" json:"description"`                     // 描述
	CreatedBy        string `gorm:"column:created_by" json:"createdBy"`                        // 创建者
	UpdatedBy        string `gorm:"column:updated_by" json:"updatedBy"`                        // 更新者
	CreatedTime      int    `gorm:"column:created_time" json:"createdTime"`                    // 创建时间戳
	UpdatedTime      int    `gorm:"column:updated_time" json:"updatedTime"`                    // 更新时间戳
	DeletedTime      int    `gorm:"column:deleted_time" json:"deletedTime"`                    // 删除时间戳
	IsDel            int8   `gorm:"column:is_del" json:"isDel"`                                // 是否删除0未删除 1已删除
}

func (t *User) TableName() string {
	return "user"
}

// Adduser insert a new user into database and returns
// last inserted Id on success.
func Adduser(logger *log.Logger, m *User) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetuserById retrieves user by Id. Returns error if
// Id doesn't exist
func GetUserById(logger *log.Logger, userId string) (v *User, err error) {
	v = &User{}
	err = dao.Where(logger, dao.IronicRdb, "user_id = ? and is_del = 0", userId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetUserByName(logger *log.Logger, userName string) (v *User, err error) {
	v = &User{}
	err = dao.Where(logger, dao.IronicRdb, "user_name = ? and is_del = 0", userName).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetuserByUuid retrieves user by Uuid. Returns error if
// Id doesn't exist
func GetUserCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, User{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAlluser retrieves all user matches certain condition. Returns empty list if
// no records exist
func GetAllUser(logger *log.Logger, query map[string]interface{}) (ml []*User, err error) {

	var db = dao.Model(logger, dao.IronicRdb, User{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiuser retrieves all user matches certain condition. Returns empty list if
// no records exist
func GetMultiUser(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*User, err error) {

	var db = dao.Model(logger, dao.IronicRdb, User{})
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

// Updateuser updates user by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(logger *log.Logger, m *User) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, User{}).Where("user_id = ?", m.UserID).Save(m).Error
}
