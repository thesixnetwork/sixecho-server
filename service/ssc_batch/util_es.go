package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

func getAssetType(assetID string) string {
	query := elastic.NewTermQuery("_id", assetID)
	searchSource := elastic.NewSearchSource().Size(2).Query(query)
	response, err := client.Search(TextAlias, ImageAlias).Type("_doc").SearchSource(searchSource).Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
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

func setDInfo(sscSetDInfo *SSCSetDInfo) {
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
		panic(err.Error())
	}
}

func updateCInfo(sscUpdateCinfo *SSCUpdateCInfo) {
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
		panic(err.Error())
	}
}

func setMdata(sscSetMdata *SSCSetMdata) {
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
		panic(err.Error())
	}
}
