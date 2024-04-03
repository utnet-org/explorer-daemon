package types

type Contract struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

type Result struct {
	BlockHash   string   `json:"block_hash"`
	BlockHeight int64    `json:"block_height"`
	Proof       []string `json:"proof"`
	Values      []Value  `json:"values"`
}

type Value struct {
	Key   string   `json:"key"`
	Proof []string `json:"proof"`
	Value string   `json:"value"`
}
