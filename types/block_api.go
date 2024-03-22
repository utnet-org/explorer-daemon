package types

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

type LastBlockRes struct {
	Height           int64  `json:"height"`
	Hash             string `json:"hash"`              //交易Hash
	ChunkHash        string `json:"chunk_hash"`        //chunkHash
	Timestamp        int64  `json:"timestamp"`         //时间
	TimestampNanoSec string `json:"timestamp_nanosec"` //时间
	Author           string `json:"author"`            //矿工
	GasPrice         int    `json:"gas_price"`         //Gas价格
	GasLimit         int    `json:"gas_limit"`         //Gas限制
	//
	//TransactionsAgg struct {
	//	Count int `json:"count"` //交易
	//} `json:"transactions_agg"`
}

type LastBlockResList struct {
	LastBlockList []LastBlockRes `json:"last_block_list"`
}

type LastBlockRes2 struct {
	Height    string `json:"height"`
	Hash      string `json:"hash"`      //交易Hash
	Time      string `json:"timestamp"` //区块时间
	AccountId string `json:"author"`    //矿工

	ChunksAgg struct {
		GasUsed int `json:"gas_used"` //使用的Gas
	} `json:"chunks_agg"`
	TransactionsAgg struct {
		Count int `json:"count"` //交易
	} `json:"transactions_agg"`
}
