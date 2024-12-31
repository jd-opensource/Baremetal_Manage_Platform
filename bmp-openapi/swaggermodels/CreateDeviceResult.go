package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response createDevice
type CreateDeviceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
