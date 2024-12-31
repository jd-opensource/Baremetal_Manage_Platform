package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A CreateRaidResult is an response struct that is used to describe create raid result.
// swagger:response createRaid
type CreateRaidResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RaidId `json:"result"`
	}
}
