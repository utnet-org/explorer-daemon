package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	log "github.com/sirupsen/logrus"
)

// Returns basic account information.
func ViewAccount(accountId string) (*types.AccountResult, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "query",
		Params: types.AccountReq{
			AccountID:   accountId,
			Finality:    "final",
			RequestType: "view_account",
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	if err != nil {
		log.Errorf("[ViewAccount] Rpc Error: %v", err)
		return nil, err
	}
	log.Debugf("[ViewAccount] Rpc Response: %s", jsonRes)
	var res types.AccountRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Error("[ViewAccount] Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debug("[ViewAccount] res:", res)
	return &res.Result, err
}

// Returns the contract code (Wasm binary) deployed to the account. Please note that the returned code will be encoded in base64.
func ViewContractCode(accountId string) (*types.ContractResult, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "query",
		Params: types.AccountReq{
			AccountID:   accountId,
			Finality:    "final",
			RequestType: "view_code",
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	if err != nil {
		log.Errorf("[ViewContractCode] Rpc Error: %v", err)
		return nil, err
	}
	log.Debugf("[ViewContractCode] Rpc Response: %s", jsonRes)
	var res types.ContractRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Error("[ViewContractCode] Error unmarshalling JSON:", err)
		return nil, err
	}
	log.Debug("[ViewContractCode] res:", res)
	return &res.Result, err
}
