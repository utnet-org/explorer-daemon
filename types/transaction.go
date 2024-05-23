package types

// Transaction Request
type SignedTransactionReq struct {
	SignedTransaction []string
}

type TransReceiptByIdReq struct {
	ReceiptId string `json:"receiptId"`
}

type TxnStatusReq struct {
	TxHash          string `json:"tx_hash"`
	SenderAccountId string `json:"sender_account_id"`
	WaitUntil       string `json:"wait_until,omitempty"`
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

type TxnStatusRes struct {
	CommonRes
	Result TxnStatusResult `json:"result"`
}

type TxnStatusResult struct {
	FinalExeStatus  string            `json:"final_execution_status"`
	ReceiptsOutcome []ReceiptsOutcome `json:"receipts_outcome"`
	//Status             ResultStatus       `json:"status"`
	Status             interface{}        `json:"status"`
	Transaction        Transaction        `json:"transaction"`
	TransactionOutcome TransactionOutcome `json:"transaction_outcome"`
}

type ReceiptsOutcome struct {
	BlockHash string                 `json:"block_hash"`
	ID        string                 `json:"id"`
	Outcome   ReceiptsOutcomeOutcome `json:"outcome"`
	Proof     []OutcomeProof         `json:"proof"`
}

type ReceiptsOutcomeOutcome struct {
	OutcomeOutcome
	Status interface{} `json:"status"`
}

type ReceiptsMetadata struct {
	GasProfile []string `json:"gas_profile"`
	Version    int64    `json:"version"`
}

type PurpleStatus struct {
	SuccessValue string `json:"successValue"`
}

type OutcomeProof struct {
	Direction string `json:"direction"`
	Hash      string `json:"hash"`
}

type ResultStatus struct {
	SuccessValue string `json:"SuccessValue"`
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
	BlockHash string                    `json:"block_hash"`
	ID        string                    `json:"id"`
	Outcome   TransactionOutcomeOutcome `json:"outcome"`
	Proof     []OutcomeProof            `json:"proof"`
}

type OutcomeOutcome struct {
	ExecutorID  string      `json:"executor_id"`
	GasBurnt    int64       `json:"gas_burnt"`
	Metadata    interface{} `json:"metadata"`
	Logs        []string    `json:"logs"`
	ReceiptIDS  []string    `json:"receipt_ids"`
	TokensBurnt string      `json:"tokens_burnt"`
}
type TransactionOutcomeOutcome struct {
	OutcomeOutcome
	Status interface{} `json:"status"`
}

type FluffyStatus struct {
	SuccessReceiptID string `json:"successReceiptId"`
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
