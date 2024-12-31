package messageLogic

import (
	// "coding.jd.com/aidc-bmp/bmp-openapi/dao/messageDao"
	"errors"
	"math/rand"
	"strconv"
	"time"

	sdkMessageParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/message"
	"coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/constant"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"

	// "coding.jd.com/aidc-bmp/bmp-openapi/util"
	"git.jd.com/cps-golang/ironic-common/exception"
	log "git.jd.com/cps-golang/log"
)

func HasUnreadMessage(logger *log.Logger) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewHasUnreadMessageParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("HasUnreadMessage openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.HasUnreadMessage(request, nil)
	logger.Info("HasUnreadMessage openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("HasUnreadMessage openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.HasUnread, nil
}

func GetPageMessages(logger *log.Logger, param *requestTypes.QueryMessagesRequest, p util.Pageable) (*response.MessageList, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewGetMessageListParams()
	userId := logger.GetPoint("userId").(string)
	request.WithPageNumber(&p.PageNumber)
	request.WithPageSize(&p.PageSize)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	if param.HasRead != "" {
		request.WithHasRead(&param.HasRead)
	}
	if param.IsAll != "" || param.ExportType != "" {
		request.WithIsAll(&param.IsAll)
	}
	if param.MessageType != "" {
		request.WithMessageType(&param.MessageType)
	}
	if param.MessageSubType != "" {
		request.WithMessageSubType(&param.MessageSubType)
	}
	if param.Detail != "" {
		request.WithDetail(&param.Detail)
	}

	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("GetMessageList openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.GetMessageList(request, nil)
	logger.Info("GetMessageList openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetMessageList openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.MessageList{}
	resp.PageNumber = responseApi.Payload.Result.PageNumber
	resp.PageSize = responseApi.Payload.Result.PageSize
	resp.TotalCount = responseApi.Payload.Result.TotalCount
	resp.Messages = responseApi.Payload.Result.Messages

	return &resp, nil

}

func GetMessageStatistic(logger *log.Logger) (int64, int64, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewGetMessageStatisticParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("GetMessageStatistic openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.GetMessageStatistic(request, nil)
	logger.Info("GetMessageStatistic openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetMessageStatistic openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return 0, 0, err
		}
		return 0, 0, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.TotalCount, responseApi.Payload.Result.UnreadCount, nil

}

func ReadMessages(logger *log.Logger, param *requestTypes.ReadMessagesRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewReadMessageParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ReadMessagesRequest{
		MessageIds: param.MessageIds,
	}
	request.WithBody(&body)

	logger.Info("ReadMessage openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.ReadMessage(request, nil)
	logger.Info("ReadMessage openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ReadMessage openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil
}

func GetMessageById(logger *log.Logger, param *requestTypes.GetMessageByIdRequest) (*sdkModels.Message, string, string, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewGetMessageByIDParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	request.WithMessageID(param.MessageId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("GetMessageByID openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.GetMessageByID(request, nil)
	logger.Info("GetMessageByID openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetMessageByID openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, "", "", err
		}
		return nil, "", "", errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Message, responseApi.Payload.Result.NextMessageID, responseApi.Payload.Result.PrevMessageID, nil

}

func GetMessageTypes(logger *log.Logger) (*sdkModels.MessageTypesResp, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewGetMessageTypesParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("GetMessageTypes openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.GetMessageTypes(request, nil)
	logger.Info("GetMessageTypes openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetMessageTypes openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return &responseApi.Payload.Result, nil

}

func DeleteMessages(logger *log.Logger, param *requestTypes.DeleteMessagesRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewDeleteMessageParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.DeleteMessagesRequest{
		MessageIds: param.MessageIds,
	}
	request.WithBody(&body)

	logger.Info("DeleteMessage openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.DeleteMessage(request, nil)
	logger.Info("DeleteMessage openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteMessage openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil
}

func ExportMessage(logger *log.Logger, result []*models.Message) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)
	tz := logger.GetPoint("timezone").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {

		if language == BaseLogic.EN_US {

		}

		sheetData = append(sheetData, []string{
			data.Detail,
			util.TimestampToString(int64(data.FinishTime), tz),
			data.MessageType,
			data.MessageSubType,
		})
	}
	exportMessageHeader := BaseLogic.ExportMessageHeader
	if language == "en_US" {
		exportMessageHeader = BaseLogic.ExportMessageHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportMessageHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "message_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}

func DialMail(logger *log.Logger, param *requestTypes.MailDialRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewDialMailParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.MailDialRequest{
		ServerAddr: &param.ServerAddr,
		ServerPort: &param.ServerPort,
		AdminAddr:  &param.AdminAddr,
		AdminPass:  &param.AdminPass,
		To:         param.To,
	}
	request.WithBody(&body)

	logger.Info("DialMail openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.DialMail(request, nil)
	logger.Info("DialMail openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Warnf("DialMail openapi error:%s", err.Error())
		panic(constant.BuildInvalidArgumentWithArgs("邮箱配置信息有误,请重新输入", "mail config error, please re input"))
	}

	return responseApi.Payload.Result.Success, nil
}

func DescribeMail(logger *log.Logger) (*sdkModels.Mail, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewMailDetailParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("MailDetail openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.MailDetail(request, nil)
	logger.Info("MailDetail openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("MailDetail openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil

}

func SaveIsPushMail(logger *log.Logger, param *requestTypes.SaveIsPushMailRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkMessageParameters.NewSaveIsPushMailParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.SaveIsPushMailRequest{
		IsPush: param.IsPush,
	}
	request.WithBody(&body)

	logger.Info("SaveIsPushMail openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Message.SaveIsPushMail(request, nil)
	logger.Info("SaveIsPushMail openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("SaveIsPushMail openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil
}
