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

func InsertBlockDetails(ctx context.Context, client *elastic.Client, bb types.BlockDetailsBody) error {
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

func InsertBlockChanges(ctx context.Context, client *elastic.Client, bb types.BlockChangesBody) error {
	createIndexIfNotExists(ctx, client, "block_change")
	startTime := time.Now()
	// insert
	bulkRequest := client.Bulk()
	blk := bb

	req := elastic.NewBulkIndexRequest().Index("block_change").Doc(blk)
	bulkRequest = bulkRequest.Add(req)
	// 每 1 条文档执行一次批量插入
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("block_changes insert success")
	duration := time.Since(startTime)
	fmt.Printf("elapsed time: %v\n", duration)
	return nil
}

func InsertChunkDetails(ctx context.Context, client *elastic.Client, bb types.ChunkDetailsBody) error {
	createIndexIfNotExists(ctx, client, "chunk")
	startTime := time.Now()
	// insert
	bulkRequest := client.Bulk()
	blk := bb

	req := elastic.NewBulkIndexRequest().Index("chunk").Doc(blk)
	bulkRequest = bulkRequest.Add(req)
	// 每 1 条文档执行一次批量插入
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("chunk details insert success")
	duration := time.Since(startTime)
	fmt.Printf("elapsed time: %v\n", duration)
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

func BlockDetailsQuery() types.BlockDetailsBody {
	//client, ctx := Init()
	// 定义要查询的用户
	userToQuery := "root"

	// 构建一个 term 查询
	termQuery := elastic.NewTermQuery("author", userToQuery)
	// 开始计时
	//startTime := time.Now()
	// 执行搜索
	searchResult, err := ECLIENT.Search().
		Index("block").
		Query(termQuery).
		//Sort("created", true). // 根据创建时间排序
		//From(0).Size(10).      // 分页参数
		Pretty(true).
		Do(ECTX)
	if err != nil {
		fmt.Println(err)
	}

	// 打印搜索结果
	fmt.Printf("查询到 %d 条数据\n", searchResult.TotalHits())
	var body types.BlockDetailsBody
	for _, hit := range searchResult.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &body)
	}
	pkg.PrintStruct(body)
	return body
}
