package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A MountDeviceResult is an response struct that is used to describe mount device result.
// swagger:response mountDevice
type MountDeviceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
