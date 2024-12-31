package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
)

// swagger:parameters deleteMessage
type DeleteMessageRequest struct {
	WriteHeader

	// in: body
	Body requestTypes.DeleteMessagesRequest
}
