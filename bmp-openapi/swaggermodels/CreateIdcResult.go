package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateIdcResult is an response struct that is used to describe create idc result.
// swagger:response createIdc
type CreateIdcResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.IdcID `json:"result"`
	}
}
