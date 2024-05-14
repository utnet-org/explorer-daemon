package types

type OverviewInfoRes struct {
	Height           int64   `json:"height"`
	LatestBlock      string  `json:"latest_block"`
	TotalPower       int64   `json:"total_power"`
	ActiveMiner      int64   `json:"active_miner"`
	BlockReward      int64   `json:"block_reward"`
	DayAveReward     float64 `json:"day_ave_reward"`
	DayProduction    int64   `json:"dayProduction"`
	DayMessages      int64   `json:"day_messages"`
	TotalAccount     int64   `json:"total_account"`
	AveBlockInterval string  `json:"aveBlockInterval"`
}
