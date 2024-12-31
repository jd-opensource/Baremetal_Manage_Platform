package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeUserProjects
type DescribeUserProjectsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryProjectsRequest
}
