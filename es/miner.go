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
	"sort"
	"time"
)

func InsertMiner(ctx context.Context, client *elastic.Client, result types.AllMinersResult) error {
	result.Timestamp = time.Now().UnixMilli()
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

func QueryMinerRange(ctx context.Context, client *elastic.Client, n int64) []types.DailyPower {
	start, end := pkg.TimeNanoRange(n)
	dateQuery := elastic.NewRangeQuery("timestamp").Gte(start).Lte(end)

	// 创建查询
	query := elastic.NewBoolQuery().Filter(dateQuery)
	//powerList := make([]float64, 0)
	powerList := make([]types.DailyPower, 0)
	// 执行查询
	searchResult, err := client.Search().Index("miner").Query(query).Do(ctx)
	if err != nil {
		fmt.Println("Error executing search: ", err)
		return powerList
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
			ts, ok := data["timestamp"].(float64)
			if !ok {
				fmt.Println("Timestamp field is not a int64")
				continue
			}
			date := pkg.NanoTimestampToDate(int64(ts), "2006-01-02")
			item := types.DailyPower{
				Date:  date,
				Power: pkg.DivisionPowerOfTen(power, 12),
			}
			powerList = append(powerList, item)
		}
	} else {
		fmt.Println("No documents found")
	}
	return powerList
}

func QueryMinerDailySum(ctx context.Context, client *elastic.Client, n int) []types.DailyPower {
	index := "miner"
	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -n+1)
	var results []types.DailyPower
	// Aggregation query
	searchService := client.Search().
		Index(index).
		Query(elastic.NewRangeQuery("timestamp").Gte(sevenDaysAgo.UnixMilli()).Lte(now.UnixMilli())).
		Aggregation("by_day", elastic.NewDateHistogramAggregation().
			Field("timestamp").
			FixedInterval("1d").
			MinDocCount(0).
			Format("yyyy-MM-dd").
			SubAggregation("total_power", elastic.NewSumAggregation().Field("total_power")),
		)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Printf("Error executing the query: %s", err)
		return results
	}

	if agg, found := searchResult.Aggregations.DateHistogram("by_day"); found {
		for _, bucket := range agg.Buckets {
			date := pkg.MilliTimestampToDate(int64(bucket.Key), time.DateOnly)
			power, _ := bucket.Sum("total_power")
			results = append(results, types.DailyPower{
				Date:  date,
				Power: pkg.DivisionPowerOfTen(*power.Value, 12),
			})
		}
	}
	// Fill in missing days with power 0
	for i := 0; i < n; i++ {
		date := sevenDaysAgo.AddDate(0, 0, i).Format(time.DateOnly)
		if !containsDate(results, date) {
			results = append(results, types.DailyPower{
				Date:  date,
				Power: 0,
			})
		}
	}
	// Sort the results by date
	sort.Slice(results, func(i, j int) bool {
		return results[i].Date < results[j].Date
	})
	return results
}

func containsDate(results []types.DailyPower, date string) bool {
	for _, result := range results {
		if result.Date == date {
			return true
		}
	}
	return false
}

func QueryMinerMonthSum(ctx context.Context, client *elastic.Client, n int) []types.DailyPower {
	index := "miner"
	now := time.Now()
	startTime := now.AddDate(0, -n+1, 0) // n months ago

	// Convert times to milliseconds
	startTimeMillis := startTime.UnixMilli()
	endTimeMillis := now.UnixMilli()

	// Aggregation query
	searchService := client.Search().
		Index(index).
		Query(elastic.NewRangeQuery("timestamp").Gte(startTimeMillis).Lte(endTimeMillis)).
		Aggregation("by_month", elastic.NewDateHistogramAggregation().
			Field("timestamp").
			CalendarInterval("month").
			Format("yyyy-MM").
			MinDocCount(0).
			SubAggregation("total_power", elastic.NewSumAggregation().Field("total_power")),
		)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		log.Fatalf("Error executing the query: %s", err)
	}

	// Parse the results
	var results []types.DailyPower
	if agg, found := searchResult.Aggregations.DateHistogram("by_month"); found {
		for _, bucket := range agg.Buckets {
			date := pkg.MilliTimestampToDate(int64(bucket.Key), "2006-01")
			power, _ := bucket.Sum("total_power")
			results = append(results, types.DailyPower{
				Date:  date,
				Power: pkg.DivisionPowerOfTen(*power.Value, 12),
			})
		}
	}

	// Ensure we have 12 months
	ensure12Months(n, &results, startTime)
	return results
}

func ensure12Months(n int, results *[]types.DailyPower, startTime time.Time) {
	months := make(map[string]bool)
	for _, result := range *results {
		months[result.Date] = true
	}

	for i := 0; i < n; i++ {
		month := startTime.AddDate(0, i, 0).Format("2006-01")
		if !months[month] {
			*results = append(*results, types.DailyPower{
				Date:  month,
				Power: 0,
			})
		}
	}
	// Sort the results by month
	sort.Slice(*results, func(i, j int) bool {
		return (*results)[i].Date < (*results)[j].Date
	})
}
