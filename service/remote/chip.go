package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func ChipsQuery() (types.ChipQueryRes, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: config.Jsonrpc,
		ID:      config.RpcId,
		Method:  "query",
		Params: types.ChipQueryReq{
			RequestType: "view_chip_list",
			Finality:    "final",
			AccountId:   "guest-book.testnet",
		},
	}
	jsonRes, err := SendRemoteCall(requestBody, url)
	log.Debugf("[ChipsQuery] Json Response:%s", jsonRes)
	var res types.ChipQueryRes
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("[ChipsQuery] Error unmarshalling JSON:", err)
	}
	log.Debugln("[ChipsQuery] res:", res)
	return res, err
}
