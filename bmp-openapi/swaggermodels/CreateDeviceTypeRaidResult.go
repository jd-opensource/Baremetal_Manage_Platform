package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateDeviceTypeRaidResult is an response struct that is used to describe create device type raid result.
// swagger:response createDeviceTypeRaid
type CreateDeviceTypeRaidResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
