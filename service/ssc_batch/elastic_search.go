package main

import (
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/olivere/elastic"
)

//SSCData struct
type SSCData struct {
	AssetID  int64
	Author   eos.Name
	Category eos.Name
	Owner    eos.Name
	IData    string
	MData    string
	AA       string
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
	Author     string `json:"author"`
	PreviewURL string `json:"preview_url"`
}

//MDataText struct
type MDataText struct {
	Author          string `json:"author"`
	Language        string `json:"language"`
	PaperBack       string `json:"paperback"`
	PublishDate     int64  `json:"publish_date"`
	Title           string `json:"title"`
	CountryOfOrigin string `json:"country_of_origin"`
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
									"asset_id":{
											"type":"keyword"
									},
									"block_num":{
											"type":"long"
									},
									"klaytn_tx_id":{
											"type":"keyword"
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
									"author":{
										"type":"keyword"
									},
									"preview_url":{
										"type":"text"
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
									"author":{
										"type":"keyword"
									},
									"title":{
										"type":"keyword"
									},
									"country_of_origin":{
										"type":"keyword"
									},
									"language":{
										"type":"keyword"
									},
									"paperback":{
										"type":"long"
									},
									"publish_date":{
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

func insertTxToES(assetID string, sscTxID string, klaytnTxID string, blockNum uint32) {
	elasticAlias := "ssc_transactions"
	type DigitalContent struct {
		AssetID    string `json:"asset_id"`
		BlockNum   int64  `json:"blokc_num"`
		KlaytnTxID string `json:"klaytn_tx_id"`
	}
	digitalContent := DigitalContent{
		AssetID:    assetID,
		KlaytnTxID: klaytnTxID,
		BlockNum:   int64(blockNum),
	}
	digitalContentJSON, _ := json.Marshal(digitalContent)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(sscTxID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
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
				fmt.Println(action.Account)
				fmt.Println(action.Name)

				klaytnTxID := submitToKlaytn(tx.Transaction.ID.String(), blockResp.BlockNum)
				if action.Account == "assets" && action.Name == "create" {
					sscData := action.Data.(*SSCData)
					json.Unmarshal([]byte(sscData.IData), &iData)
					switch typeAsset := iData.Type; typeAsset {
					case "IMAGE":
						mData := MDataImage{}
						json.Unmarshal([]byte(sscData.MData), &mData)
						insertImageToES(fmt.Sprintf("%d", sscData.AssetID), iData, mData)
					case "TEXT":
						mData := MDataText{}
						json.Unmarshal([]byte(sscData.MData), &mData)
						insertTextToES(fmt.Sprintf("%d", sscData.AssetID), iData, mData)
					}
					insertTxToES(fmt.Sprintf("%d", sscData.AssetID), tx.Transaction.ID.String(), klaytnTxID, blockResp.BlockNum)
				}
			}
		}
	}
}

func insertImageToES(assetID string, iData IData, mData MDataImage) {
	elasticAlias := "ssc_images"
	type DataImage struct {
		IData
		MDataImage
	}
	dataImage := DataImage{}
	dataImage.IData = iData
	dataImage.MDataImage = mData
	digitalContentJSON, _ := json.Marshal(dataImage)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(assetID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
}

func insertTextToES(assetID string, iData IData, mData MDataText) {
	elasticAlias := "ssc_texts"
	type DataImage struct {
		IData
		MDataText
	}
	dataImage := DataImage{}
	dataImage.IData = iData
	dataImage.MDataText = mData
	digitalContentJSON, _ := json.Marshal(dataImage)
	_, err := client.Index().Index(elasticAlias).Type("_doc").Id(assetID).BodyString(string(digitalContentJSON)).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
}
