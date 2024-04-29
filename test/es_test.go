package test

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/fetch"
	log "github.com/sirupsen/logrus"
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

func TestEsGetLastBlocks(t *testing.T) {
	es.Init()
	es.GetLastBlocks()
}

func TestEsGetLastBlock(t *testing.T) {
	es.Init()
	res, _ := es.GetLastBlock()
	pkg.PrintStruct(res)
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

func Test24NanoSecTotalAward(t *testing.T) {
	es.Init()
	sum := es.QueryBlockReward24h()
	log.Info("[Test24NanoSecTotalAward] sum:", sum)
}

func TestNetWorkInfo(t *testing.T) {
	es.Init()
	res, err := es.GetNetWorkInfo()
	if err != nil {
		return
	}
	pkg.PrintStruct(res)
}
