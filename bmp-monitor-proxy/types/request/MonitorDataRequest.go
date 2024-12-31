package request

import log "git.jd.com/cps-golang/log"

/*
{"params":{"serviceCode":"cps-vpc","resourceId":"cps-2qtflhpvmuuw5tdodatv6ofxqcek","filters":[{"name":"metric","values":["cps.process.total"]}],"startTime":"2024-06-09T00:00:00","endTime":"2024-06-12T23:59:59","timeInterval":"1h", "downSampleType":"avg"}}
*/
type GetMonitorDataRequest struct {
	MetricName     string `json:"metricName" validate:"required"`
	InstanceID     string `json:"instanceId" validate:"required"`
	Device         string `json:"device"`
	StartTime      int64  `json:"startTime"`
	EndTime        int64  `json:"endTime"`
	TimeInterval   int    `json:"timeInterval"`   //s为单位
	DownSampleType string `json:"downSampleType"` //好像不支持
}

func (req *GetMonitorDataRequest) Validate(logger *log.Logger) {

	validate(req, logger)

}
