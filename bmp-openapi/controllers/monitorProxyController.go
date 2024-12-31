package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorProxyLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// MonitorProxyController 从monitorproxy服务读取数据，包括agent状态，tag列表等
type MonitorProxyController struct {
	BaseController
}

// swagger:route GET /monitorProxy/desrcibeAgentStatus monitorProxy desrcibeAgentStatus
//
// DesrcibeAgentStatus 获取agent状态
//
//     Responses:
//       200: desrcibeAgentStatus
//       default: ErrorResponse
func (c *MonitorProxyController) DesrcibeAgentStatus() {
	req := &request.DesrcibeAgentStatusRequest{
		InstanceID: c.GetString("instanceId"),
	}
	req.Validate(c.logPoints)
	res, err := monitorProxyLogic.DescribeAgentStatus(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res

}

// swagger:route GET /monitorProxy/desrcibeTags monitorProxy desrcibeTags
//
// DesrcibeAgentStatus 获取tag列表
//
//     Responses:
//       200: desrcibeTags
//       default: ErrorResponse
func (c *MonitorProxyController) DesrcibeTags() {
	req := &request.DesrcibeTagsRequest{
		InstanceID: c.GetString("instanceId"),
		TagName:    c.GetString("tagName"),
	}
	req.Validate(c.logPoints)
	res, err := monitorProxyLogic.DescribeDeviceTags(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
