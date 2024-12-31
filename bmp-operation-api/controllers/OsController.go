package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/OsLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type OsController struct {
	BaseController
}

// DescribeOSs ...
// swagger:route GET /oss os describeOSs
//
//     Responses:
//       200: describeOSs
//       default: ErrorResponse

func (c *OsController) DescribeOSs() {
	defer c.CatchException()
	req := &requestTypes.QueryOssRequest{
		OsName:       c.GetString("osName"),
		Architecture: c.GetString("architecture"),
		OsType:       c.GetString("osType"),
		Platform:     c.GetString("platform"),
		OsVersion:    c.GetString("osVersion"),
		IsAll:        c.GetString("isAll"),
	}

	res, err := OsLogic.DescribeOSs(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	if res.Oss == nil {
		res.Oss = make([]*response.Os, 0)
	}
	c.Res.Result = res
}
func (c *OsController) DescribeOSsFilter() {
	defer c.CatchException()
	res := BaseLogic.OssFilter
	if c.logPoints.GetPoint("language") == BaseLogic.EN_US {
		res = BaseLogic.OssFilterEn
	}
	c.Res.Result = res
}
