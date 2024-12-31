package DeviceLogic

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkDeviceParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/device"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func QueryDeviceStockCount(logger *log.Logger, param requestTypes.QueryDeviceStockRequest) (int64, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkDeviceParameters.NewDescribeDevicesParams()
	userId := logger.GetPoint("userId").(string)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(&param.DeviceTypeId)
	//status和isAll等，openapi没设计好，后面修改
	status := "putaway"
	req.WithManageStatus(&status)
	isAll := "1"
	req.WithIsAll(&isAll)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	// req.WithIdcID(&param.IdcID)
	logger.Info("DescribeDevices openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Device.DescribeDevices(req, nil)
	logger.Info("DescribeDevices openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeDevices openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return 0, err
		}
		return 0, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.TotalCount, nil
}
