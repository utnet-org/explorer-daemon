package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	log "github.com/sirupsen/logrus"
)

// Update every height gas price
func HandleGasEveryHeight(height int64, err error, block *types.BlockDetailsRes) error {
	heightList := []int64{height}
	gasRes, err := remote.GasPriceByBlockHeight(heightList)
	if err != nil {
		log.Error("[HandleBlock] GasPriceByBlockHeight error:", err)
		return err
	}
	var price = ""
	if gasRes != nil {
		price = gasRes.GasPrice
	}
	result := types.GasStoreResult{
		Height:   height,
		Hash:     block.Result.Header.Hash,
		GasPrice: price,
	}
	ctx, client := es.GetESInstance()
	err = es.InsertGas(ctx, client, result)
	if err != nil {
		log.Error("[HandleBlock] InsertGas error:", err)
		return err
	}
	//err = model.CreateGasPrice(database.DB, height, block.Result.Header.Hash, price)
	//if err != nil {
	//	log.Error("[HandleBlock] CreateGas error:", err)
	//	return err
	//}
	return nil
}
