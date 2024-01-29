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

func TestEsQuery(t *testing.T) {
	es.BlockQuery()
}

func TestEsQuery2(t *testing.T) {
	es.BlockQuery2()
}
