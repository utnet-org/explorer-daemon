package remote

import (
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

// Returns current genesis configuration.
func GenesisConfig() {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_genesis_config",
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("GenesisConfig Response:%s", body)
}

// Returns most recent protocol configuration or a specific queried block. Useful for finding current storage and transaction costs.
func ProtocolConfigByFinal() {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_protocol_config",
		Params: types.BlockFinalReq{
			Finality: "final",
		},
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("ProtocolConfigByFinal Response:%s", body)
}

// Returns most recent protocol configuration or a specific queried block. Useful for finding current storage and transaction costs.
func ProtocolConfigByBlockId(blockId int) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_protocol_config",
		Params: types.BlockIdReq{
			BlockId: blockId,
		},
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("ProtocolConfigByBlockId Response:%s", body)
}
