package controllers

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/TimezoneLogic"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

type TimezoneController struct {
	BaseController
}

// swagger:route GET /user/timezones user getTimezoneList
// GetTimezoneList 获取时区列表
//     Responses:
//       200: getTimezoneList
//       default: ErrorResponse

func (c *TimezoneController) GetTimezoneList() {

	res := TimezoneLogic.GetTimeZoneList(c.logPoints)
	c.Res.Result = responseTypes.GetTimezoneListResponse{
		TimeZones: res,
	}

}
