package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyRaidResult is an response struct that is used to describe modify raid.
// swagger:response modifyRaid
type ModifyRaidResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
