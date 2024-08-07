package webserver

import (
	"GorillaWebSocket/internal/delivery"
	"GorillaWebSocket/internal/delivery/singleton"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var token = "c756a91e-cf39-4ffc-ac28-3c286f6dbcad"

func StartDataServer() {
	client := &http.Client{}

	generalInfoRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/general-info", nil,
	)
	marketChartDataRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/market-chart-data", nil,
	)
	validatorRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/validator/he1iusunGwqrNtafDtLdhsUQDFvo13z9sUa36PauBtk", nil,
	)
	validatorsAllRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/validators/all", nil,
	)
	latestBlocksRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/latest-blocks", nil,
	)

	// добавляем заголовки
	addHeaders(generalInfoRequest, marketChartDataRequest, validatorRequest, validatorsAllRequest, latestBlocksRequest)

	go func() {
		cache := singleton.GetInstance()

		generalInfo := delivery.GeneralInfo{}
		marketChart := []delivery.MarketChartData{}
		validator := delivery.Validator{}
		validatorsAll := []delivery.ValidatorsAll{}
		latestBlocks := []delivery.LatestBlock{}

		for {
			err := getJsonFromResponse(client, generalInfoRequest, &generalInfo)
			if err != nil {
				log.Println(err)
				continue
			}

			err = getJsonFromResponse(client, marketChartDataRequest, &marketChart)
			if err != nil {
				log.Println(err)
				continue
			}

			err = getJsonFromResponse(client, validatorRequest, &validator)
			if err != nil {
				log.Println(err)
				continue
			}

			err = getJsonFromResponse(client, validatorsAllRequest, &validatorsAll)
			if err != nil {
				log.Println(err)
				continue
			}

			err = getJsonFromResponse(client, latestBlocksRequest, &latestBlocks)
			if err != nil {
				log.Println(err)
				continue
			}

			temp := delivery.Response{}

			temp.Solana.Price = marketChart[len(marketChart)-1].Price
			temp.Solana.Tps = generalInfo.AvgTPS
			temp.Solana.Delta = generalInfo.DailyPriceChange

			temp.Validator.Apy = generalInfo.StakingYield
			temp.Validator.Staked = float32(validator.Validator.ActivatedStake / 1_000_000_000)
			index := 0
			for i := range validatorsAll {
				if validatorsAll[i].VotePubkey == "he1iusunGwqrNtafDtLdhsUQDFvo13z9sUa36PauBtk" {
					index = i
				}
			}
			temp.Validator.Place = uint64(index + 1)

			blocks := make([]delivery.Block, 5)
			for i := 0; i < 5; i++ {
				blocks[i].Number = latestBlocks[i].Blocknumber
				blocks[i].Producer = latestBlocks[i].Proposer
				if len(latestBlocks[i].Rewards) == 0 {
					blocks[i].Reward = 0
				} else {
					blocks[i].Reward = float32(latestBlocks[i].Rewards[0].Lamports) / 1_000_000_000
				}
			}

			temp.BlockData = blocks

			cache.Set(temp)

			//fmt.Println("validatorsAll : ")
			//fmt.Println(validatorsAll)

			time.Sleep(5 * time.Second)
		}
	}()
}

func addHeaders(requests ...*http.Request) {
	for r := range requests {
		requests[r].Header.Add("Accept", "application/json")
		requests[r].Header.Add("Authorization", token)
	}
}

func getJsonFromResponse[T delivery.ResponseJson](c *http.Client, r *http.Request, responseJson T) error {

	resp, err := c.Do(r)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
		fmt.Println(err)
	}

	err = json.Unmarshal(msg, &responseJson)

	if err != nil {
		return err
	}

	return nil
}
