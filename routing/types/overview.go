package types

type OverviewInfoRes struct {
	Height           string `json:"height"`
	LatestBlock      string `json:"latestBlock"`
	TotalPower       string `json:"totalPower"`
	ActiveMiner      string `json:"activeMiner"`
	BlockReward      string `json:"blockReward"`
	DayAveReward     string `json:"dayAveReward"`
	DayProduction    string `json:"dayProduction"`
	DayMessage       string `json:"dayMessage"`
	TotalAccount     string `json:"totalAccount"`
	AveBlockInterval string `json:"aveBlockInterval"`
}
