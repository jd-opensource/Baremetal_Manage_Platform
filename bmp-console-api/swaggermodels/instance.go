package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters createInstance
type CreateInstanceRequest struct {
	WriteHeader

	// in:body
	Body sdkModels.CreateInstanceRequest
}

// CreateInstanceResponse is an response struct
// swagger:response createInstance
type CreateInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CreateInstanceResponse `json:"result"`
		RequestId string                               `json:"requestId"`
	}
}

// swagger:parameters getInstanceList
type GetInstanceListRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryInstancesRequest
}

// GetInstanceListResponse is an response struct
// swagger:response getInstanceList
type GetInstanceListResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.InstancesResponse `json:"result"`
		RequestId string                          `json:"requestId"`
	}
}

// swagger:parameters getInstanceInfo
type GetInstanceInfoRequest struct {
	ReadHeader

	// in: path
	InstanceID string `json:"instance_id"`
}

// GetInstanceInfoResponse is an response struct
// swagger:response getInstanceInfo
type GetInstanceInfoResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.InstanceInfo `json:"result"`
		RequestId string                     `json:"requestId"`
	}
}

// swagger:parameters startInstance
type StartInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StartInstanceRequest
}

// StartInstanceResponse is an response struct
// swagger:response startInstance
type StartInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters stopInstance
type StopInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StopInstanceRequest
}

// StopInstanceResponse is an response struct
// swagger:response stopInstance
type StopInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters restartInstance
type RestartInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.RestartInstanceRequest
}

// RestartInstanceResponse is an response struct
// swagger:response restartInstance
type RestartInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters lockInstance
type LockInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.LockInstanceRequest
}

// LockInstanceResponse is an response struct
// swagger:response lockInstance
type LockInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters unlockInstance
type UnlockInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.UnLockInstanceRequest
}

// UnlockInstanceResponse is an response struct
// swagger:response unlockInstance
type UnlockInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters modifyInstance
type ModifyInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ModifyInstanceRequest
}

// ModifyInstanceResponse is an response struct
// swagger:response modifyInstance
type ModifyInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters deleteInstance
type DeleteInstanceRequest struct {
	ReadHeader

	// in: path
	InstanceID string `json:"instance_id"`
}

// DeleteInstanceResponse is an response struct
// swagger:response deleteInstance
type DeleteInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters startInstances
type StartInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StartInstancesRequest
}

// StartInstancesResponse is an response struct
// swagger:response startInstances
type StartInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters stopInstances
type StopInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StartInstancesRequest
}

// StopInstancesResponse is an response struct
// swagger:response stopInstances
type StopInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters restartInstances
type RestartInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StartInstancesRequest
}

// RestartInstancesResponse is an response struct
// swagger:response restartInstances
type RestartInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters modifyInstances
type ModifyInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ModifyInstancesRequest
}

// ModifyInstancesResponse is an response struct
// swagger:response modifyInstances
type ModifyInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters deleteInstances
type DeleteInstancesRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.StartInstancesRequest
}

// DeleteInstancesResponse is an response struct
// swagger:response deleteInstances
type DeleteInstancesResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
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

// swagger:parameters reinstallInstance
type ReinstallInstanceRequest struct {
	WriteHeader

	// in:body
	Body requestTypes.ReinstallInstanceRequest
}

// ReinstallInstanceResponse is an response struct
// swagger:response reinstallInstance
type ReinstallInstanceResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.CommonResponse `json:"result"`
		RequestId string                       `json:"requestId"`
	}
}

// swagger:parameters getInstanceListForShare
type GetInstanceListForShareRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryInstancesForShareRequest
}

// GetInstanceListForShareResponse is an response struct
// swagger:response getInstanceListForShare
type GetInstanceListForShareResponse struct {
	// in: body
	Body struct {
		Result    []*sdkModels.InstanceForShare `json:"result"`
		RequestId string                        `json:"requestId"`
	}
}
