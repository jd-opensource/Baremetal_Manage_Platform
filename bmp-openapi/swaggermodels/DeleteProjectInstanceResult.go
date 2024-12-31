package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteProjectInstanceResult is an response struct that is used to delete stop instance result.
// swagger:response deleteProjectInstance
type DeleteProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
