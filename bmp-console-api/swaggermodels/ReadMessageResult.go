package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// A ReadMessageResult is an response struct
// swagger:response readMessage
type ReadMessageResult struct {
	// in: body
	Body struct {
		Result    response.CommonResponse `json:"result"`
		RequestId string                  `json:"requestId"`
	}
}
