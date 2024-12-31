package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters describeOSs
type DescribeOSsRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryOssRequest
}

// DescribeOSsResponse is an response struct
// swagger:response describeOSs
type DescribeOSsResponse struct {
	// in: body
	Body struct {
		Result responseTypes.OsList `json:"result"`
		CommonRespBody
	}
}
