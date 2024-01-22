package types

// Genesis Config Response

type GenesisConfigRes struct {
	CommonRes CommonRes
	Body      GenesisConfigBody `json:"result"`
}

type GenesisConfigBody struct {
	AvgHiddenValidatorSeatsPerShard []int64       `json:"avgHiddenValidatorSeatsPerShard"`
	BlockProducerKickoutThreshold   int64         `json:"blockProducerKickoutThreshold"`
	ChainID                         string        `json:"chainId"`
	ChunkProducerKickoutThreshold   int64         `json:"chunkProducerKickoutThreshold"`
	DynamicResharding               bool          `json:"dynamicResharding"`
	EpochLength                     int64         `json:"epochLength"`
	FishermenThreshold              string        `json:"fishermenThreshold"`
	GasLimit                        int64         `json:"gasLimit"`
	GasPriceAdjustmentRate          []int64       `json:"gasPriceAdjustmentRate"`
	GenesisHeight                   int64         `json:"genesisHeight"`
	GenesisTime                     string        `json:"genesisTime"`
	MaxGasPrice                     string        `json:"maxGasPrice"`
	MaxInflationRate                []int64       `json:"maxInflationRate"`
	MinGasPrice                     string        `json:"minGasPrice"`
	MinimumStakeDivisor             int64         `json:"minimumStakeDivisor"`
	NumBlockProducerSeats           int64         `json:"numBlockProducerSeats"`
	NumBlockProducerSeatsPerShard   []int64       `json:"numBlockProducerSeatsPerShard"`
	NumBlocksPerYear                int64         `json:"numBlocksPerYear"`
	OnlineMaxThreshold              []int64       `json:"onlineMaxThreshold"`
	OnlineMinThreshold              []int64       `json:"onlineMinThreshold"`
	ProtocolRewardRate              []int64       `json:"protocolRewardRate"`
	ProtocolTreasuryAccount         string        `json:"protocolTreasuryAccount"`
	ProtocolUpgradeNumEpochs        int64         `json:"protocolUpgradeNumEpochs"`
	ProtocolUpgradeStakeThreshold   []int64       `json:"protocolUpgradeStakeThreshold"`
	ProtocolVersion                 int64         `json:"protocolVersion"`
	RuntimeConfig                   RuntimeConfig `json:"runtimeConfig"`
	TotalSupply                     string        `json:"totalSupply"`
	TransactionValidityPeriod       int64         `json:"transactionValidityPeriod"`
	Validators                      []Validator   `json:"validators"`
}

type RuntimeConfig struct {
	AccountCreationConfig AccountCreationConfig `json:"accountCreationConfig"`
	StorageAmountPerByte  string                `json:"storageAmountPerByte"`
	TransactionCosts      TransactionCosts      `json:"transactionCosts"`
	WASMConfig            WASMConfig            `json:"wasmConfig"`
}

type AccountCreationConfig struct {
	MinAllowedTopLevelAccountLength int64  `json:"minAllowedTopLevelAccountLength"`
	RegistrarAccountID              string `json:"registrarAccountId"`
}

type TransactionCosts struct {
	ActionCreationConfig              ActionCreationConfig        `json:"actionCreationConfig"`
	ActionReceiptCreationConfig       ActionReceiptCreationConfig `json:"actionReceiptCreationConfig"`
	BurntGasReward                    []int64                     `json:"burntGasReward"`
	DataReceiptCreationConfig         DataReceiptCreationConfig   `json:"dataReceiptCreationConfig"`
	PessimisticGasPriceInflationRatio []int64                     `json:"pessimisticGasPriceInflationRatio"`
	StorageUsageConfig                StorageUsageConfig          `json:"storageUsageConfig"`
}

type ActionCreationConfig struct {
	AddKeyCost                AddKeyCost                                  `json:"addKeyCost"`
	CreateAccountCost         CreateAccountCost                           `json:"createAccountCost"`
	DeleteAccountCost         DeleteAccountCost                           `json:"deleteAccountCost"`
	DeleteKeyCost             DeleteKeyCost                               `json:"deleteKeyCost"`
	DeployContractCost        DeployContractCost                          `json:"deployContractCost"`
	DeployContractCostPerByte DeployContractCostPerByte                   `json:"deployContractCostPerByte"`
	FunctionCallCost          ActionCreationConfigFunctionCallCost        `json:"functionCallCost"`
	FunctionCallCostPerByte   ActionCreationConfigFunctionCallCostPerByte `json:"functionCallCostPerByte"`
	StakeCost                 StakeCost                                   `json:"stakeCost"`
	TransferCost              TransferCost                                `json:"transferCost"`
}

