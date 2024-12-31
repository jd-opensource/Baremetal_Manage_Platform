package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeUserProject
type DescribeUserProjectResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ProjectInfo `json:"result"`
	}
}
