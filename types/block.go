package types

type BlockReq struct {
	BlockID int `json:"block_id"`
}

type BlockReq2 struct {
	Finality string `json:"finality"`
}

// 区块集合
type Blocks struct {
	Blocks []Block `json:"blocks"`
}

// 区块内容
type Block struct {
	Height    string `json:"block_height"`      //区块高度
	Hash      string `json:"block_hash"`        //交易Hash
	Time      string `json:"block_timestamp"`   //区块时间
	AccountId string `json:"author_account_id"` //矿工
	ChunksAgg struct {
		GasUsed int `json:"gas_used"` //使用的Gas
	} `json:"chunks_agg"`
	TransactionsAgg struct {
		Count int `json:"count"` //交易
	} `json:"transactions_agg"`
}
