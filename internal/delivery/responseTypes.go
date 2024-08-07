package delivery

type Solana struct {
	Price float64 `json:"price"`
	Delta float64 `json:"delta"`
	Tps   float64 `json:"tps"`
}

type ValidatorResponse struct {
	Place  uint64  `json:"place"`
	Apy    float64 `json:"apy"`
	Staked float32 `json:"staked"`
}

type Block struct {
	Number   uint64  `json:"number"`
	Producer string  `json:"producer"`
	Reward   float32 `json:"reward"`
}
