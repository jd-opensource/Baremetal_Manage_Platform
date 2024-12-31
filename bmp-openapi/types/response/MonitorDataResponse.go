package response

type DataEveryMetric struct {
	Query MonitorDataQuery   `json:"query"`
	Data  []*MonitorDataItem `json:"data"`
}

type MonitorDataQuery struct {
	//[bmp_monitor_counter bmp_monitor_gauge]
	TableName string `json:"tableName"`
	// MetricName string            `json:"metricName"`
	//[rate, increase]
	SampleMethod string `json:"sampleMethod"`
	//metric都放这里了
	Labels    map[string]string `json:"labels"`
	StartTime int64             `json:"startTime"`
	EndTime   int64             `json:"endTime"`
	Step      int               `json:"step"`
	//跟赛迪约定，转换单位放这里实现
	Func map[string]int `json:"func"`
}

type MonitorDataItem struct {
	Timestamp int64  `json:"timestamp"`
	Value     string `json:"value"`
}
