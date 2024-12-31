package response

type AlertList struct {
	// rules数组
	Alerts     []*Alert `json:"instances"`
	PageNumber int64    `json:"pageNumber"`
	PageSize   int64    `json:"pageSize"`
	TotalCount int64    `json:"totalCount"`
}

type Alert struct {
	// alert uuid
	AlertID string `json:"alertId"`
	// 告警时间戳
	AlertTime int64 `json:"alertTime"`
	// rule obj
	Rule *Rule `json:"rule"`
	//涉及到哪些实例
	Instance *Instance `json:"instance"`
	// 资源类型 [只支持instance一种]
	Resource string `json:"resource"`
	// 资源id,目前就是实例id
	ResourceID string `json:"resourceId"`
	// 资源名称,目前就是实例名称
	ResourceName string `json:"resourceName"`
	// 触发条件,接口需要翻译
	Trigger string `json:"trigger"`
	//触发条件描述
	TriggerDescription string `json:"triggerDescription"`
	// 报警值
	AlertValue string `json:"alertValue"`
	// 1表示一般，2表示严重，3表示紧急
	AlertLevel int64 `json:"alertLevel"`
	//级别描述
	AlertLevelDescription string `json:"alertLevelDescription"`
	// 告警持续时间(分钟为单位)
	AlertPeriod int64 `json:"alertPeriod"`
	// 通知对象 userid
	UserID string `json:"userId"`
	// 通知对象 用户名
	UserName string `json:"userName"`
	// 创建时间戳
	CreatedTime int64 `json:"createdTime"`
}
