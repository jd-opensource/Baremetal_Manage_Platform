package swaggermodels

type ReadHeader struct {
	// required: true
	// in: header
	TraceID string `json:"traceId"` //同APIKey
	// in: header
	Cookie string `json:"cookie"`
}

type WriteHeader struct {
	// required: true
	// in: header
	TraceID string `json:"traceId"`
	// in: header
	Cookie string `json:"cookie"`
}

// 返回body里面的公共字段
type ResponseHeader struct {
	// in: header
	TraceID string `json:"traceId"`
}
