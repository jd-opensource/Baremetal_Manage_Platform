package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/apikeyLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type ApikeyController struct {
	BaseController
}

// CreateAPIkey ...
// swagger:route POST /user/apikeys apiKey createUserApikey
//
// CreateUserApikey 创建apikey
//
//     Responses:
//       200: createUserApikey
//       default: ErrorResponse

func (c *ApikeyController) CreateUserApikey() {
	defer c.CatchException()
	req := &requestTypes.CreateApikeyRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateAPIkey parse pAPIkeyt body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := apikeyLogic.CreateApikey(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ApikeyId{
		ApikeyId: uuid,
	}
}

// QueryByAPIkeyIds ...
// swagger:route GET /user/apikeys apiKey describeUserAPIKeys
//
// DescribeUserAPIKeys 获取APIKey列表
//
//     Responses:
//       200: describeUserAPIKeys
//       default: ErrorResponse
func (c *ApikeyController) DescribeUserAPIKeys() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryApikeysRequest{
		Name:  c.GetString("name"),
		IsAll: c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	res, count, err := apikeyLogic.GetApikeyList(c.logPoints, *req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ApikeyList{
		Apikeys:    res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// DescribeUserAPIKey ...
// swagger:route GET /user/apikeys/{apikey_id} apiKey describeUserAPIKey
//
// DescribeUserAPIKey 获取apikey详情
//
//     Responses:
//       200: describeUserAPIKey
//       default: ErrorResponse
func (c *ApikeyController) DescribeUserAPIKey() {
	defer c.CatchException()
	apikeyId := c.Ctx.Input.Param(":apikey_id")
	res, err := apikeyLogic.GetApikeyById(c.logPoints, apikeyId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	c.Res.Result = response.ApikeyInfo{
		Apikey: res,
	}
}

// ModifyAPIkey ...
// swagger:route PUT /user/apikeys/{apikey_id} apiKey modifyUserApikey
//
// ModifyUserApikey 修改apikey(暂不启用)
//
//     Responses:
//       200: modifyUserApikey
//       default: ErrorResponse

func (c *ApikeyController) ModifyUserApikey() {
	defer c.CatchException()
	apikeyId := c.Ctx.Input.Param(":apikey_id")
	req := &requestTypes.ModifyApikeyRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyAPIkey parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := apikeyLogic.ModifyApikey(c.logPoints, req, apikeyId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteAPIkey ...
// swagger:route DELETE /user/apikeys/{apikey_id} apiKey deleteUserApikey
//
// DeleteUserApikey 删除某个apikey
//
//     Responses:
//       200: deleteUserApikey
//       default: ErrorResponse

func (c *ApikeyController) DeleteUserApikey() {
	defer c.CatchException()
	apikeyId := c.Ctx.Input.Param(":apikey_id")
	if err := apikeyLogic.DeleteApikey(c.logPoints, apikeyId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
