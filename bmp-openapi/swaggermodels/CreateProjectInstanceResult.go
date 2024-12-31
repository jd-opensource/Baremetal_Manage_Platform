package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateProjectInstanceResult is an response struct that is used to describe create instance result.
// swagger:response createProjectInstance
type CreateProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.InstanceIds `json:"result"`
	}
}
