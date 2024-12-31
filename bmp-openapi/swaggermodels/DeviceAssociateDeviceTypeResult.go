package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeviceAssociateDeviceTypeResult is an response struct
// swagger:response deviceAssociateDeviceType
type DeviceAssociateDeviceTypeResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
