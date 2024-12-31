package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteInstancesResult is an response struct that is used to describe modify instance result.
// swagger:response deleteInstances
type DeleteInstancesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
