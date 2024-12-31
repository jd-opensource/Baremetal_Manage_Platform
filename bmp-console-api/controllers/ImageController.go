package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/ImageLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// ImageController operations for image
type ImageController struct {
	BaseController
}

// swagger:route GET /image/queryImagesByDeviceType image queryImagesByDeviceType
// QueryImagesByDeviceType 根据机型获取绑定的镜像列表
//     Responses:
//       200: queryImagesByDeviceType
//       default: ErrorResponse

func (c *ImageController) QueryImagesByDeviceType() {
	defer c.CatchException()

	req := requestTypes.QueryImagesByDeviceTypeRequest{
		IdcID:        c.GetString("idc_id"),
		DeviceTypeID: c.GetString("deviceTypeId"),
	}
	// req.Validate(c.logPoints)
	res, err := ImageLogic.QueryImagesByDeviceType(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	c.Res.Result = res

}
