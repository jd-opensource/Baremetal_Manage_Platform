package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// QueryRaidsAllResult is an response struct that is used to query all raids.
// swagger:response describeRaids
type DescribeRaidsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RaidList `json:"result"`
	}
}
