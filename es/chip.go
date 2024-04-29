package es

import (
	"explorer-daemon/types"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertChip(result types.ChipQueryResult) error {
	//var cs = make([]types.Chip, 0)
	//cs = append(cs, types.Chip{
	//	MinerId:   "1",
	//	Power:     1,
	//	BusId:     "123",
	//	PublicKey: "456",
	//	ChipSN:    "789",
	//	P2Key:     "123",
	//})
	//result.Chip = cs
	createIndexIfNotExists(ECTX, ECLIENT, "chip")
	_, err := ECLIENT.Index().
		Index("chip").
		BodyJson(&result).
		Do(ECTX)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("InsertChip Success")
	return nil
}

func QueryChipsPower() (sum int64) {
	client, ctx := GetESInstance()
	// 定义嵌套聚合
	nestedAgg := elastic.NewNestedAggregation().Path("chips")
	sumAgg := elastic.NewSumAggregation().Field("chips.power")
	nestedAgg.SubAggregation("sum_power", sumAgg)

	// 执行查询
	searchResult, err := client.Search().
		Index("chip").
		Aggregation("chips_power", nestedAgg).
		Size(0).
		Do(ctx)
	if err != nil {
		log.Error("Error performing aggregation: %s", err)
	}
	// 解析聚合结果
	if agg, found := searchResult.Aggregations.Nested("chips_power"); found {
		if sum, found := agg.Aggregations.Sum("sum_power"); found && sum.Value != nil {
			log.Debug("[QueryChipsPower] Total power: %v\n", *sum.Value)
		} else {
			log.Warn("No sum calculated")
		}
	} else {
		log.Warn("No aggregation found")
	}
	return sum
}
