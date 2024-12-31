package dao

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// user user
type User struct {
	ID               int    `gorm:"primaryKey;column:id" json:"-"`
	UserID           string `gorm:"column:user_id" json:"userId"`                       // 用户uuid
	RoleID           string `gorm:"column:role_id" json:"-"`                            // 角色uuid
	UserName         string `gorm:"column:user_name" json:"userName"`                   // 用户名，唯一
	Email            string `gorm:"column:email" json:"-"`                              // 邮箱
	PhonePrefix      string `gorm:"column:phone_prefix" json:"-"`                       // 国家地区码，如86
	PhoneNumber      string `gorm:"column:phone_number" json:"-"`                       // 手机号
	DefaultProjectID string `gorm:"column:default_project_id" json:"-"`                 // 所属项目uuid
	Language         string `gorm:"column:language" json:"language"`                    // 语言（中文/English）
	Timezone         string `gorm:"column:timezone;type:varchar(64);not null" json:"-"` // timezone
	Password         string `gorm:"column:password" json:"-"`                           // 密码，sha256非对称加密后存储
	Description      string `gorm:"column:description" json:"-"`                        // 描述
	CreatedBy        string `gorm:"column:created_by" json:"-"`                         // 创建者
	UpdatedBy        string `gorm:"column:updated_by" json:"-"`                         // 更新者
	CreatedTime      int    `gorm:"column:created_time" json:"-"`                       // 创建时间戳
	UpdatedTime      int    `gorm:"column:updated_time" json:"-"`                       // 更新时间戳
	DeletedTime      int    `gorm:"column:deleted_time" json:"-"`                       // 删除时间戳
	IsDel            int8   `gorm:"column:is_del" json:"-"`                             // 是否删除0未删除 1已删除
}

func (t *User) TableName() string {
	return "user"
}

// GetuserById retrieves user by Id. Returns error if
// Id doesn't exist
func GetUserById(logger *log.Logger, userId string) (v *User, err error) {
	v = &User{}
	err = Where(logger, IronicRdb, "user_id = ? and is_del = 0", userId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetUserByName(logger *log.Logger, userName string) (v *User, err error) {
	v = &User{}
	err = Where(logger, IronicRdb, "user_name = ? and is_del = 0", userName).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}
