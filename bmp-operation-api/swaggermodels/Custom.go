package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
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
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
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
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}
