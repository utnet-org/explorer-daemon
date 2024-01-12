package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	conn, err := rpc.DialHTTP("tcp", "http://43.198.88.81:3030") // 替换成实际的地址和端口
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
	url := "http://43.198.88.81:3030" // 替换成实际的地址

	// 准备 JSON 请求体
	jsonBody := []byte(`{
		"jsonrpc": "2.0",
		"id": "dontcare",
		"method": "EXPERIMENTAL_genesis_config",
		"params": {
			"block_id": 1
		}
	}`)

	// 发起 HTTP POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal("HTTP POST error:", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read response error:", err)
	}

	// 打印响应结果
	fmt.Println("Response:", string(body))
}
