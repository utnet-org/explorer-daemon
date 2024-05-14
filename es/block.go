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
	"strconv"
	"time"
)

func InsertBlockDetailsBulk(ctx context.Context, client *elastic.Client, bb types.BlockDetailsResult) error {
	// 确保索引存在
	createIndexIfNotExists(ctx, client, "block")
	startTime := time.Now()
	// 插入
	bulkRequest := client.Bulk()
	blk := bb
	req := elastic.NewBulkIndexRequest().Index("block").Doc(blk)
	bulkRequest = bulkRequest.Add(req)
	// 每 1 条文档执行一次批量插入
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("插入完成")
	duration := time.Since(startTime)
	fmt.Printf("耗时: %v\n", duration)
	return nil
}

func InsertBlockDetails(ctx context.Context, client *elastic.Client, body types.BlockDetailsResult) error {
	gasPrice, _ := strconv.ParseInt(body.Header.GasPrice, 10, 64)
	sBody := types.BlockDetailsStoreBody{
		Author:         body.Author,
		Chunks:         body.Chunks,
		Header:         body.Header,
		Hash:           body.Header.Hash,
		ChunkHash:      body.Chunks[0].ChunkHash,
		Height:         body.Header.Height,
		TimestampMilli: pkg.ConvertNanoToMilli(body.Header.Timestamp),
		Timestamp:      body.Header.Timestamp,
		//TimestampNanoSec: body.Header.TimestampNanosec,
		PrevHash:        body.Header.PrevHash,
		PrevHeight:      body.Header.PrevHeight,
		ValidatorReward: body.Header.ValidatorReward,
		GasLimit:        body.Chunks[0].GasLimit,
		GasPrice:        gasPrice,
		GasUsed:         body.Chunks[0].GasUsed,
	}
	createIndexIfNotExists(ctx, client, "block")
	_, err := client.Index().
		Index("block").
		BodyJson(sBody).
		Id(body.Header.Hash).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InsertBlockChanges(ctx context.Context, client *elastic.Client, res types.BlockChangesResult) error {
	//createIndexIfNotExists(ctx, client, "block_changes")
	_, err := client.Index().
		Index("block_changes").
		BodyJson(res).
		Id(res.BlockHash).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Debugln("[InsertBlockChanges] block_changes insert success")
	return nil
}

func InsertChunkDetails(body types.ChunkDetailsResult, chunkHash string) error {
	ctx := ECTX
	client := ECLIENT
	createIndexIfNotExists(ctx, client, "chunk")
	// chunkHash作为唯一doc id
	_, err := client.Index().
		Index("chunk").
		BodyJson(body).
		Id(chunkHash).
		Do(ctx)
	if err != nil {
		log.Errorf("[InsertChunkDetails] ES error: %v", err)
		return err
	}
	fmt.Println("chunk details insert success")
	return nil
}

func InsertLastHeight(ctx context.Context, client *elastic.Client, height int64, hash string) error {
	createIndexIfNotExists(ctx, client, "last_height")
	_, err := client.Index().
		Index("last_height").
		Id("latest").
		BodyJson(map[string]interface{}{"height": height, "hash": hash}).
		Do(ctx)
	if err != nil {
		log.Errorln("[InsertLastHeight] Create error: ", err)
		return err
	}
	log.Infoln("[InsertLastHeight] Success")
	return nil
}

func BlockQuery2() {
	ctx, client := Init()
	// 定义要查询的用户
	userToQuery := "user0" // 假设我们要查询 user123 的推文

	// 构建一个 term 查询
	termQuery := elastic.NewTermQuery("user", userToQuery)
	// 开始计时
	//startTime := time.Now()
	// 执行搜索
	searchResult, err := client.Search().
		Index("weibo").
		Query(termQuery).
		Sort("created", true). // 根据创建时间排序
		From(0).Size(10).      // 分页参数
		Pretty(true).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// 打印搜索结果
	fmt.Printf("查询到 %d 条数据\n", searchResult.TotalHits())
	var blk types.Weibo
	for _, hit := range searchResult.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &blk)
		//fmt.Printf("用户: %s, 推文: %s, 转发数: %d\n", wb.User, wb.Message, wb.Retweets)
		// 可以进一步解析 hit.Source 以获取完整的推文数据
		//fmt.Printf("推文 ID: %s\n", hit.Id)
	}
	pkg.PrintStruct(blk)
	// 结束计时
	//duration := time.Since(startTime)
	//fmt.Printf("耗时: %v\n", duration)
	//time.Sleep(100 * time.Millisecond)
}

