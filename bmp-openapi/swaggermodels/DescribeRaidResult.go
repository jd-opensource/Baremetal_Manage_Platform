package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetRaidResult is an response struct that is used to describe raid.
// swagger:response describeRaid
type DescribeRaidResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.Raid `json:"result"`
	}
}
