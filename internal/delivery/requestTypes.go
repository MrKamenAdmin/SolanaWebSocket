package delivery

type MarketChartData struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
	Volume24H float64 `json:"volume_24h"`
	MarketCap int     `json:"market_cap"`
}

type GeneralInfo struct {
	ActivatedStake    int64   `json:"activatedStake"`
	AvgBlockTime24H   float64 `json:"avgBlockTime_24h"`
	AvgBlockTime1H    float64 `json:"avgBlockTime_1h"`
	AvgBlockTime1Min  int     `json:"avgBlockTime_1min"`
	CirculatingSupply int64   `json:"circulatingSupply"`
	DailyPriceChange  float64 `json:"dailyPriceChange"`
	DailyVolume       float64 `json:"dailyVolume"`
	DelinquentStake   int64   `json:"delinquentStake"`
	EpochInfo         struct {
		AbsoluteEpochStartSlot int `json:"absoluteEpochStartSlot"`
		AbsoluteSlot           int `json:"absoluteSlot"`
		BlockHeight            int `json:"blockHeight"`
		Epoch                  int `json:"epoch"`
		SlotIndex              int `json:"slotIndex"`
		SlotsInEpoch           int `json:"slotsInEpoch"`
		EpochStartTime         int `json:"epochStartTime"`
	} `json:"epochInfo"`
	StakingYield        float64 `json:"stakingYield"`
	TokenPrice          float64 `json:"tokenPrice"`
	TotalDelegatedStake int64   `json:"totalDelegatedStake"`
	TotalSupply         int64   `json:"totalSupply"`
	AvgLastVote         int     `json:"avgLastVote"`
	Epoch               int     `json:"epoch"`
	SkipRate            struct {
		SkipRate              float64 `json:"skipRate"`
		StakeWeightedSkipRate float64 `json:"stakeWeightedSkipRate"`
	} `json:"skipRate"`
	StakeWeightedNodeVersions []struct {
		Index   int     `json:"index"`
		Version string  `json:"version"`
		Value   float64 `json:"value"`
	} `json:"stakeWeightedNodeVersions"`
	StakingYieldAdjusted  float64 `json:"stakingYieldAdjusted"`
	AvgTPS                float64 `json:"avgTPS"`
	TotalTransactionCount int64   `json:"totalTransactionCount"`
	NrValidators          int     `json:"nrValidators"`
	NrNonValidators       int     `json:"nrNonValidators"`
	Superminority         struct {
		Stake int64 `json:"stake"`
		Nr    int   `json:"nr"`
	} `json:"superminority"`
}

