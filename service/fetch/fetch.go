package fetch

import (
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitChainData() {
	go TickerBlock()
	TickerOther()
}

func TickerBlock() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := HandleBlock(); err != nil {
			// some unknown block
			log.Errorf("[InitFetchData] HandleBlock error: %v", err)
		}
	}
}

func TickerOther() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := HandleMiner(); err != nil {
			log.Errorf("[InitFetchData] HandleMiner error: %v", err)
		}
		if err := HandleNetworkInfo(); err != nil {
			log.Error("[InitFetchData] HandleNetworkInfo error: ", err)
		}
		if err := HandleValidation(); err != nil {
			log.Error("[InitFetchData] HandleValidation error: ", err)
		}
		if err := HandleChipQuery(); err != nil {
			log.Error("[InitFetchData] HandleChipQuery error: ", err)
		}
	}
}

func InitCoinData() {
	remote.UpdatePrice()
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		remote.UpdatePrice()
	}
}
