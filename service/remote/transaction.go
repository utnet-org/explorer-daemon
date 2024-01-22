package remote

import (
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

// Sends a transaction and immediately returns transaction hash.
//
// method: broadcast_tx_async
// params: [SignedTransaction encoded in base64]
func TransactionSendAsync() {
	params := make([]string, 0)
	params = append(params, "DgAAAHNlbmRlci50ZXN0bmV0AOrmAai64SZOv9e/naX4W15pJx0GAap35wTT1T/DwcbbDwAAAAAAAAAQAAAAcmVjZWl2ZXIudGVzdG5ldNMnL7URB1cxPOu3G8jTqlEwlcasagIbKlAJlF5ywVFLAQAAAAMAAACh7czOG8LTAAAAAAAAAGQcOG03xVSFQFjoagOb4NBBqWhERnnz45LY4+52JgZhm1iQKz7qAdPByrGFDQhQ2Mfga8RlbysuQ8D8LlA6bQE=")
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "status",
		Params: types.SignedTransactionReq{
			SignedTransaction: params,
		},
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("TransactionSendAsync Response:%s", body)
}
