package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters createApikey
type CreateApikeyRequest struct {
	WriteHeader

	// in:body
	Body sdkModels.CreateApikeyRequest
}

// CreateApikeyResponse is an response struct
// swagger:response createApikey
type CreateApikeyResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.ApikeyId `json:"result"`
		RequestId string                 `json:"requestId"`
	}
}

// swagger:parameters getApikeyList
type GetApikeyListRequest struct {
	ReadHeader

	// in: query
	requestTypes.GetApikeyListRequest
}

// GetApikeyListResponse is an response struct
// swagger:response getApikeyList
type GetApikeyListResponse struct {
	// in: body
	Body struct {
		Result    sdkModels.ApikeyList `json:"result"`
		RequestId string               `json:"requestId"`
	}
}

// swagger:parameters deleteApikey
type DeleteApikeyRequest struct {
	ReadHeader

	// in: path
	ApikeyID string `json:"apikey_id"`
}

// DeleteApikeyResponse is an response struct
// swagger:response deleteApikey
type DeleteApikeyResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
