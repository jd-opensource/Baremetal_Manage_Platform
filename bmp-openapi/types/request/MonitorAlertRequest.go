package request

import log "coding.jd.com/aidc-bmp/bmp_log"

type AddAlertRequest struct {
	Alerts []*AddAlertItem `json:"alerts"`
}

type AddAlertItem struct {
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

func (req *AddAlertRequest) Validate(logger *log.Logger) {

	validate(req, logger)

}

type DescribeAlertRequest struct {
	// rule uuid
	// required: true
	AlertID string `json:"alertId" validate:"required"`
}

func (req *DescribeAlertRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type DescribeAlertsRequest struct {
	Pageable
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

func (req *DescribeAlertsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type DeleteAlertRequest struct {
	//告警id
	// required: true
	AlertId string `json:"alertId" validate:"required"`
}

func (req *DeleteAlertRequest) Validate(logger *log.Logger) {

	validate(req, logger)

}