type AddKeyCost struct {
	FullAccessCost          FullAccessCost                    `json:"fullAccessCost"`
	FunctionCallCost        AddKeyCostFunctionCallCost        `json:"functionCallCost"`
	FunctionCallCostPerByte AddKeyCostFunctionCallCostPerByte `json:"functionCallCostPerByte"`
}

type FullAccessCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type AddKeyCostFunctionCallCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type AddKeyCostFunctionCallCostPerByte struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type CreateAccountCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type DeleteAccountCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type DeleteKeyCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type DeployContractCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type DeployContractCostPerByte struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type ActionCreationConfigFunctionCallCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type ActionCreationConfigFunctionCallCostPerByte struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type StakeCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type TransferCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type ActionReceiptCreationConfig struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type DataReceiptCreationConfig struct {
	BaseCost    BaseCost    `json:"baseCost"`
	CostPerByte CostPerByte `json:"costPerByte"`
}

type BaseCost struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type CostPerByte struct {
	Execution  int64 `json:"execution"`
	SendNotSir int64 `json:"sendNotSir"`
	SendSir    int64 `json:"sendSir"`
}

type StorageUsageConfig struct {
	NumBytesAccount     int64 `json:"numBytesAccount"`
	NumExtraBytesRecord int64 `json:"numExtraBytesRecord"`
}

type WASMConfig struct {
	EXTCosts      EXTCosts    `json:"extCosts"`
	GrowMemCost   int64       `json:"growMemCost"`
	LimitConfig   LimitConfig `json:"limitConfig"`
	RegularOpCost int64       `json:"regularOpCost"`
}

type EXTCosts struct {
	Base                        int64 `json:"base"`
	ContractCompileBase         int64 `json:"contractCompileBase"`
	ContractCompileBytes        int64 `json:"contractCompileBytes"`
	Keccak256Base               int64 `json:"keccak256Base"`
	Keccak256Byte               int64 `json:"keccak256Byte"`
	Keccak512Base               int64 `json:"keccak512Base"`
	Keccak512Byte               int64 `json:"keccak512Byte"`
	LogBase                     int64 `json:"logBase"`
	LogByte                     int64 `json:"logByte"`
	PromiseAndBase              int64 `json:"promiseAndBase"`
	PromiseAndPerPromise        int64 `json:"promiseAndPerPromise"`
	PromiseReturn               int64 `json:"promiseReturn"`
	ReadMemoryBase              int64 `json:"readMemoryBase"`
	ReadMemoryByte              int64 `json:"readMemoryByte"`
	ReadRegisterBase            int64 `json:"readRegisterBase"`
	ReadRegisterByte            int64 `json:"readRegisterByte"`
	Sha256Base                  int64 `json:"sha256Base"`
	Sha256Byte                  int64 `json:"sha256Byte"`
	StorageHasKeyBase           int64 `json:"storageHasKeyBase"`
	StorageHasKeyByte           int64 `json:"storageHasKeyByte"`
	StorageIterCreateFromByte   int64 `json:"storageIterCreateFromByte"`
	StorageIterCreatePrefixBase int64 `json:"storageIterCreatePrefixBase"`
	StorageIterCreatePrefixByte int64 `json:"storageIterCreatePrefixByte"`
	StorageIterCreateRangeBase  int64 `json:"storageIterCreateRangeBase"`
	StorageIterCreateToByte     int64 `json:"storageIterCreateToByte"`
	StorageIterNextBase         int64 `json:"storageIterNextBase"`
	StorageIterNextKeyByte      int64 `json:"storageIterNextKeyByte"`
	StorageIterNextValueByte    int64 `json:"storageIterNextValueByte"`
	StorageReadBase             int64 `json:"storageReadBase"`
	StorageReadKeyByte          int64 `json:"storageReadKeyByte"`
	StorageReadValueByte        int64 `json:"storageReadValueByte"`
	StorageRemoveBase           int64 `json:"storageRemoveBase"`
	StorageRemoveKeyByte        int64 `json:"storageRemoveKeyByte"`
	StorageRemoveRetValueByte   int64 `json:"storageRemoveRetValueByte"`
	StorageWriteBase            int64 `json:"storageWriteBase"`
	StorageWriteEvictedByte     int64 `json:"storageWriteEvictedByte"`
	StorageWriteKeyByte         int64 `json:"storageWriteKeyByte"`
	StorageWriteValueByte       int64 `json:"storageWriteValueByte"`
	TouchingTrieNode            int64 `json:"touchingTrieNode"`
	Utf16DecodingBase           int64 `json:"utf16DecodingBase"`
	Utf16DecodingByte           int64 `json:"utf16DecodingByte"`
	Utf8DecodingBase            int64 `json:"utf8DecodingBase"`
	Utf8DecodingByte            int64 `json:"utf8DecodingByte"`
	ValidatorStakeBase          int64 `json:"validatorStakeBase"`
	ValidatorTotalStakeBase     int64 `json:"validatorTotalStakeBase"`
	WriteMemoryBase             int64 `json:"writeMemoryBase"`
	WriteMemoryByte             int64 `json:"writeMemoryByte"`
	WriteRegisterBase           int64 `json:"writeRegisterBase"`
	WriteRegisterByte           int64 `json:"writeRegisterByte"`
}

