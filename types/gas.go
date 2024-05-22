package types

type GasBlockHeightReq struct {
	BlockHeights []int
}

type GasBlockHashReq struct {
	BlockHash []string
}

type GasBlockNullReq struct {
	BlockNull []interface{}
}

// Gas Price Response

type GasPriceRes struct {
	CommonRes
	Result GasPriceResult `json:"result"`
}

type GasPriceResult struct {
	GasPrice string `json:"gas_price"`
}

type DailyGas struct {
	Date string  `json:"date"`
	Gas  float64 `json:"gas"`
}