// 查询Block详情
func GetBlockDetails(queryType pkg.BlockQueryType, queryValue string) (*types.BlockDetailsResult, error) {
	var queryName string
	switch queryType {
	case pkg.BlockQueryHeight:
		queryName = "height"
	case pkg.BlockQueryHash:
		queryName = "hash"
	default:
		queryName = ""
	}
	// 构建一个 term 查询
	termQuery := elastic.NewTermQuery(queryName, queryValue)
	// 执行搜索
	searchResult, err := ECLIENT.Search().
		Index("block").
		Query(termQuery).
		//Sort("created", true). // 根据创建时间排序
		//From(0).Size(10).      // 分页参数
		Pretty(true).
		Do(ECTX)
	if err != nil {
		fmt.Printf("[BlockDetailsQuery] error:%s", err)
		return nil, err
	}
	fmt.Printf("查询到 %d 条数据\n", searchResult.TotalHits())
	var body types.BlockDetailsResult
	for _, hit := range searchResult.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &body)
	}
	pkg.PrintStruct(body)
	return &body, nil
}

func GetLastBlock() (*types.BlockDetailsResult, error) {
	ctx, client := GetESInstance()
	last, err := GetLastHeightHash(client, ctx)
	if err != nil {
		log.Errorf("[GetLastBlock] GetLastHeight error: %v\n", err)
		return nil, err
	}
	query := elastic.NewTermQuery("height", last.Height)
	res, err := client.Search().
		Index("block").
		Query(query).
		Do(ctx)
	if err != nil {
		log.Errorf("[GetLastBlock] Query error: %v\n", err)
		return nil, err
	}
	var body types.BlockDetailsResult
	if res.TotalHits() == 1 {
		_ = json.Unmarshal(res.Hits.Hits[0].Source, &body)
	}
	return &body, nil
}

func GetLastBlocks() (*[]types.LastBlockRes, error) {
	ctx, client := GetESInstance()
	last, err := GetLastHeightHash(client, ctx)
	if err != nil {
		log.Errorf("[GetLastBlocks] Query error: %v\n", err)
		return nil, err
	}
	// 创建一个范围查询，查询高度小于最新高度的前10个区块
	rangeQuery := elastic.NewRangeQuery("height").Lte(last.Height)
	//rangeQuery := elastic.NewTermQuery("height", lastHeight)
	res, err := client.Search().
		Index("block").
		Query(rangeQuery).
		Sort("height", false).
		Size(10).
		Do(ctx)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	var blocks []types.LastBlockRes
	for _, hit := range res.Hits.Hits {
		var body types.LastBlockRes
		//fmt.Printf("第 %d 条数据\n", index+1)
		_ = json.Unmarshal(hit.Source, &body)
		pkg.PrintStruct(body)
		changes, err := QueryFinalBlockChanges(body.Hash)
		if err != nil {
			return nil, err
		}
		body.Messages = len(changes.Changes)
		blocks = append(blocks, body)
	}
	log.Debugf("[GetLastBlocks] total %d datas\n", res.TotalHits())
	return &blocks, nil
}

func GetLastHeightHash(client *elastic.Client, ctx context.Context) (*types.LastHeightHash, error) {
	// 查询最新 height
	latestHeightResult, err := client.Get().
		Index("last_height").
		Id("latest").
		Do(ctx)
	if err != nil {
		// 检查是否因为索引不存在而出错
		if elastic.IsNotFound(err) {
			log.Warningln("[GetLastHeight] Index not found, creating a new one...")
			if err := InsertLastHeight(ctx, client, 0, ""); err != nil {
				log.Warningf("[GetLastHeight] Error creating index: %v", err)
			}
		} else {
			log.Fatalf("[GetLastHeight] Error fetching or storing blocks: %v", err)
			return nil, err
		}
	}
	var latestHeight types.LastHeightHash
	err = json.Unmarshal(latestHeightResult.Source, &latestHeight)
	if err != nil {
		log.Errorf("[GetLastHeight] Error fetching or storing blocks: %v", err)
		return nil, err
	}
	return &latestHeight, nil
}

func UpdateLastHeight(client *elastic.Client, ctx context.Context, height int64, hash string) (int64, error) {
	// 定义在文档不存在时要插入的默认文档
	upsert := map[string]interface{}{"height": height, "hash": hash}
	latestHeightResult, err := client.Update().
		Index("last_height").
		Id("latest").
		Doc(upsert).
		Upsert(upsert).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	// 强制刷新索引，确保最新的写入立即可见
	_, err = client.Flush().Index("last_height").Do(ctx)
	if err != nil {
		return -1, err
	}
	if latestHeightResult.Result != "updated" {
		log.Warningln("[UpdateLastHeight] update result:", latestHeightResult.Result)
		return -1, errors.New("update last height failed")
	}
	log.Debug("[UpdateLastHeight] Update last height success height:", height)
	return height, nil
}

