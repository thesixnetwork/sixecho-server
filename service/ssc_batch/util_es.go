package main

import (
	"context"
	"encoding/json"

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
