package main

import (
	"context"
	"encoding/json"
	"fmt"

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
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
		var detailInfoImage DetailInfoImage
		json.Unmarshal([]byte(sscSetDInfo.DetailInfo), &detailInfoImage)
		_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(detailInfoImage).Do(context.Background())
	case "TEXT":
		elasticAlias = TextAlias
		var detailInfoText DetailInfoText
		json.Unmarshal([]byte(sscSetDInfo.DetailInfo), &detailInfoText)
		_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(detailInfoText).Do(context.Background())
	}
	if err != nil {
		panic(err.Error())
	}
}

func updateCInfo(sscUpdateCinfo *SSCUpdateCInfo) {
	assetID := fmt.Sprintf("%d", sscUpdateCinfo.AssetID)
	assetType := getAssetType(assetID)
	var err error
	var elasticAlias string
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
	case "TEXT":
		elasticAlias = TextAlias
	}
	var commonInfo CommonInfo
	json.Unmarshal([]byte(sscUpdateCinfo.DetailInfo), &commonInfo)
	_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(commonInfo).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
}
func setMdata(sscSetMdata *SSCSetMdata) {
	assetID := fmt.Sprintf("%d", sscSetMdata.AssetID)
	assetType := getAssetType(assetID)
	var err error
	var elasticAlias string
	switch assetType {
	case "IMAGE":
		elasticAlias = ImageAlias
	case "TEXT":
		elasticAlias = TextAlias
	}
	_, err = client.Update().Index(elasticAlias).Type("_doc").Id(assetID).Doc(map[string]interface{}{
		"mdata": sscSetMdata.DetailInfo,
	}).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
}
