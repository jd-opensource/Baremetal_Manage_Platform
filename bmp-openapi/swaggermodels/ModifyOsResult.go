package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyOsResult is an response struct that is used to describe modify os.
// swagger:response modifyOS
type ModifyOsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
