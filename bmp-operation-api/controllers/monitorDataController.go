package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/monitorDataLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// MonitorDataController 读取监控数据
type MonitorDataController struct {
	BaseController
}

// swagger:route GET /monitorData monitorData getMonitorData
//
// GetMonitorData 获取图表监控数据
//
//     Responses:
//       200: getMonitorData
//       default: ErrorResponse
func (c *MonitorDataController) GetMonitorData() {

	stime, _ := c.GetInt64("startTime")
	etime, _ := c.GetInt64("endTime")
	ti, _ := c.GetInt64("timeInterval")
	lastManyTime, _ := c.GetInt64("lastManyTime")
	req := requestTypes.GetMonitorDataRequest{
		MetricName:     c.GetString("metricName"),
		InstanceID:     c.GetString("instanceId"),
		Device:         c.GetString("device"),
		StartTime:      stime,
		EndTime:        etime,
		TimeInterval:   ti,
		LastManyTime:   lastManyTime,
		DownSampleType: c.GetString("downSampleType"),

		UserName: c.GetString("userName"),
		IdcID:    c.GetString("idcID"),
	}

	res, err := monitorDataLogic.GetMonitorData(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res

}
