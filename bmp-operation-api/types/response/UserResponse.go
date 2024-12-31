package response

import "time"

type UserPage struct {
	Users []User `json:"users"`
	PagingResponse
}

type User struct {
	ID                 int64     `json:"id"`
	UserID             string    `json:"userId"`                                 // 用户uuid
	RoleID             string    `json:"roleId"`                                 // 角色uuid
	RoleName           string    `json:"roleName"`                               // 角色名称
	UserName           string    `json:"userName"`                               // 用户名，唯一
	Email              string    `json:"email"`                                  // 邮箱
	PhonePrefix        string    `json:"phonePrefix"`                            // 国家地区码，如86
	PhoneNumber        string    `json:"phoneNumber"`                            // 手机号
	DefaultProjectID   string    `json:"defaultProjectId"`                       // 所属项目uuid
	DefaultProjectName string    `json:"DefaultProjectName"`                     // 默认所属项目名称
	ProjectCount       int       `json:"projectCount"`                           //用户拥有的项目数
	InstanceCount      int       `json:"instanceCount"`                          //用户拥有的实例数
	Language           string    `json:"language"`                               // 语言（中文/English）
	Timezone           time.Time `gorm:"column:timezone" json:"timezone"`        // UTC或者本时区时间
	Password           string    `json:"password"`                               // 密码，sha256非对称加密后存储
	Description        string    `json:"description"`                            // 描述
	CreatedBy          string    `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy          string    `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime        string    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime        string    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	//DeletedTime string `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	//IsDel            int8      `gorm:"column:is_del" json:"isDel"`                        // 是否删除0未删除 1已删除
}

type UserAdd struct {
	UserID string `json:"userId"` // 用户uuid
}

type LocalUser struct {
	ID                 int64  `json:"id"`
	UserID             string `json:"userId"`                                // 用户uuid
	RoleID             string `json:"roleId"`                                // 角色uuid
	RoleName           string `json:"roleName"`                              // 角色名称
	UserName           string `json:"userName"`                              // 用户名，唯一
	Email              string `json:"email"`                                 // 邮箱
	PhonePrefix        string `json:"phonePrefix"`                           // 国家地区码，如86
	PhoneNumber        string `json:"phoneNumber"`                           // 手机号
	DefaultProjectID   string `json:"defaultProjectId"`                      // 所属项目uuid
	DefaultProjectName string `json:"DefaultProjectName"`                    // 默认所属项目名称
	Language           string `json:"language"`                              // 语言（中文/English）
	Timezone           string `gorm:"column:timezone" json:"timezone"`       // UTC或者本时区时间
	Password           string `json:"password"`                              // 密码，sha256非对称加密后存储
	Description        string `json:"description"`                           // 描述
	CreatedBy          string `gorm:"column:created_by" json:"createdBy"`    // 创建者
	UpdatedBy          string `gorm:"column:updated_by" json:"updatedBy"`    // 更新者
	CreatedTime        string `gorm:"column:createdTime" json:"createdTime"` // 创建时间戳
	UpdatedTime        string `gorm:"column:updatedTime" json:"updatedTime"` // 更新时间戳
	//DeletedTime string `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	//IsDel            int8      `gorm:"column:is_del" json:"isDel"`                        // 是否删除0未删除 1已删除
	LicenseExist  bool     `json:"licenseExist"`  //系统是否做过授权
	LicenseModels []string `json:"licenseModels"` //授权过的模块
}
