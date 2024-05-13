package es

import (
	"context"
	"encoding/json"
	"errors"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

func InsertMiner(ctx context.Context, client *elastic.Client, result types.AllMinersResult) error {
	result.Timestamp = time.Now().UnixNano()
	res, err := client.Index().
		Index("miner").
		BodyJson(result).
		Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("文档Id %s, 索引名 %s\n", res.Id, res.Index)
	return nil
}

func UpdateMiner(ctx context.Context, client *elastic.Client, result types.AllMinersResult) error {
	upsert := map[string]interface{}{"timestamp": time.Now().UnixNano(), "miners": result.Miners, "total_power": result.TotalPower}
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
	query := elastic.NewMatchAllQuery()
	sortInfo := elastic.NewFieldSort("timestamp").Desc()
	res, err := client.Search().
		Index("miner").
		Query(query).
		SortBy(sortInfo).
		Size(1).
		Do(ctx)
	if err != nil {
		log.Errorf("[QueryMiner] Query error: %v\n", err)
		return nil, err
	}
	var result types.AllMinersResult
	_ = json.Unmarshal(res.Hits.Hits[0].Source, &result)
	return &result, nil
}

func QueryMinerRange(ctx context.Context, client *elastic.Client, n int64) (power int64) {
	//start, end := pkg.TimeNanoRange(n, time.Now().UnixNano())
	start, end := pkg.TimeNanoRange(n)
	dateQuery := elastic.NewRangeQuery("timestamp").Gte(start).Lte(end)

	// 创建查询
	query := elastic.NewBoolQuery().Filter(dateQuery)

	// 执行查询
	searchResult, err := client.Search().Index("miner").Query(query).Do(ctx)
	if err != nil {
		fmt.Println("Error executing search: ", err)
		return
	}
	// 处理查询结果
	if searchResult.Hits.TotalHits.Value > 0 {
		fmt.Printf("Found a total of %d documents\n", searchResult.Hits.TotalHits)
		// 遍历每个文档并输出 power 字段的值
		for _, hit := range searchResult.Hits.Hits {
			var data map[string]interface{}
			err := json.Unmarshal(hit.Source, &data)
			if err != nil {
				fmt.Println("Error unmarshalling source: ", err)
				continue
			}
			power, ok := data["total_power"].(float64)
			if !ok {
				fmt.Println("Power field is not a float64")
				continue
			}
			fmt.Printf("Power: %f\n", power)
		}
	} else {
		fmt.Println("No documents found")
	}
	return power
}
