package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDeviceDisks
type DescribeDeviceDisksResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceDisks `json:"result"`
	}
}
