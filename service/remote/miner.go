package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	log "github.com/sirupsen/logrus"
)

func AllMiners(blockHash interface{}) (*types.AllMinersRes, error) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "all_miners",
		Params: types.BlockHashReq{
			BlockHash: blockHash.(string),
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[AllMiners] Rpc Response:%v", jsonRes)
	var res types.AllMinersRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		log.Errorf("[AllMiners] Error unmarshalling JSON error: %v", err)
	}
	log.Debugln("[AllMiners] res:", res)
	return &res, err
}
