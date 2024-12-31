package util

type Pageable struct {
	PageNumber int64 `json:"page_number"`
	PageSize   int64 `json:"page_size"`
	TotalCount int64 `json:"total_count"`
}

func (p Pageable) Pageable2offset() (offset int64, limit int64) {
	limit = p.PageSize
	if p.PageNumber == 0 {
		offset = 0
	} else {
		offset = (p.PageNumber - 1) * p.PageSize
	}
	return offset, limit
}

func Offset2Pageable(total, limit, offset int64) Pageable {
	pageNumber := offset/limit + 1
	return Pageable{
		PageNumber: pageNumber,
		PageSize:   limit,
		TotalCount: total,
	}
}