type Validator struct {
	Validator struct {
		ActivatedStake   int64  `json:"activatedStake"`
		Commission       int    `json:"commission"`
		EpochCredits     []int  `json:"epochCredits"`
		EpochVoteAccount bool   `json:"epochVoteAccount"`
		LastVote         int    `json:"lastVote"`
		NodePubkey       string `json:"nodePubkey"`
		RootSlot         int    `json:"rootSlot"`
		VotePubkey       string `json:"votePubkey"`
		BlockProduction  struct {
			LeaderSlots  int `json:"leaderSlots"`
			SkippedSlots int `json:"skippedSlots"`
		} `json:"blockProduction"`
		DelegatingStakeAccounts []struct {
			Pubkey   string `json:"pubkey"`
			Lamports int    `json:"lamports"`
			Data     struct {
				State int `json:"state"`
				Meta  struct {
					RentExemptReserve int `json:"rent_exempt_reserve"`
					Authorized        struct {
						Staker     string `json:"staker"`
						Withdrawer string `json:"withdrawer"`
					} `json:"authorized"`
				} `json:"meta"`
				Lockup struct {
					UnixTimestamp int    `json:"unix_timestamp"`
					Epoch         int    `json:"epoch"`
					Custodian     string `json:"custodian"`
				} `json:"lockup"`
				Stake struct {
					Delegation struct {
						VoterPubkey        string  `json:"voter_pubkey"`
						Stake              int     `json:"stake"`
						ActivationEpoch    int     `json:"activation_epoch"`
						WarmupCooldownRate float32 `json:"warmup_cooldown_rate"`
						ValidatorInfo      struct {
							Name       string `json:"name"`
							Image      string `json:"image"`
							Website    string `json:"website"`
							NodePubkey string `json:"nodePubkey"`
						} `json:"validatorInfo"`
					} `json:"delegation"`
					CreditsObserved int `json:"credits_observed"`
				} `json:"stake"`
			} `json:"data"`
		} `json:"delegatingStakeAccounts"`
		DelegatorCount int `json:"delegatorCount"`
		Location       struct {
			Range    []int     `json:"range"`
			Country  string    `json:"country"`
			Region   string    `json:"region"`
			Eu       string    `json:"eu"`
			Timezone string    `json:"timezone"`
			City     string    `json:"city"`
			Ll       []float32 `json:"ll"`
			Metro    int       `json:"metro"`
			Area     int       `json:"area"`
		} `json:"location"`
		Moniker    string `json:"moniker"`
		Website    string `json:"website"`
		PictureURL string `json:"pictureURL"`
		Version    string `json:"version"`
		Details    string `json:"details"`
		Asn        struct {
			Organization string `json:"organization"`
		} `json:"asn"`
	} `json:"validator"`
	Slots [][]struct {
		RelativeSlot   int  `json:"relativeSlot"`
		AbsoluteSlot   int  `json:"absoluteSlot"`
		ConfirmedBlock bool `json:"confirmedBlock"`
	} `json:"slots"`
	Historic []struct {
		Stake      int    `json:"stake"`
		Delegators int    `json:"delegators"`
		Timestamp  string `json:"timestamp"`
	} `json:"historic"`
	LatestBlocks []struct {
		Blocknumber int `json:"blocknumber"`
		Blocktime   struct {
			Absolute int `json:"absolute"`
			Relative int `json:"relative"`
		} `json:"blocktime"`
		Metrics struct {
			Txcount           int `json:"txcount"`
			Failedtxs         int `json:"failedtxs"`
			Totalfees         int `json:"totalfees"`
			Instructions      int `json:"instructions"`
			Sucessfultxs      int `json:"sucessfultxs"`
			Innerinstructions int `json:"innerinstructions"`
			Totalvaluemoved   int `json:"totalvaluemoved"`
		} `json:"metrics"`
		Proposer string `json:"proposer"`
	} `json:"latestBlocks"`
}

type LatestBlock struct {
	Blocknumber       uint64 `json:"blocknumber"`
	Blockhash         string `json:"blockhash"`
	Previousblockhash string `json:"previousblockhash"`
	Parentslot        int    `json:"parentslot"`
	Blocktime         struct {
	} `json:"blocktime"`
	Metrics struct {
		Txcount           int `json:"txcount"`
		Failedtxs         int `json:"failedtxs"`
		Totalfees         int `json:"totalfees"`
		Instructions      int `json:"instructions"`
		Sucessfultxs      int `json:"sucessfultxs"`
		Innerinstructions int `json:"innerinstructions"`
	} `json:"metrics"`
	Programstats []struct {
		Count        int `json:"count"`
		Instructions struct {
		} `json:"instructions"`
	} `json:"programstats"`
	Proposer     string `json:"proposer"`
	ProposerData struct {
		Name       string `json:"name"`
		Image      string `json:"image"`
		Website    string `json:"website"`
		NodePubkey string `json:"nodePubkey"`
	} `json:"proposerData"`
	Ondemand bool `json:"ondemand"`
	Rewards  []struct {
		Lamports uint64 `json:"lamports"`
	} `json:"rewards"`
}

type ValidatorsAll struct {
	ActivatedStake int       `json:"activatedStake"`
	Commission     int       `json:"commission"`
	VotePubkey     string    `json:"votePubkey"`
	DelegatorCount int       `json:"delegatorCount"`
	Ll             []float64 `json:"ll"`
	Moniker        string    `json:"moniker"`
	Version        string    `json:"version"`
	LastVote       int       `json:"lastVote"`
	PictureURL     string    `json:"pictureURL"`
}

type ResponseJson interface {
	*GeneralInfo | *[]MarketChartData | *Validator | *[]LatestBlock | *[]ValidatorsAll
}
