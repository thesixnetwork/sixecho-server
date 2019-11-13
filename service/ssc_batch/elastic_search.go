package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eoscanada/eos-go"
	"github.com/olivere/elastic"
)

//SSCDataCreate struct
type SSCDataCreate struct {
	AssetID int64
	Creator eos.Name
	Owner   eos.Name
	IData   string
	MData   string
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

//EchoOwner stuct
type EchoOwner struct {
	EchoOwner    string `json:"echo_owner"`
	EchoRefOwner string `json:"echo_ref_owner"`
}

//IData struct
type IData struct {
	Digest   string `json:"digest"`
	Sha256   string `json:"sha256"`
	Sizefile int64  `json:"size_file"`
	Type     string `json:"type"`
}

//MDataImage struct
type MDataImage struct {
	EchoTitle      string   `json:"echo_title"`
	EchoImageURL   string   `json:"echo_image_url"`
	EchoParentID   string   `json:"echo_parent_id"`
	EchoOwner      string   `json:"echo_owner"`
	EchoRefOwner   string   `json:"echo_ref_owner"`
	EchoCreator    string   `json:"echo_creator"`
	EchoRefCreator string   `json:"echo_ref_creator"`
	EchoTags       []string `json:"echo_tags"`
}

//MDataText struct
type MDataText struct {
	EchoTitle      string   `json:"echo_title"`
	EchoImageURL   string   `json:"echo_image_url"`
	EchoParentID   string   `json:"echo_parent_id"`
	EchoOwner      string   `json:"echo_owner"`
	EchoRefOwner   string   `json:"echo_ref_owner"`
	EchoCreator    string   `json:"echo_creator"`
	EchoRefCreator string   `json:"echo_ref_creator"`
	EchoTags       []string `json:"echo_tags"`
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
	elasticAlias := "ssc_transactions"
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
									"from":{
											"type":"keyword"
									},	
									"from_user":{
											"type":"nested"
									},	
									"to":{
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
	elasticAlias := "ssc_images"
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
									"creator":{
										"type":"keyword"
									},
									"owner":{
										"type":"keyword"
									},
									"echo_title":{
										"type":"keyword"
									},
									"echo_image_url":{
										"type":"text"
									},
									"echo_creator":{
										"type":"keyword"
									},
									"echo_parent_id":{
										"type":"keyword"
									},
									"echo_owner":{
										"type":"keyword"
									},
									"echo_ref_owner":{
										"type":"keyword"
									},
									"echo_ref_creator":{
										"type":"keyword"
									},
									"echo_tags":{
										"type":"keyword"
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
	elasticAlias := "ssc_texts"
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
									"creator":{
										"type":"keyword"
									},
									"owner":{
										"type":"keyword"
									},
									"echo_title":{
										"type":"keyword"
									},
									"echo_image_url":{
										"type":"text"
									},
									"echo_creator":{
										"type":"keyword"
									},
									"echo_parent_id":{
										"type":"keyword"
									},
									"echo_owner":{
										"type":"keyword"
									},
									"echo_ref_owner":{
										"type":"keyword"
									},
									"echo_ref_creator":{
										"type":"keyword"
									},
									"echo_tags":{
										"type":"keyword"
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

func insertTxToES(authorizations []eos.PermissionLevel, from, to string, fromUser, toUser EchoOwner, assetID string, assetType string, sscTxID string, transactionAction string, transactionStatus string, klaytnTxID string, blockNum uint32, timeStamp time.Time) {
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
		From              string          `json:"from"`
		To                string          `json:"to"`
		FromUser          EchoOwner       `json:"from_user"`
		ToUser            EchoOwner       `json:"to_user"`
		CreatedTime       int64           `json:"created_time"`
		UpdatedTime       int64           `json:"updated_time"`
		CreatedAt         string          `json:"created_at"`
		UpdatedAt         string          `json:"updated_at"`
	}

	var authorizationStucts []Authorization
	for _, ele := range authorizations {
		tmp := Authorization{
			Actor:      string(ele.Actor),
			Permission: string(ele.Permission),
		}
		authorizationStucts = append(authorizationStucts, tmp)
	}

	digitalContent := DigitalContent{
		AssetID:           assetID,
		AssetType:         assetType,
		KlaytnTxID:        klaytnTxID,
		BlockNum:          int64(blockNum),
		TransactionAction: transactionAction,
		TransactionStatus: transactionStatus,
		Authorization:     authorizationStucts,
		CreatedTime:       timeStamp.Unix(),
		UpdatedTime:       timeStamp.Unix(),
		From:              from,
		To:                to,
		FromUser:          fromUser,
		ToUser:            toUser,
		CreatedAt:         timeStamp.Format("2006-01-02 15:04:05"),
		UpdatedAt:         timeStamp.Format("2006-01-02 15:04:05"),
	}
	digitalContentJSON, _ := json.Marshal(digitalContent)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(sscTxID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		fmt.Println("Error insert transaction to ES")
		panic(err.Error())
	}
}

func insertAssetToES(blockResp *eos.BlockResp) {
	iData := IData{}
	timeStamp := blockResp.Timestamp.Time
	for _, tx := range blockResp.Transactions {
		fmt.Println(blockResp.BlockNum)
		if tx.Transaction.Packed == nil {
			continue
		}
		data, _ := tx.Transaction.Packed.Unpack()
		fmt.Printf("%#v\n", data)
		if len(data.Transaction.Actions) != 0 {
			for _, action := range data.Transaction.Actions {
				// fmt.Println(action.Account)
				// fmt.Println(action.Name)
				klaytnTxID := submitToKlaytn(tx.Transaction.ID.String(), blockResp.BlockNum)
				fromUser := EchoOwner{}
				toUser := EchoOwner{}
				if action.Account == "assets" && action.Name == "create" {
					sscData := action.Data.(*SSCDataCreate)
					json.Unmarshal([]byte(sscData.IData), &iData)
					switch typeAsset := iData.Type; typeAsset {
					case "IMAGE":
						mData := MDataImage{}
						json.Unmarshal([]byte(sscData.MData), &mData)
						// fromUser.EchoRefOwner = mData.EchoRefOwner
						toUser.EchoRefOwner = mData.EchoRefOwner
						toUser.EchoOwner = mData.EchoOwner
						insertImageToES(string(sscData.Creator), string(sscData.Owner), fmt.Sprintf("%d", sscData.AssetID), iData, mData, timeStamp)
					case "TEXT":
						mData := MDataText{}
						json.Unmarshal([]byte(sscData.MData), &mData)
						toUser.EchoRefOwner = mData.EchoRefOwner
						toUser.EchoOwner = mData.EchoOwner
						insertTextToES(string(sscData.Creator), string(sscData.Owner), fmt.Sprintf("%d", sscData.AssetID), iData, mData, timeStamp)
					}
					insertTxToES(action.Authorization, string(sscData.Creator), string(sscData.Creator), fromUser, toUser, fmt.Sprintf("%d", sscData.AssetID), iData.Type, tx.Transaction.ID.String(), string(action.Name), tx.Status.String(), klaytnTxID, blockResp.BlockNum, timeStamp)
				} else if action.Account == "assets" && action.Name == "transfer" {
					sscDataTransfer := action.Data.(*SSCDataTransfer)
					updateAssetToEs(sscDataTransfer)
					json.Unmarshal([]byte(sscDataTransfer.FromJSONStr), &fromUser)
					json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &toUser)
					insertTxToES(action.Authorization, string(sscDataTransfer.From), string(sscDataTransfer.To), fromUser, toUser, fmt.Sprintf("%d", sscDataTransfer.AssetID), "", tx.Transaction.ID.String(), string(action.Name), tx.Status.String(), klaytnTxID, blockResp.BlockNum, timeStamp)
				}
			}
		}
	}
}

func updateAssetToEs(sscDataTransfer *SSCDataTransfer) {
	query := elastic.NewTermQuery("_id", sscDataTransfer.AssetID)
	var userTo EchoOwner
	json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &userTo)
	strScript := fmt.Sprintf("ctx._source.owner = '%s'; ctx._source.echo_owner = '%s'; ctx._source.echo_ref_owner = '%s'", sscDataTransfer.To, userTo.EchoOwner, userTo.EchoRefOwner)
	inScript := elastic.NewScriptInline(strScript).Lang("painless")
	_, err := client.UpdateByQuery("ssc_texts", "ssc_images").Query(query).Script(inScript).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
}

func insertImageToES(creator, owner, assetID string, iData IData, mData MDataImage, timeStamp time.Time) {
	elasticAlias := "ssc_images"
	type DataImage struct {
		IData
		MDataImage
		Creator     string `json:"creator"`
		Owner       string `json:"owner"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	dataImage := DataImage{}
	dataImage.IData = iData
	dataImage.MDataImage = mData
	dataImage.Creator = creator
	dataImage.Owner = owner
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

func insertTextToES(creator, owner, assetID string, iData IData, mData MDataText, timeStamp time.Time) {
	elasticAlias := "ssc_texts"
	type DataText struct {
		IData
		MDataText
		Creator     string `json:"creator"`
		Owner       string `json:"owner"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	dataText := DataText{}
	dataText.IData = iData
	dataText.MDataText = mData
	dataText.Creator = creator
	dataText.Owner = owner
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
