package response

type Role struct {
	ID int `json:"-"`
	// 角色uuid
	RoleID string `json:"roleId"`
	// 角色名称，唯一
	RoleNameEn string `json:"roleNameEn"`
	// 角色名称，唯一
	RoleNameCn string `json:"roleNameCn"`
	// 权限描述
	DescriptionEn string `json:"descriptionEn"`
	// description
	DescriptionCn string `json:"descriptionCn"`
	// 角色名称
	RoleName string `json:"roleName"`
	// 权限
	Permission string `json:"permission"`
	// 用户数
	UserCount int `json:"userCount"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
}
type RoleList struct {
	//角色实体列表
	Roles      []*Role `json:"roles"`
	PageNumber int64   `json:"pageNumber"`
	PageSize   int64   `json:"pageSize"`
	TotalCount int64   `json:"totalCount"`
}
type RoleInfo struct {
	//角色实体
	Role *Role `json:"role"`
}
type RoleId struct {
	//角色uuid
	RoleId string `json:"roleId"`
}
