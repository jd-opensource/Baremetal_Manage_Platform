package request

type MetricRangeQueryRequest struct {
	BaseRequest
	TableName    string                   `json:"tableName" valid:"Required"`
	SampleMethod string					  `json:"sampleMethod"`
	Labels       map[string]string        `json:"labels"`
	StartTime    int64        			  `json:"startTime" valid:"Required"`
	EndTime      int64        			  `json:"endTime" valid:"Required"`
	Step         int        			  `json:"step" valid:"Required"`
	Func         map[string]int			  `json:"func"`
}
