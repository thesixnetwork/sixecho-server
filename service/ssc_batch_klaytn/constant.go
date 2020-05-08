package main

const (
	TransactionAlias string = "ssc_transactions"
	AccountAlias     string = "ssc_accounts"
	ImageAlias       string = "ssc_images"
)

//Owner struct
type Owner struct {
	Owner    string `json:"owner"`
	RefOwner string `json:"ref_owner"`
}

//AccountKlaytn struct from klaytn
type AccountKlaytn struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

//Transaction struct
type Transaction struct {
	ID           string
	KlaytnTxID   string `json:"klaytn_tx_id"`
	BlockNumber  int64  `json:"block_num"`
	ToUser       Owner  `json:"to_user"`
	Platform     string `json:"platform"`
	FromPlatform string `json:"from_platform"`
	AssetID      string `json:"asset_id"`
	Type         string `json:"asset_type"`
}

//Transaction struct
type TransactionImage struct {
	ID    string
	Title string `json:"title"`
}

// SnapPictures struct
type SnapPictures struct {
	TxID          *string `json:"tx_id,omitempty"`
	PublicChainID *string `json:"public_chain_id"`
	SnapID        *string `json:"snap_id"`
}

// Account strcut
type Account struct {
	ID              string   `json:"_id,omitempty"`
	Platform        string   `json:"platform"`
	RefOwner        string   `json:"ref_owner"`
	PrivateKey      string   `json:"private_key"`
	WriterAddresses []string `json:"writer_addresses"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// MapAccountTx strcut
type MapAccountTx struct {
	Account     Account
	Transaction Transaction
}

//KlaytnBody struct
type KlaytnBody struct {
	Hash        string `json:"hash"`
	BlockNumber string `json:"block_number"`
	Platform    string `json:"platform"`
	Account     string `json:"account"`
	PrivateKey  string `json:"private_key"`
}

//RequestKlaytn struct
type RequestKlaytn struct {
	Name string       `json:"name"`
	Body []KlaytnBody `json:"body"`
}

//ResponseKlatyn strcut
type ResponseKlatyn struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    []Body `json:"body"`
}

//Signatures struct
type Signatures struct {
	V string `json:"V"`
	R string `json:"R"`
	S string `json:"S"`
}

//Events struct
type Events struct {
}

//Body struct
type Body struct {
	BlockHash        string       `json:"blockHash"`
	BlockNumber      int          `json:"blockNumber"`
	ContractAddress  interface{}  `json:"contractAddress"`
	From             string       `json:"from"`
	Gas              string       `json:"gas"`
	GasPrice         string       `json:"gasPrice"`
	GasUsed          int          `json:"gasUsed"`
	Input            string       `json:"input"`
	LogsBloom        string       `json:"logsBloom"`
	Nonce            string       `json:"nonce"`
	SenderTxHash     string       `json:"senderTxHash"`
	Signatures       []Signatures `json:"signatures"`
	Status           bool         `json:"status"`
	To               string       `json:"to"`
	TransactionHash  string       `json:"transactionHash"`
	TransactionIndex int          `json:"transactionIndex"`
	Type             string       `json:"type"`
	TypeInt          int          `json:"typeInt"`
	Value            string       `json:"value"`
	Events           Events       `json:"events"`
}
