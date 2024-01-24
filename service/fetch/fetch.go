package fetch

import (
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"time"
)

func InitFetchData() {
	// 定时执行RPC请求
	ticker := time.NewTicker(time.Hour) // 例如，每小时执行一次
	for range ticker.C {
		bsRes := remote.BlockDetailsByFinal()

		// 将数据存入Elasticsearch
		err := IndexDataToElasticsearch(bsRes)
		if err != nil {
			// 处理存储到Elasticsearch的错误
			continue
		}
	}
}

func IndexDataToElasticsearch(data types.BlockDetailsRes) error {
	// 使用go-elasticsearch库将数据存入Elasticsearch
	// 你需要初始化es客户端，并执行相应的索引操作
	// 参考 go-elasticsearch 文档：https://pkg.go.dev/github.com/elastic/go-elasticsearch/v8
	return nil
}
