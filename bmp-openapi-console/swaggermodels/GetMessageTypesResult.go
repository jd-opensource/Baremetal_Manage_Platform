package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response getMessageTypes
type GetMessageTypesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.MessageTypesResp `json:"result"`
	}
}
