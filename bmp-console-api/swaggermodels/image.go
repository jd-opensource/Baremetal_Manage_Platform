package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters queryImagesByDeviceType
type QueryImagesByDeviceTypeRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryImagesByDeviceTypeRequest
}

// QueryImagesByDeviceTypeResponse is an response struct
// swagger:response queryImagesByDeviceType
type QueryImagesByDeviceTypeResponse struct {
	// in: body
	Body struct {
		Result    map[string][]*sdkModels.Image `json:"result"`
		RequestId string                        `json:"requestId"`
	}
}
