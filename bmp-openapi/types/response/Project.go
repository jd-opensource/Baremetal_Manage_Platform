package response

type Project struct {
	ID int `json:"-"`
	// 项目名称
	ProjectName string `json:"projectName"`
	// 项目uuid
	ProjectID string `json:"projectId"`
	// 项目下实例数量
	InstanceCount int64 `json:"instanceCount"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
	// 共享
	ShareProjects []ShareProject `json:"shareProjects"`
	// 此项目的共享标志位，1为拥有，2为共享
	Owned int `json:"owned"`
}
type ProjectList struct {
	// 项目实体列表
	Projects []*Project `json:"projects"`
	// 页数
	PageNumber int64 `json:"pageNumber"`
	// 页大小
	PageSize int64 `json:"pageSize"`
	// 总条数
	TotalCount int64 `json:"totalCount"`
}
type ProjectInfo struct {
	// 项目实体
	Project *Project `json:"project"`
}
type ProjectId struct {
	// 项目uuid
	ProjectId string `json:"projectId"`
}
