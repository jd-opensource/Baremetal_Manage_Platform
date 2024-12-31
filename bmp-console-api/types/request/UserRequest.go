package request

// 【用户管理】【编辑用户信息】
type ModifyUserInfoRequest struct {
	UserName         *string `json:"userName"`         // 用户名，唯一
	Email            *string `json:"email"`            // 邮箱
	PhonePrefix      *string `json:"phonePrefix"`      // 国家地区码，如86
	PhoneNumber      *string `json:"phoneNumber"`      // 手机号
	DefaultProjectID *string `json:"defaultProjectId"` // 所属项目uuid
	Language         *string `json:"language"`         // 语言（中文/English）
	Timezone         *string `json:"timezone"`         // UTC或者本时区时间
	Description      *string `json:"description"`      // 描述
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"` //
	Password    string `json:"password"`    // 密码
}

type QueryUserInfoRequest struct {
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckUserConsoleAccessRequest struct {
	CheckedUserId string `json:"checkedUserId"`
}

type CheckUserConsoleAccessByNameRequest struct {
	CheckedUserName string `json:"checkedUserName"`
	ProjectID       string `json:"projectId"`
	// [move/share]
	Operation string `json:"operation"`
}
