package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteDeviceTypeRaidResult is an response struct that is used to describe delete device type raid result.
// swagger:response deleteDeviceTypeRaid
type DeleteDeviceTypeRaidResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
