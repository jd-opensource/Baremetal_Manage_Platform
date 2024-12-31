package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:parameters isDeviceStockEnough
type IsDeviceStockEnoughRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryDeviceStockRequest
}

// IsDeviceStockEnoughResponse is an response struct
// swagger:response isDeviceStockEnough
type IsDeviceStockEnoughResponse struct {
	// in: body
	Body struct {
		Result    responseTypes.QueryDeviceStockResponse `json:"result"`
		RequestId string                                 `json:"requestId"`
	}
}
