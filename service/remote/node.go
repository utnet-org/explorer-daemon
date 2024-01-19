package remote

import (
	"bytes"
	"encoding/json"
	"explorer-daemon/config"
	"explorer-daemon/types"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
)

type JSONRPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type JSONRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

func Experimental() {
	conn, err := rpc.DialHTTP("tcp", config.NodeAddress) // 替换成实际的地址和端口
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// 准备 JSON-RPC 请求参数
	request := `{
    "jsonrpc": "2.0",
    "id": "dontcare",
    "method": "EXPERIMENTAL_genesis_config",
    "params": {
        "block_id": 1
    }
}`

	var response map[string]interface{}

	err = conn.Call("rpc", request, &response)
	if err != nil {
		log.Fatal("RPC error:", err)
	}

	// 直接打印服务端返回的结果
	responseJSON, _ := json.Marshal(response)
	log.Println("Response:", string(responseJSON))
}

func ExperimentalHttp() {
	url := config.NodeAddress // 替换成实际的地址

	// 准备请求体结构体
	requestBody := types.RpcRequest{
		JsonRpc: "2.0",
		ID:      "dontcare",
		Method:  "EXPERIMENTAL_genesis_config",
		Params: types.BlockIdReq{
			BlockId: 1,
		},
	}

	// 转换结构体为 JSON 字符串
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("JSON marshal error:", err)
	}

	// 发起 HTTP POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	//resp, err := http.Post(url, "application/json", jsonBody)
	if err != nil {
		log.Fatal("HTTP POST error:", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read response error:", err)
	}

	// 打印响应结果
	fmt.Println("Response:", string(body))
}
