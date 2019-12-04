package main

const (
	TransactionAlias string = "ssc_transactions"
)

//Transaction struct
type Transaction struct {
	ID          string
	KlaytnTxID  string `json:"klaytn_tx_id"`
	BlockNumber int64  `json:"block_num"`
}

//KlaytnBody struct
type KlaytnBody struct {
	Hash        string `json:"hash"`
	BlockNumber string `json:"block_number"`
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
