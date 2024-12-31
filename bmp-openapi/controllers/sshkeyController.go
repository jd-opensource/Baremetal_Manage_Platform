package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/sshkeyLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// SshkeyController operations for sshkey
type SshkeyController struct {
	BaseController
}

// CreateSshkey ...
// swagger:route POST /user/sshkeys sshkey createUserSshkey
//
// CreateUserSshkey 创建个人sshkey
//     Responses:
//       200: createUserSshkey
//       default: ErrorResponse

func (c *SshkeyController) CreateUserSshkey() {
	defer c.CatchException()
	req := &requestTypes.CreateSshkeyRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateSshkey parse pSshkeyt body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	uuid, err := sshkeyLogic.CreateSshkey(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.SshkeyId{
		SshkeyId: uuid,
	}
}

// QueryBySshkeyIds ...
// swagger:route GET /user/sshkeys sshkey describeUserSshKeys
//
// DescribeUserSshKeys 获取个人sshkey列表
//     Responses:
//       200: describeUserSshKeys
//       default: ErrorResponse

func (c *SshkeyController) DescribeUserSshKeys() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QuerySshkeysRequest{
		Name:  c.GetString("name"),
		IsAll: c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	res, count, err := sshkeyLogic.GetSshkeyList(c.logPoints, *req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.SshkeyList{
		Sshkeys:    res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// GetInstancesBySshkey ...
// swagger:route GET /user/sshkeys/instances/{sshkey_id} sshkey getInstancesBySshkey
// GetInstancesBySshkey 根据sshkey获取实例列表
//
//     Responses:
//       200: getInstancesBySshkey
//       default: ErrorResponse

func (c *SshkeyController) GetInstancesBySshkey() {
	defer c.CatchException()
	//req := &requestTypes.QueryInstancesSshkeyRequest{
	//	SshkeyId: c.Ctx.Input.Param(":sshkey_id"),
	//}
	//req.Validate(c.logPoints)
	ids := sshkeyLogic.GetInstancesBySshkey(c.logPoints, c.Ctx.Input.Param(":sshkey_id"))
	c.Res.Result = response.InstancesSshkeyInfoResponse{
		InstanceIds: ids,
	}

}

// DescribeSshkey ...
// swagger:route GET /user/sshkeys/{sshkey_id} sshkey describeUserSshKey
//
// DescribeUserSshKey 获取sshkey详情
//     Responses:
//       200: describeUserSshKey
//       default: ErrorResponse

func (c *SshkeyController) DescribeUserSshKey() {
	defer c.CatchException()
	sshkeyId := c.Ctx.Input.Param(":sshkey_id")
	res, err := sshkeyLogic.GetSshkeyById(c.logPoints, sshkeyId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.SshkeyInfo{
		Sshkey: res,
	}
}

// ModifySshkey ...
// swagger:route PUT /user/sshkeys/{sshkey_id} sshkey modifyUserSshkey
//
// ModifyUserSshkey 修改sshkey
//     Responses:
//       200: modifyUserSshkey
//       default: ErrorResponse

func (c *SshkeyController) ModifyUserSshkey() {
	defer c.CatchException()
	sshkeyId := c.Ctx.Input.Param(":sshkey_id")
	req := &requestTypes.ModifySshkeyRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifySshkey parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := sshkeyLogic.ModifySshkey(c.logPoints, req, sshkeyId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// DeleteSshkey ...
// swagger:route DELETE /user/sshkeys/{sshkey_id} sshkey deleteUserSshkey
// DeleteUserSshkey 删除sshkey
//
//     Responses:
//       200: deleteUserSshkey
//       default: ErrorResponse

func (c *SshkeyController) DeleteUserSshkey() {
	defer c.CatchException()
	sshkeyId := c.Ctx.Input.Param(":sshkey_id")
	if err := sshkeyLogic.DeleteSshkey(c.logPoints, sshkeyId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}
