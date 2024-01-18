package remote

import (
	"bytes"
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
	"io"
	"net/http"
)

// Queries network and returns block for given height or hash. You can also use finality param to return latest block details.
func BlockDetails() {
	url := config.NodeAddress

	requestBody := types.RpcRequest{
		JsonRpc: config.JsonRpc,
		ID:      config.RpcId,
		Method:  "block",
		Params: types.BlockReq2{
			Finality: "final",
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	fmt.Println("request body:", requestBody)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
	}

	// 发起 HTTP POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	//resp, err := http.Post(url, "application/json", jsonBody)
	if err != nil {
		fmt.Println("HTTP POST error:", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response error:", err)
	}

	// 打印响应结果
	fmt.Println("Response:", string(body))
}
