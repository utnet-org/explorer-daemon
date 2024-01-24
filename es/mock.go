package es

import (
	"context"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
	"time"
)

func mockData(ctx context.Context, client *elastic.Client) {
	// 确保索引存在
	createIndexIfNotExists(ctx, client, "weibo")
	startTime := time.Now()
	// 模拟数据并批量插入
	bulkRequest := client.Bulk()
	for i := 0; i < 10; i++ { // 生成 10000 条数据
		blk := types.Weibo{
			User:     "user" + strconv.Itoa(i),
			Message:  "this is a weibo",
			Retweets: i,
			Image:    "",
			Created:  time.Time{},
			Tags:     nil,
			Location: "",
			Suggest:  nil,
		}

		req := elastic.NewBulkIndexRequest().Index("weibo").Doc(blk)
		bulkRequest = bulkRequest.Add(req)
		// 每 1000 条文档执行一次批量插入
		if i%100000 == 0 {
			_, err := bulkRequest.Do(ctx)
			if err != nil {
				log.Fatal(err)
			}
			bulkRequest = client.Bulk() // 创建新的批量请求
		}
	}

	// 执行最后一批插入（如果有剩余的文档）
	if bulkRequest.NumberOfActions() > 0 {
		_, err := bulkRequest.Do(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("批量插入完成")
	duration := time.Since(startTime)
	fmt.Printf("耗时: %v\n", duration)
}

func createIndexIfNotExists(ctx context.Context, client *elastic.Client, indexName string) {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(indexName).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
