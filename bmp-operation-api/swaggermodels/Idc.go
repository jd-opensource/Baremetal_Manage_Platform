package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters getIdcList
type GetIdcListRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryIdcListRequest
}

// GetIdcListResponse is an response struct
// swagger:response getIdcList
type GetIdcListResponse struct {
	// in: body
	Body struct {
		Result responseTypes.IdcPage `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters getIdcInfo
type GetIdcInfoRequest struct {
	ReadHeader
	// in: path
	IdcID string `json:"idc_id"`
}

// GetIdcInfoResponse is an response struct
// swagger:response getIdcInfo
type GetIdcInfoResponse struct {
	// in: body
	Body struct {
		Result responseTypes.Idc `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteIdc
type DeleteIdcRequest struct {
	WriteHeader
	// in: path
	IdcID string `json:"idc_id"`
}

// DeleteIdcResponse is an response struct
// swagger:response deleteIdc
type DeleteIdcResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters createIdc
type CreateIdcRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.CreateIdcRequest
}

// CreateIdcResponse is an response struct
// swagger:response createIdc
type CreateIdcResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CreateIdcResult `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters modifyIdcInfo
type ModifyIdcInfoRequest struct {
	WriteHeader
	// in: path
	IdcID string `json:"idc_id"`
	// in: body
	Body requestTypes.ModifyIdcRequest
}

// ModifyIdcInfoResponse is an response struct
// swagger:response modifyIdcInfo
type ModifyIdcInfoResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters verifyUser
type VerifyUserRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.VerifyUserRequest
}

// VerifyUserResponse is an response struct
// swagger:response verifyUser
type VerifyUserResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}
