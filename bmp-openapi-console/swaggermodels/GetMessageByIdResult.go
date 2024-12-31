package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response getMessageById
type GetMessageByIdResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.MessageWithNextPrev `json:"result"`
	}
}
