package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters queryRaidsAll
type QueryRaidsAllRequest struct {
	ReadHeader
}

// QueryRaidsAllResponse is an response struct
// swagger:response queryRaidsAll
type QueryRaidsAllResponse struct {
	// in: body
	Body struct {
		Result responseTypes.RaidList `json:"result"`
		CommonRespBody
	}
}
