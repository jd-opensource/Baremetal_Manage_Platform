package monitorProxyLogic

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkmonitorProxyParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/monitor_proxy"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func DescribeAgentStatus(logger *log.Logger, param *request.DesrcibeAgentStatusRequest) (*sdkModels.AgentStatusResponse, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorProxyParameters.NewDesrcibeAgentStatusParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithInstanceID(param.InstanceId)

	logger.Info("DesrcibeAgentStatus openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorProxy.DesrcibeAgentStatus(request, nil)
	logger.Info("DesrcibeAgentStatus openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DesrcibeAgentStatus openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil

}

func DescribeDeviceTags(logger *log.Logger, param *request.DesrcibeTagsRequest) (response.TagsResponse, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkmonitorProxyParameters.NewDesrcibeTagsParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	request.WithInstanceID(param.InstanceID)
	request.WithTagName(param.TagName)

	logger.Info("DesrcibeTags openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.MonitorProxy.DesrcibeTags(request, nil)
	logger.Info("DesrcibeTags openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DesrcibeTags openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Tags, nil
}
