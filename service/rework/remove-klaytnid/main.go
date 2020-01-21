package main

import (
	"Workspace/eos-go"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/olivere/elastic"
	v4 "github.com/olivere/elastic/aws/v4"
)

var (
	api           *eos.API
	elasticURL    = "https://search-sixadmin-qfgn7dggaqvur3ckwrmhwmym3u.ap-southeast-1.es.amazonaws.com"
	ctx           = context.Background()
	region        = os.Getenv("AWS_REGION")
	cred          = credentials.NewEnvCredentials()
	signingClient = v4.NewV4SigningClient(cred, region)
	sess          = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
	}))
	client, _ = elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHttpClient(signingClient),
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
)

func queryTransaction() *elastic.SearchResult {
	query := `{"query":{"bool":{"must_not":{"term":{"klaytn_tx_id":""}}}},"sort":{"created_at":"asc"},"size":1800}`
	response, err := client.Search("ssc_transactions_v2").Source(query).Sort("created_at", true).Size(30).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return response
}

func main() {
	response := queryTransaction()
	fmt.Println(response.Hits.TotalHits)
	for _, hit := range response.Hits.Hits {
		fmt.Println(hit.Id)

		update := map[string]string{
			"klaytn_tx_id": "",
		}
		_, err := client.Update().Index("ssc_transactions_v2").Type("_doc").Id(hit.Id).Doc(update).Do(context.Background())
		if err != nil {
			panic(err.Error())
		}
	}
}
