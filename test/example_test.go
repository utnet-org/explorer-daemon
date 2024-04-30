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
		TotalPower:       pkg.FakeInt(10000, 50000),
		ActiveMiner:      pkg.FakeInt(1000, 5000),
		BlockReward:      pkg.FakeInt(10, 20),
		DayAveReward:     pkg.FakeFloat(0, 1, 2),
		DayProduction:    pkg.FakeIntStr(10000, 100000),
		DayMessages:      pkg.FakeInt(10000, 20000),
		TotalAccount:     pkg.FakeInt(5000, 10000),
		AveBlockInterval: pkg.FakeIntStr(10, 60),
	}
	fmt.Println("Info:", oi)
}
