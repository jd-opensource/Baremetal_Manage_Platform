package request

type CheckKeypairNameRequest struct {
	Name string `json:"name"`
}

type CreateKeypairRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	ClientToken int    `json:"clientToken"`
}

type ModifyKeypairRequest struct {
	Name        string `json:"name"`
	ClientToken int    `json:"clientToken"`
}

type DeleteKeypairRequest struct {
	KeypairID string `json:"keypairId" valid:"Required"`
}

type QueryKeypairsRequest struct {
	PagingRequest

	Name string `json:"name"`

	KeypairIds []string `json:"keypairId"`
}

type QueryKeypairRequest struct {
	KeypairID string `json:"keypairId" valid:"Required"`
}
