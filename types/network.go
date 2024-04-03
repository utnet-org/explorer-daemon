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
	Result    NetworkInfoBody `json:"result"`
}

type NetworkInfoBody struct {
	ActivePeers         []ActivePeer    `json:"activePeers"`
	KnownProducers      []KnownProducer `json:"knownProducers"`
	NumActivePeers      int64           `json:"numActivePeers"`
	PeerMaxCount        int64           `json:"peerMaxCount"`
	ReceivedBytesPerSEC int64           `json:"receivedBytesPerSec"`
	SentBytesPerSEC     int64           `json:"sentBytesPerSec"`
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
	CommonRes CommonRes
	Result    ValidationStatusBody `json:"result"`
}

type ValidationStatusBody struct {
	CurrentFishermen  []CurrentFisherman `json:"currentFishermen"`
	CurrentProposals  []CurrentProposal  `json:"currentProposals"`
	CurrentValidators []CurrentValidator `json:"currentValidators"`
	EpochHeight       int64              `json:"epochHeight"`
	EpochStartHeight  int64              `json:"epochStartHeight"`
	NextFishermen     []NextFisherman    `json:"nextFishermen"`
	NextValidators    []NextValidator    `json:"nextValidators"`
	PrevEpochKickout  []string           `json:"prevEpochKickout"`
}

type CurrentFisherman struct {
	AccountID string `json:"accountId"`
	PublicKey string `json:"publicKey"`
	Stake     string `json:"stake"`
}

type CurrentProposal struct {
	AccountID string `json:"accountId"`
	PublicKey string `json:"publicKey"`
	Stake     string `json:"stake"`
}

type CurrentValidator struct {
	AccountID         string  `json:"accountId"`
	IsSlashed         bool    `json:"isSlashed"`
	NumExpectedBlocks int64   `json:"numExpectedBlocks"`
	NumProducedBlocks int64   `json:"numProducedBlocks"`
	PublicKey         string  `json:"publicKey"`
	Shards            []int64 `json:"shards"`
	Stake             string  `json:"stake"`
}

type NextFisherman struct {
	AccountID string `json:"accountId"`
	PublicKey string `json:"publicKey"`
	Stake     string `json:"stake"`
}

type NextValidator struct {
	AccountID string  `json:"accountId"`
	PublicKey string  `json:"publicKey"`
	Shards    []int64 `json:"shards"`
	Stake     string  `json:"stake"`
}
