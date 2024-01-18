package types

// Request 结构体用于表示 JSON 请求体的结构
type RpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}
