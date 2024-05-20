package es

import (
	"encoding/json"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var ECLIENT *elastic.Client
var ECTX context.Context

func Init() (context.Context, *elastic.Client) {
	client, err := elastic.NewClient(
		//es.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		//elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetURL("http://localhost:9200"),
		// 禁用嗅探器用于兼容内网ip
		elastic.SetSniff(false))
	//elastic.SetBasicAuth("user", "nvUt974rcNeg==*k0W3W"))
	if err != nil {
		log.Panicln("[ESInit] Elastic connect error:", err)
	}
	ctx := context.Background()
	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		log.Fatalf("[ESInit] Error pinging the Elasticsearch server: %s", err)
	}
	log.Printf("[ESInit] Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	ECTX = ctx
	ECLIENT = client
	CreateCheckIndex(ctx, client)
	log.Infoln("ES Connected")
	return ECTX, ECLIENT
	//mockData(ctx, client)
	//crud(client, ctx)
}

func GetESInstance() (context.Context, *elastic.Client) {
	if ECTX == nil || ECLIENT == nil {
		log.Fatalln("[GetESInstance] Failed")
	}
	return ECTX, ECLIENT
}

func CreateCheckIndex(ctx context.Context, client *elastic.Client) {
	createIndexIfNotExists(ctx, client, "block")
	createIndexIfNotExists(ctx, client, "chunk")
	createIndexIfNotExists(ctx, client, "block_changes")
	createIndexIfNotExists(ctx, client, "chip")
	createIndexIfNotExists(ctx, client, "network_info")
	createIndexIfNotExists(ctx, client, "last_height")
	createIndexIfNotExists(ctx, client, "miner")
	createIndexIfNotExists(ctx, client, "validator")
	createIndexIfNotExists(ctx, client, "transaction")
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

func crudDemo(client *elastic.Client, ctx context.Context) {
	// 索引mapping定义，这里仿微博消息结构定义
	const mapping = `
	{
	  "mappings": {
		"properties": {
		  "user": {
			"type": "keyword"
		  },
		  "message": {
			"type": "text"
		  },
		  "image": {
			"type": "keyword"
		  },
		  "created": {
			"type": "date"
		  },
		  "tags": {
			"type": "keyword"
		  },
		  "location": {
			"type": "geo_point"
		  },
		  "suggest_field": {
			"type": "completion"
		  }
		}
	  }
	}`
	// 首先检测下weibo索引是否存在
	exists, err := client.IndexExists("weibo").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
		//fmt.Println("weibo索引不存在")
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := client.CreateIndex("weibo").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	// 创建创建一条微博
	msg1 := types.Weibo{User: "olivere", Message: "打酱油的一天", Retweets: 0}
	// 使用client创建一个新的文档
	put1, err := client.Index().
		Index("weibo"). // 设置索引名称
		Id("1").        // 设置文档id
		BodyJson(msg1). // 指定前面声明的微博内容
		Do(ctx)         // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		panic(err)
	}
	log.Debugf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)

	// 根据id查询文档
	get1, err := client.Get().
		Index("weibo"). // 指定索引名
		Id("1").        // 设置文档id
		Do(ctx)         // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}
	// 手动将文档内容转换成go struct对象
	msg2 := types.Weibo{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	_ = json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2.Message)

	//根据文档id更新内容
	_, err = client.Update().
		Index("weibo").                             // 设置索引名
		Id("1").                                    // 文档id
		Doc(map[string]interface{}{"retweets": 0}). // 更新retweets=0，支持传入键值结构
		Do(ctx)                                     // 执行ES查询
	if err != nil {
		// Handle error
		panic(err)
	}

	// 根据id删除一条数据
	_, err = client.Delete().
		Index("weibo").
		Id("1").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
}
