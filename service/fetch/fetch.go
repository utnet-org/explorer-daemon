package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"fmt"
)

func InitFetchData() {
	// 定时执行RPC请求
	//ticker := time.NewTicker(time.Hour) // 例如，每小时执行一次
	//for range ticker.C {
	res, err := remote.BlockDetailsByFinal()
	if err != nil {
		fmt.Println("rpc error")
		//continue
	}
	// 将数据存入Elasticsearch
	err = es.InsertData(es.ECTX, es.ECLIENT, res.Body)
	pkg.PrintStruct(res.Body)
	if err != nil {
		fmt.Println("InsertData error:", err)
		// 处理存储到Elasticsearch的错误
		//continue
	}
	//}
}
