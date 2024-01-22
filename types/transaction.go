package types

type SignedTransactionReq struct {
	SignedTransaction []string
}

// Transaction Send Response

type TransactionSendRes struct {
	CommonRes CommonRes
	Body      TransactionSendBody `json:"result"`
}

type TransactionSendBody struct {
	Result string
}
