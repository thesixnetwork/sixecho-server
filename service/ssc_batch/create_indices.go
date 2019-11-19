package main

import "github.com/olivere/elastic"

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
									"memo":{
											"type":"text"
									},
									"detail_values":{
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
									"revoked":{
										"type":"boolean"
									},
									"mdata":{
										"type":"text"
									},
									"photographer":{
										"type":"keyword"
									},
									"width":{
										"type":"integer"
									},
									"hight":{
										"type":"integer"
									},
									"dpi":{
										"type":"integer"
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
									"revoked":{
										"type":"boolean"
									},
									"mdata":{
										"type":"text"
									},
									"isbn":{
										"type":"keyword"
									},
									"author":{
										"type":"keyword"
									},
									"publisher":{
										"type":"keyword"
									},
									"language":{
										"type":"keyword"
									},
									"published_date":{
										"type":"integer"
									},
									"number_of_pages":{
										"type":"integer"
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
