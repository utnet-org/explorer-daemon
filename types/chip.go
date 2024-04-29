package types

type ChipQueryReq struct {
	RequestType string `json:"request_type"`
	Finality    string `json:"finality"`
	AccountId   string `json:"account_id"`
}

type ChipQueryRes struct {
	CommonRes CommonRes
	Result    ChipQueryResult `json:"result"`
}

type ChipQueryResult struct {
	BlockHash   string `json:"block_hash"`
	BlockHeight int64  `json:"block_height"`
	Chip        []Chip `json:"chips"`
	TotalPower  int64  `json:"total_power"`
}

type Chip struct {
	MinerId   string `json:"miner_id"`
	Power     int64  `json:"power"`
	BusId     string `json:"bus_id"`
	PublicKey string `json:"public_key"`
	ChipSN    string `json:"sn"`
	P2Key     string `json:"p2key"`
}