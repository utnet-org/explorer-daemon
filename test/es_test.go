package test

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/fetch"
	"testing"
)

func TestEsInsert(t *testing.T) {
	es.Init()
	fetch.InitFetchData()
}

func TestEsBlockDetailsQuery(t *testing.T) {
	es.Init()
	//es.BlockDetailsQuery()
}

func TestEsLastBlockQuery(t *testing.T) {
	es.Init()
	es.LastBlockQuery()
}

func TestEsQuery2(t *testing.T) {
	es.BlockQuery2()
}

func TestBlockChanges(t *testing.T) {
	es.Init()
	res, err := es.QueryFinalBlockChanges("B2Cg2d1YipzztLmXDx8W9Zn6nENU1M8tEwKtvsorj3R9")
	if err != nil {
		return
	}
	pkg.PrintStruct(res)
}
