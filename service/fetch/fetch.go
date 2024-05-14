package fetch

import (
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
		if err := HandleBlock(); err != nil {
			// TODO 某些block没有数据
			log.Errorf("[InitFetchData] HandleBlock error: %v", err)
		}
		if err := HandleMiner(); err != nil {
			log.Errorf("[InitFetchData] HandleMiner error: %v", err)
		}
		if err := HandleNetworkInfo(); err != nil {
			log.Error("[InitFetchData] HandleNetworkInfo error: ", err)
		}
		if err := HandleChipQuery(); err != nil {
			log.Error("[InitFetchData] HandleChipQuery error: ", err)
		}
	}
}
