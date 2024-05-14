package types

// Network Node Status Response

type NetworkNodeStatusRes struct {
	CommonRes CommonRes
	Result    NetworkNodeStatusBody `json:"result"`
}

type NetworkNodeStatusBody struct {
	ChainID               string             `json:"chainId"`
	LatestProtocolVersion int64              `json:"latestProtocolVersion"`
	ProtocolVersion       int64              `json:"protocolVersion"`
	RPCAddr               string             `json:"rpcAddr"`
	SyncInfo              SyncInfo           `json:"syncInfo"`
	ValidatorAccountID    string             `json:"validatorAccountId"`
	Validators            []NetworkValidator `json:"validators"`
	Version               Version            `json:"version"`
}

type SyncInfo struct {
	LatestBlockHash   string `json:"latestBlockHash"`
	LatestBlockHeight int64  `json:"latestBlockHeight"`
	LatestBlockTime   string `json:"latestBlockTime"`
	LatestStateRoot   string `json:"latestStateRoot"`
	Syncing           bool   `json:"syncing"`
}

type NetworkValidator struct {
	AccountID string `json:"accountId"`
	IsSlashed bool   `json:"isSlashed"`
}

type Version struct {
	Build   string `json:"build"`
	Version string `json:"version"`
}

// Network Info Response
type NetworkInfoRes struct {
	CommonRes CommonRes
	Result    NetworkInfoResult `json:"result"`
}

type NetworkInfoResult struct {
	ActivePeers         []ActivePeer    `json:"active_peers"`           // 活跃的节点列表
	KnownProducers      []KnownProducer `json:"known_producers"`        // 已知的生产者列表
	NumActivePeers      int64           `json:"num_active_peers"`       // 活跃节点的数量
	PeerMaxCount        int64           `json:"peer_max_count"`         // 节点的最大数量限制
	ReceivedBytesPerSEC int64           `json:"received_bytes_per_sec"` // 每秒接收的字节数
	SentBytesPerSEC     int64           `json:"sent_bytes_per_sec"`     // 每秒发送的字节数
}

type ActivePeer struct {
	AccountID interface{} `json:"accountId"`
	Addr      *string     `json:"addr,omitempty"`
	ID        *string     `json:"id,omitempty"`
}

type KnownProducer struct {
	AccountID *string     `json:"accountId,omitempty"`
	Addr      interface{} `json:"addr"`
	PeerID    *string     `json:"peerId,omitempty"`
}

// Validation Status Response

type ValidationStatusRes struct {
	CommonRes
	Result ValidationStatusResult `json:"result"`
}

type ValidationStatusResult struct {
	CurrentFishermen       []string           `json:"current_fishermen"`
	CurrentPledgeProposals []string           `json:"current_pledge_proposals"`
	CurrentPowerProposals  []string           `json:"current_power_proposals"`
	CurrentValidators      []CurrentValidator `json:"current_validators"`
	EpochHeight            int64              `json:"epoch_height"`
	EpochStartHeight       int64              `json:"epoch_start_height"`
	NextFishermen          []string           `json:"next_fishermen"`
	NextValidators         []NextValidator    `json:"next_validators"`
	PrevEpochKickout       []string           `json:"prev_epoch_kickout"`
}

type CurrentValidator struct {
	AccountID                 string  `json:"account_id"`
	IsSlashed                 bool    `json:"is_slashed"`
	NumExpectedBlocks         int64   `json:"num_expected_blocks"`
	NumExpectedChunks         int64   `json:"num_expected_chunks"`
	NumExpectedChunksPerShard []int64 `json:"num_expected_chunks_per_shard"`
	NumProducedBlocks         int64   `json:"num_produced_blocks"`
	NumProducedChunks         int64   `json:"num_produced_chunks"`
	NumProducedChunksPerShard []int64 `json:"num_produced_chunks_per_shard"`
	Pledge                    string  `json:"pledge"`
	Power                     string  `json:"power"`
	PublicKey                 string  `json:"public_key"`
	Shards                    []int64 `json:"shards"`
}

type NextValidator struct {
	AccountID string  `json:"account_id"`
	Pledge    string  `json:"pledge"`
	Power     string  `json:"power"`
	PublicKey string  `json:"public_key"`
	Shards    []int64 `json:"shards"`
}
