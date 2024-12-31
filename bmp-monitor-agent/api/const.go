package api

const (
	API_OK = iota
	API_NO_PARAMETER
	API_PARAMETER
	API_NORESULT
	API_REMOTE_ERROR
	API_UNKNOWN   = 99
	API_URL_ERROR = 403
)

var APICodeMap = map[int]string{
	API_OK:           "Success",
	API_NO_PARAMETER: "No Parameter",
	API_PARAMETER:    "Parameter Error",
	API_NORESULT:     "No Result",
	API_REMOTE_ERROR: "Remote Error",
	API_UNKNOWN:      "Unknown",
	API_URL_ERROR:    "url error",
}

type ErrorResponse struct {
	// 错误码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 错误状态
	Status string `json:"status"`
}

type ProxyResponse struct {
	// 操作失败结果。成功时有此结构
	Result interface{} `json:"result,omitempty"`
	// 操作失败结果。失败时有此结构
	Error *ErrorResponse `json:"error,omitempty"`
	// 请求traceId
	// required: true
	RequestId string `json:"requestId"`
}
