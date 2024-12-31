package request

type HeaderRequest struct {
	Region string `json:"region"` // 必填
	Az     string `json:"az"`
	Pin    string `json:"pin"`
	Tenant string `json:"tenant"`
}
