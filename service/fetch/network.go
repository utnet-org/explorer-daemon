package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	"fmt"
)

func HandleNetworkInfo() error {
	res, err := remote.NetworkInfo()
	if err != nil {
		fmt.Println("rpc error")
		return err
	}
	if res == nil {
		fmt.Println("rpc error res nil")
		return err
	}
	// insert Elasticsearch
	err = es.InsertNetWorkInfo(res.Result)
	if err != nil {
		fmt.Println("[BlockDetailsByFinal] insert error:", err)
		return err
	}
	return nil
}
