package response

type Role struct {
	ID            int    `json:"-"`
	RoleID        string `json:"roleId"`        // 角色uuid
	RoleNameEn    string `json:"roleNameEn"`    // 角色名称，唯一
	RoleNameCn    string `json:"roleNameCn"`    // 角色名称，唯一
	DescriptionEn string `json:"descriptionEn"` // 权限描述
	DescriptionCn string `json:"descriptionCn"` // 权限描述
	RoleName      string `json:"roleName"`
	Permission    string `json:"permission"`
	UserCount     int    `json:"userCount"`
	CreatedBy     string `json:"createdBy"`   // 创建者
	UpdatedBy     string `json:"updatedBy"`   // 更新者
	CreatedTime   string `json:"createdTime"` // 创建时间戳
	UpdatedTime   string `json:"updatedTime"` // 更新时间戳
}
type RoleList struct {
	Roles      []*Role `json:"roles"`
	PageNumber int64   `json:"pageNumber"`
	PageSize   int64   `json:"pageSize"`
	TotalCount int64   `json:"totalCount"`
}
type RoleInfo struct {
	Role *Role `json:"role"`
}
type RoleId struct {
	RoleId string `json:"roleId"`
}
