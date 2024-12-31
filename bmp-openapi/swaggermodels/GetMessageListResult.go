package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response getMessageList
type GetMessageListResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.MessageList `json:"result"`
	}
}
