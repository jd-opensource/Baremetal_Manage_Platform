package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDevice
type DescribeDeviceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceInfo `json:"result"`
	}
}
