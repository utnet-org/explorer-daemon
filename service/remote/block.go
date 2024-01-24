package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

/**
The RPC API enables you to query the network and get details about specific blocks or chunks.
*/

var url = config.EnvLoad(config.NodeHostKey) + ":" + config.EnvLoad(config.NodePortKey)

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetailsByFinal() types.BlockDetailsRes {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockFinalReq{
			Finality: "final",
		},
	}

	jsonRes := SendRemoteCall(requestBody, url)

	fmt.Printf("BlockDetailsByFinal Json Response:%s", jsonRes)
	var bdRes types.BlockDetailsRes
	err := json.Unmarshal(jsonRes, &bdRes)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("BlockDetailsByFinal bdRes:", bdRes)
	return bdRes
}

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetailsByBlockId(blockId int) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockIdReq{
			BlockId: blockId,
		},
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("BlockDetailsByBlockId Response:%s", body)
}

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetailsByBlockHash(blockHash string) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockHashReq{
			BlockHash: blockHash,
		},
	}

	body := SendRemoteCall(requestBody, url)

	fmt.Printf("BlockDetailsByBlockHash Response:%s", body)
}
