package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A RemoveDeviceResult is an response struct that is used to describe restart instance result.
// swagger:response removeDevice
type RemoveDeviceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
