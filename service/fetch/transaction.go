package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
	"time"
)

// CompleteTxnDetails complete transaction details
func CompleteTxnDetails() error {
	//completeHeights := make([]int64, 0)
	//completeHeights = append(completeHeights, 74691)
	_, client := es.Init()
	// Get needful heights
	completeHeights := es.QueryTxnHeights(client)
	for _, height := range completeHeights {
		//for height := last.Height; height <= rpcHeight; height++ {
		currBlk, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			if err.Error() == "UNKNOWN_BLOCK" {
				log.Warningf("[CompleteTxnDetails] Continue UNKNOWN_BLOCK Height: %v", height)
				continue
			}
			log.Error("[CompleteTxnDetails] BlockDetailsByBlockId error:", err)
			return err
		}
		if currBlk == nil {
			log.Error("[CompleteTxnDetails] HandleBlock rpc res nil")
			return err
		}

		//err = HandleGasEveryHeight(height, err, currBlk)
		//if err != nil {
		//	return err
		//}
		blkHeader := currBlk.Result.Header
		// Get chunk hash
		chunkHash := currBlk.Result.Chunks[0].ChunkHash
		// Get final block's chunk details
		if currBlk.Result.Header.Hash == "" {
			log.Errorf("[CompleteTxnDetails] Hash null")
			return err
		}
		err = HandleChunkDetailsByChunkId(chunkHash, blkHeader)
		if err != nil {
			log.Errorf("[CompleteTxnDetails] HandleChunkDetailsByChunkId error: %v", err)
			return err
		}
		time.Sleep(200 * time.Millisecond)
		log.Infof("[CompleteTxnDetails] success height: %v", height)
	}
	return nil
}
