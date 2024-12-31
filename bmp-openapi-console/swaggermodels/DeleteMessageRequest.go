package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters deleteMessage
type DeleteMessageRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.DeleteMessagesRequest
}
