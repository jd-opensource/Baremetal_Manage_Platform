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

type ProxyResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
}
