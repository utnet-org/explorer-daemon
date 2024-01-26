package test

import (
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"testing"
)

func TestRPC(t *testing.T) {
	remote.Experimental()
}

func TestHttp(t *testing.T) {
	remote.ExperimentalHttp()
}

func TestBlockDetailsByFinal(t *testing.T) {
	res, _ := remote.BlockDetailsByFinal()
	pkg.PrintStruct(res.Body.Chunks)
	//fmt.Println("BlockDetailsByFinal bdRes:", res)
}

func TestBlockDetailsByBlockId(t *testing.T) {
	remote.BlockDetailsByBlockId(17821130)
}
func TestBlockDetailsByBlockHash(t *testing.T) {
	remote.BlockDetailsByBlockHash("81k9ked5s34zh13EjJt26mxw5npa485SY4UNoPi6yYLo")
}

func TestGasPriceByBlockHeight(t *testing.T) {
	heights := []int{1}
	remote.GasPriceByBlockHeight(heights)
}

func TestGasPriceByBlockHash(t *testing.T) {
	hashs := []string{"AXa8CHDQSA8RdFCt12rtpFraVq4fDUgJbLPxwbaZcZrj"}
	remote.GasPriceByBlockHash(hashs)
}

func TestGasPriceByNull(t *testing.T) {
	remote.GasPriceByNull()
}

// Protocol

func TestGenesisConfig(t *testing.T) {
	remote.GenesisConfig()
}

func TestProtocolConfigByFinal(t *testing.T) {
	remote.ProtocolConfigByFinal()
}

func TestProtocolConfigByBlockId(t *testing.T) {
	remote.ProtocolConfigByBlockId(1)
}

// Network

func TestNetworkNodeStatus(t *testing.T) {
	remote.NetworkNodeStatus()
}

func TestNetworkInfo(t *testing.T) {
	remote.NetworkInfo()
}

func TestNetworkValidationStatusByBlockNumber(t *testing.T) {
	remote.ValidationStatusByBlockNumber(17791098)
}

func TestNetworkValidationStatusByNull(t *testing.T) {
	remote.ValidationStatusByNull()
}

// Transaction

func TestTransactionSendAsync(t *testing.T) {
	remote.TransactionSendAsync()
}

func TestTransactionSendAwait(t *testing.T) {
	remote.TransactionSendAwait()
}

func TestTransactionStatus(t *testing.T) {
	remote.TransactionStatus()
}

func TestTransactionStatusReceipts(t *testing.T) {
	remote.TransactionStatusReceipts()
}

func TestTransactionReceiptById(t *testing.T) {
	remote.TransactionReceiptsById("2EbembRPJhREPtmHCrGv3Xtdm3xoc5BMVYHm3b2kjvMY")
}
