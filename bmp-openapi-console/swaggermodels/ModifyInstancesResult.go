package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyInstancesResult is an response struct that is used to describe modify instance result.
// swagger:response modifyInstances
type ModifyInstancesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
