package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"

// swagger:response getMessageById
type GetMessageByIdResult struct {
	// in: body
	Body struct {
		Result    response.MessageWithNextPrev `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
