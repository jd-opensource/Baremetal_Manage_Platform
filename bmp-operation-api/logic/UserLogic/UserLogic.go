package UserLoigc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"

	sdkUserParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/user"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

const OPERATION_ROLE_ID = "role-operator-uuid-01"

// 【用户管理】【用户列表】
func GetUserList(logger *log.Logger, param *requestTypes.QueryUsersRequest) (*response.UserPage, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDescribeUsersParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithDefaultProjectID(&param.ProjectID)
	req.WithUserName(&param.UserName)
	req.WithRoleID(&param.RoleID)
	req.WithIsAll(&param.IsAll)
	req.WithPageSize(&param.PageSize)
	req.WithPageNumber(&param.PageNumber)
	logger.Info("DescribeUsers openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.DescribeUsers(req, nil)
	logger.Info("DescribeUsers openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.GetUserList:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.GetUserList error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	resp := response.UserPage{}
	obj := response.User{}
	if len(client.Payload.Result.Users) == 0 {
		resp.Users = make([]response.User, 0)
		return &resp, nil
	}
	for _, v := range client.Payload.Result.Users {
		obj.ID = v.ID
		obj.UserID = v.UserID
		obj.UserName = v.UserName
		obj.Email = v.Email
		// obj.Timezone = time.Time(v.Timezone)
		obj.Language = v.Language
		obj.DefaultProjectID = v.DefaultProjectID
		obj.DefaultProjectName = v.DefaultProjectName
		obj.ProjectCount = int(v.ProjectCount)
		obj.InstanceCount = int(v.InstanceCount)
		obj.Description = v.Description
		obj.PhoneNumber = v.PhoneNumber
		obj.PhonePrefix = v.PhonePrefix
		obj.RoleID = v.RoleID
		obj.RoleName = v.RoleName
		obj.RoleName = v.RoleName
		obj.CreatedBy = v.CreatedBy
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedBy = v.UpdatedBy
		obj.UpdatedBy = util.UtcToTime(v.UpdatedTime)

		resp.Users = append(resp.Users, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

// 导出Excel文件
func ExportUserExcel(logger *log.Logger, result []response.User) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {
		sheetData = append(sheetData, []string{
			data.UserName,
			data.UserID,
			data.RoleName,
			data.PhoneNumber,
			data.Email,
			data.Description,
			data.CreatedTime,
		})
	}
	exportExcelHeader := BaseLogic.ExportUserHeader
	if language == BaseLogic.EN_US {
		exportExcelHeader = BaseLogic.ExportUserHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportExcelHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	// 生成100000到100000+900000之间的随机数，左闭右开
	downloadFileName := "device_list_" + time.Now().Format(BaseLogic.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx"
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

// 【用户管理】【添加用户】
func AddUser(logger *log.Logger, param *requestTypes.AddUserRequest) (*response.UserAdd, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewCreateUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.CreateUserRequest{
		RoleID:      &param.RoleID,
		UserName:    &param.UserName,
		Email:       param.Email,
		PhonePrefix: param.PhonePrefix,
		PhoneNumber: param.PhoneNumber,
		Language:    param.Language,
		Password:    &param.Password,
		Description: param.Description,
	}
	req.WithBody(&body)
	logger.Info("CreateUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.CreateUser(req, nil)
	logger.Info("CreateUser openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.CreateUser:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.CreateUser error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	return &response.UserAdd{
		UserID: client.Payload.Result.UserID,
	}, nil
}

// 【用户管理】【用户详情】
func GetUserInfo(logger *log.Logger, param *requestTypes.QueryUserInfoRequest) *response.User {
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDescribeUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithUserID(param.UserID)
	logger.Info("DescribeUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.DescribeUser(req, nil)
	logger.Info("DescribeUser openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.DescribeUser:", err.Error())
			return nil
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.DescribeUser error:", respErr.Message)
		return nil
	}

	var v = client.Payload.Result.User
	obj := response.User{}
	obj.ID = v.ID
	obj.UserID = v.UserID
	obj.UserName = v.UserName
	obj.Email = v.Email
	// obj.Timezone = time.Time(v.Timezone)
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
	return &obj
}

func GetUserByName(logger *log.Logger, username string) (*sdkModels.User, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDescribeUserByNameParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	// userId := logger.GetPoint("userId").(string)
	// req.WithBmpUserID(&userId)
	req.WithUserName(username)

	logger.Info("DescribeUserByName openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.DescribeUserByName(req, nil)
	logger.Info("DescribeUserByName openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		fmt.Println("GetUserByName resp.err:", err.Error())
	}
	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			fmt.Println("unknown error in User.GetUserByName:", err.Error())
			logger.Warn("unknown error in User.GetUserByName:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.GetUserByName error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	return client.Payload.Result.User, nil

}

// 【用户管理】【编辑用户信息】
func ModifyUserInfo(logger *log.Logger, param *requestTypes.ModifyUserInfoRequest, userIdModify string) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewModifyUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithUserID(userIdModify)

	body := sdkModels.ModifyUserRequest{
		//RoleID:      param.RoleID,
		//UserName:    param.UserName,
		Email:       param.Email,
		PhonePrefix: param.PhonePrefix,
		PhoneNumber: param.PhoneNumber,
		//Language:    param.Language,
		//Password:    param.Password,
		Description: param.Description,
	}
	req.WithBody(&body)
	logger.Info("ModifyUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.ModifyUser(req, nil)
	logger.Info("ModifyUser openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.ModifyUser:", err.Error())
			return err
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.ModifyUser error:", respErr.Message)
		return errors.New(respErr.Message)
	}
	return nil
}

// 【用户管理】【删除用户】
func DeleteUserInfo(logger *log.Logger, param *requestTypes.DeleteUserInfoRequest) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewDeleteUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithUserID(param.UserID)
	logger.Info("DeleteUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.DeleteUser(req, nil)
	logger.Info("DeleteUser openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.VerifyUser:", err.Error())
			return err
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.VerifyUser error:", respErr.Message)
		return errors.New(respErr.Message)
	}
	return nil

}

func VerifyUser(logger *log.Logger, username, password string) (*response.CommonResponse, error) {

	fmt.Println("[DEBUG] ORIGIN PASSWORD IS:", password)
	//前端传过来的是加密并base64过的
	pv, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, errors.New("parse password error:" + err.Error())
	}
	pv, err = util.AesDecrypt(pv)
	if err != nil {
		return nil, errors.New("password aes decode error:" + err.Error())
	}
	password = string(pv)
	fmt.Println("[DEBUG] PARSED PASSWORD IS:", password)

	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewVerifyUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	roleId := OPERATION_ROLE_ID
	body := &sdkModels.VerifyUserRequest{
		UserName: &username,
		Password: &password,
		RoleID:   &roleId,
	}
	req.WithBody(body)

	logger.Info("VerifyUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.VerifyUser(req, nil)
	logger.Info("VerifyUser openapi sdk resp:", util.InterfaceToJson(client))
	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("unknown error in User.VerifyUser:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warn("User.VerifyUser error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("VerifyUser failed")
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【用户管理】【用户详情】
func GetLocalUserInfo(logger *log.Logger) (*response.LocalUser, error) {
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
	obj := response.LocalUser{}
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
func ModifyLocalUserPassword(logger *log.Logger, userID string, param *requestTypes.UpdateLocalPasswordRequest) (bool, error) {

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
func ModifyLocalUserInfo(logger *log.Logger, userID string, param *requestTypes.ModifyLocalUserInfoRequest) (bool, error) {
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
