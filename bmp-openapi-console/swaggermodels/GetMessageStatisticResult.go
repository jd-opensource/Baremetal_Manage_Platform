package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response getMessageStatistic
type GetMessageStatisticResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.MessageStatistic `json:"result"`
	}
}