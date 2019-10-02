package main

import (
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/eoscanada/eos-go"
)

var (
	api             *eos.API
	es              *elasticsearch.Client
	currentBlockNum uint32
)

func tailBlock(block chan *eos.BlockResp, blockNum chan uint32) {
	for {
		select {
		case num := <-blockNum:
			blockResp, _ := api.GetBlockByNum(num)
			// fmt.Println(blockResp.BlockNum)
			block <- blockResp
		}
	}
}

func updateBlockNumToES() {
	fmt.Println("Update Elasticsearch BlockNum")
}

func invokeLambda(data *eos.BlockResp) {
	fmt.Printf("%#v\n", data.ID.String())
	// fmt.Println("Invokde lambda")
}

func loadAllBackgroundProcess(block chan *eos.BlockResp, blockNum chan uint32) {
	go func() {
		for range time.Tick(time.Second * 10) {
			updateBlockNumToES()
		}
	}()
	go func() {
		for range time.Tick(time.Second * 1) {
			infoResp, _ := api.GetInfo()
			lastBlockNum := infoResp.HeadBlockNum
			for i := currentBlockNum; i < lastBlockNum; i++ {
				blockNum <- i
			}
			currentBlockNum = lastBlockNum
		}
	}()
	go func() {
		for {
			select {
			case data := <-block:
				invokeLambda(data)
			}
		}
	}()
}

func getCurrentBlockNum() {
	infoResp, _ := api.GetInfo()
	currentBlockNum = infoResp.HeadBlockNum
}

func createIndexElastic() {
	fmt.Println("Load elasticsearch...")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://search-es-six-zunsizmfamv7eawswgdvwmyd6u.ap-southeast-1.es.amazonaws.com",
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)
	createSSCBlockNumIndex(es)
	createSSCDigitalContentIndex(es)
}

func main() {
	createIndexElastic()
	block := make(chan *eos.BlockResp)
	blockNum := make(chan uint32)
	api = eos.New("http://ec2-3-0-89-218.ap-southeast-1.compute.amazonaws.com:8888")
	getCurrentBlockNum()
	loadAllBackgroundProcess(block, blockNum)
	tailBlock(block, blockNum)
}
