package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
)

// swagger:parameters readMessage
type ReadMessageRequest struct {
	WriteHeader

	// in: body
	Body requestTypes.ReadMessagesRequest
}
