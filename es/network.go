package es

import (
	"encoding/json"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func InsertNetWorkInfo(result types.NetworkInfoResult) error {
	createIndexIfNotExists(ECTX, ECLIENT, "network_info")
	_, err := ECLIENT.Index().
		Index("network_info").
		BodyJson(result).
		Do(ECTX)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("InsertNetWorkInfo Success")
	return nil
}

func GetNetWorkInfo() (*types.NetworkInfoResult, error) {
	ctx, client := GetESInstance()
	query := elastic.NewMatchAllQuery()
	res, err := client.Search().
		Index("network_info").
		Query(query).
		Size(1).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var body types.NetworkInfoResult
	_ = json.Unmarshal(res.Hits.Hits[0].Source, &body)
	return &body, nil
}
