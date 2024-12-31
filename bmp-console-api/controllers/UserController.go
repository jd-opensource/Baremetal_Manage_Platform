package controllers

import (
	"encoding/json"
	"fmt"

	UserLoigc "coding.jd.com/aidc-bmp/bmp-console-api/logic/UserLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

type UserController struct {
	BaseController
}

// Login 登录
// swagger:route POST /login user login
// Login 登录
//     Responses:
//       200: login
//       default: ErrorResponse

func (c *UserController) Login() {

	// defer c.CatchException()
	req := &requestTypes.LoginRequest{}
	// req.Validate()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	_, err := UserLoigc.VerifyUser(c.logPoints, req.Username, req.Password)

	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}

	jdc := c.Ctx.GetCookie(sessionName)
	if err := c.SetSession(jdc, req.Username); err != nil {
		c.logPoints.Warnf("SetSession error, session_id:%s, username:%s, error:%s", jdc, req.Username, err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Ctx.Output.Header("X-Jdcloud-Username", req.Username)
	c.logPoints.Infof("SetSession success, session_id:%s, username:%s:", jdc, req.Username)
	c.Res.Result = responseTypes.CommonResponse{
		Success: true,
	}
}

// Logout 登出
// swagger:route POST /logout user logout
// Logout 登出
//     Responses:
//       200: logout
//       default: ErrorResponse

func (c *UserController) Logout() {
	//username := c.GetString("username")

	// fmt.Println("所有cookie：", c.Ctx.Request.Cookies())
	jdc := c.Ctx.GetCookie(sessionName)
	username := c.GetSession(jdc)

	if err := c.DelSession(jdc); err != nil {
		c.logPoints.Warnf("del session error, session_id:%s, username:%s, error:%s", jdc, username, err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.logPoints.Infof("del session success, session_id:%s, username:%s", jdc, username)
	c.Res.Result = responseTypes.CommonResponse{
		Success: true,
	}
}

// swagger:route GET /user user getUserInfo
//GetUserInfo 获取用户详情
//     Responses:
//       200: getUserInfo
//       default: ErrorResponse

func (c *UserController) GetUserInfo() {
	defer c.CatchException()
	result, err := UserLoigc.GetUserInfo(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = result
}

// swagger:route PUT /user/password user updatePassword
// UpdatePassword 修改密码
//     Responses:
//       200: updatePassword
//       default: ErrorResponse

func (c *UserController) UpdatePassword() {
	req := &requestTypes.UpdatePasswordRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	userId := c.logPoints.GetPoint("userId").(string)
	res, err := UserLoigc.ModifyUserPassword(c.logPoints, userId, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	//修改完成后登出
	if res {
		jdc := c.Ctx.GetCookie(sessionName)
		username := c.GetSession(jdc)

		if err := c.DelSession(jdc); err != nil {
			c.logPoints.Warnf("UpdatePassword del session error, session_id:%s, username:%s, error:%s", jdc, username, err.Error())
		} else {
			c.Ctx.ResponseWriter.WriteHeader(402)
			fmt.Fprintln(c.Ctx.ResponseWriter, "请重新登录")
			return
		}
	}
	c.Res.Result = responseTypes.CommonResponse{
		Success: res,
	}

}

// swagger:route PUT /user user updateUserInfo
// UpdateUserInfo 修改个人信息(除密码)
//     Responses:
//       200: updateUserInfo
//       default: ErrorResponse

func (c *UserController) UpdateUserInfo() {
	userId := c.logPoints.GetPoint("userId").(string)
	req := &requestTypes.ModifyUserInfoRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := UserLoigc.ModifyUserInfo(c.logPoints, userId, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = responseTypes.CommonResponse{
		Success: res,
	}

}

// swagger:route GET /user/getConsoleAccessUserList user getConsoleAccessUserList
// GetConsoleAccessUserList 获取有控制台权限的用户列表
//     Responses:
//       200: getConsoleAccessUserList
//       default: ErrorResponse
func (c *UserController) GetConsoleAccessUserList() {
	userId := c.logPoints.GetPoint("userId").(string)

	res, err := UserLoigc.GetConsoleAccessUserList(c.logPoints, userId)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = res

}

// swagger:route GET /user/checkUserConsoleAccess user checkUserConsoleAccess
// CheckUserConsoleAccess 校验被转移的用户有没有控制台权限
//     Responses:
//       200: checkUserConsoleAccess
//       default: ErrorResponse
func (c *UserController) CheckUserConsoleAccess() {

	userId := c.logPoints.GetPoint("userId").(string)
	req := requestTypes.CheckUserConsoleAccessRequest{
		CheckedUserId: c.GetString("checkedUserId"),
	}
	res, err := UserLoigc.CheckUserConsoleAccess(c.logPoints, userId, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = responseTypes.CommonResponse{
		Success: res,
	}
}

// swagger:route GET /user/checkUserConsoleAccessByName user checkUserConsoleAccessByName
// CheckUserConsoleAccess 按用户名校验被转移的用户有没有控制台权限
//     Responses:
//       200: checkUserConsoleAccessByName
//       default: ErrorResponse
func (c *UserController) CheckUserConsoleAccessByName() {

	userId := c.logPoints.GetPoint("userId").(string)
	req := requestTypes.CheckUserConsoleAccessByNameRequest{
		CheckedUserName: c.GetString("checkedUserName"),
		ProjectID:       c.GetString("projectId"),
		Operation:       c.GetString("operation"),
	}
	res, err := UserLoigc.CheckUserConsoleAccessByName(c.logPoints, userId, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = responseTypes.CommonResponse{
		Success: res,
	}
}
