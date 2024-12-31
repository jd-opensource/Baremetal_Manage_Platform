package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-console-api/types/response"

// swagger:response getMessageStatistic
type GetMessageStatisticResult struct {
	// in: body
	Body struct {
		Result    response.MessageStatistic `json:"result"`
		RequestId string                    `json:"requestId"`
	}
}
