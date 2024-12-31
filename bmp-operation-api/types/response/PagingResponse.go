package response

type PagingResponse struct {
	PageNumber int64 `json:"pageNumber"`

	PageSize int64 `json:"pageSize"`

	TotalCount int64 `json:"totalCount"`
}
