package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetInstancesBySshkeyResult is an response struct that is used to get instance list by sshkey.
// swagger:response getInstancesBySshkey
type GetInstancesBySshkeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.InstancesSshkeyInfoResponse `json:"result"`
	}
}
