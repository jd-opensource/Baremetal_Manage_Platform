package response

type ShareProject struct {
	ID int `json:"-"`
	// 项目名称
	ProjectName string `json:"projectName"`
	// 项目uuid
	ProjectID string `json:"projectId"`
	// 项目拥有者用户id
	OwnerUserID string `json:"ownerUserId"`
	// 项目共享者用户id
	SharedUserID string `json:"sharedUserId"`
	// 项目拥有者用户名
	OwnerUserName string `json:"ownerUserName"`
	// 项目拥有者用户名
	SharedUserName string `json:"sharedUserName"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`

	ShareProjects []ShareProject `json:"shareProjects"`
}
