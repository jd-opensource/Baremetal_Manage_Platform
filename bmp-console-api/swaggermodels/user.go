package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters login
type LoginRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.LoginRequest
}

// LoginResponse is an response struct
// swagger:response login
type LoginResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters logout
type LogoutRequest struct {
	WriteHeader
}

// LogoutResponse is an response struct
// swagger:response logout
type LogoutResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters getUserInfo
type GetUserInfoRequest struct {
	ReadHeader
}

// GetUserInfoResponse is an response struct
// swagger:response getUserInfo
type GetUserInfoResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.User `json:"result"`
		RequestId string             `json:"requestId"`
	}
}

// swagger:parameters updatePassword
type UpdatePasswordRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.UpdatePasswordRequest
}

// UpdatePasswordResponse is an response struct
// swagger:response updatePassword
type UpdatePasswordResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters updateUserInfo
type UpdateUserInfoRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ModifyUserInfoRequest
}

// UpdateUserInfoResponse is an response struct
// swagger:response updateUserInfo
type UpdateUserInfoResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters getTimezoneList
type GetTimezoneListRequest struct {
	ReadHeader
}

// GetTimezoneListResponse is an response struct
// swagger:response getTimezoneList
type GetTimezoneListResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.GetTimezoneListResponse `json:"result"`
		RequestId string                                `json:"requestId"`
	}
}

// swagger:parameters getConsoleAccessUserList
type GetConsoleAccessUserListRequest struct {
	ReadHeader
}

// GetConsoleAccessUserListResponse is an response struct
// swagger:response getConsoleAccessUserList
type GetConsoleAccessUserListResponse struct {
	// in: body
	Body struct {
		Result    []*models.User `json:"result"`
		RequestId string         `json:"requestId"`
	}
}

// swagger:parameters checkUserConsoleAccess
type CheckUserConsoleAccessRequest struct {
	ReadHeader

	// in: query
	requestTypes.CheckUserConsoleAccessRequest
}

// CheckUserConsoleAccessResponse is an response struct
// swagger:response checkUserConsoleAccess
type CheckUserConsoleAccessResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters checkUserConsoleAccessByName
type CheckUserConsoleAccessByNameRequest struct {
	ReadHeader

	// in: query
	requestTypes.CheckUserConsoleAccessByNameRequest
}

// CheckUserConsoleAccessByNameResponse is an response struct
// swagger:response checkUserConsoleAccessByName
type CheckUserConsoleAccessByNameResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
