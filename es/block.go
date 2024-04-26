package es

import (
	"context"
	"encoding/json"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func InsertBlockDetailsBulk(ctx context.Context, client *elastic.Client, bb types.BlockDetailsBody) error {
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

func InsertBlockDetails(ctx context.Context, client *elastic.Client, body types.BlockDetailsBody) error {
	sBody := types.BlockDetailsStoreBody{
		Author:           body.Author,
		Chunks:           body.Chunks,
		Header:           body.Header,
		Hash:             body.Header.Hash,
		ChunkHash:        body.Chunks[0].ChunkHash,
		Height:           body.Header.Height,
		Timestamp:        body.Header.Timestamp,
		TimestampNanoSec: body.Header.TimestampNanosec,
		PrevHash:         body.Header.PrevHash,
		PrevHeight:       body.Header.PrevHeight,
		GasLimit:         body.Chunks[0].GasLimit,
		GasPrice:         body.Header.GasPrice,
	}
	// Ensure the index exists
	createIndexIfNotExists(ctx, client, "block")
	_, err := client.Index().
		Index("block").
		Id(sBody.Hash).
		BodyJson(sBody).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("InsertBlockDetails Success")
	return nil
}

func InsertBlockChanges(ctx context.Context, client *elastic.Client, bb types.BlockChangesBody) error {
	createIndexIfNotExists(ctx, client, "block_changes")
	fmt.Println("[InsertBlockChanges] BlockHash:", bb.BlockHash)
	_, err := client.Index().
		Index("block_changes").
		BodyJson(bb).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("block_changes insert success")
	return nil
}

func InsertChunkDetails(body types.ChunkDetailsBody, chunkHash string) error {
	ctx := ECTX
	client := ECLIENT
	createIndexIfNotExists(ctx, client, "chunk")
	startTime := time.Now()
	// chunkHash作为唯一doc id
	_, err := client.Index().Index("chunk").BodyJson(body).Id(chunkHash).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("chunk details insert success")
	duration := time.Since(startTime)
	fmt.Printf("elapsed time: %v\n", duration)
	return nil
}

func InsertLastHeight(height int64) error {
	// 确保索引存在
	createIndexIfNotExists(ECTX, ECLIENT, "last_height")
	// 插入
	_, err := ECLIENT.Index().
		Index("last_height").
		Id("latest").
		BodyJson(map[string]interface{}{"latest_height": height}).
		Do(ECTX)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("InsertLastHeight Success")
	return nil
}

func BlockQuery2() {
	client, ctx := Init()
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
func BlockDetailsQuery(queryValue string, queryType pkg.BlockQueryType) (*types.BlockDetailsBody, error) {
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
	var body types.BlockDetailsBody
	for _, hit := range searchResult.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &body)
	}
	pkg.PrintStruct(body)
	return &body, nil
}

func LastBlockQuery() (*[]types.LastBlockRes, error) {
	client := ECLIENT
	ctx := ECTX
	lastHeight, err := LastHeightQuery(ctx, client)
	if err != nil {
		fmt.Println("LastHeightQuery error:", err)
		return nil, err
	}
	// 创建一个范围查询，查询高度小于最新高度的前10个区块
	rangeQuery := elastic.NewRangeQuery("height").Lte(lastHeight)
	//rangeQuery := elastic.NewTermQuery("height", lastHeight)
	res, err := client.Search().
		Index("block").
		Query(rangeQuery).
		Sort("height", false).
		Size(10).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var blocks []types.LastBlockRes
	for index, hit := range res.Hits.Hits {
		var body types.LastBlockRes
		fmt.Printf("第 %d 条数据\n", index+1)
		_ = json.Unmarshal(hit.Source, &body)
		pkg.PrintStruct(body)
		changes, err := QueryFinalBlockChanges(body.Hash)
		if err != nil {
			return nil, err
		}
		body.Messages = len(changes.Changes)
		blocks = append(blocks, body)
	}
	fmt.Printf("共查询到 %d 条数据\n", res.TotalHits())
	return &blocks, nil
}

func LastHeightQuery(ctx context.Context, client *elastic.Client) (int, error) {
	// 查询最新 height
	latestHeightResult, err := client.Get().
		Index("last_height").
		Id("latest").
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	type LatestHeight struct {
		Height int `json:"latest_height"`
	}

	var latestHeight LatestHeight
	err = json.Unmarshal(latestHeightResult.Source, &latestHeight)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return latestHeight.Height, nil
}

func QueryFinalBlockChanges(hash string) (*types.BlockChangesBody, error) {
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
	var body types.BlockChangesBody
	for _, hit := range res.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &body)
	}
	pkg.PrintStruct(body)
	return &body, nil
}
