package test

import (
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"testing"
	"time"
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
		DayProduction:    pkg.FakeInt(10000, 100000),
		DayMessages:      pkg.FakeInt(10000, 20000),
		TotalAccount:     pkg.FakeInt(5000, 10000),
		AveBlockInterval: pkg.FakeIntStr(10, 60),
	}
	fmt.Println("Info:", oi)
}

func TestNano(t *testing.T) {
	currentTime := time.Now().UnixNano()
	oneDay := int64(24 * time.Hour)
	sevenDays := oneDay * 7
	startTime := currentTime - sevenDays
	fmt.Printf("Start Time: %d, End Time: %d\n", startTime, currentTime)
}

func TestLen(t *testing.T) {
	str := "441d580e88c60f78b1a08f32be137f58198d1522a3d323c383715447e29e1e5c"
	fmt.Println(len(str))
}
