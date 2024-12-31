package monitorDataLogic

import (
	"errors"

	sdkmonitorDataParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/monitor_data"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

func GetMonitorData(logger *log.Logger, param request.GetMonitorDataRequest) ([]*sdkModels.DataEveryMetric, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorDataParameters.NewGetMonitorDataParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithMetricName(&param.MetricName)
	request.WithInstanceID(&param.InstanceID)
	request.WithDevice(&param.Device)
	request.WithStartTime(&param.StartTime)
	request.WithEndTime(&param.EndTime)
	request.WithTimeInterval(&param.TimeInterval)
	request.WithDownSampleType(&param.DownSampleType)
	request.WithLastManyTime(&param.LastManyTime)
	request.WithUserName(&param.UserName)
	request.WithIdcID(&param.IdcID)
	logger.Info("GetMonitorData openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorData.GetMonitorData(request, nil)
	logger.Info("GetMonitorData openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetMonitorData openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil
}
