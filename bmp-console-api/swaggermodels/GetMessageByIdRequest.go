package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
)

// swagger:parameters getMessageById
type GetMessageByIdRequest struct {
	ReadHeader

	// in: query
	requestTypes.GetMessageByIdRequest
}