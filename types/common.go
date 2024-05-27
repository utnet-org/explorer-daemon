package types

type RpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type CommonRes struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type RpcErrRes struct {
	Error RpcErr `json:"error"`
	CommonRes
}

type RpcErr struct {
	Cause   RpcErrCause `json:"cause"`
	Code    int64       `json:"code"`
	Data    string      `json:"data"`
	Message string      `json:"message"`
	Name    string      `json:"name"`
}

type RpcErrCause struct {
	Info map[string]interface{} `json:"info"`
	Name string                 `json:"name"`
}

type Page struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}
