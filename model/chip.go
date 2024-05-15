package model

import "gorm.io/gorm"

// Chip struct
type Chip struct {
	gorm.Model
	SearchKey    string // 搜索key
	ChipType     string // 1684, 1684x, 1686
	Power        int64  //  power
	SerialNumber string //  序列号
	BusId        string //  bus id
	P2Key        string //
	PubKey       string //
	Flag         string // 启用标志(1-启用 0-失效)

}

// GetChipBySerialBus 通过 serial bus 寻找芯片
func GetChipBySerialBus(db *gorm.DB, serialNumber string, busId string) (*Chip, error) {
	u := Chip{}
	if err := db.Model(&u).Where("serial_number = ? AND bus_id = ?", serialNumber, busId).Take(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// GetChipBySearchKey 通过 searchKey 寻找芯片
func GetChipBySearchKey(db *gorm.DB, searchKey string) (chips []Chip, err error) {
	query := db.Model(Chip{})
	if err := query.Where("search_key = ?", searchKey).Find(&chips).Error; err != nil {
		return nil, err
	}
	return chips, nil
}

// InsertNewChip 新增芯片
func (u *Chip) InsertNewChip(db *gorm.DB) (uint, error) {
	if err := db.Create(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}
