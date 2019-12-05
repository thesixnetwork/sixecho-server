package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	l "github.com/aws/aws-sdk-go/service/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/olivere/elastic"
	v4 "github.com/olivere/elastic/aws/v4"
)

var (
	// elasticURL      = os.Getenv("ELASTIC_URL")
	elasticURL     = "https://search-sixadmin-qfgn7dggaqvur3ckwrmhwmym3u.ap-southeast-1.es.amazonaws.com"
	lambdaFunction = "SixEchoFunction-ContractClient-PNSG31FO5WH3"
	ctx            = context.Background()
	region         = os.Getenv("AWS_REGION")
	cred           = credentials.NewEnvCredentials()
	signingClient  = v4.NewV4SigningClient(cred, region)
	// sess            = session.Must(session.NewSession())
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
		// Profile: "default",
	}))
	client, _ = elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetHttpClient(signingClient),
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags))
	)
)

func queryTransactoin() []*Transaction {
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("klaytn_tx_id", ""), elastic.NewTermQuery("transaction_action", "create"))
	response, err := client.Search(TransactionAlias).Query(query).Sort("created_at", true).Size(30).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	var transactions []*Transaction
	for _, hit := range response.Hits.Hits {
		transaction := Transaction{}
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			panic(err.Error())
		}
		err = json.Unmarshal(data, &transaction)
		transaction.ID = hit.Id
		transactions = append(transactions, &transaction)
	}
	return transactions
}

func submitToKlaytn(txs []*Transaction) ResponseKlatyn {
	var kReqs []KlaytnBody
	for _, t := range txs {
		tmp := KlaytnBody{
			Hash:        t.ID,
			BlockNumber: fmt.Sprintf("%d", t.BlockNumber),
		}
		kReqs = append(kReqs, tmp)
	}
	payload := RequestKlaytn{
		Name: "new-assets",
		Body: kReqs,
	}
	payloadJSON, _ := json.Marshal(payload)
	lambdaClient := l.New(sess)
	input := &l.InvokeInput{
		FunctionName: aws.String(lambdaFunction),
		Payload:      payloadJSON,
	}
	result, err := lambdaClient.Invoke(input)
	if err != nil {
		panic(err.Error())
	}
	var response ResponseKlatyn
	err = json.Unmarshal(result.Payload, &response)
	if err != nil {
		panic(err.Error())
	}
	if len(response.Body) == 0 {
		fmt.Println(string(result.Payload))
	}
	return response
}

func matching(txs []*Transaction, klaynTxs []Body) {
	for index, tx := range txs {
		tx.KlaytnTxID = klaynTxs[index].TransactionHash
	}
}

func updateElastBatch(txs []*Transaction) {
	bulk := client.Bulk()
	for _, tx := range txs {
		req := elastic.NewBulkUpdateRequest()
		req.Index(TransactionAlias)
		req.Doc(map[string]interface{}{"klaytn_tx_id": tx.KlaytnTxID})
		req.Id(tx.ID)
		req.Type("_doc")
		bulk = bulk.Add(req)
	}
	bulkResp, err := bulk.Refresh("wait_for").Do(ctx)
	if err != nil {
		// panic(err.Error())
		fmt.Println(err.Error())
	}
	fmt.Println(bulkResp.Took)
}

func backGround() {
	fmt.Println("Runing...")
	for range time.Tick(time.Second * 3) {
		allProcess()
	}
}

func allProcess() {
	txs := queryTransactoin()
	if len(txs) > 0 {
		responseKlatyn := submitToKlaytn(txs)
		if len(responseKlatyn.Body) > 0 {
			fmt.Println("Submit Klaytn Suucess")
			matching(txs, responseKlatyn.Body)
			updateElastBatch(txs)
			//doc, _ := json.Marshal(txs)
			//fmt.Println(string(doc))
		} else {
			fmt.Println("Submit Klaytn is null")
			time.Sleep(time.Second * 10)
		}
	}
}

func main() {
	backGround()
}
