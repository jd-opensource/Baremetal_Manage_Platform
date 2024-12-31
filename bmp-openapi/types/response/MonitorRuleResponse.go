package response

type RuleList struct {
	// rules数组
	Rules      []*Rule `json:"rules"`
	PageNumber int64   `json:"pageNumber"`
	PageSize   int64   `json:"pageSize"`
	TotalCount int64   `json:"totalCount"`
}

type Rule struct {
	//规则uuid
	RuleID string `json:"ruleId"`
	//规则名称
	RuleName string `json:"ruleName"`
	//维度 [instance、disk、mountpoint、nic]
	Dimension string `json:"dimension"`
	//维度名称 [实例、盘符、挂载点、网卡]
	DimensionName string `json:"dimensionName"`
	//资源类型,目前只支持instance
	Resource string `json:"resource"`
	//资源类型名称 实例
	ResourceName string `json:"resourceName"`
	//触发条件
	TriggerOption []*RuleTrigger `json:"triggerOption"`
	//触发条件描述
	TriggerDescription []string `json:"triggerDescription"`
	//通知策略
	NoticeOption *RuleNotice `json:"noticeOption"`
	//实例id列表
	InstanceIds []string `json:"instanceIds"`
	// 关联实例个数
	InstanceCount int64 `json:"instanceCount"`
	//实例详细信息
	Instances []*Instance `json:"instances"`
	//盘符、挂载点、网口列表
	DeviceTag string `json:"deviceTag"`
	//此规则关联的资源数
	RelatedResourceCount int64 `json:"relatedResourceCount"`
	//规则状态[1->正常，2->已禁用，3->报警]
	Status int64 `json:"status"`
	//规则状态名称
	StatusName string `json:"statusName"`
	// user_id
	UserID string `json:"userId"`
	// user_name
	UserName string `json:"userName"`
}

//单个触发条件
type RuleTrigger struct {
	// [bmp_monitor_counter, bmp_monitor_gauge]
	TableName string `json:"tableName"`
	//监控指标 cps.cpu.util
	Metric string `json:"metric"`
	//监控指标名称 CPU使用率
	MetricName string `json:"metricName"`
	//周期 5(分钟)
	Period int `json:"period"`
	//计算方式 [min max avg sum]
	Calculation string `json:"calculation"`
	//计算结果单位 [对于使用量，有Bytes,KB,MB,GB,TB，对于使用率，是%，对于连接个数，是count，对于网络包量，是pps,Kpps,Mpps,Gpps,Tpps，对于网络速率，是bps,Kbps,Mbps,Gbps,Tbps 对于负载，没有单位。。。。。。产品需详细列出]
	CalculationUnit string `json:"calculationUnit"`
	//比较方式 [> >= < <= == !=]或者[gt gte lt lte eq neq]
	Operation string `json:"operation"`
	//阈值
	Threshold float64 `json:"threshold"`
	//持续周期数 [1, 2, 3, 5, 10, 15, 30, 60]
	Times int `json:"times"`
	//告警级别 [1表示一般，2表示严重，3表示紧急]
	NoticeLevel int `json:"noticeLevel"`
	//对此rule的描述
	Description string `json:"description"`
}

//通知策略
type RuleNotice struct {
	//通知周期(分钟) [5 10 15 30 60 180 360 720 1440]
	NoticePeriod int `json:"noticePeriod"`
	// 有效时段开始时间 ["00:00:00"]
	EffectiveIntervalStart string `json:"effectiveIntervalStart"`
	// 有效时段结束时间 ["23:59:59"]
	EffectiveIntervalEnd string `json:"effectiveIntervalEnd"`
	//通知条件，可多选 [1表示报警， 2表示恢复正常]
	NoticeCondition []int `json:"noticeCondition"`
	//接收渠道，可多选 [1表示站内信， 2表示邮件]
	NoticeWay []int `json:"noticeWay"`
	//通知对象用户Id
	UserID string `json:"userId"`
	//通知对象用户
	UserName string `json:"userName"`
}
