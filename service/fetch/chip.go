package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	"fmt"
)

func HandleChipQuery() error {
	res, err := remote.ChipsQuery()
	if err != nil {
		fmt.Println("rpc error")
		return err
	}
	//if res == nil {
	//	fmt.Println("rpc error res nil")
	//	return err
	//}
	// insert Elasticsearch
	err = es.InsertChip(res.Result)
	if err != nil {
		fmt.Println("[HandleChipQuery] insert error:", err)
		return err
	}
	return nil
}
