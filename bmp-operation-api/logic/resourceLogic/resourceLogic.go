package resourceLogic

import (
	"errors"

	sdkClient "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/resource"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【设备管理】【设备列表】
func DescribeResources(logger *log.Logger, param *requestTypes.QueryResourcesRequest) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkClient.NewDescribeResourcesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithImageName(&param.ImageName)
	req.WithName(&param.Name)
	req.WithDeviceType(&param.DeviceType)
	req.WithUserName(&param.UserName)

	logger.Info("DescribeResources openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Resource.DescribeResources(req, nil)
	logger.Info("DescribeResources openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error:", err.Error())
			return true, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("openapi error:", respErr.Message)
		return true, errors.New(respErr.Message)
	}
	return client.Payload.Result.Success, nil
}
