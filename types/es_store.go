package types

type BlockDetailsStoreBody struct {
	Author           string             `json:"author"`
	Chunks           []Chunk            `json:"chunks"`
	Header           BlockDetailsHeader `json:"header"`
	Hash             string             `json:"hash"`
	ChunkHash        string             `json:"chunk_hash"`
	Height           int64              `json:"height"`
	TimestampMilli   int64              `json:"timestamp_milli"`
	Timestamp        int64              `json:"timestamp"`
	TimestampNanoSec string             `json:"timestamp_nanosec"`
	PrevHash         string             `json:"prev_hash"`        // 父哈希
	PrevHeight       int64              `json:"prev_height"`      // 父高度
	GasLimit         int64              `json:"gas_limit"`        // Gas 限制
	GasPrice         int64              `json:"gas_price"`        // Gas 价格
	GasUsed          int64              `json:"gas_used"`         // Gas 消耗
	ValidatorReward  string             `json:"validator_reward"` // 奖励
	TotalSupply      string             `json:"total_supply"`     // 总奖励
}

type ChunkDetailsStoreResult struct {
	Author       string             `json:"author"`
	Timestamp    int64              `json:"timestamp"`
	Header       ChunkDetailsHeader `json:"header"`
	Receipts     []ReceiptElement   `json:"receipts"`
	Transactions []Transaction      `json:"transactions"`
	ChunkHash    string             `json:"chunk_hash"`
	BlockHash    string             `json:"block_hash"`
	Height       int64              `json:"height"`
}

type TxnStoreResult struct {
	Height    int64 `json:"height"`
	Timestamp int64 `json:"timestamp"`
	TxnStatusReceiptsResult
	//TxnStatusResult
	//FinalExeStatus     string             `json:"final_execution_status"`
	//ReceiptsOutcome    []ReceiptsOutcome  `json:"receipts_outcome"`
	//Status             interface{}        `json:"status"`
	//Transaction        `json:"transaction"`
	//TransactionOutcome `json:"transaction_outcome"`
}

type GasStoreResult struct {
	Height   int64  `json:"height"`
	Hash     string `json:"hash"`
	GasPrice string `json:"gas_price"`
}
