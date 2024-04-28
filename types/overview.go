package types

type OverviewInfoRes struct {
	Height           int64  `json:"height"`
	LatestBlock      string `json:"latest_block"`
	TotalPower       string `json:"totalPower"`
	ActiveMiner      int64  `json:"active_miner"`
	BlockReward      string `json:"blockReward"`
	DayAveReward     string `json:"dayAveReward"`
	DayProduction    string `json:"dayProduction"`
	DayMessage       string `json:"dayMessage"`
	TotalAccount     int64  `json:"total_account"`
	AveBlockInterval string `json:"aveBlockInterval"`
}
