package dao

type QueryCustomInfoRequest struct {
	PageKey string `json:"page_key"`
	Reload  string `json:"reload"`
}
type SetCustomInfoRequest struct {
	PageKey   string                     `json:"page_key"`
	PageValue map[string]map[string]bool `json:"page_value"`
}
