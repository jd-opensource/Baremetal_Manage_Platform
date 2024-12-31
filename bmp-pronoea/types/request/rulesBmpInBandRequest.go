package request

type RulesRequest struct {
	BaseRequest
	Source        string    	 `json:"source" valid:"Required"`        //规则来源
	Rules         []*RulesDetail `json:"rules" valid:"Required"`         //规则来源
}

type RulesDetail struct {
	BaseRequest
	RuleId        string         `json:"ruleId" valid:"Required"`        //规则uuid
	RuleName      string         `json:"ruleName" valid:"Required"`      //规则名称
	Dimension     string         `json:"dimension"`                      //维度 [instance、disk、mountpoint、nic]
	DimensionName string         `json:"dimensionName"`                  //维度名称 [实例、盘符、挂载点、网卡]
	Resource      string         `json:"resource"`                       //资源类型,目前只支持instance
	ResourceName  string         `json:"resourceName"`                   //资源类型名称 实例
	TriggerOption []*RuleTrigger `json:"triggerOption" valid:"Required"` //触发条件
	NoticeOption  *RuleNotice    `json:"noticeOption" valid:"Required"`  //通知策略
	InstanceIds   []string       `json:"instanceIds" valid:"Required"`   //实例id列表
	DeviceTag     string         `json:"deviceTag"`                      //盘符、挂载点、网口列表
}

// RuleTrigger 单个触发条件
type RuleTrigger struct {
	BaseRequest
	TableName       string  `json:"tableName" valid:"Required"`
	Metric          string  `json:"metric" valid:"Required"`      //监控指标 cps.cpu.util
	Period          int     `json:"period" valid:"Required"`      //周期 5(分钟)
	Calculation     string  `json:"calculation" valid:"Required;Match(/^(min|max|avg|sum)$/)"` //计算方式 [min max avg sum]
	CalculationUnit string  `json:"calculationUnit"`              //计算结果单位 [对于使用量，有Bytes,KB,MB,GB,TB，对于使用率，是%，对于连接个数，是count，对于网络包量，是pps,Kpps,Mpps,Gpps,Tpps，对于网络速率，是bps,Kbps,Mbps,Gbps,Tbps 对于负载，没有单位。。。。。。产品需详细列出]
	Operation       string  `json:"operation" valid:"Required;Match(/^(>|>=|<|<=|==|!=|gt|gte|lt|lte|eq|neq)$/)"`   //比较方式 [> >= < <= == !=]或者[gt gte lt lte eq neq]
	Threshold       float64 `json:"threshold" valid:"Required"`   //阈值
	Times           int     `json:"times" valid:"Required;Match(/^(1|2|3|5|10|15|30|60)$/)"`       //持续周期数 [1, 2, 3, 5, 10, 15, 30, 60]
	NoticeLevel     int     `json:"noticeLevel" valid:"Required"` //告警级别 [1表示一般，2表示严重，3表示紧急]
}

// RuleNotice 通知策略
type RuleNotice struct {
	BaseRequest
	NoticePeriod           int    `json:"noticePeriod" valid:"Required;Match(/^(5|10|15|30|60|180|360|720|1440)$/)"`           //通知周期(分钟) [5 10 15 30 60 180 360 720 1440]
	/*
	EffectiveIntervalStart string `json:"effectiveIntervalStart"` // 有效时段开始时间 ["00:00:00"]
	EffectiveIntervalEnd   string `json:"effectiveIntervalEnd"`   // 有效时段结束时间 ["23:59:59"]
	NoticeCondition        []int  `json:"noticeCondition"`        //通知条件，可多选 [1表示报警， 2表示恢复正常]
	NoticeWay              []int  `json:"noticeWay"`              //接收渠道，可多选 [1表示站内信， 2表示邮件]
	UserID                 string `json:"userId"`                 //通知对象
	 */
}



