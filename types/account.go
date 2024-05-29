package types

type AccountRes struct {
	CommonRes
	Result AccountResult `json:"result"`
}

type AccountResult struct {
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

type AccountChange struct {
	ID      string              `json:"id"`
	Jsonrpc string              `json:"jsonrpc"`
	Result  AccountChangeResult `json:"result"`
}

type AccountChangeResult struct {
	BlockHash string          `json:"block_hash"`
	Changes   []ChangeElement `json:"changes"`
}

type ChangeElement struct {
	Cause  Cause        `json:"cause"`
	Change ChangeChange `json:"change"`
	Type   string       `json:"type"`
}

type Cause struct {
	ReceiptHash string  `json:"receipt_hash"`
	TxHash      *string `json:"tx_hash,omitempty"`
	Type        string  `json:"type"`
}

type ChangeChange struct {
	AccountID     string `json:"account_id"`
	Amount        string `json:"amount"`
	CodeHash      string `json:"code_hash"`
	Locked        string `json:"locked"`
	StoragePaidAt int64  `json:"storage_paid_at"`
	StorageUsage  int64  `json:"storage_usage"`
}

type AccountReq struct {
	AccountID   string `json:"account_id"`
	Finality    string `json:"finality"`
	RequestType string `json:"request_type"`
}

type ContractRes struct {
	CommonRes
	Result ContractResult `json:"result"`
}

type ContractResult struct {
	BlockHash   string `json:"block_hash"`
	BlockHeight int64  `json:"block_height"`
	CodeBase64  string `json:"code_base64"`
	Hash        string `json:"hash"`
}
