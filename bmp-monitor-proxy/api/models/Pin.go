package models

type PinResult struct {
	RequestID string `json:"requestId"`
	Result    struct {
		Pin    string `json:"pin,omitempty"`
		Tenant string `json:"tenant,omitempty"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}
