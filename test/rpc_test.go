package test

import (
	"explorer-daemon/pkg"
	"explorer-daemon/service/fetch"
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
	pkg.PrintStruct(res.Result.Chunks)
	//fmt.Println("BlockDetailsByFinal bdRes:", res)
}

func TestBlockDetailsByBlockId(t *testing.T) {
	remote.BlockDetailsByBlockId(17821130)
}
func TestBlockDetailsByBlockHash(t *testing.T) {
	remote.BlockDetailsByBlockId("81k9ked5s34zh13EjJt26mxw5npa485SY4UNoPi6yYLo")
}

func TestChangeInBlockByFinal(t *testing.T) {
	//res, _ := remote.ChangeInBlockByFinal()
	//pkg.PrintStruct(res.Result)
}

func TestGasPriceByBlockHeight(t *testing.T) {
	heights := []int64{1}
	_, _ = remote.GasPriceByBlockHeight(heights)
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

func TestTransactionReceiptById(t *testing.T) {
	remote.TransactionReceiptsById("2EbembRPJhREPtmHCrGv3Xtdm3xoc5BMVYHm3b2kjvMY")
}

func TestContract(t *testing.T) {
	fetch.HandleContract()
}
