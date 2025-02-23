package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:response hasUnreadMessage
type HasUnreadMessageResult struct {
	// in: body
	Body struct {
		Result    response.HasUnreadMessage `json:"result"`
		RequestId string                    `json:"requestId"`
	}
}
