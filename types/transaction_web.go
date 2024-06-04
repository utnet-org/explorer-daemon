package types

type TxnResWeb struct {
	Height    int64  `json:"height"`
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
	TxnType   string `json:"txn_type"`
	//PublicKey  string `json:"public_key"`
	ReceiverID string `json:"receiver_id"`
	//Signature  string `json:"signature"`
	SignerID string  `json:"signer_id"`
	Deposit  string  `json:"deposit"`
	TxnFee   float64 `json:"txn_fee"`
}

type TxnListResWeb struct {
	Total   int64       `json:"total"`
	TxnList []TxnResWeb `json:"txn_list"`
}

type TxnDetailResWeb struct {
	Hash      string `json:"hash"`
	Status    string `json:"status"`
	Height    int64  `json:"height"`
	Timestamp int64  `json:"timestamp"`
	TimeUTC   string `json:"time_utc"`
	//TxnType   string `json:"txn_type"`
	SignerID         string        `json:"signer_id"`
	ReceiverID       string        `json:"receiver_id"`
	TokenTransferred []interface{} `json:"token_transferred"`
	Deposit          string        `json:"deposit"`
	TxnFee           float64       `json:"txn_fee"`
}

type TxnDeployContractResWeb struct {
	TxnHash   string `json:"txn_hash"`
	Timestamp int64  `json:"timestamp"`
	TimeUTC   string `json:"time_utc"`
	CodeHash  string `json:"code_hash"`
	Height    int64  `json:"height"`
}
