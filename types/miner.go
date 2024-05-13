package types

type AllMinersRes struct {
	CommonRes CommonRes
	Result    AllMinersResult `json:"result"`
}

type AllMinersResult struct {
	Miners     []Miner `json:"miners"`
	TotalPower int64   `json:"total_power"`
}

type Miner struct {
	AccountID                   string `json:"account_id"`
	Power                       string `json:"power"`
	PublicKey                   string `json:"public_key"`
	ValidatorPowerStructVersion string `json:"validator_power_struct_version"`
}
