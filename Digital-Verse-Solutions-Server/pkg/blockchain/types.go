package blockchain

type NftJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type NeoStorageResponse struct {
	Ok    bool
	Value ValueField
}

type ValueField struct {
	Cid     string
	Created string
}

type Resp struct {
	FileURL         string                 `json:"file_url"`
	TransactionHash string                 `json:"tx_hash"`
	TransactionURL  string                 `json:"url"`
	Extra           map[string]interface{} `json:"extra"`
}
