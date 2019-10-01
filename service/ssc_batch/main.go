package main

import (
	"fmt"
	"time"

	"github.com/eoscanada/eos-go"
)

var (
	api             *eos.API
	currentBlockNum uint32
)

func tailBlock(block chan *eos.BlockResp, blockNum chan uint32) {
	for {
		select {
		case num := <-blockNum:
			blockResp, _ := api.GetBlockByNum(num)
			fmt.Println(blockResp.BlockNum)
			block <- blockResp
		}
	}
}

func updateBlockNumToES() {
	fmt.Println("Update Elasticsearch BlockNum")
}

func invokeLambda(data *eos.BlockResp) {
	fmt.Println("Invokde lambda")
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

func main() {
	block := make(chan *eos.BlockResp)
	blockNum := make(chan uint32)
	api = eos.New("http://ec2-3-0-89-218.ap-southeast-1.compute.amazonaws.com:8888")
	getCurrentBlockNum()
	loadAllBackgroundProcess(block, blockNum)
	tailBlock(block, blockNum)
}
