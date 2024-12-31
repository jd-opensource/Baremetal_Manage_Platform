package controllers

import (
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
