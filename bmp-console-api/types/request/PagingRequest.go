package request

type PagingRequest struct {
	PageNumber int64 `json:"pageNumber"`
	PageSize   int64 `json:"pageSize"`
}
