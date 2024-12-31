package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/idcLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type IdcController struct {
	BaseController
}

// DescribeIdcs ...
// swagger:route GET /idcs idc describeIdcs
//	DescribeIdcs 获取idc列表
//     Responses:
//       200: describeIdcs
//       default: ErrorResponse

func (c *IdcController) DescribeIdcs() {
	defer c.CatchException()
	req := &requestTypes.QueryIdcsRequest{
		Name:  c.GetString("name"),
		Level: c.GetString("level"),
		IsAll: c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	c.logPoints.Point("QueryIdcsRequest req", req)
	res, count, err := idcLogic.QueryIdcs(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.IdcList{
		Idcs:       res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}
}

// swagger:route GET /idcs/{idc_id} idc describeIdc
// DescribeIdc 获取idc详情
//     Responses:
//       200: describeIdc
//       default: ErrorResponse

func (c *IdcController) DescribeIdc() {
	defer c.CatchException()
	req := &requestTypes.QueryIdcRequest{
		IDcID: c.Ctx.Input.Param(":idc_id"),
	}
	req.Validate(c.logPoints)
	c.logPoints.Point("QueryIdcRequest req", req)
	res, err := idcLogic.QueryIdcById(c.logPoints, req.IDcID)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res
}
