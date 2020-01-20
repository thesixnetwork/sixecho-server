package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/eoscanada/eos-go"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/olivere/elastic"
	v4 "github.com/olivere/elastic/aws/v4"
)

var (
	elasticURL    = "https://search-es-six-zunsizmfamv7eawswgdvwmyd6u.ap-southeast-1.es.amazonaws.com"
	ctx           = context.Background()
	region        = os.Getenv("AWS_REGION")
	cred          = credentials.NewEnvCredentials()
	signingClient = v4.NewV4SigningClient(cred, region)
	// sess            = session.Must(session.NewSession())
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
	}))
	client, _ = elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		// elastic.SetHttpClient(signingClient),
		//elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	)
	concurrent     = 10
	eosURL         = "http://ec2-3-0-89-218.ap-southeast-1.compute.amazonaws.com:8888"
	apiEOS         = eos.New(eosURL)
	currentAssetID = ""
	sizeQuery      = 100
	startAssetID   = "15770573341233142"
	stopAssetID    = "15773381420735946"
	accountName    = eos.AccountName("ookbee")
	insertDigest   = make(chan InsertDigest)
)

func updateCurrent() {
	d1 := []byte(currentAssetID)
	err := ioutil.WriteFile("./current.txt", d1, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func getText(assetID string) int64 {
	response, err := client.Get().Index("ssc_texts").Type("_doc").Id(assetID).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	data, err := response.Source.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}
	var text Text
	json.Unmarshal(data, &text)
	return text.CreatedTime
}

func start() {
	startTime := getText(startAssetID)
	endTime := getText(stopAssetID)

	query := elastic.NewRangeQuery("created_time").Gte(startTime).Lte(endTime)
	startQuery := client.Search().Index("ssc_texts").Query(query).SortBy(elastic.NewFieldSort("created_time").Asc(), elastic.NewFieldSort("_id").Asc()).Size(sizeQuery)

	response, err := startQuery.Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Hits.TotalHits)
	checkDub := make(map[int64]bool)
	i := 0
	for {
		var sortValues []interface{}
		if len(response.Hits.Hits) == 0 {
			break
		}

		for _, hit := range response.Hits.Hits {
			i = i + 1
			data, err := hit.Source.MarshalJSON()
			if err != nil {
				panic(err.Error())
			}
			var idata IData
			var submitedBy SubmittedBy
			json.Unmarshal(data, &idata)
			json.Unmarshal(data, &submitedBy)

			idataJSON, err := json.Marshal(idata)
			if err != nil {
				panic(err)
			}
			assetID, err := strconv.ParseInt(hit.Id, 10, 64)
			reworkData := InsertDigest{
				SubmittedBy: eos.Name(submitedBy.SubmittedBy),
				AssetID:     assetID,
				IData:       string(idataJSON),
			}
			insertDigest <- reworkData
			fmt.Printf("%d-%d\n", reworkData.AssetID, i)
			if checkDub[reworkData.AssetID] == false {
				checkDub[reworkData.AssetID] = true
			} else {
				panic("Dub Loop")
			}
			sortValues = hit.Sort
		}
		response, err = client.Search().
			Index("ssc_texts").
			Query(query).
			Size(sizeQuery).
			SearchAfter(sortValues...).
			SortBy(elastic.NewFieldSort("created_time").Asc(), elastic.NewFieldSort("_id").Asc()).
			Do(context.Background())
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Finished")
	fmt.Println(i)
}

func loadProcess() {
	for i := 0; i < concurrent; i++ {
		go func() {
			for {
				select {
				case data := <-insertDigest:
					count := 0
					for i := 0; i <= 3; i++ {
						err := submitAction(data)
						if err == nil {
							break
						}
						count = count + 1
					}
					if count == 3 {
						panic(fmt.Sprintf("Error at : %d", data.AssetID))
					}
				}
			}
		}()
	}

	go func() {
		for range time.Tick(time.Second * 3) {
			updateCurrent()
		}
	}()

	start()
}

func getPrivateKey() string {
	return "5KEcvztdsZu9SP8XkeFkXadyJtUwkMYh4j8F2PFi3j1C9kBgLL5"
}

func submitAction(data InsertDigest) error {
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(getPrivateKey())
	if err != nil {
		panic(err.Error())
	}
	apiEOS.SetSigner(keyBag)
	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(apiEOS); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	action := &eos.Action{
		Account: eos.AccountName("assets"),
		Name:    eos.ActionName("insertdigest"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      accountName,
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(data),
	}

	tx := eos.NewTransaction([]*eos.Action{action}, txOpts)
	_, packedTx, err := apiEOS.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	// if err != nil {
	// panic(fmt.Errorf("sign transaction: %s", err))
	// }
	// content, err := json.MarshalIndent(signedTx, "", "  ")
	// if err != nil {
	// panic(fmt.Errorf("json marshalling transaction: %s", err))
	// }

	// fmt.Println(string(content))
	response, err := apiEOS.PushTransactionRaw(packedTx)
	if err != nil {
		return err
	}
	result, err := response.MarshalJSON()
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(result))
	var responseEOS ResponseEOS
	json.Unmarshal(result, &responseEOS)
	currentAssetID = fmt.Sprintf("%d", data.AssetID)
	fmt.Println(fmt.Sprintf("%s %d", responseEOS.TransactionID, responseEOS.Processed.BlockNum))
	// fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
	return nil
}

func main() {
	eos.RegisterAction(eos.AccountName("assets"), eos.ActionName("insertdigest"), InsertDigest{})
	loadProcess()
}
