package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters uploadDevices
type UploadDevicesRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.UploadDeviceRequest
}

// swagger:response uploadDevices
type UploadDevicesResponse struct {
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.UploadDevice `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters createDevices
type CreateDevicesRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.CreateDeviceRequest
}

// swagger:response createDevices
type CreateDevicesResponse struct {
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.DeviceIds `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters queryDevices
type QueryDevicesRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryDeviceListRequest
}

// QueryDevicesResponse is an response struct
// swagger:response queryDevices
type QueryDevicesResponse struct {
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.DevicePage `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters getDevice
type GetDeviceRequest struct {
	ReadHeader
	// in: path
	DeviceID string `json:"device_id"`
}

// GetDeviceResponse is an response struct
// swagger:response getDevice
type GetDeviceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.Device `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters unPutaway
type UnPutawayRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.UnPutawayDeviceRequest
}

// UnPutawayResponse is an response struct
// swagger:response unPutaway
type UnPutawayResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters putaway
type PutawayRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.PutawayDeviceRequest
}

// PutawayResponse is an response struct
// swagger:response putaway
type PutawayResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteDevice
type DeleteDeviceRequest struct {
	WriteHeader
	// in: path
	DeviceID string `json:"device_id"`
}

// DeleteDeviceResponse is an response struct
// swagger:response deleteDevice
type DeleteDeviceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters removeDevice
type RemoveDeviceRequest struct {
	WriteHeader
	// in: path
	DeviceID string `json:"device_id"`
}

// RemoveDeviceResponse is an response struct
// swagger:response removeDevice
type RemoveDeviceResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}
