package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:parameters getAvailableDeviceTypes
type GetAvailableDeviceTypesRequest struct {
	ReadHeader

	// in: query
	requestTypes.GetAvailableDeviceTypesRequest
}

// GetAvailableDeviceTypesResponse is an response struct
// swagger:response getAvailableDeviceTypes
type GetAvailableDeviceTypesResponse struct {
	// in: body
	Body struct {
		Result    map[string][]responseTypes.DeviceType `json:"result"`
		RequestId string                                `json:"requestId"`
	}
}
