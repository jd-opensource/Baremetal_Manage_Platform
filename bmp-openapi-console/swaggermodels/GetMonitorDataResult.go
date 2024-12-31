package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response getMonitorData
type GetMonitorDataResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []*responseTypes.DataEveryMetric `json:"result"`
	}
}
