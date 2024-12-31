package response

type GetMonitorDataResponse struct {
	MetricName string            `json:"metricName" `
	InstanceID string            `json:"instanceId" `
	Device     string            `json:"device"`
	Datas      []MonitorDataItem `json:"datas"`
}

type MonitorDataItem struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}
