package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/osLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type OsController struct {
	BaseController
}

// CreateOs ...
// swagger:route POST /oss os createOS
// CreateOS 添加操作系统(暂不启用)
//
//     Responses:
//       200: createOS
//       default: ErrorResponse

func (c *OsController) CreateOS() {
	defer c.CatchException()
	req := &requestTypes.CreateOSRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateOs parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := osLogic.CreateOS(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.OsId{
		OsId: uuid,
	}
}

// DescribeOS ...
// swagger:route GET /oss/{os_id} os describeOS
// DescribeOS 获取os系统详情(暂不启用)
//
//     Responses:
//       200: describeOS
//       default: ErrorResponse

func (c *OsController) DescribeOS() {
	defer c.CatchException()
	osId := c.Ctx.Input.Param(":os_id")
	res, err := osLogic.GetByOsId(c.logPoints, osId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.OsInfo{
		Os: res,
	}
}

// ModifyOS ...
// swagger:route PUT /oss/{os_id} os modifyOS
//
// ModifyOS 修改os信息(暂不启用)
//     Responses:
//       200: modifyOS
//       default: ErrorResponse

func (c *OsController) ModifyOS() {
	defer c.CatchException()
	os_id := c.Ctx.Input.Param(":os_id")
	req := &requestTypes.ModifyOSRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyOS parse post body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := osLogic.ModifyOS(c.logPoints, req, os_id); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteOS ...
// swagger:route DELETE /oss/{os_id} os deleteOS
//
// DeleteOS 删除os(暂不启用)
//     Responses:
//       200: deleteOS
//       default: ErrorResponse

func (c *OsController) DeleteOS() {
	defer c.CatchException()
	osId := c.Ctx.Input.Param(":os_id")
	if err := osLogic.DeleteOS(c.logPoints, osId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DescribeOSs ...
// swagger:route GET /oss os describeOSs
// DescribeOSs 获取os列表(暂不启用)
//
//     Responses:
//       200: describeOSs
//       default: ErrorResponse

func (c *OsController) DescribeOSs() {
	defer c.CatchException()
	//ids := strings.Split(c.GetString("ids"), ",")
	req := &requestTypes.QueryOssRequest{
		OsName:    c.GetString("osName"),
		OsType:    c.GetString("osType"),
		OsVersion: c.GetString("osVersion"),
		IsAll:     c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	res, err := osLogic.DescribeOSs(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.OsList{
		Oss: res,
		//PageNumber: c.pageable.PageNumber,
		//PageSize:   c.pageable.PageSize,
		//TotalCount: count,
	}

}
