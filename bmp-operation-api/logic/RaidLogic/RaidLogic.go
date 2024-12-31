package RaidLogic

import (
	"errors"

	sdkParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/raid"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"

	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device_type"
	sdkRaidParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/raid"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// 【raid列表】
func QueryRaidsAll(logger *log.Logger) (*response.RaidList, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeRaidsParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	logger.Info("DescribeRaids openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Raid.DescribeRaids(req, nil)
	logger.Info("DescribeRaids openapi sdk resp:", util.InterfaceToJson(client))

	resp := response.RaidList{}
	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	for _, v := range client.Payload.Result.Raids {
		obj := &response.Raid{}
		obj.RaidID = v.RaidID
		obj.Name = v.Name

		obj.DescriptionZh = v.DescriptionZh
		obj.DescriptionEn = v.DescriptionEn

		//obj.AvailableValue = int(v.AvailableValue)

		resp.Raids = append(resp.Raids, obj)
	}
	return &resp, nil
}

func QueryRaids(logger *log.Logger, param requestTypes.QueryRaidsRequest) ([]responseTypes.DeviceTypeRaid, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkDeviceTypeParameters.NewDescribeDeviceTypeRaidsParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithDeviceTypeID(&param.DeviceTypeID)
	request.WithVolumeType(&param.VolumeType)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	//request.SetOsType(param.ImageType)
	logger.Info("DescribeDeviceTypeRaids openapi request:", util.InterfaceToJson(request))
	responseApi, err := sdk.SdkClient.DeviceType.DescribeDeviceTypeRaids(request, nil)
	logger.Info("DescribeDeviceTypeRaids openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeDeviceTypeRaids openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := []responseTypes.DeviceTypeRaid{}

	rdeviceTypeRaids := responseApi.Payload.Result
	for _, v := range rdeviceTypeRaids {
		request := sdkRaidParameters.NewDescribeRaidParams()
		request.WithBmpUserID(&userId)
		request.WithTraceID(logid)
		request.WithRaidID(v.RaidID)
		language := logger.GetPoint("language").(string)
		request.WithBmpLanguage(&language)
		responseApi, err := sdk.SdkClient.Raid.DescribeRaid(request, nil)
		if err != nil {
			logger.Info("DescribeRaid openapi error:", err.Error())
			respErr, b := util.GetOpenapiError(err)
			if !b {
				return nil, err
			}
			return nil, errors.New(respErr.Message)
		}
		res = append(res, SdkRDeviceTypeRaid2Raid(v, responseApi.Payload.Result.Name))

	}
	return res, nil
}

func SdkRDeviceTypeRaid2Raid(v *sdkModels.RDeviceTypeRaid, raidName string) responseTypes.DeviceTypeRaid {
	res := responseTypes.DeviceTypeRaid{
		RDeviceTypeRaid: *v,
		RaidName:        raidName,
	}

	return res
}
