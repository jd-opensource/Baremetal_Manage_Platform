package response

type User struct {
	// 用户id
	ID int `json:"id"`
	// 用户uuid
	UserID string `json:"userId"`
	// 角色uuid
	RoleID string `json:"roleId"`
	// 角色名称
	RoleName string `json:"roleName"`
	// 用户名，唯一
	UserName string `json:"userName"`
	// 邮箱
	Email string `json:"email"`
	// 国家地区码，如86
	PhonePrefix string `json:"phonePrefix"`
	// 手机号
	PhoneNumber string `json:"phoneNumber"`
	// 用户默认项目uuid
	DefaultProjectID string `json:"defaultProjectId"`
	// 用户默认项目名称
	DefaultProjectName string `json:"defaultProjectName"`
	// 用户拥有的项目数量
	ProjectCount int `json:"projectCount"`
	// 用户拥有的实例数量
	InstanceCount int `json:"instanceCount"`
	// 默认语言（en_US/zh_CN）
	Language string `json:"language"`
	// 时区
	Timezone string `gorm:"column:timezone" json:"timezone"`
	Password string `json:"-"`
	// 描述
	Description string `json:"description"`
	// 创建者
	CreatedBy string `gorm:"column:created_by" json:"createdBy"`
	// 更新者
	UpdatedBy string `gorm:"column:updated_by" json:"updatedBy"`
	// 创建时间
	CreatedTime string `gorm:"column:created_time" json:"createdTime"`
	// 更新时间
	UpdatedTime string `gorm:"column:updated_time" json:"updatedTime"`
}
type UserList struct {
	// user实体列表
	Users []*User `json:"users"`
	// 页数
	PageNumber int64 `json:"pageNumber"`
	// 页大小
	PageSize int64 `json:"pageSize"`
	// 总条数
	TotalCount int64 `json:"totalCount"`
}
type UserInfo struct {
	// 用户实体
	User *User `json:"user"`
}
type UserId struct {
	// 用户uuid
	UserId string `json:"userId"`
}
