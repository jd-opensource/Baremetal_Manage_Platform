package request

type QueryCustomInfoRequest struct {
	PageKey string `json:"pageKey"`
	Reload  string `json:"reload"`
}
type SetCustomInfoRequest struct {
	PageKey   string                     `json:"pageKey"`
	PageValue map[string]map[string]bool `json:"pageValue"`
}
type QueryPrivilegeNewRequest struct {
	Pin string `json:"pin"`
}
