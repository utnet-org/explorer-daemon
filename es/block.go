package es

import (
	"encoding/json"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
)

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

func BlockQuery() {
	client, ctx := Init()
	// 定义要查询的用户
	userToQuery := "root" // 假设我们要查询 user123 的推文

	// 构建一个 term 查询
	termQuery := elastic.NewTermQuery("author", userToQuery)
	// 开始计时
	//startTime := time.Now()
	// 执行搜索
	searchResult, err := client.Search().
		Index("block").
		Query(termQuery).
		//Sort("created", true). // 根据创建时间排序
		//From(0).Size(10).      // 分页参数
		Pretty(true).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// 打印搜索结果
	fmt.Printf("查询到 %d 条数据\n", searchResult.TotalHits())
	var blk types.BlockDetailsBody
	for _, hit := range searchResult.Hits.Hits {
		_ = json.Unmarshal(hit.Source, &blk)
		//fmt.Printf("用户: %s, 推文: %s, 转发数: %d\n", wb.User, wb.Message, wb.Retweets)
		// 可以进一步解析 hit.Source 以获取完整的推文数据
		//fmt.Printf("推文 ID: %s\n", hit.Id)
	}
	pkg.PrintStruct(blk)
}
