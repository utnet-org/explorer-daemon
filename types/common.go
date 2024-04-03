package types

type RpcRequest struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type CommonRes struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	//Result  interface{} `json:"result"`
}
