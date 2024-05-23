package remote

import (
	"bytes"
	"encoding/json"
	"errors"
	"explorer-daemon/config"
	"explorer-daemon/types"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var url = config.EnvLoad(config.NodeHostKey) + ":" + config.EnvLoad(config.NodePortKey)

func SendRemoteCall(requestBody types.RpcRequest, url string) ([]byte, error) {
	var body []byte
	var err error
	var maxRetries = 3 // Retry 3 times
	for i := 1; i <= maxRetries; i++ {
		body, err = sendRequest(requestBody, url)
		if err == nil {
			return body, nil
		}
		if err.Error() == "UNKNOWN_BLOCK" {
			return nil, err
		}
		log.Errorf("[SendRemoteCall] Attempt %d failed error: %v", i, err)
		time.Sleep(200 * time.Millisecond)
	}
	return nil, errors.New("max retries reached")
}

func sendRequest(requestBody types.RpcRequest, url string) ([]byte, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Error("[SendRemoteCall] JSON marshal error:", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		log.Errorf("[SendRemoteCall] Remote error: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("[SendRemoteCall] Error closing body error %v:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("[SendRemoteCall] Read body error: %v", err)
		return nil, err
	}

	var rpcErrRes types.RpcErrRes
	err = json.Unmarshal(body, &rpcErrRes)
	if err != nil {
		log.Errorf("[SendRemoteCall] JSON unmarshal error: %v", err)
		return nil, err
	}

	if rpcErrRes.Error.Code != 0 {
		log.Errorf("[SendRemoteCall] Rpc Error: %v, Code: %v", rpcErrRes.Error.Name, rpcErrRes.Error.Code)
		log.Errorf("[SendRemoteCall] Rpc Error Data: %v", rpcErrRes.Error.Data)
		return nil, errors.New(rpcErrRes.Error.Cause.Name)
	}
	return body, nil
}
