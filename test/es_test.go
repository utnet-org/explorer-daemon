package test

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/fetch"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestEsInsert(t *testing.T) {
	es.Init()
	fetch.InitChainData()
}

func TestEsBlockDetailsQuery(t *testing.T) {
	es.Init()
	//es.BlockDetailsQuery()
}

func TestEsGetLastBlocks(t *testing.T) {
	es.Init()
	es.GetLastBlocksByLastHeight()
}

func TestEsGetLastBlock(t *testing.T) {
	es.Init()
	ctx, client := es.Init()
	res, _ := es.GetLastBlockByLastHeight(ctx, client)
	pkg.PrintStruct(res)
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

func Test24NanoSecTotalChanges(t *testing.T) {
	es.Init()
	sum := es.QueryBlockChangeMsg24h()
	log.Info("[Test24NanoSecTotalChanges] sum=", sum)
}

func TestChipsPower(t *testing.T) {
	ctx, client := es.Init()
	sum := es.QueryChipsPower(ctx, client)
	log.Info("[TestChipsPower] sum=", sum)
}

func TestNetWorkInfo(t *testing.T) {
	ctx, client := es.Init()
	res, err := es.GetNetworkInfo(ctx, client)
	if err != nil {
		return
	}
	pkg.PrintStruct(res)
}

func TestCheckHeight(t *testing.T) {
	ctx, client := es.Init()
	index := "chunk"
	res, err := es.CheckMissingHeights(ctx, client, index, 0, 250)
	if err != nil {
		return
	}
	fmt.Println(res)
	fmt.Println(len(res))
}

func TestChunkDetail(t *testing.T) {
	ctx, client := es.Init()
	res, err := es.QueryChunkDetails(ctx, client, pkg.ChunkQueryBlockHeight, 74691)
	if err != nil {
		return
	}
	pkg.PrintStruct(res)
}
