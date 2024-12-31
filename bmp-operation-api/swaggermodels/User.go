package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
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
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
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
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters getUserList
type GetUserListRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryUsersRequest
}

// GetUserListResponse is an response struct
// swagger:response getUserList
type GetUserListResponse struct {
	// in: body
	Body struct {
		Result responseTypes.UserPage `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters addUser
type AddUserRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.AddUserRequest
}

// AddUserResponse is an response struct
// swagger:response addUser
type AddUserResponse struct {
	// in: body
	Body struct {
		Result responseTypes.UserAdd `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters getUserInfo
type GetUserInfoRequest struct {
	WriteHeader
	// in: path
	UserId string `json:"user_id"`
}

// GetUserInfoResponse is an response struct
// swagger:response getUserInfo
type GetUserInfoResponse struct {
	// in: body
	Body struct {
		Result responseTypes.User `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters modifyUserInfo
type ModifyUserInfoRequest struct {
	WriteHeader
	// in: path
	UserId string `json:"user_id"`
	// in: body
	Body requestTypes.ModifyUserInfoRequest
}

// ModifyUserInfoResponse is an response struct
// swagger:response modifyUserInfo
type ModifyUserInfoResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteUserInfo
type DeleteUserInfoRequest struct {
	WriteHeader
	// in: path
	UserId string `json:"user_id"`
}

// DeleteUserInfoResponse is an response struct
// swagger:response deleteUserInfo
type DeleteUserInfoResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters verifyLoginUser
type VerifyLoginUserRequest struct {
	ReadHeader
	// in:body
	Body requestTypes.LoginRequest
}

// VerifyLoginUserResponse is an response struct
// swagger:response verifyLoginUser
type VerifyLoginUserResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}
