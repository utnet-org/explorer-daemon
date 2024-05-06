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

func Init() (*elastic.Client, context.Context) {
	// 创建ES client用于后续操作ES
	//client, err := es.NewClient(
	//	// 设置ES服务地址，支持多个地址
	//	es.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
	//	// 设置基于http base auth验证的账号和密码
	//	es.SetBasicAuth("user", "secret"))
	//if err != nil {
	//	// Handle error
	//	fmt.Printf("连接失败: %v\n", err)
	//} else {
	//	fmt.Println("连接成功")
	//}

	// 创建client
	client, err := elastic.NewClient(
		//es.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		elastic.SetURL("http://127.0.0.1:9200"),
		// 禁用嗅探器用于兼容内网ip
		elastic.SetSniff(false),
		elastic.SetBasicAuth("user", "nvUt974rcNeg==*k0W3W"))
	if err != nil {
		log.Panicln("Elastic connect error:", err)
	}
	ECLIENT = client
	ECTX = context.Background()
	log.Infoln("Elastic connected")
	return ECLIENT, ECTX
	//mockData(ctx, client)
	//crud(client, ctx)
}

func GetESInstance() (context.Context, *elastic.Client) {
	return ECTX, ECLIENT
}

func crud(client *elastic.Client, ctx context.Context) {
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
	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)

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
