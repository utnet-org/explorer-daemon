package types

type BlockIdReq struct {
	BlockId int `json:"block_id"`
}

type BlockFinalReq struct {
	Finality string `json:"finality"`
}

type BlockHashReq struct {
	BlockHash string `json:"block_id"`
}

// Block Details Response

type BlockDetailsRes struct {
	CommonRes CommonRes
	Body      BlockDetailsBody `json:"result"`
}

type BlockDetailsBody struct {
	Author string             `json:"author"`
	Chunks []Chunk            `json:"chunks"`
	Header BlockDetailsHeader `json:"header"`
}

type Chunk struct {
	BalanceBurnt         *string  `json:"balance_burnt,omitempty"`
	ChunkHash            *string  `json:"chunk_hash,omitempty"`
	EncodedLength        *int64   `json:"encoded_length,omitempty"`
	EncodedMerkleRoot    *string  `json:"encoded_merkle_root,omitempty"`
	GasLimit             *int64   `json:"gas_limit,omitempty"`
	GasUsed              *int64   `json:"gas_used,omitempty"`
	HeightCreated        *int64   `json:"height_created,omitempty"`
	HeightIncluded       *int64   `json:"height_included,omitempty"`
	OutcomeRoot          *string  `json:"outcome_root,omitempty"`
	OutgoingReceiptsRoot *string  `json:"outgoing_receipts_root,omitempty"`
	PrevBlockHash        *string  `json:"prev_block_hash,omitempty"`
	PrevStateRoot        *string  `json:"prev_state_root,omitempty"`
	RentPaid             *string  `json:"rent_paid,omitempty"`
	ShardID              *int64   `json:"shard_id,omitempty"`
	Signature            *string  `json:"signature,omitempty"`
	TxRoot               *string  `json:"tx_root,omitempty"`
	ValidatorProposals   []string `json:"validator_proposals,omitempty"`
	ValidatorReward      *string  `json:"validator_reward,omitempty"`
}

type BlockDetailsHeader struct {
	Approvals             []*string   `json:"approvals"`
	BlockBodyHash         interface{} `json:"block_body_hash"`
	BlockMerkleRoot       string      `json:"block_merkle_root"`
	BlockOrdinal          interface{} `json:"block_ordinal"`
	ChallengesResult      []string    `json:"challenges_result"`
	ChallengesRoot        string      `json:"challenges_root"`
	ChunkHeadersRoot      string      `json:"chunk_headers_root"`
	ChunkMask             []bool      `json:"chunk_mask"`
	ChunkReceiptsRoot     string      `json:"chunk_receipts_root"`
	ChunkTxRoot           string      `json:"chunk_tx_root"`
	ChunksIncluded        int64       `json:"chunks_included"`
	EpochID               string      `json:"epoch_id"`
	EpochSyncDataHash     interface{} `json:"epoch_sync_data_hash"`
	GasPrice              string      `json:"gas_price"`
	Hash                  string      `json:"hash"`
	Height                int64       `json:"height"`
	LastDsFinalBlock      string      `json:"last_ds_final_block"`
	LastFinalBlock        string      `json:"last_final_block"`
	LatestProtocolVersion int64       `json:"latest_protocol_version"`
	NextBpHash            string      `json:"next_bp_hash"`
	NextEpochID           string      `json:"next_epoch_id"`
	OutcomeRoot           string      `json:"outcome_root"`
	PrevHash              string      `json:"prev_hash"`
	PrevHeight            interface{} `json:"prev_height"`
	PrevStateRoot         string      `json:"prev_state_root"`
	RandomValue           string      `json:"random_value"`
	RentPaid              string      `json:"rent_paid"`
	Signature             string      `json:"signature"`
	Timestamp             int64       `json:"timestamp"`
	TimestampNanosec      string      `json:"timestamp_nanosec"`
	TotalSupply           string      `json:"total_supply"`
	ValidatorProposals    []string    `json:"validator_proposals"`
	ValidatorReward       string      `json:"validator_reward"`
}

// Block Changes Response

type BlockChangesRes struct {
	CommonRes CommonRes
	Body      BlockChangesBody `json:"result"`
}

type BlockChangesBody struct {
	BlockHash string   `json:"blockHash"`
	Changes   []Change `json:"changes"`
}

type Change struct {
	AccountID string `json:"accountId"`
	Type      string `json:"type"`
}

// Chunk Details Response

type ChunkDetailsRes struct {
	CommonRes CommonRes
	Body      ChunkDetailsBody `json:"result"`
}

type ChunkDetailsBody struct {
	Author       string             `json:"author"`
	Header       ChunkDetailsHeader `json:"header"`
	Receipts     []string           `json:"receipts"`
	Transactions []string           `json:"transactions"`
}

type ChunkDetailsHeader struct {
	BalanceBurnt         string   `json:"balanceBurnt"`
	ChunkHash            string   `json:"chunkHash"`
	EncodedLength        int64    `json:"encodedLength"`
	EncodedMerkleRoot    string   `json:"encodedMerkleRoot"`
	GasLimit             int64    `json:"gasLimit"`
	GasUsed              int64    `json:"gasUsed"`
	HeightCreated        int64    `json:"heightCreated"`
	HeightIncluded       int64    `json:"heightIncluded"`
	OutcomeRoot          string   `json:"outcomeRoot"`
	OutgoingReceiptsRoot string   `json:"outgoingReceiptsRoot"`
	PrevBlockHash        string   `json:"prevBlockHash"`
	PrevStateRoot        string   `json:"prevStateRoot"`
	RentPaid             string   `json:"rentPaid"`
	ShardID              int64    `json:"shardId"`
	Signature            string   `json:"signature"`
	TxRoot               string   `json:"txRoot"`
	ValidatorProposals   []string `json:"validatorProposals"`
	ValidatorReward      string   `json:"validatorReward"`
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
