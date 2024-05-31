package types

type ContractDetailResultWeb struct {
	BlockHash   string `json:"block_hash"`
	BlockHeight int64  `json:"block_height"`
	TimeStamp   string `json:"timestamp"`
	TxnHash     string `json:"txn_hash"`
	Locked      string `json:"locked"`
	CodeHash    string `json:"code_hash"`
	CodeBase64  string `json:"code_base64"`
}
