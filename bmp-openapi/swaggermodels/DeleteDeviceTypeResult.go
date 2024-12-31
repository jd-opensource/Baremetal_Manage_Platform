package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteDeviceTypeResult is an response struct that is used to delete device type.
// swagger:response deleteDeviceType
type DeleteDeviceTypeResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
