package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/**
The RPC API enables you to query the network and get details about specific blocks or chunks.
*/

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetailsByFinal() (*types.BlockDetailsRes, error) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockFinalReq{
			Finality: "final",
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	if err != nil {
		log.Error("[BlockDetailsByFinal] Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debug("[BlockDetailsByFinal] Rpc Response:%s", jsonRes)
	var res types.BlockDetailsRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Error("[BlockDetailsByFinal] Error unmarshalling JSON:", err)
	}
	log.Debug("[BlockDetailsByFinal] res:", res)
	return &res, err
}

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetailsByBlockId(blockId interface{}) (*types.BlockDetailsRes, error) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockIdReq{
			BlockId: blockId,
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debug("[BlockDetailsByBlockId] Rpc Response:%s", jsonRes)
	var res types.BlockDetailsRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Error("[BlockDetailsByBlockId] Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debug("[BlockDetailsByBlockId] res:", res)
	return &res, err
}

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
//func BlockDetailsByBlockHash(blockHash string) {
//	requestBody := types.RpcRequest{
//		JsonRpc: config.JsonRpc,
//		ID:      config.RpcId,
//		Method:  "block",
//		Params: types.BlockHashReq{
//			BlockHash: blockHash,
//		},
//	}
//	body,_ := SendRemoteCall(requestBody, url)
//	fmt.Printf("BlockDetailsByBlockHash Response:%s", body)
//}

// Returns changes in block for given block height or hash. You can also use finality param to return latest block details.
// rpcType 0 final 1 block_id
func ChangesInBlock(rpcType pkg.BlockChangeRpcType, value interface{}) (types.BlockChangesRes, error) {
	var params types.BlockChangesReq
	if rpcType == pkg.BlockChangeRpcFinal {
		params.Finality = "final"
	} else {
		params.BlockId = value
	}
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "EXPERIMENTAL_changes_in_block",
		Params:  params,
	}

	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[ChangesInBlock] Json Response:%s", jsonRes)
	var res types.BlockChangesRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Errorf("Error unmarshalling JSON: %v", err)
	}
	log.Debugf("ChangeInBlockByFinal res: %v", res)
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
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("ChunkDetailsByChunkId Json Response:%s", jsonRes)
	var res types.ChunkDetailsRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	log.Debugln("ChunkDetailsByChunkId res:", res)
	return res, err
}

func ChunkDetailsByBlockId(chunkId string) (types.ChunkDetailsRes, error) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "chunk",
		Params: types.ChunkId{
			ChunkId: chunkId,
		},
	}

	jsonRes, err := SendRemoteCall(requestBody, url)

	fmt.Printf("ChunkDetailsByBlockId Json Response:%s", jsonRes)
	var res types.ChunkDetailsRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("ChunkDetailsByBlockId res:", res)
	return res, err
}
