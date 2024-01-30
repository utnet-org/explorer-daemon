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
func BlockDetailsByFinal() (types.BlockDetailsRes, error) {

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
	return bdRes, err
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

// Returns changes in block for given block height or hash. You can also use finality param to return latest block details.
func ChangeInBlockByFinal() (types.BlockChangesRes, error) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_changes_in_block",
		Params: types.BlockFinalReq{
			Finality: "final",
		},
	}

	jsonRes := SendRemoteCall(requestBody, url)

	fmt.Printf("ChangeInBlockByFinal Json Response:%s", jsonRes)
	var res types.BlockChangesRes
	err := json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("ChangeInBlockByFinal res:", res)
	return res, err
}

// Returns details of a specific chunk. You can run a block details query to get a valid chunk hash.
func ChunkDetailsByChunkId(chunkId string) (types.ChunkDetailsRes, error) {

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "chunk",
		Params: types.ChunkId{
			ChunkId: chunkId,
		},
	}

	jsonRes := SendRemoteCall(requestBody, url)

	fmt.Printf("ChunkDetailsByChunkId Json Response:%s", jsonRes)
	var res types.ChunkDetailsRes
	err := json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("ChunkDetailsByChunkId res:", res)
	return res, err
}
