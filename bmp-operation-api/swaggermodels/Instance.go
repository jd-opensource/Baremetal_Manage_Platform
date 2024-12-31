package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters startInstance
type StartInstanceRequest struct {
	WriteHeader
	// in: path
	InstanceID string `json:"instance_id"`
}

// StartInstanceResponse is an response struct
// swagger:response startInstance
type StartInstanceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters restartInstance
type RestartInstanceRequest struct {
	WriteHeader
	// in: path
	InstanceID string `json:"instance_id"`
}

// RestartInstanceResponse is an response struct
// swagger:response restartInstance
type RestartInstanceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters stopInstance
type StopInstanceRequest struct {
	WriteHeader
	// in: path
	InstanceID string `json:"instance_id"`
}

// StopInstanceResponse is an response struct
// swagger:response stopInstance
type StopInstanceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters resetInstanceStatus
type ResetInstanceStatusRequest struct {
	WriteHeader
	// in: path
	InstanceID string `json:"instance_id"`
}

// ResetInstanceStatus is an response struct
// swagger:response resetInstanceStatus
type ResetInstanceStatusResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteInstance
type DeleteInstanceRequest struct {
	WriteHeader
	// in: path
	InstanceID string `json:"instance_id"`
}

// DeleteInstanceResponse is an response struct
// swagger:response deleteInstance
type DeleteInstanceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters resetPasswd
type ResetPasswdRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ResetPasswdRequest
}

// ResetPasswdResponse is an response struct
// swagger:response resetPasswd
type ResetPasswdResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters resetInstancesPasswd
type ResetInstancesPasswdRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ResetInstancesPasswdRequest
}

// ResetInstancesPasswdResponse is an response struct
// swagger:response resetInstancesPasswd
type ResetInstancesPasswdResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}
