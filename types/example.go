package types

import (
	"github.com/olivere/elastic/v7"
	"time"
)

type Example struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ExampleRes struct {
	Phone string `json:"phone"`
}

type Weibo struct {
	User     string                `json:"user"`               // 用户
	Message  string                `json:"message"`            // 微博内容
	Retweets int                   `json:"retweets"`           // 转发数
	Image    string                `json:"image,omitempty"`    // 图片
	Created  time.Time             `json:"created,omitempty"`  // 创建时间
	Tags     []string              `json:"tags,omitempty"`     // 标签
	Location string                `json:"location,omitempty"` // 位置
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}
