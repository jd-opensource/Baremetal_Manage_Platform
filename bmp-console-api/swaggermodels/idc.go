package swaggermodels

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters getIdcList
type GetIdcListRequest struct {
	ReadHeader
	//in: query
	ExportType string `json:"exportType"`
}

// GetIdcListResponse is an response struct
// swagger:response getIdcList
type GetIdcListResponse struct {
	// in: body
	Body struct {
		Result    sdkModels.IdcList `json:"result"`
		RequestId string            `json:"requestId"`
	}
}
