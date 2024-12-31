package AuditLogLogic

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkAuditLogParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/audit_log"
	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	"git.jd.com/cps-golang/ironic-common/exception"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

func QueryAuditLogs(logger *log.Logger, param *requestTypes.DescribeAuditLogsRequest) (*response.AuditLogsList, error) {
	// language := logger.GetPoint("language").(string)
	logid := logger.GetPoint("logid").(string)
	request := sdkAuditLogParameters.NewDescribeAuditLogsParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithInstanceID(&param.InstanceID)

	request.WithPageNumber(&param.PageNumber)
	request.WithPageSize(&param.PageSize)
	if param.Operation != "" {
		request.WithOperation(&param.Operation)
	}
	if param.UserName != "" {
		request.WithUserName(&param.UserName)
	}
	if param.Result != "" {
		request.WithResult(&param.Result)
	}
	if param.StartTime != 0 {
		request.WithStartTime(&param.StartTime)
	}
	if param.EndTime != 0 {
		request.WithEndTime(&param.EndTime)
	}

	if param.ExportType != "" || param.IsAll == "1" {
		isAll := "1"
		request.IsAll = &isAll
	}

	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	logger.Info("DescribeAuditLogs openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.AuditLog.DescribeAuditLogs(request, nil)
	logger.Info("DescribeAuditLogs openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeAuditLogs openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	resp := &response.AuditLogsList{
		AuditLogs:  responseApi.Payload.Result.AuditLogs,
		PageNumber: responseApi.Payload.Result.PageNumber,
		PageSize:   responseApi.Payload.Result.PageSize,
		TotalCount: responseApi.Payload.Result.TotalCount,
	}

	return resp, nil

}

var resultMapEn map[string]string = map[string]string{
	"success": "Operation Success",
	"fail":    "Operation Fail",
	"doing":   "Doing",
}

var resultMapZh map[string]string = map[string]string{
	"success": "操作成功",
	"fail":    "操作失败",
	"doing":   "操作中",
}

func ExportAuditLogs(logger *log.Logger, result []*models.AuditLog) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)

	for _, data := range result {
		var result string = resultMapZh[data.Result]
		if language == "en_US" {
			result = resultMapEn[data.Result]
		}
		sheetData = append(sheetData, []string{
			data.LogID,
			data.OperationName,
			data.UserName,
			time.Unix(data.OperateTime, 0).Format("2006-01-02 15:04:05"),
			result,
			data.FailReason,
		})
	}
	exportHeader := BaseLogic.ExportAuditLogsHeader
	if language == "en_US" {
		exportHeader = BaseLogic.ExportAuditLogsHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "auditlogs_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}

func DescribeAuditLogsTypes(logger *log.Logger) ([]*models.AuditLogsType, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkAuditLogParameters.NewDescribeAuditLogTypesParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)

	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	logger.Info("DescribeAuditLogTypes openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.AuditLog.DescribeAuditLogTypes(request, nil)
	logger.Info("DescribeAuditLogTypes openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeAuditLogTypes openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil
}
