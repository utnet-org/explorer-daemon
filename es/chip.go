package es

import (
	"context"
	"explorer-daemon/types"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertChip(ctx context.Context, client *elastic.Client, result types.ChipQueryResult) error {
	createIndexIfNotExists(ctx, client, "chip")
	_, err := client.Index().
		Index("chip").
		Id("chip").
		BodyJson(&result).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Debugln("InsertChip Success")
	return nil
}

func QueryChipByHeight(ctx context.Context, client *elastic.Client, height int64) (*elastic.SearchResult, error) {
	existsQuery := elastic.NewTermQuery("block_height", height)
	searchResult, err := client.Search().
		Index("chip").
		Query(existsQuery).
		//Size(0). // check exist only
		Do(ctx)
	if err != nil {
		log.Errorf("[QueryChipByHeight] Error searching for height: %s", err)
		return nil, err
	}
	return searchResult, nil
}

func QueryChipsPower(ctx context.Context, client *elastic.Client) int64 {
	sumAgg := elastic.NewSumAggregation().Field("total_power")
	searchResult, err := client.Search().
		Index("chip").
		Aggregation("total_power_sum", sumAgg).
		Do(ctx)
	if err != nil {
		log.Error("[QueryChipsPower] Error performing aggregation: %s", err)
		return -1
	}
	agg, found := searchResult.Aggregations.Sum("total_power_sum")
	if !found {
		log.Errorln("[QueryChipsPower] Aggregation not found")
	}
	log.Debugf("[QueryChipsPower] Total Power Sum: %v\n", *agg.Value)
	sum := int64(*agg.Value)
	return sum
}
