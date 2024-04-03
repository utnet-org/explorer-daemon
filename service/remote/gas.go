package remote

import (
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

// Returns gas price for a specific block_height or block_hash.
func GasPriceByBlockHeight(blockHeights []int) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  blockHeights,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("GasPriceByBlockHeight Response:%s", body)
}

// Returns gas price for a specific block_height or block_hash.
func GasPriceByBlockHash(blockHash []string) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  blockHash,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("GasPriceByBlockHeight Response:%s", body)
}

// Returns gas price for a specific block_height or block_hash.
func GasPriceByNull() {
	params := []interface{}{nil}

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  params,
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("GasPriceByBlockHeight Response:%s", body)
}
