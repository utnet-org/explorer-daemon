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
	Header       ChunkDetailsHeader `json:"header"`
	Receipts     []string           `json:"receipts"`
	Transactions []string           `json:"transactions"`
	ChunkHash    string             `json:"chunk_hash"`
	BlockHash    string             `json:"block_hash"`
	Height       int64              `json:"height"`
}
