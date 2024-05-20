package types

// Transaction Request
type SignedTransactionReq struct {
	SignedTransaction []string
}

type TransReceiptByIdReq struct {
	ReceiptId string `json:"receiptId"`
}

// Transaction Send Response

type TransSendRes struct {
	CommonRes CommonRes
	Body      TransSendBody `json:"result"`
}

type TransSendBody struct {
	Result string
}

// Transaction Status Response

type TransStatusRes struct {
	CommonRes CommonRes
	Result    TransStatusBody `json:"result"`
}

type TransStatusBody struct {
	ReceiptsOutcome    []ReceiptsOutcome  `json:"receiptsOutcome"`
	Status             ResultStatus       `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transactionOutcome"`
}

type ReceiptsOutcome struct {
	BlockHash string                 `json:"blockHash"`
	ID        string                 `json:"id"`
	Outcome   ReceiptsOutcomeOutcome `json:"outcome"`
	Proof     []ReceiptsOutcomeProof `json:"proof"`
}

type ReceiptsOutcomeOutcome struct {
	ExecutorID  string       `json:"executorId"`
	GasBurnt    int64        `json:"gasBurnt"`
	Logs        []string     `json:"logs"`
	ReceiptIDS  []string     `json:"receiptIds"`
	Status      PurpleStatus `json:"status"`
	TokensBurnt string       `json:"tokensBurnt"`
}

type PurpleStatus struct {
	SuccessValue string `json:"successValue"`
}

type ReceiptsOutcomeProof struct {
	Direction string `json:"direction"`
	Hash      string `json:"hash"`
}

type ResultStatus struct {
	SuccessValue string `json:"successValue"`
}

type Transaction struct {
	// TODO many type actions
	//Actions    []Action `json:"actions"`
	Actions    []interface{} `json:"actions"`
	Hash       string        `json:"hash"`
	Nonce      int64         `json:"nonce"`
	PublicKey  string        `json:"public_key"`
	ReceiverID string        `json:"receiver_id"`
	Signature  string        `json:"signature"`
	SignerID   string        `json:"signer_id"`
}

type Action struct {
	Transfer *Transfer `json:"transfer,omitempty"`
}

type Transfer struct {
	Deposit string `json:"deposit"`
}

type TransactionOutcome struct {
	BlockHash string                    `json:"blockHash"`
	ID        string                    `json:"id"`
	Outcome   TransactionOutcomeOutcome `json:"outcome"`
	Proof     []TransactionOutcomeProof `json:"proof"`
}

type TransactionOutcomeOutcome struct {
	ExecutorID  string       `json:"executorId"`
	GasBurnt    int64        `json:"gasBurnt"`
	Logs        []string     `json:"logs"`
	ReceiptIDS  []string     `json:"receiptIds"`
	Status      FluffyStatus `json:"status"`
	TokensBurnt string       `json:"tokensBurnt"`
}

type FluffyStatus struct {
	SuccessReceiptID string `json:"successReceiptId"`
}

type TransactionOutcomeProof struct {
	Direction *string `json:"direction,omitempty"`
	Hash      *string `json:"hash,omitempty"`
}

// Transaction Status Response With Receipts

type TransStatusReceiptsRes struct {
	ID      string                  `json:"id"`
	Jsonrpc string                  `json:"jsonrpc"`
	Result  TransStatusReceiptsBody `json:"result"`
}

type TransStatusReceiptsBody struct {
	Receipts           []ReceiptElement   `json:"receipts"`
	ReceiptsOutcome    []ReceiptsOutcome  `json:"receiptsOutcome"`
	Status             ResultStatus       `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transactionOutcome"`
}

type ReceiptElement struct {
	PredecessorID string         `json:"predecessor_id"`
	Receipt       ReceiptReceipt `json:"receipt"`
	ReceiptID     string         `json:"receipt_id"`
	ReceiverID    string         `json:"receiver_id"`
}

type ReceiptReceipt struct {
	Action ReceiptAction `json:"Action"`
}

type ReceiptAction struct {
	// TODO many type actions
	//Actions             []ActionAction `json:"actions"`
	Actions             []interface{} `json:"actions"`
	GasPrice            string        `json:"gas_price"`
	InputDataIDS        []string      `json:"input_data_ids"`
	OutputDataReceivers []string      `json:"output_data_receivers"`
	SignerID            string        `json:"signer_id"`
	SignerPublicKey     string        `json:"signer_public_key"`
}

type ActionAction struct {
	FunctionCall *PurpleFunctionCall `json:"functionCall,omitempty"`
	Transfer     Transfer            `json:"transfer"`
}

type PurpleFunctionCall struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"methodName"`
}

type TransactionAction struct {
	FunctionCall *FluffyFunctionCall `json:"functionCall,omitempty"`
}

type FluffyFunctionCall struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"methodName"`
}

// Transaction Receipt By ID Response

type TransReceiptIdRes struct {
	CommonRes CommonRes
	Result    TransReceiptIdBody `json:"result"`
}

type TransReceiptIdBody struct {
	PredecessorID string  `json:"predecessorId"`
	Receipt       Receipt `json:"receipt"`
	ReceiptID     string  `json:"receiptId"`
	ReceiverID    string  `json:"receiverId"`
}

type Receipt struct {
	Action ReceiptAction `json:"action"`
}

type ActionElement struct {
	Transfer *Transfer `json:"transfer,omitempty"`
}
