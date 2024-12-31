package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
)

// swagger:parameters readMessage
type ReadMessageRequest struct {
	WriteHeader

	// in: body
	Body requestTypes.ReadMessagesRequest
}
