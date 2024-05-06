package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func BlockDetailsByFinal() {
	ctx, client := es.GetESInstance()
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		fmt.Println("rpc error")
	}
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

func HandleBlock() error {
	ctx, client := es.GetESInstance()
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		return err
	}
	rpcHeight := res.Result.Header.Height
	lastHeight, err := es.GetLastHeight(client, ctx)
	if err != nil {
		log.Errorln("[HandleBlock] GetLastHeight error:", err)
		return err
	}
	if lastHeight == rpcHeight {
		log.Infof("[HandleBlock] No new blocks to fetch, height: %v", rpcHeight)
		return nil
	}
	// TODO 待完善初始height逻辑
	// init last height
	if lastHeight == 0 {
		lastHeight = rpcHeight - 1
	}
	// 存储从 lastHeight+1 到最新高度的所有区块
	for height := lastHeight + 1; height <= rpcHeight; height++ {
		block, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			log.Error("[HandleBlock] BlockDetailsByBlockId error:", err)
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, block.Result); err != nil {
			log.Error("[HandleBlock] InsertBlockDetails error:", err)
			return err
		}
		if _, err = es.UpdateLastHeight(client, ctx, height); err != nil {
			log.Error("[HandleBlock] UpdateLastHeight error:", err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
