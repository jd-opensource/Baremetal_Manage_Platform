package controllers

import (
	RaidLogic "coding.jd.com/aidc-bmp/bmp-operation-api/logic/RaidLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTypeController operations for Device
type RaidController struct {
	BaseController
}

// QueryRaidsAll ...
// swagger:route GET /raids raid queryRaidsAll
//
//     Responses:
//       200: queryRaidsAll
//       default: ErrorResponse

func (c *RaidController) QueryRaidsAll() {
	defer c.CatchException()
	res, err := RaidLogic.QueryRaidsAll(c.logPoints)
	if err != nil {
		c.SetErrorResponse(httpStatus.INTERNAL_SERVER_ERROR, err.Error(), errorCode.INTERNAL)
		return
	}
	if res.Raids == nil {
		res.Raids = make([]*response.Raid, 0)
	}
	c.Res.Result = res
}

// swagger:route GET /raid/queryRaidsByDeviceTypeIDAndVolumeType raid queryRaidsByDeviceTypeIDAndVolumeType
// QueryRaidsByDeviceTypeIDAndVolumeType 根据机型查询系统盘/数据盘可用raid列表
//     Responses:
//       200: queryRaidsByDeviceTypeIDAndVolumeType
//       default: ErrorResponse

func (c *RaidController) QueryRaidsByDeviceTypeIDAndVolumeType() {
	defer c.CatchException()

	req := requestTypes.QueryRaidsRequest{
		DeviceTypeID: c.GetString("deviceTypeId"),
		VolumeType:   c.GetString("volumeType"),
	}

	raids, err := RaidLogic.QueryRaids(c.logPoints, req)
	if err != nil {
		c.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.PARAMETER_ERROR)
		return
	}

	/*dt, _ := DeviceTypeLogic.GetDeviceTypeInfo(c.logPoints, &requestTypes.QueryDeviceTypeRequest{DeviceTypeID: req.DeviceTypeID})
	if req.VolumeType == "system" {
		for idx, raid := range raids {
			if raid.RaidName == "NORAID" { //产品坚持认为，两块盘的noraid要叫单盘RAID0
				if dt.SystemVolumeAmount == 2 && dt.RaidCan == "RAID" {
					raids[idx].RaidName = "单盘RAID0"
					if c.logPoints.GetPoint("language").(string) == BaseLogic.EN_US {
						raids[idx].RaidName = "RAID0-stripping"
					}
				} else if dt.RaidCan == "NORAID" {
					raids[idx].RaidName = "未配置RAID"
					if c.logPoints.GetPoint("language").(string) == BaseLogic.EN_US {
						raids[idx].RaidName = "RAID is not configured"
					}
				}

			}
		}
	}*/
	c.Res.Result = raids
	return
}