type LimitConfig struct {
	InitialMemoryPages               int64 `json:"initialMemoryPages"`
	MaxActionsPerReceipt             int64 `json:"maxActionsPerReceipt"`
	MaxArgumentsLength               int64 `json:"maxArgumentsLength"`
	MaxContractSize                  int64 `json:"maxContractSize"`
	MaxGasBurnt                      int64 `json:"maxGasBurnt"`
	MaxGasBurntView                  int64 `json:"maxGasBurntView"`
	MaxLengthMethodName              int64 `json:"maxLengthMethodName"`
	MaxLengthReturnedData            int64 `json:"maxLengthReturnedData"`
	MaxLengthStorageKey              int64 `json:"maxLengthStorageKey"`
	MaxLengthStorageValue            int64 `json:"maxLengthStorageValue"`
	MaxMemoryPages                   int64 `json:"maxMemoryPages"`
	MaxNumberBytesMethodNames        int64 `json:"maxNumberBytesMethodNames"`
	MaxNumberInputDataDependencies   int64 `json:"maxNumberInputDataDependencies"`
	MaxNumberLogs                    int64 `json:"maxNumberLogs"`
	MaxNumberRegisters               int64 `json:"maxNumberRegisters"`
	MaxPromisesPerFunctionCallAction int64 `json:"maxPromisesPerFunctionCallAction"`
	MaxRegisterSize                  int64 `json:"maxRegisterSize"`
	MaxStackHeight                   int64 `json:"maxStackHeight"`
	MaxTotalLogLength                int64 `json:"maxTotalLogLength"`
	MaxTotalPrepaidGas               int64 `json:"maxTotalPrepaidGas"`
	RegistersMemoryLimit             int64 `json:"registersMemoryLimit"`
}

type Validator struct {
	AccountID string `json:"accountId"`
	Amount    string `json:"amount"`
	PublicKey string `json:"publicKey"`
}

// Protocol Config Response

type ProtocolConfigRes struct {
	CommonRes CommonRes
	Result    ProtocolConfigBody `json:"result"`
}

type ProtocolConfigBody struct {
	AvgHiddenValidatorSeatsPerShard []int64       `json:"avgHiddenValidatorSeatsPerShard"`
	BlockProducerKickoutThreshold   int64         `json:"blockProducerKickoutThreshold"`
	ChainID                         string        `json:"chainId"`
	ChunkProducerKickoutThreshold   int64         `json:"chunkProducerKickoutThreshold"`
	DynamicResharding               bool          `json:"dynamicResharding"`
	EpochLength                     int64         `json:"epochLength"`
	FishermenThreshold              string        `json:"fishermenThreshold"`
	GasLimit                        int64         `json:"gasLimit"`
	GasPriceAdjustmentRate          []int64       `json:"gasPriceAdjustmentRate"`
	GenesisHeight                   int64         `json:"genesisHeight"`
	GenesisTime                     string        `json:"genesisTime"`
	MaxGasPrice                     string        `json:"maxGasPrice"`
	MaxInflationRate                []int64       `json:"maxInflationRate"`
	MinGasPrice                     string        `json:"minGasPrice"`
	MinimumStakeDivisor             int64         `json:"minimumStakeDivisor"`
	NumBlockProducerSeats           int64         `json:"numBlockProducerSeats"`
	NumBlockProducerSeatsPerShard   []int64       `json:"numBlockProducerSeatsPerShard"`
	NumBlocksPerYear                int64         `json:"numBlocksPerYear"`
	OnlineMaxThreshold              []int64       `json:"onlineMaxThreshold"`
	OnlineMinThreshold              []int64       `json:"onlineMinThreshold"`
	ProtocolRewardRate              []int64       `json:"protocolRewardRate"`
	ProtocolTreasuryAccount         string        `json:"protocolTreasuryAccount"`
	ProtocolUpgradeStakeThreshold   []int64       `json:"protocolUpgradeStakeThreshold"`
	ProtocolVersion                 int64         `json:"protocolVersion"`
	RuntimeConfig                   RuntimeConfig `json:"runtimeConfig"`
	TransactionValidityPeriod       int64         `json:"transactionValidityPeriod"`
}
