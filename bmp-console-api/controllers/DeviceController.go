package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/DeviceLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceController operations for Device
type DeviceController struct {
	BaseController
}

// swagger:route GET /device/isDeviceStockEnough device isDeviceStockEnough
// IsDeviceStockEnough 查询设备库存是否足够
//     Responses:
//       200: isDeviceStockEnough
//       default: ErrorResponse

func (c *DeviceController) IsDeviceStockEnough() {
	defer c.CatchException()

	cnt, _ := c.GetInt("count")
	req := requestTypes.QueryDeviceStockRequest{
		IdcID:        c.GetString("idc_id"),
		DeviceTypeId: c.GetString("deviceTypeId"),
		Count:        int64(cnt),
	}

	// req.Validate(c.logPoints)
	res, err := DeviceLogic.QueryDeviceStockCount(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = responseTypes.QueryDeviceStockResponse{
		Success:        res >= int64(cnt),
		AvaliableCount: res,
	}

}
