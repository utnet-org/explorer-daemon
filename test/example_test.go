package test

import (
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("test name is kobe")
}

func TestOverview(t *testing.T) {
	oi := types.OverviewInfoRes{
		Height:           pkg.FakeInt(0, 100000),
		LatestBlock:      pkg.FakeIntStr(10, 120),
		TotalPower:       pkg.FakeIntStr(10000, 50000),
		ActiveMiner:      pkg.FakeIntStr(1000, 5000),
		BlockReward:      pkg.FakeFloat(0, 1, 3),
		DayAveReward:     pkg.FakeFloat(0, 1, 2),
		DayProduction:    pkg.FakeIntStr(10000, 100000),
		DayMessage:       pkg.FakeIntStr(10000, 20000),
		TotalAccount:     pkg.FakeIntStr(5000, 10000),
		AveBlockInterval: pkg.FakeIntStr(10, 60),
	}
	fmt.Println("Info:", oi)
}
