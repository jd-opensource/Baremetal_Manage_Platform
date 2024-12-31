package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response hasUnreadMessage
type HasUnreadMessageResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.HasUnreadMessage `json:"result"`
	}
}
