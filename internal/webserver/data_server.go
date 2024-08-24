package webserver

import (
	"GorillaWebSocket/internal/delivery"
	"GorillaWebSocket/internal/delivery/singleton"
	"GorillaWebSocket/pkg/psql"
	history_repo "GorillaWebSocket/pkg/psql/repos/history.repo"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"time"
)

var token = "c756a91e-cf39-4ffc-ac28-3c286f6dbcad"
var votePubkey = "he1iusunGwqrNtafDtLdhsUQDFvo13z9sUa36PauBtk"

func StartDataServer() {

	pool := psql.NewPool(context.Background(),
		"postgresql://user:password@localhost:5432/gorilla?sslmode=disable",
	)
	repo := history_repo.New(pool)

	client := &http.Client{}

	generalInfoRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/general-info", nil,
	)
	marketChartDataRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/market-chart-data", nil,
	)
	validatorRequest, _ := http.NewRequest(
		"GET", "https://api.solanabeach.io/v1/validator/"+votePubkey, nil,
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

		requests := map[*http.Request]any{
			generalInfoRequest:     &generalInfo,
			marketChartDataRequest: &marketChart,
			validatorRequest:       &validator,
			validatorsAllRequest:   &validatorsAll,
			latestBlocksRequest:    &latestBlocks,
		}

		for {
			for k, v := range requests {
				err := getJsonFromResponse(client, k, v)
				if err != nil {
					log.Println(err)
					continue
				}
			}

			temp := delivery.Response{}

			temp.Solana.Price = marketChart[len(marketChart)-1].Price
			temp.Solana.Tps = generalInfo.AvgTPS
			temp.Solana.Delta = generalInfo.DailyPriceChange

			temp.Validator.Apy = generalInfo.StakingYield
			temp.Validator.Staked = float64(validator.Validator.ActivatedStake / 1_000_000_000)

			idx := slices.IndexFunc(validatorsAll, func(v delivery.ValidatorsAll) bool { return v.VotePubkey == votePubkey })
			temp.Validator.Place = uint64(idx + 1)

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

			history, err := historyUpload(cache, repo, uint64(temp.Validator.Staked))
			if err != nil {
				log.Println(err)
			}
			temp.History = history

			cache.Set(temp)
			time.Sleep(5 * time.Second)
		}
	}()
}

func historyUpload(cache *singleton.Cache, repo *history_repo.Repo, stake uint64) ([]delivery.History, error) {
	ctx := context.Background()
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	localHistory := cache.Get().History

	if len(localHistory) == 0 {
		history, err := repo.GetHistory(ctx)
		if err != nil {
			fmt.Println("Не смог получить историю из БД")
			return nil, err
		}

		return history, nil
	} else if localHistory[0].CaptureDate.Sub(now) != 0 {
		err := repo.AddStake(ctx, stake, now)
		if err != nil {
			return nil, err
		}

		history, err := repo.GetHistory(ctx)
		if err != nil {
			fmt.Println("Не смог получить историю из БД")
			return nil, err
		}

		return history, nil

	} else {
		return cache.Get().History, nil
	}
}

func addHeaders(requests ...*http.Request) {
	for r := range requests {
		requests[r].Header.Add("Accept", "application/json")
		requests[r].Header.Add("Authorization", token)
	}
}

func getJsonFromResponse(c *http.Client, r *http.Request, responseJson any) error {

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
