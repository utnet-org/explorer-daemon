package types

type AccountResultWeb struct {
	Amount        string `json:"amount"`
	BlockHash     string `json:"block_hash"`
	BlockHeight   int64  `json:"block_height"`
	CodeHash      string `json:"code_hash"`
	Locked        string `json:"locked"`
	Pledging      string `json:"pledging"`
	Power         string `json:"power"`
	StoragePaidAt int64  `json:"storage_paid_at"`
	StorageUsage  int64  `json:"storage_usage"`
}
