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

func HandleChunkDetailsByChunkId(chunkHash string, header types.BlockDetailsHeader) error {
	blkHash := header.Hash
	ts := header.Timestamp
	res, err := remote.ChunkDetailsByChunkId(chunkHash)
	if err != nil {
		log.Errorf("[HandleChunkDetailsByChunkId] ChunkDetailsByChunkId error: %v", err)
		return err
	}
	ctx, client := es.GetESInstance()
	// Handle transactions not nil
	if len(res.Result.Transactions) != 0 {
		for _, txn := range res.Result.Transactions {
			txnHash := txn.Hash
			senderId := txn.SignerID
			// Remote txns data
			tRes, err := remote.TxnStatus(txnHash, senderId, "")
			if err != nil {
				log.Errorf("[HandleChunkDetailsByChunkId] Remote TransactionStatus error: %v", err)
				return err
			}
			result := types.TxnStoreResult{
				Height:          res.Result.Header.HeightCreated,
				Timestamp:       ts,
				TxnStatusResult: *tRes,
			}
			// Store Es data
			err = es.InsertTxnStatus(ctx, client, result)
			if err != nil {
				log.Errorf("[HandleChunkDetailsByChunkId] InsertTxnStatus error: %v", err)
				return err
			}
		}
	}
	err = es.InsertChunkDetails(ctx, client, res.Result, blkHash, ts)
	if err != nil {
		log.Errorf("[HandleChunkDetailsByChunkId] InsertChunkDetails error: %v", err)
		return err
	}
	log.Debugf("[HandleChunkDetailsByChunkId] Success Height: %v", header.Height)
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
		log.Errorln("[HandleBlock] BlockDetailsByFinal Error:", err)
		return err
	}
	rpcHeight := res.Result.Header.Height
	last, err := es.GetLastHeightHash(client, ctx)
	if err != nil {
		log.Errorln("[HandleBlock] GetLastHeight error:", err)
		return err
	}
	if last.Height == rpcHeight {
		log.Infof("[HandleBlock] No New Blocks, Height: %v, RpcHeight: %v", last.Height, rpcHeight)
		return nil
	}
	// TODO 待完善初始height逻辑
	// init last height
	//if last.Height == 0 {
	//	last.Height = rpcHeight - 1
	//}
	// 存储从 lastHeight+1 到最新高度的所有区块
	for height := last.Height; height <= rpcHeight; height++ {
		log.Infof("[HandleBlock] Start Height: %v, RpcHeight: %v", height, rpcHeight)
		block, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			if err.Error() == "UNKNOWN_BLOCK" {
				log.Warningf("[HandleBlock] Continue UNKNOWN_BLOCK Height: %v", height)
				continue
			}
			log.Error("[HandleBlock] BlockDetailsByBlockId error:", err)
			return err
		}
		if block == nil {
			log.Error("[HandleBlock] HandleBlock rpc res nil")
			return err
		}

		err = HandleGasEveryHeight(height, err, block)
		if err != nil {
			return err
		}
		blkHeader := block.Result.Header
		blkHash := block.Result.Header.Hash
		// Get chunk hash
		chunkHash := block.Result.Chunks[0].ChunkHash
		// 获取最新block中的chunk details
		if block.Result.Header.Hash == "" {
			log.Errorf("[HandleBlock] Hash null")
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, block.Result); err != nil {
			log.Error("[HandleBlock] InsertBlockDetails error:", err)
			return err
		}
		err = HandleChunkDetailsByChunkId(chunkHash, blkHeader)
		if err != nil {
			log.Errorf("[HandleBlock] HandleChunkDetailsByChunkId error: %v", err)
			return err
		}
		err = HandleBlockChanges(2, res.Result.Header)
		if err != nil {
			log.Errorf("[HandleBlock] HandleBlockChanges error: %v", err)
			return err
		}
		// Update the height after all operations have been processed
		if _, err = es.UpdateLastHeight(client, ctx, height, blkHash); err != nil {
			log.Error("[HandleBlock] UpdateLastHeight error:", err)
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}
