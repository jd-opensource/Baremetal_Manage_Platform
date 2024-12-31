package request

type DescribeAuditLogsRequest struct {
	PagingRequest
	// instance_id
	InstanceID string `json:"instanceId" validate:"required"`
	// 操作名称
	Operation string `json:"operation"`
	// 操作人
	UserName string `json:"userName"`
	// result
	Result string `json:"result"`
	// 操作时间下限
	StartTime int64 `json:"startTime"`
	// 操作时间上限
	EndTime int64 `json:"endTime"`
	// 是否显示所有
	IsAll      string `json:"isAll"`
	ExportType string `json:"exportType"` // v2.27.2   0：导出全部实力；1：导出选中实例；2：导出搜索结果
}
