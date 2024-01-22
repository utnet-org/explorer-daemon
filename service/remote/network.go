package remote

import (
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

// Returns general status of a given node (sync status, utility core node version, protocol version, etc), and the current set of validators.
func NetworkNodeStatus() {
	params := make([]interface{}, 0)
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "status",
		Params:  params,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("NetworkNodeStatus Response:%s", body)
}

// Returns the current state of node network connections (active peers, transmitted data, etc.)
func NetworkInfo() {
	params := make([]interface{}, 0)
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "network_info",
		Params:  params,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("NetworkInfo Response:%s", body)
}

//Queries active validators on the network returning details and the state of validation on the blockchain.

// method: validators
// params: ["block hash"], [block number], {"epoch_id": "epoch id"}, {"block_id": block number}, {"block_id": "block hash"}, or [null] for the latest block
func ValidationStatusByBlockNumber(blockNumber int) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "validators",
		Params:  blockNumber,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("NetworkValidationStatus Response:%s", body)
}

func ValidationStatusByNull() {
	params := []interface{}{nil}
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "validators",
		Params:  params,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("ValidationStatusByNull Response:%s", body)
}
