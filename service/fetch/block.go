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
	err = es.InsertLastHeight(ctx, client, res.Result.Header.Height, "")
	HandleBlockChanges(1, res.Result.Header)
	pkg.PrintStruct(res.Result)
	// 获取chunk hash
	cHash := res.Result.Chunks[0].ChunkHash
	// 获取最新block中的chunk details
	err = HandleChunkDetailsByChunkId(cHash)
	if err != nil {
		fmt.Println("InsertData error:", err)
	}
}

func HandleBlockChanges(rpcType pkg.BlockChangeRpcType, header types.BlockDetailsHeader) error {
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
		log.Errorf("[BlockChangesRpc] ChangesInBlock rpc error: %v", err)
		return err
	}
	res.Result.Height = header.Height
	res.Result.Timestamp = header.Timestamp
	res.Result.TimestampNanosec = header.TimestampNanosec
	ctx, client := es.GetESInstance()
	err = es.InsertBlockChanges(ctx, client, res.Result)
	if err != nil {
		log.Errorln("[BlockChangesRpc] InsertData error:", err)
		return err
	}
	log.Debugln("[BlockChangesRpc] HandleBlockChanges success")
	return nil
}

func HandleChunkDetailsByChunkId(chunkHash string) error {
	log.Infof("[HandleChunkDetailsByChunkId] chunkHash %v", chunkHash)
	// 获取最新block中的chunk details
	res, err := remote.ChunkDetailsByChunkId(chunkHash)
	if err != nil {
		log.Errorf("[HandleChunkDetailsByChunkId] ChunkDetailsByChunkId error: %v", err)
		return err
	}
	err = es.InsertChunkDetails(res.Result, chunkHash)
	if err != nil {
		log.Errorf("[HandleChunkDetailsByChunkId] InsertChunkDetails error: %v", err)
		return err
	}
	return nil
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
	last, err := es.GetLastHeightHash(client, ctx)
	if err != nil {
		log.Errorln("[HandleBlock] GetLastHeight error:", err)
		return err
	}
	if last.Height == rpcHeight {
		log.Infof("[HandleBlock] No new blocks to fetch, height: %v", rpcHeight)
		return nil
	}
	// TODO 待完善初始height逻辑
	// init last height
	if last.Height == 0 {
		last.Height = rpcHeight - 1
	}
	// 存储从 lastHeight+1 到最新高度的所有区块
	for height := last.Height + 1; height <= rpcHeight; height++ {
		block, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			log.Error("[HandleBlock] BlockDetailsByBlockId error:", err)
			return err
		}
		if block == nil {
			log.Error("[HandleBlock] HandleBlock rpc res nil")
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, block.Result); err != nil {
			log.Error("[HandleBlock] InsertBlockDetails error:", err)
			return err
		}
		if _, err = es.UpdateLastHeight(client, ctx, height, block.Result.Header.Hash); err != nil {
			log.Error("[HandleBlock] UpdateLastHeight error:", err)
			return err
		}
		// get chunk hash
		chunkHash := res.Result.Chunks[0].ChunkHash
		// 获取最新block中的chunk details
		err = HandleChunkDetailsByChunkId(chunkHash)
		if err != nil {
			log.Errorf("[HandleBlock] HandleChunkDetailsByChunkId error: %v", err)
			return err
		}
		err = HandleBlockChanges(2, res.Result.Header)
		if err != nil {
			log.Errorf("[HandleBlock] HandleBlockChanges error: %v", err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
