package UserLoigc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkProjectParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/project"
	sdkUserParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/user"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

const CONSOLE_ROLE_ID = "role-user-uuid-01"
const ADMIN_ROLE_ID = "role-admin-uuid-01"

func GetUserByName(logger *log.Logger, username string) (*sdkModels.User, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDescribeUserByNameParams()
	req.WithTraceID(logid)
	// userId := logger.GetPoint("userId").(string)
	// req.WithBmpUserID(&userId)
	req.WithUserName(username)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	resp, err := openApi.SdkClient.User.DescribeUserByName(req, nil)
	if err != nil {

		logger.Warn("DescribeUserByName openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			fmt.Println("DescribeUserByName error:", err.Error())
			return nil, err
		}
		fmt.Println("DescribeUserByName error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	return resp.Payload.Result.User, nil

}

// 【用户管理】【用户详情】
func GetUserInfo(logger *log.Logger) (*response.User, error) {
	userId := logger.GetPoint("userId").(string)
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDescribeLocalUserParams()
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	responseApi, err := openApi.SdkClient.User.DescribeLocalUser(req, nil)
	if err != nil {
		logger.Info("DescribeLocalUser openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	var v = responseApi.Payload.Result.User
	obj := response.User{}
	obj.ID = int64(v.ID)
	obj.UserID = v.UserID
	obj.UserName = v.UserName
	obj.Email = v.Email
	obj.Timezone = v.Timezone
	obj.Language = v.Language
	obj.DefaultProjectID = v.DefaultProjectID
	obj.DefaultProjectName = v.DefaultProjectName
	obj.Description = v.Description
	obj.PhoneNumber = v.PhoneNumber
	obj.PhonePrefix = v.PhonePrefix
	obj.RoleID = v.RoleID
	obj.RoleName = v.RoleName
	obj.CreatedBy = v.CreatedBy
	obj.CreatedTime = v.CreatedTime
	obj.UpdatedBy = v.UpdatedBy
	obj.UpdatedBy = v.UpdatedTime
	obj.LicenseExist = true
	obj.LicenseModels = []string{"带内监控管理", "带外监控管理"}

	return &obj, nil
}

// 【用户管理】【修改密码】
func ModifyUserPassword(logger *log.Logger, userID string, param *requestTypes.UpdatePasswordRequest) (bool, error) {

	fmt.Println("[DEBUG] ORIGIN OLD PASSWORD IS:", param.OldPassword)
	//前端传过来的是加密并base64过的
	pv, err := base64.StdEncoding.DecodeString(param.OldPassword)
	if err != nil {
		return false, errors.New("parse old password error:" + err.Error())
	}
	pv, err = util.AesDecrypt(pv)
	if err != nil {
		return false, errors.New("old password aes decode error:" + err.Error())
	}
	param.OldPassword = string(pv)
	fmt.Println("[DEBUG] PARSED OLD PASSWORD IS:", param.OldPassword)
	////
	fmt.Println("[DEBUG] ORIGIN PASSWORD IS:", param.Password)
	//前端传过来的是加密并base64过的
	pv, err = base64.StdEncoding.DecodeString(param.Password)
	if err != nil {
		return false, errors.New("parse password error:" + err.Error())
	}
	pv, err = util.AesDecrypt(pv)
	if err != nil {
		return false, errors.New("password aes decode error:" + err.Error())
	}
	param.Password = string(pv)
	fmt.Println("[DEBUG] PARSED PASSWORD IS:", param.Password)

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewModifyLocalUserPasswordParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)

	body := sdkModels.ModifyUserPasswordRequest{
		Password:    &param.Password,
		OldPassword: &param.OldPassword,
	}
	req.WithBody(&body)

	res, err := openApi.SdkClient.User.ModifyLocalUserPassword(req, nil)
	if err != nil {
		logger.Info("ModifyLocalUserPassword openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return res.Payload.Result.Success, nil
}

// 【用户管理】【编辑用户信息】
func ModifyUserInfo(logger *log.Logger, userID string, param *requestTypes.ModifyUserInfoRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewModifyLocalUserParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)

	body := sdkModels.ModifyLocalUserRequest{
		DefaultProjectID: param.DefaultProjectID,
		//Description:      param.Description,
		Email:       param.Email,
		Language:    param.Language,
		PhonePrefix: param.PhonePrefix,
		PhoneNumber: param.PhoneNumber,
		Timezone:    param.Timezone,
	}
	req.WithBody(&body)

	res, err := openApi.SdkClient.User.ModifyLocalUser(req, nil)
	if err != nil {
		logger.Info("ModifyLocalUserPassword openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return res.Payload.Result.Success, nil
}

func VerifyUser(logger *log.Logger, username, password string) (bool, error) {

	fmt.Println("[DEBUG] ORIGIN PASSWORD IS:", password)
	//前端传过来的是加密并base64过的
	pv, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return false, errors.New("parse password error:" + err.Error())
	}
	pv, err = util.AesDecrypt(pv)
	if err != nil {
		return false, errors.New("password aes decode error:" + err.Error())
	}
	password = string(pv)
	fmt.Println("[DEBUG] PARSED PASSWORD IS:", password)

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewVerifyUserParams()
	req.WithTraceID(logid)

	var roleId string = CONSOLE_ROLE_ID

	body := &sdkModels.VerifyUserRequest{
		UserName: &username,
		Password: &password,
		RoleID:   &roleId,
	}
	req.WithBody(body)

	logger.Info("VerifyUser req:", util.InterfaceToJson(req))
	resp, err := openApi.SdkClient.User.VerifyUser(req, nil)
	logger.Info("VerifyUser resp:", util.InterfaceToJson(resp))
	if err != nil {
		logger.Info("VerifyUser openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return resp.Payload.Result.Success, nil
}

func GetConsoleAccessUserList(logger *log.Logger, userId string) ([]*sdkModels.User, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkUserParameters.NewDescribeUsersParams()
	request.WithTraceID(logid)
	request.WithBmpUserID(&userId)
	consoleAccessRoles := strings.Join([]string{ADMIN_ROLE_ID, CONSOLE_ROLE_ID}, ",")
	request.WithRoleID(&consoleAccessRoles)
	isAll := "1"
	request.WithIsAll(&isAll)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	//request.SetOsType(param.ImageType)
	logger.Info("DescribeUsers openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.User.DescribeUsers(request, nil)
	logger.Info("DescribeUsers openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeUsers openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	users := []*sdkModels.User{}
	for _, v := range responseApi.Payload.Result.Users {
		if v.UserID != userId {
			users = append(users, v)
		}
	}
	return users, nil
}

func CheckUserConsoleAccess(logger *log.Logger, userId string, req requestTypes.CheckUserConsoleAccessRequest) (bool, error) {

	//不能校验自己
	if userId == req.CheckedUserId {
		return false, fmt.Errorf("param invalided")
	}

	users, err := GetConsoleAccessUserList(logger, userId)
	if err != nil {
		return false, err
	}
	for _, v := range users {
		if v.UserID == req.CheckedUserId {
			return true, nil
		}

	}
	return false, nil
}

func CheckUserConsoleAccessByName(logger *log.Logger, userId string, req requestTypes.CheckUserConsoleAccessByNameRequest) (bool, error) {

	isEn := logger.GetPoint("language").(string) == BaseLogic.EN_US

	var errMsg string
	u, _ := GetUserInfo(logger)
	if u == nil {
		if isEn {
			errMsg = "no operator user info found"
		} else {
			errMsg = "未获取到操作用户信息"
		}
		return false, fmt.Errorf(errMsg)
	}
	if u.UserName == req.CheckedUserName {
		if isEn {
			errMsg = "cannot be project owner"
		} else {
			errMsg = "不能是项目拥有者本人"
		}
		return false, fmt.Errorf(errMsg)
	}

	u1, _ := GetUserByName(logger, req.CheckedUserName)
	if u1 == nil {
		if isEn {
			errMsg = "user input not exist, re input please"
		} else {
			errMsg = "该用户不存在，请重新输入"
		}
		return false, fmt.Errorf(errMsg)
	}

	if req.Operation == "move" { //项目转移

	} else if req.Operation == "share" { //项目共享，可以多次共享
		// s, _ := GetShareProjectInfo(logger, req.ProjectID)
		// for _, v := range s.Shares {
		// 	if v.SharedUserName == req.CheckedUserName {
		// 		if isEn {
		// 			errMsg = "project already shared to the user"
		// 		} else {
		// 			errMsg = "该用户已经添加共享"
		// 		}
		// 		return false, fmt.Errorf(errMsg)
		// 	}
		// }
	}
	users, err := GetConsoleAccessUserList(logger, userId)
	if err != nil {
		logger.Warnf("CheckUserConsoleAccessByName.GetConsoleAccessUserList error:%s", err.Error())
	}

	for _, v := range users {
		if v.UserName == req.CheckedUserName {
			return true, nil
		}
	}

	if isEn {
		errMsg = "the user is non-console role, re input please"
	} else {
		errMsg = "该用户不是 user 角色，请重新输入"
	}
	return false, fmt.Errorf(errMsg)

}

func GetShareProjectInfo(logger *log.Logger, ProjectId string) (*sdkModels.ShareProjectInfo, error) {
	logid := logger.GetPoint("logid").(string)

	req := sdkProjectParameters.NewDescribeShareProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DescribeShareProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.DescribeShareProject(req, nil)
	logger.Info("DescribeShareProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeShareProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil
}
