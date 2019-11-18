package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/olivere/elastic"
)

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

//SSCDataUpdate struct
type SSCDataUpdate struct {
	Owner   eos.Name `json:"owner"`
	AssetID int64    `json:"asset_id"`
}

//EchoOwner stuct
type EchoOwner struct {
	Owner    string `json:"owner"`
	RefOwner string `json:"ref_owner"`
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
	SubmittedBy string     `json:"submitted_by,omitempty"`
	Platform    string     `json:"platform,omitempty"`
	FromUser    *EchoOwner `json:"from_user,omitempty"`
	ToUser      *EchoOwner `json:"to_user,omitempty"`
}

//CommonInfo struct
type CommonInfo struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	ParentID string   `json:"parent_id"`
	Tags     []string `json:"tags"`
}

//DetailInfo strcut
type DetailInfo struct {
}

func createSSCBlockNumIndex(client *elastic.Client) {
	elasticIndex := "ssc_blocknum_v1"
	elasticAlias := "ssc_blocknum"
	exists, err := client.IndexExists(elasticIndex).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	mapping := `
	{
		"settings": {
			"number_of_shards": 5,
			"number_of_replicas": 0
		},
		"mappings": {		
			"_doc": {
							"properties": {
									"block_num":{
											"type":"long"
									}
							}
					}
		}
	}`
	if !exists {
		createIndex, err := client.CreateIndex(elasticIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
		_, err = client.Alias().Add(elasticIndex, elasticAlias).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func createSSCDigitalContentIndex(client *elastic.Client) {
	elasticIndex := "ssc_transactions_v1"
	elasticAlias := TransactionAlias
	exists, err := client.IndexExists(elasticIndex).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	mapping := `
	{
		"settings": {
			"number_of_shards": 5,
			"number_of_replicas": 0
		},
		"mappings": {		
			"_doc": {
							"properties": {
									"submitted_by":{
											"type":"keyword"
									},	
									"from_user":{
											"type":"nested"
									},	
									"platform":{
											"type":"keyword"
									},	
									"to_user":{
											"type":"nested"
									},	
									"asset_id":{
											"type":"keyword"
									},
									"asset_type":{
											"type":"keyword"
									},
									"block_num":{
											"type":"long"
									},
									"klaytn_tx_id":{
											"type":"keyword"
									},
									"transaction_action":{
											"type":"keyword"
									},
									"transaction_status":{
											"type":"keyword"
									},
									"authorization":{
											"type":"nested"
									},
									"created_time":{
											"type":"integer"
									},
									"updated_time":{
											"type":"integer"
									},
									"created_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									},
									"updated_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									}
							}
					}
		}
	}`
	if !exists {
		createIndex, err := client.CreateIndex(elasticIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
		_, err = client.Alias().Add(elasticIndex, elasticAlias).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func createSSCImageIndex(client *elastic.Client) {
	elasticIndex := "ssc_images_v1"
	elasticAlias := ImageAlias
	exists, err := client.IndexExists(elasticIndex).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	mapping := `
	{
		"settings": {
			"number_of_shards": 5,
			"number_of_replicas": 0
		},
		"mappings": {		
			"_doc": {
							"properties": {
									"digest":{
											"type":"keyword"
									},
									"sha256":{
											"type":"keyword"
									},
									"size_file":{
											"type":"keyword"
									},
									"submitted_by":{
										"type":"keyword"
									},
									"platform":{
										"type":"keyword"
									},
									"title":{
										"type":"keyword"
									},
									"image_url":{
										"type":"text"
									},
									"creator":{
										"type":"keyword"
									},
									"parent_id":{
										"type":"keyword"
									},
									"owner":{
										"type":"keyword"
									},
									"ref_owner":{
										"type":"keyword"
									},
									"ref_creator":{
										"type":"keyword"
									},
									"tags":{
										"type":"keyword"
									},
									"status":{
										"type":"keyword"
									},
									"mdata":{
										"type":"text"
									},
									"created_time":{
											"type":"integer"
									},
									"updated_time":{
											"type":"integer"
									},
									"created_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									},
									"updated_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									}
							}
					}
		}
	}`
	if !exists {
		createIndex, err := client.CreateIndex(elasticIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
		_, err = client.Alias().Add(elasticIndex, elasticAlias).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func createSSCTextIndex(client *elastic.Client) {
	elasticIndex := "ssc_texts_v1"
	elasticAlias := TextAlias
	exists, err := client.IndexExists(elasticIndex).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	mapping := `
	{
		"settings": {
			"number_of_shards": 5,
			"number_of_replicas": 0
		},
		"mappings": {		
			"_doc": {
							"properties": {
									"digest":{
											"type":"keyword"
									},
									"sha256":{
											"type":"keyword"
									},
									"size_file":{
											"type":"keyword"
									},
									"submitted_by":{
										"type":"keyword"
									},
									"platform":{
										"type":"keyword"
									},
									"title":{
										"type":"keyword"
									},
									"image_url":{
										"type":"text"
									},
									"creator":{
										"type":"keyword"
									},
									"parent_id":{
										"type":"keyword"
									},
									"owner":{
										"type":"keyword"
									},
									"ref_owner":{
										"type":"keyword"
									},
									"ref_creator":{
										"type":"keyword"
									},
									"tags":{
										"type":"keyword"
									},
									"status":{
										"type":"keyword"
									},
									"mdata":{
										"type":"text"
									},
									"created_time":{
											"type":"integer"
									},
									"updated_time":{
											"type":"integer"
									},
									"created_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									},
									"updated_at":{
											"type":"date",
											"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"
									}
							}
					}
		}
	}`
	if !exists {
		createIndex, err := client.CreateIndex(elasticIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
		_, err = client.Alias().Add(elasticIndex, elasticAlias).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
func getCurrentBlockNumFromES(client *elastic.Client, blockNumber uint32) uint32 {
	elasticAlias := "ssc_blocknum"
	doc, err := client.Get().Index(elasticAlias).Type("_doc").Id("1").Pretty(true).Do(ctx)
	if err != nil {
		doc := map[string]interface{}{
			"block_num": blockNumber,
		}
		docJSON, _ := json.Marshal(doc)
		_, err := client.Index().Index("ssc_blocknum").Type("_doc").Id("1").BodyString(string(docJSON)).Do(ctx)
		if err != nil {
			panic(err.Error())
		}
		return blockNumber
	}
	data, _ := doc.Source.MarshalJSON()
	type BlockNum struct {
		BlockNum int64 `json:"block_num"`
	}
	var num BlockNum
	json.Unmarshal(data, &num)
	blockNumber = uint32(num.BlockNum)
	return blockNumber
}

func insertTxToES(blockResp *eos.BlockResp, tx eos.TransactionReceipt, action *eos.Action, assetID string, iData *IData, klaytnTxID string, fromto FromToTransaction) {
	elasticAlias := "ssc_transactions"
	type Authorization struct {
		Actor      string `json:"actor"`
		Permission string `json:"permission"`
	}
	type DigitalContent struct {
		AssetID           string          `json:"asset_id"`
		AssetType         string          `json:"asset_type"`
		BlockNum          int64           `json:"block_num"`
		KlaytnTxID        string          `json:"klaytn_tx_id"`
		TransactionAction string          `json:"transaction_action"`
		TransactionStatus string          `json:"transaction_status"`
		Authorization     []Authorization `json:"authorization"`
		FromToTransaction
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}

	var authorizationStucts []Authorization
	var authorizations []eos.PermissionLevel
	authorizations = action.Authorization
	for _, ele := range authorizations {
		tmp := Authorization{
			Actor:      string(ele.Actor),
			Permission: string(ele.Permission),
		}
		authorizationStucts = append(authorizationStucts, tmp)
	}
	var assetType string
	assetType = iData.Type
	// assetID = fmt.Sprintf("%d", sscData.AssetID)
	if assetType == "" {
		assetType = getAssetType(assetID)
	}

	timeStamp := blockResp.Timestamp.Time
	digitalContent := DigitalContent{
		AssetID:           assetID,
		AssetType:         assetType,
		KlaytnTxID:        klaytnTxID,
		BlockNum:          int64(blockResp.BlockNum),
		TransactionAction: string(action.Name),
		TransactionStatus: tx.Status.String(),
		Authorization:     authorizationStucts,
		FromToTransaction: fromto,
		CreatedTime:       timeStamp.Unix(),
		UpdatedTime:       timeStamp.Unix(),
		CreatedAt:         timeStamp.Format("2006-01-02 15:04:05"),
		UpdatedAt:         timeStamp.Format("2006-01-02 15:04:05"),
	}
	digitalContentJSON, _ := json.Marshal(digitalContent)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(tx.Transaction.ID.String()).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		fmt.Println("Error insert transaction to ES")
		panic(err.Error())
	}
}

func insertAssetToES(blockResp *eos.BlockResp) {
	iData := IData{}
	for _, tx := range blockResp.Transactions {
		fmt.Println(blockResp.BlockNum)
		if tx.Transaction.Packed == nil {
			continue
		}
		data, _ := tx.Transaction.Packed.Unpack()
		if len(data.Transaction.Actions) != 0 {
			for _, action := range data.Transaction.Actions {
				klaytnTxID := submitToKlaytn(tx.Transaction.ID.String(), blockResp.BlockNum)
				if action.Account == "assets" && action.Name == "create" {
					sscData := action.Data.(*SSCDataCreate)
					json.Unmarshal([]byte(sscData.IData), &iData)
					switch typeAsset := iData.Type; typeAsset {
					case "IMAGE":
						insertImageToES(blockResp, sscData, &iData)
					case "TEXT":
						insertTextToES(blockResp, sscData, &iData)
					}
					var refInfo *RefInfo
					json.Unmarshal([]byte(sscData.RefInfo), &refInfo)
					fromto := FromToTransaction{
						SubmittedBy: string(sscData.SubmittedBy),
						Platform:    string(sscData.SubmittedBy),
						FromUser:    &refInfo.EchoOwner,
						ToUser:      &refInfo.EchoOwner,
					}

					insertTxToES(blockResp, tx, action, fmt.Sprintf("%d", sscData.AssetID), &iData, klaytnTxID, fromto)

				} else if action.Account == "assets" && action.Name == "transfer" {
					sscDataTransfer := action.Data.(*SSCDataTransfer)
					updateAssetToEs(sscDataTransfer)
					var fromUser, toUser *EchoOwner
					json.Unmarshal([]byte(sscDataTransfer.FromJSONStr), &fromUser)
					json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &toUser)
					fromto := FromToTransaction{
						SubmittedBy: string(sscDataTransfer.From),
						Platform:    string(sscDataTransfer.To),
						FromUser:    fromUser,
						ToUser:      toUser,
					}
					insertTxToES(blockResp, tx, action, fmt.Sprintf("%d", sscDataTransfer.AssetID), &iData, klaytnTxID, fromto)
				} else if action.Account == "assets" && action.Name == "update" {
					// sscDataUpdate := action.Data.(*SSCDataUpdate)
					// insertTxToES(action.Authorization, "", "", nil, nil, fmt.Sprintf("%d", sscDataUpdate.AssetID), "", tx.Transaction.ID.String(), string(action.Name), tx.Status.String(), klaytnTxID, blockResp.BlockNum, timeStamp)
				} else if action.Account == "assets" && action.Name == "revoke" {
					// insertTxToES(action.Authorization, "", "", nil, nil, fmt.Sprintf("%d", sscDataUpdate.AssetID), "", tx.Transaction.ID.String(), string(action.Name), tx.Status.String(), klaytnTxID, blockResp.BlockNum, timeStamp)
				}
			}
		}
	}
}

func updateAssetToEs(sscDataTransfer *SSCDataTransfer) {
	query := elastic.NewTermQuery("_id", sscDataTransfer.AssetID)
	var userTo EchoOwner
	json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &userTo)
	strScript := fmt.Sprintf("ctx._source.platform = '%s'; ctx._source.owner = '%s'; ctx._source.ref_owner = '%s'", sscDataTransfer.To, userTo.Owner, userTo.RefOwner)
	inScript := elastic.NewScriptInline(strScript).Lang("painless")
	_, err := client.UpdateByQuery("ssc_texts", "ssc_images").Query(query).Script(inScript).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
}

func insertImageToES(blockResp *eos.BlockResp, sscData *SSCDataCreate, iData *IData) {
	elasticAlias := ImageAlias
	type DetailInfoImage struct {
		Photographer *string `json:"photographer,omitempty"`
		Width        int64   `json:"width,omitempty"`
		Hight        int64   `json:"hight,omitempty"`
		Dpi          int64   `json:"dpi,omitempty"`
	}

	type DataImage struct {
		*IData
		*DetailInfoImage
		*CommonInfo
		*RefInfo
		Platform    string `json:"platform"`
		SubmittedBy string `json:"submitted_by"`
		MData       string `json:"mdata,omitempty"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	var detailInfo *DetailInfoImage
	var refInfo *RefInfo
	var commonInfo *CommonInfo
	json.Unmarshal([]byte(sscData.DetailInfo), &detailInfo)
	json.Unmarshal([]byte(sscData.RefInfo), &refInfo)
	json.Unmarshal([]byte(sscData.CommonInfo), &commonInfo)
	assetID := fmt.Sprintf("%d", sscData.AssetID)
	dataImage := DataImage{}
	dataImage.IData = iData
	timeStamp := blockResp.Timestamp.Time
	dataImage.DetailInfoImage = detailInfo
	dataImage.RefInfo = refInfo
	dataImage.MData = sscData.MData
	dataImage.Platform = string(sscData.SubmittedBy)
	dataImage.SubmittedBy = string(sscData.SubmittedBy)
	dataImage.CreatedTime = timeStamp.Unix()
	dataImage.UpdatedTime = timeStamp.Unix()
	dataImage.CreatedAt = timeStamp.Format("2006-01-02 15:04:05")
	dataImage.UpdatedAt = timeStamp.Format("2006-01-02 15:04:05")
	digitalContentJSON, _ := json.Marshal(dataImage)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(assetID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		fmt.Println("Error Insert Image To ES : ")
		panic(err.Error())
	}
}

func insertTextToES(blockResp *eos.BlockResp, sscData *SSCDataCreate, iData *IData) {
	elasticAlias := TextAlias
	type DetailInfoText struct {
		ISBN          string `json:"isbn,omitempty"`
		Author        string `json:"author,omitempty"`
		Publisher     string `json:"publisher,omitempty"`
		PublishedDate string `json:"published_date,omitempty"`
		Language      string `json:"language,omitempty"`
		NumberOfpages string `json:"number_of_pages,omitempty"`
	}
	type DataText struct {
		*IData
		*DetailInfoText
		*CommonInfo
		*RefInfo
		Platform    string `json:"platform"`
		SubmittedBy string `json:"submitted_by"`
		MData       string `json:"mdata,omitempty"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	var detailInfo *DetailInfoText
	var refInfo *RefInfo
	var commonInfo *CommonInfo
	json.Unmarshal([]byte(sscData.DetailInfo), &detailInfo)
	json.Unmarshal([]byte(sscData.RefInfo), &refInfo)
	json.Unmarshal([]byte(sscData.CommonInfo), &commonInfo)
	dataText := DataText{}
	assetID := fmt.Sprintf("%d", sscData.AssetID)
	timeStamp := blockResp.Timestamp.Time
	dataText.IData = iData
	dataText.DetailInfoText = detailInfo
	dataText.CommonInfo = commonInfo
	dataText.RefInfo = refInfo
	fmt.Printf("%#v\n", sscData.MData)
	dataText.MData = sscData.MData
	dataText.Platform = string(sscData.SubmittedBy)
	dataText.SubmittedBy = string(sscData.SubmittedBy)
	dataText.CreatedTime = timeStamp.Unix()
	dataText.UpdatedTime = timeStamp.Unix()
	dataText.CreatedAt = timeStamp.Format("2006-01-02 15:04:05")
	dataText.UpdatedAt = timeStamp.Format("2006-01-02 15:04:05")

	digitalContentJSON, _ := json.Marshal(dataText)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(assetID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		fmt.Println("Error Insert Text To ES : ")
		panic(err.Error())
	}
}
