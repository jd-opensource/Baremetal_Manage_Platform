package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CancelShareUserProjectResult is an response struct
// swagger:response cancelShareUserProject
type CancelShareUserProjectResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
