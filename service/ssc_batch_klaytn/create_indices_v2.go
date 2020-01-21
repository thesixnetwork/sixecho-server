package main

import "github.com/olivere/elastic"

func createSSCAccountIndexV2(client *elastic.Client) {
	elasticIndex := "ssc_accounts_v2"
	exists, err := client.IndexExists(elasticIndex).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	mapping := `
	{
		"settings": {
			"number_of_shards": 5,
			"number_of_replicas": 1
		},
		"mappings": {		
			"_doc": {
							"properties": {
									"platform":{
										"type":"keyword"
									},
									"ref_owner":{
										"type":"keyword"
									},
									"private_key":{
										"type":"text"
									},
									"writer_addresses":{
										"type":"text",
										"fields":{
											"keyword":{
												"type":"keyword"
											}
										}
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
	}
}
