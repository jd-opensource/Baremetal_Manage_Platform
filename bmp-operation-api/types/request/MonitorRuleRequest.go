package request

type AddRuleRequest struct {
	//规则名称
	// required: true
	RuleName string `json:"ruleName" validate:"required"`
	//维度 [instance、disk、mountpoint、nic]
	// required: true
	Dimension string `json:"dimension" validate:"required,oneof=instance disk mountpoint nic"`
	//资源类型,目前只支持instance
	// required: true
	Resource string `json:"resource" validate:"required,oneof=instance"`
	//触发条件
	// required: true
	TriggerOption []RuleTrigger `json:"triggerOption"`
	//通知策略
	// required: true
	NoticeOption RuleNotice `json:"noticeOption"`
	//实例id列表
	// required: true
	InstanceIds []string `json:"instanceIds" validate:"required"`
	//盘符、挂载点、网口列表
	DeviceTag string `json:"deviceTag"`
}

//单个触发条件
type RuleTrigger struct {
	//监控指标 cps.cpu.util
	// required: true
	Metric string `json:"metric" validate:"required"`
	//周期(分钟) [1, 2, 5, 15, 30, 60]
	// required: true
	Period int64 `json:"period" validate:"required,oneof=1 2 5 15 30 60"`
	//计算方式 [min max avg sum]
	// required: true
	Calculation string `json:"calculation" validate:"required,oneof=min max avg sum"`
	//计算结果单位 [对于使用量，有Bytes,KB,MB,GB,TB，对于使用率，是%，对于连接个数，是count，对于网络包量，是pps,Kpps,Mpps,Gpps,Tpps，对于网络速率，是bps,Kbps,Mbps,Gbps,Tbps 对于负载，没有单位。。。。。。产品需详细列出]
	// required: true
	CalculationUnit string `json:"calculationUnit"`
	//比较方式 [> >= < <= == !=]或者[gt gte lt lte eq neq]
	// required: true
	Operation string `json:"operation" validate:"required"`
	//阈值
	// required: true
	Threshold float64 `json:"threshold"`
	//持续周期数 [1, 2, 3, 4, 5, 10, 15, 30, 60]
	// required: true
	Times int64 `json:"times" validate:"required,oneof=1 2 3 4 5 10 15 30 60"`
	//告警级别 [1表示一般，2表示严重，3表示紧急]
	// required: true
	NoticeLevel int64 `json:"noticeLevel" validate:"required,oneof=1 2 3"`
}

//通知策略
type RuleNotice struct {
	//通知周期(分钟) [5 10 15 30 60 180 360 720 1440]
	// required: true
	NoticePeriod int64 `json:"noticePeriod" validate:"required,oneof=5 10 15 30 60 180 360 720 1440"`
	// 有效时段开始时间 ["00:00:00"]
	// required: true
	EffectiveIntervalStart string `json:"effectiveIntervalStart" validate:"required"`
	// 有效时段结束时间 ["23:59:59"]
	// required: true
	EffectiveIntervalEnd string `json:"effectiveIntervalEnd" validate:"required"`
	//通知条件，可多选 [1表示报警， 2表示恢复正常]
	// required: true
	NoticeCondition []int64 `json:"noticeCondition" validate:"required"`
	//接收渠道，可多选 [1表示站内信， 2表示邮件]
	// required: true
	NoticeWay []int64 `json:"noticeWay" validate:"required"`
	//通知对象
	// required: true
	UserID string `json:"userId" validate:"required"`
}

type DescribeRuleRequest struct {
	// rule uuid
	// required: true
	RuleID string `json:"ruleId" validate:"required"`
}

type DescribeRulesRequest struct {
	PagingRequest
	ExportType string `json:"exportType"`
	// 是否显示全部，取值为1时表示全部
	IsAll string `json:"isAll"`
	//user uuid
	UserId string `json:"userId"`
	//username
	UserName string `json:"userName"`
	//规则名称
	RuleName string `json:"ruleName"`
	//规则uuid
	RuleId string `json:"ruleId"`
	Status int    `json:"status"`
}

type EditRuleRequest struct {
	//规则id
	// required: true
	RuleId string `json:"ruleId" validate:"required"`
	//规则名称
	// required: true
	RuleName string `json:"ruleName" validate:"required"`
	//维度 [instance、disk、mountpoint、nic]
	// required: true
	Dimension string `json:"dimension" validate:"required,oneof=instance disk mountpoint nic"`
	//资源类型,目前只支持instance
	// required: true
	Resource string `json:"resource" validate:"required,oneof=instance"`
	//触发条件
	// required: true
	TriggerOption []RuleTrigger `json:"triggerOption"`
	//通知策略
	// required: true
	NoticeOption RuleNotice `json:"noticeOption"`
	//实例id列表
	// required: true
	InstanceIds []string `json:"instanceIds" validate:"required"`
	//盘符、挂载点、网口列表
	DeviceTag string `json:"deviceTag"`
}

type EnableRuleRequest struct {
	//规则id
	// required: true
	RuleId string `json:"ruleId" validate:"required"`
}

type DisableRuleRequest struct {
	//规则id
	// required: true
	RuleId string `json:"ruleId" validate:"required"`
}

type DeleteRuleRequest struct {
	//规则id
	// required: true
	RuleId string `json:"ruleId" validate:"required"`
}
