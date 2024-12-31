package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// A DeleteMessageResult is an response struct
// swagger:response deleteMessage
type DeleteMessageResult struct {
	// in: body
	Body struct {
		Result    response.CommonResponse `json:"result"`
		RequestId string                  `json:"requestId"`
	}
}
