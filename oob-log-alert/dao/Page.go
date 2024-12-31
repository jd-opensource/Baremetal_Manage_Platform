package dao

type PageResult struct {
	Total      int64 `json:"total_count"`
	PageNumber int64 `json:"page_number"`
	PageSize   int64 `json:"page_size"`
}

type PageRequest struct {
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
	IsAll    string `json:"isAll"` //非空表示全部
}
