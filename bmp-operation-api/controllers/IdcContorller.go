package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	IdcLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/IdcLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type IdcController struct {
	BaseController
}

// GetIdcList【机房管理】【机房列表】
// swagger:route GET /idc idc getIdcList
// GetIdcList【机房管理】【机房列表】
//     Responses:
//       200: getIdcList
//       default: ErrorResponse

func (c *IdcController) GetIdcList() {

	fmt.Println("IdcController.GetIdcList traffic...")
	// defer c.CatchException()
	page := requestTypes.PagingRequest{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryIdcListRequest{
		Name:          c.GetString("name"),
		Level:         c.GetString("level"),
		IsAll:         c.GetString("isAll"),
		Show:          c.GetString("show"),
		PagingRequest: page,
	}
	result, err := IdcLogic.GetIdcList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := IdcLogic.ExportIdc(c.logPoints, result.Idcs)
		if err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "导出excel错误"+err.Error(), errorCode.INTERNAL)
			return
		}
		c.Ctx.Output.Download(fileName, downloadFileName)
		if err := os.Remove(fileName); err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "删除上传文件失败"+err.Error(), errorCode.INTERNAL)
			return
		}
		return

	}
	c.Res.Result = result
}

// GetIdcInfo【机房管理】【机房详情】
// swagger:route GET /idc/{idc_id} idc getIdcInfo
// GetIdcInfo【机房管理】【机房详情】
//     Responses:
//       200: getIdcInfo
//       default: ErrorResponse

func (c *IdcController) GetIdcInfo() {
	defer c.CatchException()
	idcID := c.Ctx.Input.Param(":idc_id")
	req := &requestTypes.QueryIdcRequest{
		IdcID: idcID,
		Show:  c.GetString("show"),
	}
	result, err := IdcLogic.GetIdcInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	c.Res.Result = result
}

// DeleteIdc【机房管理】【删除机房】
// swagger:route DELETE /idc/{idc_id} idc deleteIdc
// DeleteIdc【机房管理】【删除机房】
//     Responses:
//       200: deleteIdc
//       default: ErrorResponse

func (c *IdcController) DeleteIdc() {
	defer c.CatchException()
	idcId := c.Ctx.Input.Param(":idc_id")
	req := &requestTypes.DeleteIdcRequest{
		IdcID: idcId,
	}
	result, err := IdcLogic.DeleteIdc(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// CreateIdc【机房管理】【添加机房】
// swagger:route PUT /idc idc createIdc
// CreateIdc【机房管理】【添加机房】
//     Responses:
//       200: createIdc
//       default: ErrorResponse

func (c *IdcController) CreateIdc() {
	defer c.CatchException()
	req := &requestTypes.CreateIdcRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	result, err := IdcLogic.CreateIdc(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// ModifyIdcInfo【机房管理】【编辑机房】
// swagger:route POST /idc/{idc_id} idc modifyIdcInfo
// ModifyIdcInfo【机房管理】【编辑机房】
//     Responses:
//       200: modifyIdcInfo
//       default: ErrorResponse

func (c *IdcController) ModifyIdcInfo() {
	defer c.CatchException()
	idcId := c.Ctx.Input.Param(":idc_id")
	req := &requestTypes.ModifyIdcRequest{
		IdcID: idcId,
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	result, err := IdcLogic.ModifyIdcInfo(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// VerifyUser【机房管理】【验证用户名密码】
// swagger:route POST /idc/user/validate idc verifyUser
// VerifyUser【机房管理】【验证用户名密码】
//     Responses:
//       200: verifyUser
//       default: ErrorResponse

func (c *IdcController) VerifyUser() {
	defer c.CatchException()
	req := &requestTypes.VerifyUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	result, err := IdcLogic.VerifyUser(c.logPoints, req)
	if err != nil || !result.Success {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}
