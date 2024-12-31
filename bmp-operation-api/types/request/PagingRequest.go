package request

type PagingRequest struct {
	PageNumber int64 `json:"pageNumber" validate:"gte=1"`
	PageSize   int64 `json:"pageSize" validate:"lte=1000,gte=1"`
}
