package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Sends a transaction and immediately returns transaction hash.
//
// method: broadcast_tx_async
// params: [SignedTransaction encoded in base64]
func TransactionSendAsync() {
	params := make([]string, 0)
	params = append(params, "DgAAAHNlbmRlci50ZXN0bmV0AOrmAai64SZOv9e/naX4W15pJx0GAap35wTT1T/DwcbbDwAAAAAAAAAQAAAAcmVjZWl2ZXIudGVzdG5ldNMnL7URB1cxPOu3G8jTqlEwlcasagIbKlAJlF5ywVFLAQAAAAMAAACh7czOG8LTAAAAAAAAAGQcOG03xVSFQFjoagOb4NBBqWhERnnz45LY4+52JgZhm1iQKz7qAdPByrGFDQhQ2Mfga8RlbysuQ8D8LlA6bQE=")
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "broadcast_tx_async",
		Params: types.SignedTransactionReq{
			SignedTransaction: params,
		},
	}

	body, _ := SendRemoteCall(requestBody, url)

	fmt.Printf("TransactionSendAsync Response:%s", body)
}

// Sends a transaction and waits until transaction is fully complete. (Has a 10 second timeout)
//
// method: broadcast_tx_commit
// params: [SignedTransaction encoded in base64]
func TransactionSendAwait() {
	params := make([]string, 0)
	params = append(params, "DgAAAHNlbmRlci50ZXN0bmV0AOrmAai64SZOv9e/naX4W15pJx0GAap35wTT1T/DwcbbDwAAAAAAAAAQAAAAcmVjZWl2ZXIudGVzdG5ldNMnL7URB1cxPOu3G8jTqlEwlcasagIbKlAJlF5ywVFLAQAAAAMAAACh7czOG8LTAAAAAAAAAGQcOG03xVSFQFjoagOb4NBBqWhERnnz45LY4+52JgZhm1iQKz7qAdPByrGFDQhQ2Mfga8RlbysuQ8D8LlA6bQE=")
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "broadcast_tx_commit",
		Params: types.SignedTransactionReq{
			SignedTransaction: params,
		},
	}

	body, _ := SendRemoteCall(requestBody, url)

	fmt.Printf("TransactionSendAwait Response:%s", body)
}

// Queries status of a transaction by hash and returns the final transaction result.
//
// method: tx
// params:
// transaction hash (see UtilityBlocks Explorer for a valid transaction hash)
// sender account id
func TxnStatus(hash, accountId, waitUntil string) (*types.TxnStatusResult, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "tx",
		Params: types.TxnStatusReq{
			TxHash:          hash,
			SenderAccountId: accountId,
			WaitUntil:       waitUntil,
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[TransactionStatus] Json Response:%s", jsonRes)
	var res types.TxnStatusRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	log.Debugln("[TransactionStatus] res:", res)
	return &res.Result, nil
}

// Queries status of a transaction by hash, returning the final transaction result and details of all receipts.
//
// method: EXPERIMENTAL_tx_status
// params:
// transaction hash (see UtilityBlocks Explorer for a valid transaction hash)
// sender account id (used to determine which shard to query for transaction)
func TransactionStatusReceipts() {
	params := make([]string, 0)
	params = append(params, "HEgnVQZfs9uJzrqTob4g2Xmebqodq9waZvApSkrbcAhd")
	params = append(params, "bowen")
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_tx_status",
		Params: types.SignedTransactionReq{
			SignedTransaction: params,
		},
	}

	body, _ := SendRemoteCall(requestBody, url)

	fmt.Printf("TransactionStatusReceipts Response:%s", body)
}

// Fetches a receipt by it's ID (as is, without a status or execution outcome)
//
// method: EXPERIMENTAL_receipt
// params:
// receipt_id (see Utility Explorer for a valid receipt id)
func TransactionReceiptsById(receiptId string) {

	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_receipt",
		Params: types.TransReceiptByIdReq{
			ReceiptId: receiptId,
		},
	}

	body, _ := SendRemoteCall(requestBody, url)

	fmt.Printf("TransactionReceiptsById Response:%s", body)
}
