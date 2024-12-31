package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters getMessageById
type GetMessageByIdRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.GetMessageByIdRequest
}
