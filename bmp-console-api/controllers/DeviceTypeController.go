package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/DeviceTypeLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for DeviceType
type DeviceTypeController struct {
	BaseController
}

// swagger:route GET /deviceType/getAvailableDeviceTypes deviceType getAvailableDeviceTypes
// GetAvailableDeviceTypes 获取该机房内可用机型列表
//     Responses:
//       200: getAvailableDeviceTypes
//       default: ErrorResponse

func (c *DeviceTypeController) GetAvailableDeviceTypes() {
	defer c.CatchException()
	req := requestTypes.GetAvailableDeviceTypesRequest{
		IdcID: c.GetString("idc_id"),
	}

	// req.Validate(c.logPoints)
	res, err := DeviceTypeLogic.GetAvailableDeviceTypes(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res

}
