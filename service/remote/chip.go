package remote

import (
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
)

func ChipsQuery() (types.ChipQueryRes, error) {
	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "query",
		Params: types.ChipQueryReq{
			RequestType: "view_chip_list",
			Finality:    "final",
			AccountId:   "guest-book.testnet",
		},
	}
	jsonRes := SendRemoteCall(requestBody, url)
	fmt.Printf("[ChipsQuery] Json Response:%s", jsonRes)
	var res types.ChipQueryRes
	err := json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("[ChipsQuery] Error unmarshalling JSON:", err)
	}
	fmt.Println("[ChipsQuery] res:", res)
	return res, err
}
