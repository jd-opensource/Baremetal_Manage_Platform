package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribIdcsResult is an response struct that is used to describe idcs.
// swagger:response describeIdcs
type DescribIdcsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.IdcList `json:"result"`
	}
}
