package IdcLogic

import (
	"errors"
	"strconv"

	"math/rand"
	"time"

	sdkIdcParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/idc"
	sdkUserParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/user"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【设备管理】【设备列表】
func GetIdcList(logger *log.Logger, param *requestTypes.QueryIdcListRequest) (*response.IdcPage, error) {
	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkIdcParameters.NewDescribeIdcsParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithLevel(&param.Level)
	req.WithName(&param.Name)
	req.WithIsAll(&param.IsAll)
	req.WithPageNumber(&param.PageNumber)
	req.WithPageSize(&param.PageSize)

	logger.Info("DescribeIdcs openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Idc.DescribeIdcs(req, nil)
	logger.Info("DescribeIdcs openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.GetIdcList:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.GetIdcList error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}

	resp := response.IdcPage{}
	obj := response.Idc{}
	if len(client.Payload.Result.Idcs) == 0 {
		resp.Idcs = make([]response.Idc, 0)
		return &resp, nil
	}

	for _, v := range client.Payload.Result.Idcs {
		obj = buildEntity(v, param.Show)
		resp.Idcs = append(resp.Idcs, obj)
	}

	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

func ExportIdc(logger *log.Logger, result []response.Idc) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {
		sheetData = append(sheetData, []string{
			//strconv.FormatInt(data.CabinetNum, 10),
			data.Name,
			data.NameEn,
			data.Level,
			data.Address,
			data.CreateTime,
			data.CreatedBy,
			data.UpdateTime,
			data.UpdatedBy,
		})
	}
	exportIdcHeader := BaseLogic.ExportIdcHeader
	if language == BaseLogic.EN_US {
		exportIdcHeader = BaseLogic.ExportIdcHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportIdcHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "idc_list_" + time.Now().Format(BaseLogic.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

// 【设备管理】【设备详情】
func GetIdcInfo(logger *log.Logger, param *requestTypes.QueryIdcRequest) (*response.Idc, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkIdcParameters.NewDescribeIdcParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithIdcID(param.IdcID)
	logger.Info("DescribeIdc openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Idc.DescribeIdc(req, nil)
	logger.Info("DescribeIdc openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.GetIdcInfo:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.GetIdcInfo error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	if client.Payload.Result == nil || client.Payload.Result.Idc == nil {
		return nil, errors.New("idc not found")
	}

	obj := buildEntity(client.Payload.Result.Idc, param.Show)
	return &obj, nil
}

// 列表和详情共用方法
func buildEntity(v *sdkModels.Idc, show string) response.Idc {
	obj := response.Idc{}

	obj.Address = v.Address
	obj.ID = v.ID
	obj.IDcID = v.IDcID
	obj.IloUser = v.IloUser
	obj.Level = v.Level
	obj.Name = v.Name
	obj.NameEn = v.NameEn
	obj.Shortname = v.Shortname
	obj.SwitchUser1 = v.SwitchUser1
	obj.SwitchUser2 = v.SwitchUser2
	obj.CreateTime = util.UtcToTime(v.CreateTime)
	obj.CreatedBy = v.CreatedBy
	obj.UpdateTime = util.UtcToTime(v.UpdateTime)
	obj.UpdatedBy = v.UpdatedBy

	if show == "iloPassword" {
		obj.IloPassword = v.IloPassword
		obj.SwitchPassword1 = ""
		obj.SwitchPassword2 = ""
	}
	if show == "switchPassword1" {
		obj.IloPassword = ""
		obj.SwitchPassword1 = v.SwitchPassword1
		obj.SwitchPassword2 = ""
	}
	if show == "switchPassword2" {
		obj.IloPassword = ""
		obj.SwitchPassword1 = ""
		obj.SwitchPassword2 = v.SwitchPassword2
	}
	if show == "iloPassword,switchPassword1,switchPassword2" {
		obj.IloPassword = v.IloPassword
		obj.SwitchPassword1 = v.SwitchPassword1
		obj.SwitchPassword2 = v.SwitchPassword2
	}
	return obj
}

// 【设备管理】【删除设备】
func DeleteIdc(logger *log.Logger, param *requestTypes.DeleteIdcRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkIdcParameters.NewDeleteIdcParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithIdcID(param.IdcID)
	logger.Info("DeleteIdc openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Idc.DeleteIdc(req, nil)
	logger.Info("DeleteIdc openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.DeleteIdc:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.DeleteIdc error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【机型管理】【添加机型】
func CreateIdc(logger *log.Logger, param *requestTypes.CreateIdcRequest) (*response.CreateIdcResult, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkIdcParameters.NewCreateIdcParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	body := sdkModels.CreateIdcRequest{
		Address:         param.Address,
		NameEn:          &param.NameEn,
		IloPassword:     param.IloPassword,
		IloUser:         param.IloUser,
		Level:           &param.Level,
		Name:            &param.Name,
		Shortname:       &param.Shortname,
		SwitchPassword1: param.SwitchPassword1,
		SwitchPassword2: param.SwitchPassword2,
		SwitchUser1:     param.SwitchUser1,
		SwitchUser2:     param.SwitchUser2,
	}
	req.WithBody(&body)
	logger.Info("CreateIdc openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Idc.CreateIdc(req, nil)
	logger.Info("CreateIdc openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.CreateIdc:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.CreateIdc error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("add idc info failed")
	}

	res := &response.CreateIdcResult{
		IdcId: client.Payload.Result.IdcID,
	}

	return res, nil
}

// 【机型管理】【编辑机型】
func ModifyIdcInfo(logger *log.Logger, param *requestTypes.ModifyIdcRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkIdcParameters.NewModifyIdcParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithIdcID(param.IdcID)
	body := sdkModels.ModifyIdcRequest{
		Address: param.Address,
		// CabinetNum:      int64(param.CabinetNum),
		IloPassword:     param.IloPassword,
		IloUser:         param.IloUser,
		Level:           param.Level,
		Name:            param.Name,
		NameEn:          param.NameEn,
		Shortname:       param.Shortname,
		SwitchPassword1: param.SwitchPassword1,
		SwitchPassword2: param.SwitchPassword2,
		SwitchUser1:     param.SwitchUser1,
		SwitchUser2:     param.SwitchUser2,
	}
	req.WithBody(&body)
	logger.Info("ModifyIdc openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Idc.ModifyIdc(req, nil)
	logger.Info("ModifyIdc openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.ModifyIdc:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.ModifyIdc error:", respErr.Message)
		return nil, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return nil, errors.New("modify idc info failed")
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【机型管理】【安全验证】
func VerifyUser(logger *log.Logger, param *requestTypes.VerifyUserRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkUserParameters.NewVerifyUserParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	body := sdkModels.VerifyUserRequest{
		Password: &param.Password,
		UserName: &param.Username,
	}
	req.WithBody(&body)
	logger.Info("VerifyUser openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.User.VerifyUser(req, nil)
	logger.Info("VerifyUser openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in Idc.VerifyUser:", err.Error())
			return nil, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("Idc.VerifyUser error:", respErr.Message)
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
