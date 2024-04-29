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

var url = config.EnvLoad(config.NodeHostKey) + ":" + config.EnvLoad(config.NodePortKey)

func SendRemoteCall(requestBody types.RpcRequest, url string) []byte {
	jsonBody, err := json.Marshal(requestBody)
	fmt.Println("request body:", requestBody)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println("BlockDetails POST error:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	return body
}
