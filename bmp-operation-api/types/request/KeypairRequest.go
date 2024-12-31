package request

type QueryKeypairsRequest struct {
	PagingRequest

	Name string `json:"name"`

	KeypairIds []string `json:"keypairId"`
}
