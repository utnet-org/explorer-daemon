package es

import (
	"context"
	"encoding/json"
	"errors"
	"explorer-daemon/types"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertMiner(ctx context.Context, client *elastic.Client, result types.AllMinersResult) error {
	createIndexIfNotExists(ctx, client, "miner")
	_, err := client.Index().
		Index("miner").
		BodyJson(result).
		Id("miner").
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMiner(ctx context.Context, client *elastic.Client, result types.AllMinersResult) error {
	upsert := map[string]interface{}{"miners": result.Miners, "total_power": result.TotalPower}
	latestHeightResult, err := client.Update().
		Index("miner").
		Id("miner").
		Doc(upsert).
		Upsert(upsert).
		Do(ctx)
	if err != nil {
		log.Errorf("[UpdateMiners] update result error: %v", err)
		return err
	}
	// 强制刷新索引，确保最新的写入立即可见
	if _, err = client.Flush().Index("miner").Do(ctx); err != nil {
		return err
	}
	if latestHeightResult.Result != "updated" {
		log.Warningln("[UpdateMiners] update result:", latestHeightResult.Result)
		return errors.New("update miner failed")
	}
	log.Debug("[UpdateMiners] Update success")
	return nil
}

func QueryMiner(ctx context.Context, client *elastic.Client) (*types.AllMinersResult, error) {
	res, err := client.Get().
		Index("miner").
		Id("miner").
		Do(ctx)
	if err != nil {
		log.Errorf("[QueryMiner] Query error: %v\n", err)
		return nil, err
	}
	var result types.AllMinersResult
	_ = json.Unmarshal(res.Source, &result)
	return &result, nil
}
