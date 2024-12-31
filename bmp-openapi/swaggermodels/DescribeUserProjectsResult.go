package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeUserProjects
type DescribeUserProjectsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ProjectList `json:"result"`
	}
}
