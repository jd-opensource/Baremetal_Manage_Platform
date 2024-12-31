package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeUsers
type DescribeUsersRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryUsersRequest
}
