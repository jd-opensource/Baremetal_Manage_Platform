package request

type AddAlertRequest struct {
	//规则名称
	RuleName string `json:"ruleName" validate:"required"`
	//规则ID
	// required: true
	RuleID string `json:"ruleId" validate:"required"`
	//实例id
	// required: true
	InstanceID string `json:"instanceId" validate:"required"`
	//触发条件
	// required: true
	Trigger string `json:"trigger" validate:"required"`
	//告警值
	// required: true
	AlertValue string `json:"alertValue" validate:"required"`
	//告警时间戳
	// required: true
	AlertTimestamp int64 `json:"alertTimestamp" validate:"required"`
	//告警持续时间
	// required: true
	AlertPeriod int `json:"alertPeriod" validate:"required"`
}

type DescribeAlertRequest struct {
	// rule uuid
	// required: true
	AlertID string `json:"alertId" validate:"required"`
}

type DescribeAlertsRequest struct {
	PagingRequest
	ExportType string `json:"exportType"`

	// 是否显示全部，取值为1时表示全部
	IsAll string `json:"isAll"`
	//user uuid
	UserID string `json:"userId"`
	//username
	UserName string `json:"userName"`
	//规则名称
	RuleName string `json:"ruleName"`
	//规则uuid
	RuleID string `json:"ruleId"`
	//资源id,目前就是实例id
	ResourceID string `json:"resourceId"`
	//报警时间筛选stime
	StartTime int64 `json:"startTime"`
	//报警时间筛选etime
	EndTime int64 `json:"endTime"`
}

type DeleteAlertRequest struct {
	//告警id
	// required: true
	AlertId string `json:"alertId" validate:"required"`
}
