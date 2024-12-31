package response

type QueryDeviceStockResponse struct {
	Success        bool  `json:"success"`
	AvaliableCount int64 `json:"avaliableCount"`
}
