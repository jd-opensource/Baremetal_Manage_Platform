package request

type HeaderRequest struct {
	IdcID string `json:"idc_id"` // 必填
	Token string `json:"token"`
}
