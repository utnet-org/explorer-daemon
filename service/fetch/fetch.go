package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitFetchData() {
	// 定时执行RPC请求
	//ticker := time.NewTicker(time.Hour) // 例如，每小时执行一次
	//for range ticker.C {
	//BlockDetailsByFinal()
	//BlockChangesRpc()
	//HandleNetworkInfo()
	//HandleChipQuery()
	//}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := fetchAndStoreBlocks()
		if err != nil {
			log.Error("[InitFetchData] Error fetching or storing blocks: ", err)
		}
	}
}

func fetchAndStoreBlocks() error {
	ctx, client := es.GetESInstance()
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		return err
	}
	rpcHeight := res.Result.Header.Height
	lastHeight, err := es.GetLastHeight(client, ctx)
	if err != nil {
		log.Errorln("[fetchAndStoreBlocks] GetLastHeight error:", err)
		return err
	}
	if lastHeight == rpcHeight {
		log.Info("[fetchAndStoreBlocks] No new blocks to fetch,height:", rpcHeight)
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
			log.Error("[fetchAndStoreBlocks] BlockDetailsByBlockId error:", err)
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, block.Result); err != nil {
			log.Error("[fetchAndStoreBlocks] InsertBlockDetails error:", err)
			return err
		}
		if _, err = es.UpdateLastHeight(client, ctx, height); err != nil {
			log.Error("[fetchAndStoreBlocks] UpdateLastHeight error:", err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
	//blocKRes, err := remote.BlockDetailsByFinal()
	//if err != nil {
	//	log.Error("[fetchAndStoreBlocks] BlockDetailsByFinal error:", err)
	//	return err
	//}
	//blocKResHeight := blocKRes.Result.Header.Height
	return nil
}
