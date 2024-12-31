package monitorRuleLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkmonitorRuleParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/monitor_rule"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
)

func AddRule(logger *log.Logger, param *request.AddRuleRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorRuleParameters.NewAddRuleParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	trOption := []*sdkModels.RuleTrigger{}
	for _, vv := range param.TriggerOption {
		v := vv
		item := &sdkModels.RuleTrigger{
			Calculation:     &v.Calculation,
			CalculationUnit: v.CalculationUnit,
			Metric:          &v.Metric,
			NoticeLevel:     &v.NoticeLevel,
			Operation:       &v.Operation,
			Period:          &v.Period,
			Threshold:       &v.Threshold,
			Times:           &v.Times,
		}
		trOption = append(trOption, item)
	}
	notifyOption := &sdkModels.RuleNotice{
		EffectiveIntervalEnd:   &param.NoticeOption.EffectiveIntervalEnd,
		EffectiveIntervalStart: &param.NoticeOption.EffectiveIntervalStart,
		NoticeCondition:        param.NoticeOption.NoticeCondition,
		NoticePeriod:           &param.NoticeOption.NoticePeriod,
		NoticeWay:              param.NoticeOption.NoticeWay,
		UserID:                 &param.NoticeOption.UserID,
	}
	body := &sdkModels.AddRuleRequest{
		DeviceTag:     param.DeviceTag,
		Dimension:     &param.Dimension,
		InstanceIds:   param.InstanceIds,
		Resource:      &param.Resource,
		RuleName:      &param.RuleName,
		TriggerOption: trOption,
		NoticeOption:  notifyOption,
		ProjectID:     &param.ProjectID,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("AddRule openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorRule.AddRule(req, nil)
	logger.Info("AddRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("AddRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}

func DescribeRule(logger *log.Logger, ruleId string) (*sdkModels.Rule, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorRuleParameters.NewDescribeRuleParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithRuleID(ruleId)
	logger.Info("DescribeRule openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorRule.DescribeRule(request, nil)
	logger.Info("DescribeRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil

}

func DescribeRules(logger *log.Logger, param *request.DescribeRulesRequest, p util.Pageable) (*sdkModels.RuleList, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorRuleParameters.NewDescribeRulesParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithIsAll(&param.IsAll)
	request.WithPageNumber(&p.PageNumber)
	request.WithPageSize(&p.PageSize)
	request.WithUserID(&param.UserId)
	request.WithUserName(&param.UserName)
	request.WithRuleID(&param.RuleId)
	request.WithRuleName(&param.RuleName)
	request.WithProjectID(&param.ProjectID)
	status := int64(param.Status)
	request.WithStatus(&status)

	logger.Info("DescribeRules openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorRule.DescribeRules(request, nil)
	logger.Info("DescribeRules openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeRules openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil
}

func ExportRules(logger *log.Logger, result []*sdkModels.Rule) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {

		sheetData = append(sheetData, []string{
			data.RuleName,
			data.RuleID,
			data.ResourceName,
			fmt.Sprintf("%d", len(data.InstanceIds)),
			strings.Join(data.TriggerDescription, ";"),
			data.StatusName,
		})
	}
	exportRuleHeader := BaseLogic.ExportRulesHeader
	if language == "en_US" {
		exportRuleHeader = BaseLogic.ExportRulesHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportRuleHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "rule_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}

func EditRule(logger *log.Logger, param *request.EditRuleRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorRuleParameters.NewEditRuleParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	trOption := []*sdkModels.RuleTrigger{}
	for _, vv := range param.TriggerOption {
		v := vv
		item := &sdkModels.RuleTrigger{
			Calculation:     &v.Calculation,
			CalculationUnit: v.CalculationUnit,
			Metric:          &v.Metric,
			NoticeLevel:     &v.NoticeLevel,
			Operation:       &v.Operation,
			Period:          &v.Period,
			Threshold:       &v.Threshold,
			Times:           &v.Times,
		}
		trOption = append(trOption, item)
	}
	notifyOption := &sdkModels.RuleNotice{
		EffectiveIntervalEnd:   &param.NoticeOption.EffectiveIntervalEnd,
		EffectiveIntervalStart: &param.NoticeOption.EffectiveIntervalStart,
		NoticeCondition:        param.NoticeOption.NoticeCondition,
		NoticePeriod:           &param.NoticeOption.NoticePeriod,
		NoticeWay:              param.NoticeOption.NoticeWay,
		UserID:                 &param.NoticeOption.UserID,
	}
	body := &sdkModels.EditRuleRequest{
		DeviceTag:     param.DeviceTag,
		Dimension:     &param.Dimension,
		InstanceIds:   param.InstanceIds,
		Resource:      &param.Resource,
		RuleID:        &param.RuleId,
		RuleName:      &param.RuleName,
		TriggerOption: trOption,
		NoticeOption:  notifyOption,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("EditRule openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorRule.EditRule(req, nil)
	logger.Info("EditRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("EditRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}

func EnableRule(logger *log.Logger, param *request.EnableRuleRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorRuleParameters.NewEnableRuleParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := &sdkModels.EnableRuleRequest{
		RuleID: &param.RuleId,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("EnableRule openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorRule.EnableRule(req, nil)
	logger.Info("EnableRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("EnableRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}

func DisableRule(logger *log.Logger, param *request.DisableRuleRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorRuleParameters.NewDisableRuleParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := &sdkModels.DisableRuleRequest{
		RuleID: &param.RuleId,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DisableRule openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorRule.DisableRule(req, nil)
	logger.Info("DisableRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DisableRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}

func DeleteRule(logger *log.Logger, param *request.DeleteRuleRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkmonitorRuleParameters.NewDeleteRuleParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := &sdkModels.DeleteRuleRequest{
		RuleID: &param.RuleId,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DeleteRule openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.MonitorRule.DeleteRule(req, nil)
	logger.Info("DeleteRule openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteRule openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil

}
