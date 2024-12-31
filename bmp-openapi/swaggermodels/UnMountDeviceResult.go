package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A UnMountDeviceResult is an response struct that is used to unmount device.
// swagger:response unMountDevice
type UnMountDeviceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
