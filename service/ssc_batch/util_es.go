package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eoscanada/eos-go"
	"github.com/olivere/elastic"
)

func getAssetType(assetID string) string {
	query := elastic.NewTermQuery("_id", assetID)
	searchSource := elastic.NewSearchSource().Size(2).Query(query)
	response, err := client.Search(TextAlias, ImageAlias).Type("_doc").SearchSource(searchSource).Pretty(true).Do(context.Background())
	if err != nil {
		fmt.Println("GET ASSET Error")
	}
	if response.Hits.TotalHits > 1 {
		panic("Source more than one")
	}
	var idata IData
	for _, hit := range response.Hits.Hits {
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			panic(err.Error())
		}
		err = json.Unmarshal(data, &idata)
		if err != nil {
			panic(err.Error())
		}
		break
	}
	return idata.Type
}

func setDInfo(blockResp *eos.BlockResp, sscSetDInfo *SSCSetDInfo) {
	assetID := fmt.Sprintf("%d", sscSetDInfo.AssetID)
	assetType := getAssetType(assetID)
	var err error
	var elasticAlias string
	timeStamp := time.Now()
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
		var detailInfoImage DetailInfoImage
		json.Unmarshal([]byte(sscSetDInfo.DetailInfo), &detailInfoImage)
		type Update struct {
			DetailInfoImage
			UpdatedTime int64  `json:"updated_time"`
			UpdatedAt   string `json:"updated_at"`
		}
		update := Update{
			UpdatedTime:     timeStamp.Unix(),
			UpdatedAt:       timeStamp.Format("2006-01-02 15:04:05"),
			DetailInfoImage: detailInfoImage,
		}
		_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(update).Do(context.Background())
	case "TEXT":
		elasticAlias = TextAlias
		var detailInfoText DetailInfoText
		json.Unmarshal([]byte(sscSetDInfo.DetailInfo), &detailInfoText)
		type Update struct {
			DetailInfoText
			UpdatedTime int64  `json:"updated_time"`
			UpdatedAt   string `json:"updated_at"`
		}
		update := Update{
			UpdatedTime:    timeStamp.Unix(),
			UpdatedAt:      timeStamp.Format("2006-01-02 15:04:05"),
			DetailInfoText: detailInfoText,
		}
		_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(update).Do(context.Background())
	}
	if err != nil {
		insertError(blockResp.BlockNum, SETDinfoError, err.Error())
		// panic(err.Error())
	}
}

func updateTransferES(blockResp *eos.BlockResp, sscDataTransfer *SSCDataTransfer) {
	assetID := fmt.Sprintf("%d", sscDataTransfer.AssetID)
	assetType := getAssetType(assetID)
	timeStamp := time.Now()
	var err error
	var elasticAlias string
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
	case "TEXT":
		elasticAlias = TextAlias
	}
	var userTo EchoOwner
	json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &userTo)
	type Update struct {
		EchoOwner
		Platform    string `json:"platform"`
		UpdatedTime int64  `json:"updated_time"`
		UpdatedAt   string `json:"updated_at"`
	}
	update := Update{
		UpdatedTime: timeStamp.Unix(),
		UpdatedAt:   timeStamp.Format("2006-01-02 15:04:05"),
		EchoOwner:   userTo,
		Platform:    string(sscDataTransfer.To),
	}
	_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(update).Do(context.Background())
	if err != nil {
		insertError(blockResp.BlockNum, TransferError, err.Error())
		//panic(err.Error())
	}
	//query := elastic.NewTermQuery("_id", sscDataTransfer.AssetID)
	//var userTo EchoOwner
	//json.Unmarshal([]byte(sscDataTransfer.ToJSONStr), &userTo)
	//now := time.Now()
	//strScript := fmt.Sprintf("ctx._source.platform = '%s'; ctx._source.owner = '%s'; ctx._source.ref_owner = '%s'; ctx._source.updated_time = %d; ctx._source.updated_at = '%s'", sscDataTransfer.To, userTo.Owner, userTo.RefOwner, now.Unix(), now.Format("2006-01-02 15:04:05"))
	//inScript := elastic.NewScriptInline(strScript).Lang("painless")
	//_, err := client.UpdateByQuery("ssc_texts", "ssc_images").Query(query).Script(inScript).Do(context.Background())
	//if err != nil {
	//insertError(blockResp.BlockNum, TransferError, err.Error())
	//panic(err.Error())
	//}
}

func updateCInfo(blockResp *eos.BlockResp, sscUpdateCinfo *SSCUpdateCInfo) {
	assetID := fmt.Sprintf("%d", sscUpdateCinfo.AssetID)
	assetType := getAssetType(assetID)
	timeStamp := time.Now()
	var err error
	var elasticAlias string
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
	case "TEXT":
		elasticAlias = TextAlias
	}
	var commonInfo CommonInfo
	type Update struct {
		CommonInfo
		UpdatedTime int64  `json:"updated_time"`
		UpdatedAt   string `json:"updated_at"`
	}
	json.Unmarshal([]byte(sscUpdateCinfo.DetailInfo), &commonInfo)
	update := Update{
		UpdatedTime: timeStamp.Unix(),
		UpdatedAt:   timeStamp.Format("2006-01-02 15:04:05"),
		CommonInfo:  commonInfo,
	}
	_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(update).Do(context.Background())
	if err != nil {
		insertError(blockResp.BlockNum, UCinfoError, err.Error())
		// panic(err.Error())
	}
}

func setMdata(blockResp *eos.BlockResp, sscSetMdata *SSCSetMdata) {
	assetID := fmt.Sprintf("%d", sscSetMdata.AssetID)
	assetType := getAssetType(assetID)
	timeStamp := time.Now()
	var err error
	var elasticAlias string
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
	case "TEXT":
		elasticAlias = TextAlias
	}
	_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(map[string]interface{}{
		"mdata":        sscSetMdata.DetailInfo,
		"updated_time": timeStamp.Unix(),
		"updated_at":   timeStamp.Format("2006-01-02 15:04:05"),
	}).Do(context.Background())
	if err != nil {
		insertError(blockResp.BlockNum, SETMdataError, err.Error())
		// panic(err.Error())
	}
}

func revoke(blockResp *eos.BlockResp, sscRevoke *SSCRevoke) {
	query := elastic.NewTermQuery("_id", sscRevoke.AssetID)
	now := time.Now()
	strScript := fmt.Sprintf("ctx._source.revoked = %t ; ctx._source.updated_time = %d; ctx._source.updated_at = '%s'", true, now.Unix(), now.Format("2006-01-02 15:04:05"))
	inScript := elastic.NewScriptInline(strScript).Lang("painless")
	_, err := client.UpdateByQuery(TextAlias, ImageAlias).Query(query).Script(inScript).Do(context.Background())
	if err != nil {
		insertError(blockResp.BlockNum, RevokeError, err.Error())
		// panic(err.Error())
	}
}

func insertError(blockNum uint32, errorType string, errorMsg string) {
	errorMessage := ErrorMessage{
		BlockNum:     int64(blockNum),
		ErrorType:    errorType,
		ErrorMessage: errorMsg,
	}
	errorMessageJSON, _ := json.Marshal(errorMessage)
	_, err := client.Index().Index(ErrorAlias).Type("_doc").BodyString(string(errorMessageJSON)).Do(ctx)
	if err != nil {
		_, err = client.Index().Index(ErrorAlias).Type("_doc").BodyString(string(errorMessageJSON)).Do(ctx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
