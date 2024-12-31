package response

import "git.jd.com/bmp-pronoea/types/request"

type MetricRangeQueryResponse struct {
	Data   []*MetricRangeQueryData                    `json:"data"`
	Query  *request.MetricRangeQueryRequest           `json:"query"`
}

type MetricRangeQueryData struct {
	Timestamp   int64 	`json:"timestamp"`
	Value       string 	`json:"value"`
}

