package main

import "github.com/eoscanada/eos-go"

//InsertDigest
type InsertDigest struct {
	SubmittedBy eos.Name
	AssetID     int64
	IData       string
}

//SubmittedBy struct
type SubmittedBy struct {
	SubmittedBy string `json:"submitted_by"`
}

//Asset struct
type Asset struct {
	ID     string `json:"id"`
	Digest string `json:"digest"`
}

//IData struct
type IData struct {
	Digest   string `json:"digest"`
	Sha256   string `json:"sha256"`
	Sizefile int64  `json:"size_file"`
	Type     string `json:"type"`
}

//Text stuct
type Text struct {
	CreatedTime int64 `json:"created_time"`
}

//ResponseEOS struct
type ResponseEOS struct {
	TransactionID string `json:"transaction_id"`
	Processed     struct {
		BlockNum int64 `json:"block_num"`
	} `json:"processed"`
}
