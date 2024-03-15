package test

import (
	"explorer-daemon/es"
	"explorer-daemon/service/fetch"
	"testing"
)

func TestEsInsert(t *testing.T) {
	es.Init()
	fetch.InitFetchData()
}

func TestEsBlockDetailsQuery(t *testing.T) {
	es.Init()
	es.BlockDetailsQuery()
}

func TestEsLastBlockQuery(t *testing.T) {
	es.Init()
	es.LastBlockQuery()
}

func TestEsQuery2(t *testing.T) {
	es.BlockQuery2()
}
