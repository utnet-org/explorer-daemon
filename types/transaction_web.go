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
