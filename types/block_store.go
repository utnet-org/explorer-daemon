package types

type BlockDetailsStoreBody struct {
	Author           string             `json:"author"`
	Chunks           []Chunk            `json:"chunks"`
	Header           BlockDetailsHeader `json:"header"`
	Hash             string             `json:"hash"`
	ChunkHash        string             `json:"chunk_hash"`
	Height           int64              `json:"height"`
	Timestamp        int64              `json:"timestamp"`
	TimestampNanoSec string             `json:"timestamp_nanosec"`
	PrevHash         string             `json:"prev_hash"`   // 父哈希
	PrevHeight       int64              `json:"prev_height"` // 父高度
	GasLimit         int64              `json:"gas_limit"`   // Gas 限制
	GasPrice         string             `json:"gas_price"`   // Gas 价格
}
