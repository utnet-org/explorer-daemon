package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"fmt"
)

func InitFetchData() {
	// 定时执行RPC请求
	//ticker := time.NewTicker(time.Hour) // 例如，每小时执行一次
	//for range ticker.C {
	BlockDetailsByFinal()
	//BlockChangesByFinal()
	//}
}

func BlockDetailsByFinal() {
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		fmt.Println("rpc error")
	}
	// insert Elasticsearch
	err = es.InsertBlockDetails(es.ECTX, es.ECLIENT, res.Body)
	pkg.PrintStruct(res.Body)
	if err != nil {
		fmt.Println("InsertData error:", err)
		// 处理存储到Elasticsearch的错误
		//continue
	}
}

func BlockChangesByFinal() {
	res, err := remote.ChangeInBlockByFinal()
	if err != nil {
		fmt.Println("rpc error")
	}
	// insert Elasticsearch
	err = es.InsertBlockChanges(es.ECTX, es.ECLIENT, res.Body)
	pkg.PrintStruct(res.Body)
	if err != nil {
		fmt.Println("InsertData error:", err)
		// 处理存储到Elasticsearch的错误
		//continue
	}
}

func ChunkDetailsByChunkId() {
	_, err := remote.ChunkDetailsByChunkId("")
	if err != nil {
		fmt.Println("rpc error")
	}
}
