package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyDeviceTypeResult is an response struct that is used to modify device type.
// swagger:response modifyDeviceType
type ModifyDeviceTypeResult struct {
	// in: header
	ResponseHeader

	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
