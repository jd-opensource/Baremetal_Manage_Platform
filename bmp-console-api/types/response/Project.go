package response

type Project struct {
	ID            int    `json:"-"`
	ProjectName   string `json:"projectName"` // 角色名称，唯一
	ProjectID     string `json:"projectId"`   // 权限
	InstanceCount int    `json:"instanceCount"`
	CreatedBy     string `json:"createdBy"`   // 创建者
	UpdatedBy     string `json:"updatedBy"`   // 更新者
	CreatedTime   string `json:"createdTime"` // 创建时间戳
	UpdatedTime   string `json:"updatedTime"` // 更新时间戳
}
type ProjectList struct {
	Projects   []*Project `json:"projects"`
	PageNumber int64      `json:"pageNumber"`
	PageSize   int64      `json:"pageSize"`
	TotalCount int64      `json:"totalCount"`
}
type ProjectInfo struct {
	Project *Project `json:"project"`
}
type ProjectId struct {
	ProjectId string `json:"projectId"`
}
