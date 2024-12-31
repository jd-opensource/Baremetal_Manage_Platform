package controllers

import (
	"encoding/json"

	commandLogic "coding.jd.com/aidc-bmp/bmp-openapi/logic/commandLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceController operations for Device
type CommandController struct {
	BaseController
}

func (c *CommandController) Query() {
	// defer c.CatchException()
	req := &requestTypes.QueryCommandsRequest{}
	req.RequestId = c.GetString("request_id") //用户请求带过来的参数，不是header里面标识流量的request_id
	req.Sn = c.GetString("sn")
	req.InstanceId = c.GetString("instance_id")
	data, count, err := commandLogic.Query(c.logPoints, req, c.pageable)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.QueryCommandsResult{
		Commands:   data,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

func (c *CommandController) CancelCommand() {
	// defer c.CatchException()
	req := &requestTypes.CancelCommandsRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := commandLogic.CancelCommand(c.logPoints, req.Sn); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

func (c *CommandController) RetryCommand() {
	defer c.CatchException()
	req := &requestTypes.RetryCommandRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := commandLogic.RetryCommand(c.logPoints, *req.OffsetCommandId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
