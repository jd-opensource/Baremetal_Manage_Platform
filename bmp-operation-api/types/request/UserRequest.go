package request

// 【用户管理】【用户列表】
type QueryUsersRequest struct {
	BaseRequest
	PagingRequest
	RoleID    string `json:"roleId"`
	ProjectID string `json:"defaultProjectId"`
	UserName  string `json:"userName"`
	IsAll     string `json:"isAll"`
}

// 【用户管理】【添加用户】
type AddUserRequest struct {
	// BaseRequest
	RoleID      string `json:"roleId"`      // 角色uuid
	UserName    string `json:"userName"`    // 用户名，唯一
	Email       string `json:"email"`       // 邮箱
	PhonePrefix string `json:"phonePrefix"` // 国家地区码，如86
	PhoneNumber string `json:"phoneNumber"` // 手机号
	Language    string `json:"language"`    // 语言（中文/English）
	Password    string `json:"password"`    // 密码，sha256非对称加密后存储
	Description string `json:"description"` // 描述
}

type QueryUserInfoRequest struct {
	UserID string `json:"userId"` // 用户uuid
}

// 【用户管理】【编辑用户信息】
type ModifyUserInfoRequest struct {
	//UserID           string  `json:"userId"`
	RoleID           *string `json:"roleId"`           // 角色uuid
	UserName         *string `json:"userName"`         // 用户名，唯一
	Email            *string `json:"email"`            // 邮箱
	PhonePrefix      *string `json:"phonePrefix"`      // 国家地区码，如86
	PhoneNumber      *string `json:"phoneNumber"`      // 手机号
	DefaultProjectID string  `json:"defaultProjectId"` // 所属项目uuid
	Language         *string `json:"language"`         // 语言（中文/English）
	//Timezone         time.Time `gorm:"column:timezone" json:"timezone"`                   // UTC或者本时区时间
	Password    *string `json:"password"`    // 密码，sha256非对称加密后存储
	Description *string `json:"description"` // 描述
}

type DeleteUserInfoRequest struct {
	UserID string `json:"userId"` // 用户uuid
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerifyUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateLocalPasswordRequest struct {
	OldPassword string `json:"oldPassword"` //
	Password    string `json:"password"`    // 密码
}

// 【用户管理】【编辑用户信息】
type ModifyLocalUserInfoRequest struct {
	UserName         *string `json:"userName"`         // 用户名，唯一
	Email            *string `json:"email"`            // 邮箱
	PhonePrefix      *string `json:"phonePrefix"`      // 国家地区码，如86
	PhoneNumber      *string `json:"phoneNumber"`      // 手机号
	DefaultProjectID *string `json:"defaultProjectId"` // 所属项目uuid
	Language         *string `json:"language"`         // 语言（中文/English）
	Timezone         *string `json:"timezone"`         // UTC或者本时区时间
	Description      *string `json:"description"`      // 描述
}
