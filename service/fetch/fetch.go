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
			log.Error("Error fetching or storing blocks:", err)
		}
	}
}

func fetchAndStoreBlocks() error {
	client, ctx := es.GetESInstance()

	// 获取当前最新区块数据
	res, err := remote.BlockDetailsByFinal()

	latestBlock := res.Result.Header
	if err != nil {
		return err
	}

	// 从 ES 获取 last_height
	lastHeight, err := es.GetLastHeight(client, ctx)
	if err != nil {
		return err
	}
	if latestBlock.Height == lastHeight {
		log.Info("[fetchAndStoreBlocks] No new blocks to fetch,height:", latestBlock.Height)
		return nil
	}

	// 存储从 lastHeight+1 到 latestBlock.Height 的所有区块
	for height := lastHeight + 1; height <= latestBlock.Height; height++ {
		block, err := remote.BlockDetailsByBlockId(height)
		if err != nil {
			log.Error("[fetchAndStoreBlocks] BlockDetailsByBlockId error:", err)
			return err
		}
		if err = es.InsertBlockDetails(ctx, client, block.Result); err != nil {
			log.Error("[fetchAndStoreBlocks] InsertBlockDetails error:", err)
			return err
		}
	}
	// 查询最新 height
	blocKRes, err := remote.BlockDetailsByFinal()
	if err != nil {
		log.Error("[fetchAndStoreBlocks] BlockDetailsByFinal error:", err)
		return err
	}
	blocKResHeight := blocKRes.Result.Header.Height
	// 更新 last_height
	if _, err = es.UpdateLastHeight(client, ctx, blocKResHeight); err != nil {
		log.Error("[fetchAndStoreBlocks] UpdateLastHeight error:", err)
		return err
	}
	return nil
}
