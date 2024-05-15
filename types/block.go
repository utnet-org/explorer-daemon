package types

type BlockIdReq struct {
	BlockId interface{} `json:"block_id"`
}

type BlockFinalReq struct {
	Finality string `json:"finality"`
}

type BlockHashReq struct {
	BlockHash string `json:"block_hash"`
}

type ChunkId struct {
	ChunkId string `json:"chunk_id"`
}

type LastHeightHash struct {
	Height int64  `json:"height"`
	Hash   string `json:"hash"`
}

type BlockDetailsReq struct {
	//BlockId string `json:"block_id"`
	//BlockHash string `json:"block_hash"`
	//Finality  string `json:"finality"`
	QueryType int         `json:"query_type"`
	QueryWord interface{} `json:"query_word"`
}

// Block Details Response

type BlockDetailsRes struct {
	CommonRes
	Result BlockDetailsResult `json:"result"`
}

type BlockDetailsResult struct {
	Author string             `json:"author"`
	Chunks []Chunk            `json:"chunks"`
	Header BlockDetailsHeader `json:"header"`
}

type Chunk struct {
	BalanceBurnt             string   `json:"balance_burnt,omitempty"`              // 燃烧的余额
	ChunkHash                string   `json:"chunk_hash,omitempty"`                 // 区块哈希
	EncodedLength            int64    `json:"encoded_length,omitempty"`             // 编码长度
	EncodedMerkleRoot        string   `json:"encoded_merkle_root,omitempty"`        // 编码默克尔根
	GasLimit                 int64    `json:"gas_limit,omitempty"`                  // Gas 限制
	GasUsed                  int64    `json:"gas_used,omitempty"`                   // 已使用的 Gas
	HeightCreated            int64    `json:"height_created,omitempty"`             // 创建高度
	HeightIncluded           int64    `json:"height_included,omitempty"`            // 包含高度
	OutcomeRoot              string   `json:"outcome_root,omitempty"`               // 结果根
	OutgoingReceiptsRoot     string   `json:"outgoing_receipts_root,omitempty"`     // 发出的收据根
	PrevBlockHash            string   `json:"prev_block_hash,omitempty"`            // 上一个区块哈希
	PrevStateRoot            string   `json:"prev_state_root,omitempty"`            // 上一个状态根
	RentPaid                 string   `json:"rent_paid,omitempty"`                  // 租金支付
	ShardID                  int64    `json:"shard_id,omitempty"`                   // 分片ID
	Signature                string   `json:"signature,omitempty"`                  // 签名
	TxRoot                   string   `json:"tx_root,omitempty"`                    // 交易根
	ValidatorFrozenProposals []string `json:"validator_frozen_proposals,omitempty"` // 验证者冻结提案
	ValidatorPowerProposals  []string `json:"validator_power_proposals,omitempty"`  // 验证者权力提案
	ValidatorReward          string   `json:"validator_reward,omitempty"`           // 验证者奖励
}

