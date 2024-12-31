package monitorAlertLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkmonitorAlertParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/monitor_alert"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
)

func DescribeAlert(logger *log.Logger, AlertId string) (*sdkModels.Alert, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorAlertParameters.NewDescribeAlertParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithAlertID(AlertId)
	logger.Info("DescribeAlert openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorAlert.DescribeAlert(request, nil)
	logger.Info("DescribeAlert openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeAlert openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil

}

func DescribeAlerts(logger *log.Logger, param *request.DescribeAlertsRequest, p util.Pageable) (*sdkModels.AlertList, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorAlertParameters.NewDescribeAlertsParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithIsAll(&param.IsAll)
	request.WithPageNumber(&p.PageNumber)
	request.WithPageSize(&p.PageSize)
	request.WithUserID(&param.UserID)
	request.WithUserName(&param.UserName)
	request.WithRuleID(&param.RuleID)
	request.WithRuleName(&param.RuleName)
	request.WithStartTime(&param.StartTime)
	request.WithEndTime(&param.EndTime)
	request.WithResourceID(&param.ResourceID)
	request.WithProjectID(&param.ProjectID)

	logger.Info("DescribeAlerts openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorAlert.DescribeAlerts(request, nil)
	logger.Info("DescribeAlerts openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeAlerts openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil
}

func ExportAlerts(logger *log.Logger, result []*sdkModels.Alert) (fileName, downLoadFileName string) {

	tz := logger.GetPoint("timezone").(string)
	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {
		alertPeriod := fmt.Sprintf("%d分钟", data.AlertPeriod)
		if language == "en_US" {
			alertPeriod = fmt.Sprintf("%dmin", data.AlertPeriod)
		}
		sheetData = append(sheetData, []string{
			util.TimestampToString(int64(data.AlertTime), tz),
			data.ResourceID,
			data.TriggerDescription,
			data.AlertValue,
			data.AlertLevelDescription,
			alertPeriod,
			data.UserName,
		})
	}
	exportAlertHeader := BaseLogic.ExportAlertsHeader
	if language == "en_US" {
		exportAlertHeader = BaseLogic.ExportAlertsHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportAlertHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "alert_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}

func DeleteAlert(logger *log.Logger, param *request.DeleteAlertRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorAlertParameters.NewDeleteAlertParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := &sdkModels.DeleteAlertRequest{
		AlertID: &param.AlertId,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DeleteAlert openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorAlert.DeleteAlert(req, nil)
	logger.Info("DeleteAlert openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteAlert openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}
