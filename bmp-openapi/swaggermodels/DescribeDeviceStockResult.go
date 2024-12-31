package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetDeviceStockResult is an response struct that is used to get device stock.
// swagger:response describeDeviceStock
type DescribeDeviceStockResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceStock `json:"result"`
	}
}
