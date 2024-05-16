package types

type CoinPriceRes struct {
	CoinPrice map[string]CoinPrice
}

type CoinPrice struct {
	USD          float64 `json:"usd"`
	USD24hChange float64 `json:"usd_24h_change"`
}

type CoinPriceResWeb struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}
