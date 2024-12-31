package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateDeviceTypeResult is an response struct that is used to describe create deviceType result.
// swagger:response createDeviceType
type CreateDeviceTypeResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceTypeId `json:"result"`
	}
}
