package bmpApi

type BmpAlertResponse struct {
	ReqiestId     string                   `json:"reqiestId"`
	Error         *BmpAlertErrorResponse   `json:"error"`
	Result        map[string]interface{}   `json:"result"`
}

type BmpAlertErrorResponse struct {
	Code       int      `json:"code"`
	Message    string	`json:"message"`
	Status     string   `json:"result"`
}
