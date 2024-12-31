package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDevices
type DescribeDevicesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceList `json:"result"`
	}
}
