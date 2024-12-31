package ImageLogic

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/device_type"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func QueryImagesByDeviceType(logger *log.Logger, param requestTypes.QueryImagesByDeviceTypeRequest) (map[string][]*sdkModels.Image, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypeImagesParams()
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	// req.WithIdcID(&param.IdcID)
	logger.Info("DescribeImages openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.DeviceType.DescribeDeviceTypeImages(req, nil)
	logger.Info("DescribeImages openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeImages openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	res := map[string][]*sdkModels.Image{}
	images := responseApi.Payload.Result.Images
	for _, v := range images {
		if _, ok := res[v.OsType]; !ok {
			res[v.OsType] = []*sdkModels.Image{v}
		} else {
			res[v.OsType] = append(res[v.OsType], v)
		}
	}
	return res, nil
}
