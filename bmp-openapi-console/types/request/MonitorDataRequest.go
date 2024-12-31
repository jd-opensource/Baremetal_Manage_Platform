package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

/*
{"params":{"serviceCode":"cps-vpc","resourceId":"cps-2qtflhpvmuuw5tdodatv6ofxqcek","filters":[{"name":"metric","values":["cps.process.total"]}],"startTime":"2024-06-09T00:00:00","endTime":"2024-06-12T23:59:59","timeInterval":"1h", "downSampleType":"avg"}}
*/
type GetMonitorDataRequest struct {
	//指标名称，支持多个，用逗号隔离开
	MetricName string `json:"metricName" validate:"required"`
	//实例id
	InstanceID string `json:"instanceId" validate:"required"`
	//标签值
	Device string `json:"device"`
	//开始时间戳
	StartTime int64 `json:"startTime"`
	//结束时间戳
	EndTime int64 `json:"endTime"`
	// 间隔多久一个点 s为单位
	TimeInterval int `json:"timeInterval"`
	//最近多少时间，小时为单位
	LastManyTime int `json:"lastManyTime"`
	//avg, min, max, sum
	DownSampleType string `json:"downSampleType"`
}

func (req *GetMonitorDataRequest) Validate(logger *log.Logger) {

	validate(req, logger)

}
