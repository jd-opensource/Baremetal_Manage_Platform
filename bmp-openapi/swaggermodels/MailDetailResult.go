package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response mailDetail
type MailDetailResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.Mail `json:"result"`
	}
}
