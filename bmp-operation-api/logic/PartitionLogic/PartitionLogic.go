package PartitionLogic

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"

	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device_type"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

func QueryDefaultSystemPartitions(logger *log.Logger, param requestTypes.QueryDefaultSystemPartitionRequest) ([]*sdkModels.Partition, error) {

	// language := logger.GetPoint("language").(string)
	logid := logger.GetPoint("logid").(string)
	request := sdkDeviceTypeParameters.NewDescribeDeviceTypeImagePartitionsParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	request.WithImageID(&param.ImageID)
	request.WithDeviceTypeID(&param.DeviceTypeID)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("QueryDeviceTypeImagePartition openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.DeviceType.DescribeDeviceTypeImagePartitions(request, nil)
	logger.Info("QueryDeviceTypeImagePartition openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("QueryDeviceTypeImagePartition openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.SystemPartition, nil

}
