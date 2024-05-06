package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"fmt"
	"strconv"
)

// 定时获取最新的block details
func BlockDetailsByFinal() {
	ctx, client := es.GetESInstance()
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		fmt.Println("rpc error")
	}
	// insert Elasticsearch
	err = es.InsertBlockDetails(ctx, client, res.Result)
	err = es.InsertLastHeight(ctx, client, res.Result.Header.Height)
	BlockChangesRpc(1, res.Result.Header)
	pkg.PrintStruct(res.Result)
	// 获取chunk hash
	cHash := res.Result.Chunks[0].ChunkHash
	// 获取最新block中的chunk details
	ChunkDetailsByChunkId(cHash)
	fmt.Printf("chunk hash:%s", cHash)
	if err != nil {
		fmt.Println("InsertData error:", err)
	}
}

func BlockChangesRpc(rpcType pkg.BlockChangeRpcType, header types.BlockDetailsHeader) {
	var queryValue interface{}
	if rpcType == pkg.BlockChangeRpcFinal {
		queryValue = strconv.Itoa(int(header.Height))

	} else if rpcType == pkg.BlockChangeRpcHeight {
		queryValue = header.Height
	} else if rpcType == pkg.BlockChangeRpcHash {
		queryValue = header.Hash
	}
	res, err := remote.ChangesInBlock(rpcType, queryValue)
	if err != nil {
		fmt.Println("ChangesInBlock rpc error")
	}
	res.Result.Height = header.Height
	res.Result.Timestamp = header.Timestamp
	res.Result.TimestampNanosec = header.TimestampNanosec

	err = es.InsertBlockChanges(res.Result)
	pkg.PrintStruct(res.Result)
	if err != nil {
		fmt.Println("InsertData error:", err)
	}
}

func ChunkDetailsByChunkId(chunkHash string) {
	res, err := remote.ChunkDetailsByChunkId(chunkHash)
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
