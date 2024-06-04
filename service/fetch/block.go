package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func HandleBlockChanges(rpcType pkg.BlockChangeRpcType, header types.BlockDetailsHeader) error {
	var queryValue interface{}
	switch rpcType {
	case pkg.BlockChangeRpcFinal:
		queryValue = "final"
	case pkg.BlockChangeRpcHeight:
		queryValue = header.Height
	case pkg.BlockChangeRpcHash:
		queryValue = header.Hash
	default:
		queryValue = "final"
	}
	res, err := remote.ChangesInBlock(rpcType, queryValue)
	if err != nil {
		log.Errorf("[BlockChangesRpc] ChangesInBlock rpc error: %v", err)
		return err
	}
	res.Height = header.Height
	res.Timestamp = header.Timestamp
	res.TimestampNanosec = header.TimestampNanosec
	ctx, client := es.GetESInstance()
	err = es.InsertBlockChanges(ctx, client, res)
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
	if len(res.Receipts) != 0 {
		for i, rec := range res.Receipts {
			res.Receipts[i].Receipt.Action.Actions = convertActions(rec.Receipt.Action.Actions)
		}
	}
	if len(res.Transactions) != 0 {
		for i, txn := range res.Transactions {
			txnHash := txn.Hash
			senderId := txn.SignerID
			// Remote txns data
			tRes, err := remote.TxnStatus(txnHash, senderId, "")
			if err != nil {
				log.Errorf("[HandleChunkDetailsByChunkId] Remote TransactionStatus error: %v", err)
				return err
			}
			tRes.Transaction.Actions = convertActions(tRes.Transaction.Actions)
			res.Transactions[i].Actions = convertActions(txn.Actions)
			result := types.TxnStoreResult{
				Height:          res.Header.HeightCreated,
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
	err = es.InsertChunkDetails(ctx, client, res, blkHash, ts)
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
	// TODO 待完善初始Height逻辑
	// init last height
	//if last.Height == 0 {
	//	last.Height = rpcHeight - 1
	//}
	for height := last.Height; height <= rpcHeight; height++ {
		log.Infof("[HandleBlock] Start Height: %v, RpcHeight: %v", height, rpcHeight)
		currBlk, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			if err.Error() == "UNKNOWN_BLOCK" {
				log.Warningf("[HandleBlock] Continue UNKNOWN_BLOCK Height: %v", height)
				continue
			}
			log.Error("[HandleBlock] BlockDetailsByBlockId error:", err)
			return err
		}
		if currBlk == nil {
			log.Error("[HandleBlock] HandleBlock rpc res nil")
			return err
		}

		err = HandleGasEveryHeight(height, err, currBlk)
		if err != nil {
			return err
		}
		blkHeader := currBlk.Result.Header
		blkHash := currBlk.Result.Header.Hash
		// Get chunk hash
		chunkHash := currBlk.Result.Chunks[0].ChunkHash
		// 获取最新block中的chunk details
		if currBlk.Result.Header.Hash == "" {
			log.Errorf("[HandleBlock] Hash null")
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, currBlk.Result); err != nil {
			log.Error("[HandleBlock] InsertBlockDetails error:", err)
			return err
		}
		err = HandleChunkDetailsByChunkId(chunkHash, blkHeader)
		if err != nil {
			log.Errorf("[HandleBlock] HandleChunkDetailsByChunkId error: %v", err)
			return err
		}
		// Handle current height block detail
		//err = HandleBlockChanges(pkg.BlockChangeRpcHeight, res.Result.Header)
		err = HandleBlockChanges(pkg.BlockChangeRpcHeight, currBlk.Result.Header)
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

func convertActions(actions []interface{}) []interface{} {
	var convertedActions []interface{}
	for _, action := range actions {
		switch v := action.(type) {
		case string:
			convertedActions = append(convertedActions, map[string]interface{}{v: v})
		default:
			convertedActions = append(convertedActions, v)
		}
	}
	return convertedActions
}
