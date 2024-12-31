package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:parameters setCustomInfo
type SetCustomInfoRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.SetCustomInfoRequest
}

// SetCustomInfoResponse is an response struct
// swagger:response setCustomInfo
type SetCustomInfoResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters getCustomInfo
type GetCustomInfoRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryCustomInfoRequest
}

// GetCustomInfoResponse is an response struct
// swagger:response getCustomInfo
type GetCustomInfoResponse struct {

	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
