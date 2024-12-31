package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeRoles
type DescribeRolesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryRolesRequest
}
