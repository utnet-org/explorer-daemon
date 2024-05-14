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
	CommonRes CommonRes
	Body      GasPriceBody `json:"result"`
}

type GasPriceBody struct {
	GasPrice string `json:"gasPrice"`
}

type DailyGas struct {
	Date string  `json:"date"`
	Gas  float64 `json:"gas"`
}
