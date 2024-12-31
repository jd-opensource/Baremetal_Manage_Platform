package response

type KeyPairsResponse struct {
	PagingResponse
	Keypairs []KeyPair `json:"keypairs"`
}

//给列表用的
type KeyPair struct {
	KeypairId     string `json:"keypairId"`
	Name          string `json:"name"`
	IsCheckDelete string `json:"isCheckDelete"`
	PublicKey     string `json:"publicKey"`
	PrivateKey    string `json:"privateKey"`
	FingerPrint   string `json:"fingerPrint"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}

//给详情用的
type KeyPairInfo struct {
	Keypair struct {
		KeypairId     string `json:"keypairId"`
		Name          string `json:"name"`
		IsCheckDelete string `json:"isCheckDelete"`
		PublicKey     string `json:"publicKey"`
		PrivateKey    string `json:"privateKey"`
		FingerPrint   string `json:"fingerPrint"`
		CreatedTime   string `json:"createdTime"`
		UpdatedTime   string `json:"updatedTime"`
	} `json:"keypair"`
}
