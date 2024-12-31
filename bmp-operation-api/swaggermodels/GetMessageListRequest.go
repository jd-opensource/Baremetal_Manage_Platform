package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
)

// swagger:parameters getMessageList
type GetMessageListRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryMessagesRequest
}
