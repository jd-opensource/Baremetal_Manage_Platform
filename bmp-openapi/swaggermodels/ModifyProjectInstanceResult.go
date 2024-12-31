package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyProjectInstanceResult is an response struct that is used to describe modify instance result.
// swagger:response modifyProjectInstance
type ModifyProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
