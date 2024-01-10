package test

import (
	"explorer-daemon/pkg"
	"explorer-daemon/routing/types"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("test name is kobe")
}

func TestOverview(t *testing.T) {
	oi := types.OverviewInfoRes{
		Height:           pkg.FakeInt(0, 100000),
		LatestBlock:      pkg.FakeInt(10, 120),
		TotalPower:       pkg.FakeInt(10000, 50000),
		ActiveMiner:      pkg.FakeInt(1000, 5000),
		BlockReward:      pkg.FakeFloat(0, 1, 3),
		DayAveReward:     pkg.FakeFloat(0, 1, 2),
		DayProduction:    pkg.FakeInt(10000, 100000),
		DayMessage:       pkg.FakeInt(10000, 20000),
		TotalAccount:     pkg.FakeInt(5000, 10000),
		AveBlockInterval: pkg.FakeInt(10, 60),
	}
	fmt.Println("Info:", oi)
}
