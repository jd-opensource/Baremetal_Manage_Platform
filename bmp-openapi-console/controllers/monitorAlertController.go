package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorAlertLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// MonitorAlertController 告警
type MonitorAlertController struct {
	BaseController
}

// swagger:route POST /monitorAlert/addAlert monitorAlert addAlert
//
// AddAlert 添加(接收)告警
//
//     Responses:
//       200: addAlert
//       default: ErrorResponse
func (c *MonitorAlertController) AddAlert() {
	defer c.CatchException()
	c.logPoints.Infof("AddAlert body:%s", string(c.Ctx.Input.RequestBody))

	body := string(c.Ctx.Input.RequestBody)
	//赛迪改成trigger单独urlencode
	// body, err := url.QueryUnescape(body)
	// if err != nil {
	// 	c.logPoints.Warn("AddAlert.QueryUnescape error:", err.Error())
	// 	c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
	// 	return
	// }
	req := &request.AddAlertRequest{}
	if err := json.Unmarshal([]byte(body), req); err != nil {
		c.logPoints.Warn("AddAlert parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	for _, v := range req.Alerts {
		c.logPoints.Infof("alert.trigger is:%s", v.Trigger)
	}

	res, err := monitorAlertLogic.AddAlert(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}

// swagger:route GET /monitorAlert/describeAlert monitorAlert describeAlert
//
// DescribeAlert 告警详情
//
//     Responses:
//       200: describeAlert
//       default: ErrorResponse
func (c *MonitorAlertController) DescribeAlert() {
	defer c.CatchException()
	req := request.DescribeAlertRequest{
		AlertID: c.GetString("alertId"),
	}
	req.Validate(c.logPoints)
	res, err := monitorAlertLogic.DescribeAlert(c.logPoints, req.AlertID)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route GET /monitorAlert/describeAlerts monitorAlert describeAlerts
//
// DescribeAlerts 告警列表
//
//     Responses:
//       200: describeAlerts
//       default: ErrorResponse
func (c *MonitorAlertController) DescribeAlerts() {
	defer c.CatchException()

	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}

	stime, _ := c.GetInt64("startTime")
	etime, _ := c.GetInt64("endTime")
	req := &request.DescribeAlertsRequest{
		IsAll:      c.GetString("isAll"),
		UserName:   c.GetString("userName"),
		UserID:     c.GetString("userId"),
		RuleName:   c.GetString("ruleName"),
		RuleID:     c.GetString("ruleId"),
		ResourceID: c.GetString("resourceId"),
		StartTime:  stime,
		EndTime:    etime,
		ProjectID:  c.GetString("projectId"),
	}
	req.Validate(c.logPoints)

	res, err := monitorAlertLogic.DescribeAlerts(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}

// swagger:route DELETE /monitorAlert/deleteAlert monitorAlert deleteAlert
//
// DeleteAlert 删除告警
//
//     Responses:
//       200: deleteAlert
//       default: ErrorResponse
func (c *MonitorAlertController) DeleteAlert() {
	defer c.CatchException()
	c.logPoints.Infof("DeleteAlert body:%s", string(c.Ctx.Input.RequestBody))
	req := &request.DeleteAlertRequest{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("DeleteAlert parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	res, err := monitorAlertLogic.DeleteAlert(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}
}
