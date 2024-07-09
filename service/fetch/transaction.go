package fetch

import (
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
	"time"
)

// CompleteTransactionDetails complete transaction details
func CompleteTransactionDetails() error {
	completeHeights := make([]int64, 0)
	completeHeights = append(completeHeights, 74691)
	//_, client := es.Init()
	// Get needful heights
	//completeHeights := es.QueryTxnHeights(client)
	for _, height := range completeHeights {
		//for height := last.Height; height <= rpcHeight; height++ {
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

		//err = HandleGasEveryHeight(height, err, currBlk)
		//if err != nil {
		//	return err
		//}
		blkHeader := currBlk.Result.Header
		// Get chunk hash
		chunkHash := currBlk.Result.Chunks[0].ChunkHash
		// Get final block's chunk details
		if currBlk.Result.Header.Hash == "" {
			log.Errorf("[HandleBlock] Hash null")
			return err
		}
		err = HandleChunkDetailsByChunkId(chunkHash, blkHeader)
		if err != nil {
			log.Errorf("[HandleBlock] HandleChunkDetailsByChunkId error: %v", err)
			return err
		}
		time.Sleep(200 * time.Millisecond)
		log.Infof("CompleteTransactionDetails height: %v", height)
	}
	return nil
}
