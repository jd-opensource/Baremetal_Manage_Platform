package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/RaidLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// RaidController operations for raid
type RaidController struct {
	BaseController
}

// swagger:route GET /raid/queryRaidsByDeviceTypeIDAndVolumeType raid queryVolumeRaids
// QueryVolumeRaids 根据机型查询系统卷/数据卷和可用raid列表
//     Responses:
//       200: queryVolumeRaids
//       default: ErrorResponse

func (c *RaidController) QueryVolumeRaids() {
	defer c.CatchException()

	req := requestTypes.QueryVolumeRaidsRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		VolumeType:   c.GetString("volumeType"),
	}
	if err := req.Validate(c.logPoints); err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.PARAMETER_ERROR)
		return
	}
	raids, err := RaidLogic.QueryVolumesRaids(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.PARAMETER_ERROR)
		return
	}

	c.Res.Result = raids
	return
}
