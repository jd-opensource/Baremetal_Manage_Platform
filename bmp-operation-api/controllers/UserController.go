package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	UserLoigc "coding.jd.com/aidc-bmp/bmp-operation-api/logic/UserLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"git.jd.com/cps-golang/ironic-common/exception"
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
		panic(exception.Exception{Msg: err.Error()})
	}

	_, err := UserLoigc.VerifyUser(c.logPoints, req.Username, req.Password)

	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	jdc := c.Ctx.GetCookie(sessionName)
	if err := c.SetSession(jdc, req.Username); err != nil {
		c.logPoints.Warnf("SetSession error, session_id:%s, username:%s, error:%s", jdc, req.Username, err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, errorCode.HTTP_FORBIDDEN, err.Error())
		return
	}
	c.Ctx.Output.Header("X-Jdcloud-Username", req.Username)
	c.logPoints.Infof("SetSession success, session_id:%s, username:%s:", jdc, req.Username)
	c.Res.Result = noQueryResponse
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
		c.SetErrorResponse(httpStatus.BAD_REQUEST, errorCode.HTTP_FORBIDDEN, err.Error())
		return
	}
	c.logPoints.Infof("del session success, session_id:%s, username:%s", jdc, username)
	c.Res.Result = noQueryResponse
}

// GetUserList 获取用户列表
// swagger:route GET /users user getUserList
// GetUserList 获取用户列表
//     Responses:
//       200: getUserList
//       default: ErrorResponse

func (c *UserController) GetUserList() {
	defer c.CatchException()
	p := requestTypes.PagingRequest{
		PageSize:   c.getPageSize(),
		PageNumber: c.getPageNumber(),
	}
	req := &requestTypes.QueryUsersRequest{
		UserName:      c.GetString("userName"),
		RoleID:        c.GetString("roleId"),
		ProjectID:     c.GetString("projectId"),
		IsAll:         c.GetString("isAll"),
		PagingRequest: p,
	}
	result, err := UserLoigc.GetUserList(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}

	if c.GetString("exportType") != "" {
		fileName, downloadFileName, err := UserLoigc.ExportUserExcel(c.logPoints, result.Users)
		if err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "导出设备列表excel错误 "+err.Error(), errorCode.INTERNAL)
			return
		}
		//c.Ctx.Input.SetData("download", "yes")
		c.Ctx.Output.Download(fileName, downloadFileName)
		if err := os.Remove(fileName); err != nil {
			c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, "删除上传文件失败"+err.Error(), errorCode.INTERNAL)
			return
		}
		return
	}
	c.Res.Result = result
}

// AddUser 【用户管理】【添加用户】
// swagger:route PUT /users user addUser
// AddUser 【用户管理】【添加用户】
//     Responses:
//       200: addUser
//       default: ErrorResponse

func (c *UserController) AddUser() {
	defer c.CatchException()

	req := &requestTypes.AddUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	result, err := UserLoigc.AddUser(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// GetUserInfo 【用户管理】【用户详情】
// swagger:route GET /users/{user_id} user getUserInfo
// GetUserInfo 【用户管理】【用户详情】
//     Responses:
//       200: getUserInfo
//       default: ErrorResponse

func (c *UserController) GetUserInfo() {
	defer c.CatchException()
	userId := c.Ctx.Input.Param(":user_id")
	req := &requestTypes.QueryUserInfoRequest{
		UserID: userId,
	}
	result := UserLoigc.GetUserInfo(c.logPoints, req)
	c.Res.Result = result
}

// ModifyUserInfo 【用户管理】【编辑用户信息】
// swagger:route POST /users/{user_id} user modifyUserInfo
// ModifyUserInfo 【用户管理】【编辑用户信息】
//     Responses:
//       200: modifyUserInfo
//       default: ErrorResponse

func (c *UserController) ModifyUserInfo() {
	defer c.CatchException()
	userId := c.Ctx.Input.Param(":user_id")
	req := &requestTypes.ModifyUserInfoRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.logPoints.Warn("ModifyUser parse pUsert body error:", err.Error())
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	if err := UserLoigc.ModifyUserInfo(c.logPoints, req, userId); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// DeleteUserInfo 【用户管理】【删除用户】
// swagger:route DELETE /users/{user_id} user deleteUserInfo
// DeleteUserInfo 【用户管理】【删除用户】
//     Responses:
//       200: deleteUserInfo
//       default: ErrorResponse

func (c *UserController) DeleteUserInfo() {
	defer c.CatchException()
	userId := c.Ctx.Input.Param(":user_id")
	req := &requestTypes.DeleteUserInfoRequest{
		UserID: userId,
	}
	if err := UserLoigc.DeleteUserInfo(c.logPoints, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: true,
	}
}

// VerifyLoginUser【验证用户名密码】
// swagger:route POST /users/login/validate user verifyLoginUser
// VerifyLoginUser【验证用户名密码】
//     Responses:
//       200: verifyLoginUser
//       default: ErrorResponse

func (c *UserController) VerifyLoginUser() {
	defer c.CatchException()
	req := &requestTypes.VerifyUserRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	result, err := UserLoigc.VerifyUser(c.logPoints, req.Username, req.Password)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.ABORTED)
		return
	}
	c.Res.Result = result
}

// swagger:route GET /user/local user getLocalUserInfo
//GetUserInfo 获取当前用户详情
//     Responses:
//       200: getLocalUserInfo
//       default: ErrorResponse

func (c *UserController) GetLocalUserInfo() {
	defer c.CatchException()
	result, err := UserLoigc.GetLocalUserInfo(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = result
}

// swagger:route PUT /user/local/password user updateLocalPassword
// UpdatePassword 修改当前登录用户密码
//     Responses:
//       200: updateLocalPassword
//       default: ErrorResponse

func (c *UserController) UpdateLocalPassword() {
	req := &requestTypes.UpdateLocalPasswordRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	userId := c.logPoints.GetPoint("userId").(string)
	res, err := UserLoigc.ModifyLocalUserPassword(c.logPoints, userId, req)
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
	c.Res.Result = response.CommonResponse{
		Success: res,
	}

}

// swagger:route PUT /user/local user updateLocalUserInfo
// UpdateLocalUserInfo 修改个人信息(除密码)
//     Responses:
//       200: updateLocalUserInfo
//       default: ErrorResponse

func (c *UserController) UpdateLocalUserInfo() {
	userId := c.logPoints.GetPoint("userId").(string)
	req := &requestTypes.ModifyLocalUserInfoRequest{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}
	res, err := UserLoigc.ModifyLocalUserInfo(c.logPoints, userId, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.HTTP_FORBIDDEN)
		return
	}
	c.Res.Result = response.CommonResponse{
		Success: res,
	}

}
