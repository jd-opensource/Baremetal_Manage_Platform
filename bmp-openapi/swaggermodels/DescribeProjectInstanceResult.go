package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetInstanceResult is an response struct that is used to describe instance.
// swagger:response describeProjectInstance
type GetInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.InstanceInfo `json:"result"`
	}
}
