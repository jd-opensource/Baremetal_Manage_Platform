package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

// swagger:parameters queryDefaultSystemPartitions
type QueryDefaultSystemPartitionsRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryDefaultSystemPartitionRequest
}

// QueryDefaultSystemPartitionsResponse is an response struct
// swagger:response queryDefaultSystemPartitions
type QueryDefaultSystemPartitionsResponse struct {
	// in: body
	Body struct {
		Result    []*sdkModels.Partition `json:"result"`
		RequestId string                 `json:"requestId"`
	}
}
