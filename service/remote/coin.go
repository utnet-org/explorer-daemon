package remote

import (
	"encoding/json"
	"explorer-daemon/types"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
)

const (
	contractAddr = "0x35Da89A339DE2c78F8FB1c5e1A9a9C6539e2FA8A"
	apiURL       = "https://api.coingecko.com/api/v3/simple/token_price/binance-smart-chain?contract_addresses=%s&vs_currencies=usd&include_24hr_change=true"
)

var (
	currentPrice  float64
	currentChange float64
	mu            sync.RWMutex
)

func getCoinPrice() (float64, float64, error) {
	url := fmt.Sprintf(apiURL, contractAddr)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var priceResponse map[string]types.CoinPrice
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return 0, 0, err
	}

	price, ok := priceResponse[strings.ToLower(contractAddr)]
	if !ok {
		return 0, 0, fmt.Errorf("[getCoinPrice] could not find price for contract address %s", contractAddr)
	}

	return price.USD, price.USD24hChange, nil
}

func UpdatePrice() {
	price, change, err := getCoinPrice()
	if err != nil {
		log.Errorln("[UpdatePrice] Error getting token price:", err)
	} else {
		mu.Lock()
		currentPrice = price
		currentChange = change
		mu.Unlock()
	}
}

func PriceHandler() (float64, float64) {
	mu.RLock()
	price := currentPrice
	change := currentChange
	mu.RUnlock()
	return price, change
}
