package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:parameters getKeypairList
type GetKeypairListRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryKeypairsRequest
}

// GetKeypairListResponse is an response struct
// swagger:response getKeypairList
type GetKeypairListResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.KeyPairsResponse `json:"result"`
		RequestId string                         `json:"requestId"`
	}
}

// swagger:parameters getKeypairInfo
type GetKeypairInfoRequest struct {
	ReadHeader

	// in: path
	KeypairID string `json:"keypair_id"`
}

// GetKeypairInfoResponse is an response struct
// swagger:response getKeypairInfo
type GetKeypairInfoResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.KeyPairInfo `json:"result"`
		RequestId string                    `json:"requestId"`
	}
}

// swagger:parameters createKeypair
type CreateKeypairRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.CreateKeypairRequest
}

// CreateKeypairResponse is an response struct
// swagger:response createKeypair
type CreateKeypairResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.KeyPair `json:"result"`
		RequestId string                `json:"requestId"`
	}
}

// swagger:parameters deleteKeypair
type DeleteKeypairRequest struct {
	ReadHeader

	// in: path
	KeypairID string `json:"keypair_id"`
}

// DeleteKeypairResponse is an response struct
// swagger:response deleteKeypair
type DeleteKeypairResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters modifyKeypair
type ModifyKeypairRequest struct {
	ReadHeader

	// in: path
	KeypairID string `json:"keypair_id"`

	// in:body
	Body requestTypes.ModifyKeypairRequest
}

// ModifyKeypairResponse is an response struct
// swagger:response modifyKeypair
type ModifyKeypairResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters checkKeypairName
type CheckKeypairNameRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryKeypairsRequest
}

// CheckKeypairNameResponse is an response struct
// swagger:response checkKeypairName
type CheckKeypairNameResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
