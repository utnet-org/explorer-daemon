package model

import "gorm.io/gorm"

type GasPrice struct {
	gorm.Model
	Height    int64 `gorm:"uniqueIndex"`
	BlockHash string
	Price     string
}

func CreateGasPrice(db *gorm.DB, height int64, blockHash, price string) error {
	gas := &GasPrice{
		Height:    height,
		BlockHash: blockHash,
		Price:     price,
	}
	if err := db.FirstOrCreate(gas).Error; err != nil {
		return err
	}
	return nil
}

// query gas price by height
func QueryGasPriceByHeight(db *gorm.DB, height int64) (string, error) {
	gas := &GasPrice{}
	if err := db.Where("height = ?", height).First(gas).Error; err != nil {
		return "", err
	}
	return gas.Price, nil
}
