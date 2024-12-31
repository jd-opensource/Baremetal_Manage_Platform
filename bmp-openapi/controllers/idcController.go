package controllers

import (
	"encoding/json"
	"path"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/idcLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"

	deviceLogic "coding.jd.com/aidc-bmp/bmp-openapi/logic/deviceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type IdcController struct {
	BaseController
}

func (c *IdcController) CreateDevicesFromIdc() {
	defer c.CatchException()
	req := &requestTypes.CreateIdcDevicesRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateDevicesFromIdc parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.logPoints.Point("createDevicesFromIdc sns", req)
	if err := deviceLogic.SaveDevicesFromIdc(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// ImportBms ...
func (c *IdcController) DeviceBatchImport() {
	defer c.CatchException()
	file, information, err := c.GetFile("file")
	if err != nil {
		c.logPoints.Warn("ImportBms getfile error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
	defer file.Close() //关闭上传的文件，否则出现临时文件不清除的情况
	filename := information.Filename
	if !util.InArray(path.Ext(filename), []string{".xls", ".xlsx"}) {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, "only support excel", errorCode.INVALID_ARGUMENT)
		return
	}
	if err := deviceLogic.ImportBms(c.logPoints, file, information.Size); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// CreateIdc ...
// swagger:route POST /idcs idc createIdc
// CreateIdc 创建idc(接口废弃，不对外开放)
//     Responses:
//       200: createIdc
//       default: ErrorResponse

func (c *IdcController) CreateIdc() {
	defer c.CatchException()
	req := &requestTypes.CreateIdcRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateIdc parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	c.logPoints.Point("createIdc req:", req)
	idcId, err := idcLogic.CreateIdc(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.IdcID{
		IdcID: idcId,
	}
}

// ModifyIdc ...
// swagger:route PUT /idcs/{idc_id} idc modifyIdc
// ModifyIdc 修改idc信息
//     Responses:
//       200: modifyIdc
//       default: ErrorResponse

func (c *IdcController) ModifyIdc() {
	defer c.CatchException()
	req := &requestTypes.ModifyIdcRequest{}
	idcId := c.Ctx.Input.Param(":idc_id")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyIdc parse post body err:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := idcLogic.ModifyIdc(c.logPoints, req, idcId); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
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

// swagger:route DELETE /idcs/{idc_id} idc deleteIdc
// DeleteIdc 删除idc
//     Responses:
//       200: deleteIdc
//       default: ErrorResponse

func (c *IdcController) DeleteIdc() {
	defer c.CatchException()
	req := &requestTypes.QueryIdcRequest{
		IDcID: c.Ctx.Input.Param(":idc_id"),
	}
	req.Validate(c.logPoints)
	if err := idcLogic.DeleteIdc(c.logPoints, req.IDcID); err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
