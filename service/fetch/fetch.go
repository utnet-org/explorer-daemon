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

// 定时获取最新的block details
func BlockDetailsByFinal() {
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		fmt.Println("rpc error")
	}
	// insert Elasticsearch
	err = es.InsertBlockDetails(es.ECTX, es.ECLIENT, res.Body)
	pkg.PrintStruct(res.Body)
	// 获取chunk hash
	cHash := res.Body.Chunks[0].ChunkHash
	// 获取最新block中的chunk details
	ChunkDetailsByChunkId(cHash)
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

func ChunkDetailsByChunkId(chunkHash string) {
	res, err := remote.ChunkDetailsByChunkId(chunkHash)
	// 存入es
	err = es.InsertChunkDetails(res.Body, chunkHash)
	if err != nil {
		fmt.Println("InsertChunkDetails error:", err)
	}
	pkg.PrintStruct(res.Body)
}

func ChunkDetailsByBlockId() {
	_, err := remote.ChunkDetailsByBlockId("")
	if err != nil {
		fmt.Println("rpc error")
	}
}
