package controllers

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/userLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type UserController struct {
	BaseController
}

// swagger:route GET /users user describeUsers
// DescribeUsers 获取用户列表
//     Responses:
//       200: describeUsers
//       default: ErrorResponse

func (c *UserController) DescribeUsers() {
	defer c.CatchException()
	p := util.Pageable{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryUsersRequest{
		UserName:         c.GetString("userName"),
		RoleID:           c.GetString("roleId"),
		DefaultProjectID: c.GetString("defaultProjectId"),
		IsAll:            c.GetString("isAll"),
	}
	req.Validate(c.logPoints)
	res, count, err := userLogic.GetUserList(c.logPoints, req, p)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.UserList{
		Users:      res,
		PageNumber: c.pageable.PageNumber,
		PageSize:   c.pageable.PageSize,
		TotalCount: count,
	}

}

// swagger:route GET /users/{user_id} user describeUser
// DescribeUser 获取用户详情
//     Responses:
//       200: describeUser
//       default: ErrorResponse

func (c *UserController) DescribeUser() {
	defer c.CatchException()
	userId := c.Ctx.Input.Param(":user_id")
	res, err := userLogic.GetUserById(c.logPoints, userId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.UserInfo{
		User: res,
	}
}

// GetUserByName ...
// swagger:route GET /users/getUserByName user describeUserByName
// DescribeUserByName 根据用户名获取用户详情
//     Responses:
//       200: describeUserByName
//       default: ErrorResponse
func (c *UserController) DescribeUserByName() {
	defer c.CatchException()
	req := &requestTypes.GetUserByNameRequest{
		UserName: c.GetString("userName"),
	}
	req.Validate(c.logPoints)

	res, err := userLogic.GetUserByName(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.UserInfo{
		User: res,
	}
}

// ModifyUser ...
// swagger:route PUT /users/{user_id} user modifyUser
// ModifyUser 修改用户信息
//     Responses:
//       200: modifyUser
//       default: ErrorResponse

func (c *UserController) ModifyUser() {
	defer c.CatchException()
	userId := c.Ctx.Input.Param(":user_id")
	req := &requestTypes.ModifyUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyUser parse pUsert body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := userLogic.ModifyUser(c.logPoints, req, userId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// VerifyUser ...
// swagger:route POST /users/verify user verifyUser
//
// VerifyUser 鉴定用户
//
//     Responses:
//       200: verifyUser
//       default: ErrorResponse

func (c *UserController) VerifyUser() {
	defer c.CatchException()
	req := &requestTypes.VerifyUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("CreateUser parse pUsert body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	isTrue := userLogic.VerifyUser(c.logPoints, req)

	c.Res.Result = response.CommonResponse{
		Success: isTrue,
	}
}

// DescribeLocalUser ...
// swagger:route GET /local/users user describeLocalUser
// DescribeLocalUser 控制台获取用户详情
//     Responses:
//       200: describeLocalUser
//       default: ErrorResponse

func (c *UserController) DescribeLocalUser() {
	defer c.CatchException()
	userId := c.logPoints.GetPoint("userId").(string)
	res, err := userLogic.GetUserById(c.logPoints, userId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.UserInfo{
		User: res,
	}
}

// swagger:route PUT /local/users user modifyLocalUser
//
// ModifyLocalUser 控制台修改除密码外的个人信息
//     Responses:
//       200: modifyLocalUser
//       default: ErrorResponse

func (c *UserController) ModifyLocalUser() {
	defer c.CatchException()
	req := &requestTypes.ModifyLocalUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyUser parse pUsert body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := userLogic.ModifyLocalUser(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}

// swagger:route PUT /local/users/password user modifyLocalUserPassword
//
// ModifyLocalUserPassword 控制台修改个人密码
//     Responses:
//       200: modifyLocalUserPassword
//       default: ErrorResponse

func (c *UserController) ModifyLocalUserPassword() {
	defer c.CatchException()
	req := &requestTypes.ModifyUserPasswordRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyLocalUserPassword parse body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	req.Validate(c.logPoints)
	if err := userLogic.ModifyUserPassword(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}

}