func QueryFinalBlockChanges(hash string) (*types.BlockChangesResult, error) {
	client := ECLIENT
	ctx := ECTX
	// Elasticsearch会默认分词，使用block_hash.keyword
	query := elastic.NewMatchQuery("block_hash.keyword", hash)
	res, err := client.Search().
		Index("block_changes").
		Query(query).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var body types.BlockChangesResult
	for _, hit := range res.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &body)
	}
	pkg.PrintStruct(body)
	return &body, nil
}

// 查询24小时所有区块的总产出
func QueryBlockReward24h() (sum int64) {
	// 计算24小时前的时间戳（纳秒）
	nanoSecAgo := pkg.TimeNanoSecAgo()
	// 创建一个范围查询
	rangeQuery := elastic.NewRangeQuery("timestamp_nanosec").Gte(nanoSecAgo)
	ctx, client := GetESInstance()
	// 创建一个求和聚合
	sumAgg := elastic.NewSumAggregation().Field("award")
	searchResult, err := client.Search().
		Index("block").
		Query(rangeQuery).
		Aggregation("total_award", sumAgg).
		Size(0).
		Do(ctx)
	if err != nil {
		log.Error("Error performing aggregation: %s", err)
	}
	// 解析聚合结果
	if agg, found := searchResult.Aggregations.Sum("total_award"); found && agg.Value != nil {
		log.Debugf("[QueryBlockReward24h] Total award: %v\n", *agg.Value)
		sum = int64(*agg.Value)
	} else {
		log.Debug("No aggregation found or sum is nil")
	}
	return sum
}

// 查询24小时消息数量
func QueryBlockChangeMsg24h() (sum int64) {
	ctx, client := GetESInstance()
	nanoSecAgo := pkg.TimeNanoSecAgo()
	rangeQuery := elastic.NewRangeQuery("timestamp_nanosec").Gte(nanoSecAgo)
	// 使用 Painless 脚本计算数组长度的总和
	script := "params._source.changes.size()"
	sumAgg := elastic.NewSumAggregation().Script(elastic.NewScript(script))
	searchResult, err := client.Search().
		Index("block_changes").
		Query(rangeQuery).
		Aggregation("total_changes", sumAgg).
		Size(0).
		Do(ctx)
	if err != nil {
		log.Error("Error performing aggregation: %s", err)
	}
	// 解析聚合结果
	if agg, found := searchResult.Aggregations.Sum("total_changes"); found && agg.Value != nil {
		log.Debugf("[QueryBlockChangeMsg24h] Total messages: %v\n", *agg.Value)
		sum = int64(*agg.Value)
	} else {
		log.Debug("No aggregation found or sum is nil")
	}
	return sum
}

func QueryGasRange(ctx context.Context, client *elastic.Client, n int64) []types.DailyGas {
	index := "block"
	now := time.Now().UTC()
	startTime := now.Add(-24 * time.Hour)

	// Convert times to nanoseconds
	startTimeNano := startTime.UnixMilli()
	endTimeNano := now.UnixMilli()

	// Aggregation query
	searchService := client.Search().
		Index(index).
		Query(elastic.NewRangeQuery("timestamp_milli").Gte(startTimeNano).Lte(endTimeNano)).
		Aggregation("by_4h", elastic.NewDateHistogramAggregation().
			Field("timestamp_milli").
			FixedInterval("4h").
			MinDocCount(0).
			Format("yyyy-MM-dd HH:mm:ss").
			//SubAggregation("total_gas", elastic.NewSumAggregation().Field("gas_price")),
			SubAggregation("total_gas", elastic.NewSumAggregation().Script(elastic.NewScript("doc['gas_used'].value * doc['gas_price'].value"))),
		)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		log.Fatalf("Error executing the query: %s", err)
	}

	// Parse the results
	var results []types.DailyGas
	if agg, found := searchResult.Aggregations.DateHistogram("by_4h"); found {
		for _, bucket := range agg.Buckets {
			date := pkg.MilliTimestampToDate(int64(bucket.Key), time.DateTime)
			gas, _ := bucket.Sum("total_gas")
			results = append(results, types.DailyGas{
				Date: date,
				Gas:  pkg.DivisionPowerOfTen(*gas.Value, 9),
			})
		}
	}

	// Fill in missing intervals with gas 0
	for i := 0; i < 6; i++ { // 24 hours / 4 hours per interval = 6 intervals
		expectedTime := startTime.Add(time.Duration(i) * 4 * time.Hour).Format(time.DateTime)
		if !containsTimestamp(results, expectedTime) {
			results = append(results, types.DailyGas{
				Date: expectedTime,
				Gas:  0,
			})
		}
	}

	// Sort the results by timestamp
	sort.Slice(results, func(i, j int) bool {
		return results[i].Date < results[j].Date
	})

	fmt.Println(results)
	return results
}

func containsTimestamp(results []types.DailyGas, timestamp string) bool {
	for _, result := range results {
		if result.Date == timestamp {
			return true
		}
	}
	return false
}
