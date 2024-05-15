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
)

var url = config.EnvLoad(config.NodeHostKey) + ":" + config.EnvLoad(config.NodePortKey)

func SendRemoteCall(requestBody types.RpcRequest, url string) ([]byte, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Error("[SendRemoteCall] JSON marshal error:", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		log.Errorf("[SendRemoteCall] POST remote error: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("[SendRemoteCall] Error closing body error %v:", err)
		}
	}(resp.Body)
	var rpcErrRes types.RpcErrRes
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &rpcErrRes)
	if rpcErrRes.Error.Code != 0 {
		log.Errorf("[SendRemoteCall] Rpc Error: %v, Code: %v", rpcErrRes.Error.Name, rpcErrRes.Error.Code)
		log.Errorf("[SendRemoteCall] Rpc Error Data: %v", rpcErrRes.Error.Data)
		return nil, errors.New(rpcErrRes.Error.Cause.Name)
	}
	return body, nil
}
