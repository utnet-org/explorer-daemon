package test

import (
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
	remote.BlockDetailsByFinal()
}
func TestBlockDetailsByBlockId(t *testing.T) {
	remote.BlockDetailsByBlockId(17821130)
}
func TestBlockDetailsByBlockHash(t *testing.T) {
	remote.BlockDetailsByBlockHash("81k9ked5s34zh13EjJt26mxw5npa485SY4UNoPi6yYLo")
}