type BlockDetailsHeader struct {
	Approvals                []string    `json:"approvals"`                  // 确认
	BlockBodyHash            string      `json:"block_body_hash"`            // 区块体哈希
	BlockMerkleRoot          string      `json:"block_merkle_root"`          // 区块默克尔根
	BlockOrdinal             int64       `json:"block_ordinal"`              // 区块序号
	ChallengesResult         []string    `json:"challenges_result"`          // 挑战结果
	ChallengesRoot           string      `json:"challenges_root"`            // 挑战根
	ChunkHeadersRoot         string      `json:"chunk_headers_root"`         // 区块头根
	ChunkMask                []bool      `json:"chunk_mask"`                 // 区块掩码
	ChunkReceiptsRoot        string      `json:"chunk_receipts_root"`        // 区块收据根
	ChunkTxRoot              string      `json:"chunk_tx_root"`              // 区块交易根
	ChunksIncluded           int64       `json:"chunks_included"`            // 包含的区块数
	EpochID                  string      `json:"epoch_id"`                   // 时代ID
	EpochSyncDataHash        interface{} `json:"epoch_sync_data_hash"`       // 时代同步数据哈希
	GasPrice                 string      `json:"gas_price"`                  // Gas 价格
	Hash                     string      `json:"hash"`                       // 哈希
	Height                   int64       `json:"height"`                     // 高度
	LastDsFinalBlock         string      `json:"last_ds_final_block"`        // 上一个DS最终区块
	LastFinalBlock           string      `json:"last_final_block"`           // 上一个最终区块
	LatestProtocolVersion    int64       `json:"latest_protocol_version"`    // 最新协议版本
	NextBpHash               string      `json:"next_bp_hash"`               // 下一个BP哈希
	NextEpochID              string      `json:"next_epoch_id"`              // 下一个时代ID
	OutcomeRoot              string      `json:"outcome_root"`               // 结果根
	PrevHash                 string      `json:"prev_hash"`                  // 上一个哈希
	PrevHeight               int64       `json:"prev_height"`                // 上一个高度
	PrevStateRoot            string      `json:"prev_state_root"`            // 上一个状态根
	RandomValue              string      `json:"random_value"`               // 随机值
	RentPaid                 string      `json:"rent_paid"`                  // 租金支付
	Signature                string      `json:"signature"`                  // 签名
	Timestamp                int64       `json:"timestamp"`                  // 时间戳
	TimestampNanosec         string      `json:"timestamp_nanosec"`          // 时间戳纳秒
	TotalSupply              string      `json:"total_supply"`               // 总供应量
	ValidatorFrozenProposals []string    `json:"validator_frozen_proposals"` // 验证者冻结提案
	ValidatorPowerProposals  []string    `json:"validator_power_proposals"`  // 验证者权力提案
	ValidatorReward          string      `json:"validator_reward"`           // 验证者奖励
}

type BlockChangesReq struct {
	Finality string      `json:"finality"`
	BlockId  interface{} `json:"block_id"`
}

// Block Changes Response

type BlockChangesRes struct {
	CommonRes CommonRes
	Result    BlockChangesResult `json:"result"`
}

type BlockChangesResult struct {
	Height           int64    `json:"height"`
	BlockHash        string   `json:"block_hash"`
	Timestamp        int64    `json:"timestamp"`         // 时间戳
	TimestampNanosec string   `json:"timestamp_nanosec"` // 时间戳纳秒
	Changes          []Change `json:"changes"`
}

type Change struct {
	AccountID string `json:"accountId"`
	Type      string `json:"type"`
}

// Chunk Details Response

type ChunkDetailsRes struct {
	CommonRes CommonRes
	Result    ChunkDetailsResult `json:"result"`
}

type ChunkDetailsResult struct {
	Author       string             `json:"author"`
	Header       ChunkDetailsHeader `json:"header"`
	Receipts     []string           `json:"receipts"`
	Transactions []string           `json:"transactions"`
}

type ChunkDetailsHeader struct {
	BalanceBurnt         string   `json:"balance_burnt"`
	ChunkHash            string   `json:"chunk_hash"`
	EncodedLength        int64    `json:"encoded_length"`
	EncodedMerkleRoot    string   `json:"encoded_merkle_root"`
	GasLimit             int64    `json:"gas_limit"`
	GasUsed              int64    `json:"gas_used"`
	HeightCreated        int64    `json:"height_created"`
	HeightIncluded       int64    `json:"height_included"`
	OutcomeRoot          string   `json:"outcome_root"`
	OutgoingReceiptsRoot string   `json:"outgoing_receipts_root"`
	PrevBlockHash        string   `json:"prev_block_hash"`
	PrevStateRoot        string   `json:"prev_state_root"`
	RentPaid             string   `json:"rent_paid"`
	ShardID              int64    `json:"shard_id"`
	Signature            string   `json:"signature"`
	TxRoot               string   `json:"tx_root"`
	ValidatorProposals   []string `json:"validator_proposals"`
	ValidatorReward      string   `json:"validator_reward"`
}
