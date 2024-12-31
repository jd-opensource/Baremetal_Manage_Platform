package prometheusApi


type BasePrometheusResponse struct {
	Status     string         `json:"status"`
	ErrorMsg   string		  `json:"error"`
	ErrorType  string		  `json:"errorType"`
	Data      interface{}     `json:"data"`
}

type MetricDataListResponse struct {
	BasePrometheusResponse
	RequestId string     `json:"requestId"`
	Data *MetricDataInfo `json:"data"`
}


type MetricDataInfo struct {
	ResultType string          `json:"resultType"`
	Result []*MetricDataDetail `json:"result"`
}

type MetricDataDetail struct {
	Metric map[string]string  `json:"metric"`
	Values [][]interface{}    `json:"values"`
}
