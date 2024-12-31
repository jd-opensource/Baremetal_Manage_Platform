package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters getMessageList
type GetMessageListRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryMessagesRequest
}
