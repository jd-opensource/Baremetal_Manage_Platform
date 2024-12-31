package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetAssociatedDisksResult is an response struct that is used to get instance list by sshkey.
// swagger:response getAssociatedDisks
type GetAssociatedDisksResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []*response.Disk `json:"result"`
	}
}
