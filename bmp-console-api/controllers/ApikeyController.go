package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/ApikeyLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	response "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type ApikeyController struct {
	BaseController
}

// swagger:route PUT /apikey keypair createApikey
// CreateAPIkey 创建apikey
//     Responses:
//       200: createApikey
//       default: ErrorResponse

func (c *ApikeyController) CreateApikey() {
	req := &sdkModels.CreateApikeyRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateAPIkey parse pAPIkeyt body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	// req.Validate(c.logPoints)
	uuid, err := ApikeyLogic.CreateApikey(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.ApikeyId{
		ApikeyId: uuid,
	}
}

// swagger:route GET /apikey keypair getApikeyList
// GetApikeyList 获取apikey列表
//     Responses:
//       200: getApikeyList
//       default: ErrorResponse

func (c *ApikeyController) GetApikeyList() {
	defer c.CatchException()

	req := requestTypes.GetApikeyListRequest{
		PagingRequest: requestTypes.PagingRequest{
			PageNumber: c.getPageNumber(),
			PageSize:   c.getPageSize(),
		},
	}

	// req.Validate(c.logPoints)
	res, err := ApikeyLogic.GetApikeyList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = res

}

// swagger:route DELETE /apikey/{apikey_id} keypair deleteApikey
// DeleteApikey 删除apikey
//     Responses:
//       200: deleteApikey
//       default: ErrorResponse

func (c *ApikeyController) DeleteApikey() {
	defer c.CatchException()
	apikeyId := c.Ctx.Input.Param(":apikey_id")
	if err := ApikeyLogic.DeleteApikey(c.logPoints, apikeyId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
