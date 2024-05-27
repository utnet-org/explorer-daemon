package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Returns gas price for a specific block_height or block_hash.
func GasPriceByBlockHeight(blockHeights []int64) (*types.GasPriceResult, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  blockHeights,
	}

	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[GasPriceByBlockHeight] Json Response:%s", jsonRes)
	var res types.GasPriceRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debugln("[GasPriceByBlockHeight] res:", res)
	return &res.Result, nil
}

// Returns gas price for a specific block_height or block_hash.
func GasPriceByBlockHash(blockHash []string) (*types.GasPriceResult, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  blockHash,
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[GasPriceByBlockHash] Json Response:%s", jsonRes)
	var res types.GasPriceResult
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debugln("[GasPriceByBlockHash] res:", res)
	return &res, nil
}

// Returns gas price for a specific block_height or block_hash.
func GasPriceByNull() (*types.GasPriceResult, error) {
	params := []interface{}{nil}
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "gas_price",
		Params:  params,
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[GasPriceByBlockHash] Json Response:%s", jsonRes)
	var res types.GasPriceResult
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debugln("[GasPriceByBlockHash] res:", res)
	return &res, nil
}
