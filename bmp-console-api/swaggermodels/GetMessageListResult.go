package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-console-api/types/response"

// swagger:response getMessageList
type GetMessageListResult struct {
	// in: body
	Body struct {
		Result    response.MessageList `json:"result"`
		RequestId string               `json:"requestId"`
	}
}
