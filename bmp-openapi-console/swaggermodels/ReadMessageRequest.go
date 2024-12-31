package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters readMessage
type ReadMessageRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.ReadMessagesRequest
}
