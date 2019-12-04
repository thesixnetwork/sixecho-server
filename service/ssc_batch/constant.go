package main

import "github.com/eoscanada/eos-go"

const (
	TextAlias        string = "ssc_texts"
	ImageAlias       string = "ssc_images"
	TransactionAlias string = "ssc_transactions"
	ErrorAlias       string = "ssc_errors"
	TXCreateError    string = "TX_CREATE"
	TransferError    string = "TRANSFER"
	ImgCreateError   string = "IMG_CREATE"
	TextCreateError  string = "TEXT_CREATE"
	RevokeError      string = "REVOKED"
	SETMdataError    string = "MDATA_UPDATE"
	UCinfoError      string = "CINFO_UPDATE"
	SETDinfoError    string = "DINFO_UPDATE"
)

//DetailInfoImage struct
type DetailInfoImage struct {
	Photographer *string `json:"photographer,omitempty"`
	Width        int64   `json:"width,omitempty"`
	Height       int64   `json:"height,omitempty"`
	Dpi          int64   `json:"dpi,omitempty"`
}

//DetailInfoText struct
type DetailInfoText struct {
	ISBN          string `json:"isbn,omitempty"`
	Author        string `json:"author,omitempty"`
	Publisher     string `json:"publisher,omitempty"`
	PublishedDate int64  `json:"published_date,omitempty"`
	Language      string `json:"language,omitempty"`
	NumberOfpages int64  `json:"number_of_pages,omitempty"`
}

//SSCDataCreate struct
type SSCDataCreate struct {
	SubmittedBy eos.Name
	AssetID     int64
	IData       string
	MData       string
	CommonInfo  string
	DetailInfo  string
	RefInfo     string
}

//SSCDataTransfer struct
type SSCDataTransfer struct {
	From        eos.Name `json:"from"`
	To          eos.Name `json:"to"`
	FromJSONStr string   `json:"from_json_str"`
	ToJSONStr   string   `json:"to_json_str"`
	AssetID     int64    `json:"asset_id"`
	Memo        string   `json:"memo"`
}

//SSCSetDInfo struct
type SSCSetDInfo struct {
	Platform   eos.Name `json:"platform"`
	AssetID    int64    `json:"asset_id"`
	DetailInfo string   `json:"detail_info"`
}

//SSCSetMdata struct
type SSCSetMdata struct {
	Platform   eos.Name `json:"platform"`
	AssetID    int64    `json:"asset_id"`
	DetailInfo string   `json:"detail_info"`
}

//SSCUpdateCInfo struct
type SSCUpdateCInfo struct {
	Platform   eos.Name `json:"platform"`
	AssetID    int64    `json:"asset_id"`
	DetailInfo string   `json:"detail_info"`
}

//SSCRevoke struct
type SSCRevoke struct {
	Platform eos.Name `json:"platform"`
	AssetID  int64    `json:"asset_id"`
	Memo     string   `json:"memo"`
}

//EchoOwner stuct
type EchoOwner struct {
	Owner    string `json:"owner,omitempty"`
	RefOwner string `json:"ref_owner,omitempty"`
}

//IData struct
type IData struct {
	Digest   string `json:"digest"`
	Sha256   string `json:"sha256"`
	Sizefile int64  `json:"size_file"`
	Type     string `json:"type"`
}

//RefInfo struct
type RefInfo struct {
	EchoOwner
	Creator    string `json:"creator"`
	RefCreator string `json:"ref_creator"`
}

//FromToTransaction struct
type FromToTransaction struct {
	FromPlatform string     `json:"from_platform,omitempty"`
	ToPlatform   string     `json:"to_platform,omitempty"`
	SubmittedBy  string     `json:"submitted_by,omitempty"`
	Platform     string     `json:"platform,omitempty"`
	FromUser     *EchoOwner `json:"from_user,omitempty"`
	ToUser       *EchoOwner `json:"to_user,omitempty"`
	Memo         string     `json:"memo,omitempty"`
}

//CommonInfo struct
type CommonInfo struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	ParentID string   `json:"parent_id"`
	Tags     []string `json:"tags"`
}

//ErrorMessage struct
type ErrorMessage struct {
	BlockNum     int64  `json:"block_num"`
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}
