package es

import (
	"context"
	"encoding/json"
	"errors"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertNetworkInfo(ctx context.Context, client *elastic.Client, result types.NetworkInfoResult) error {
	createIndexIfNotExists(ctx, client, "network_info")
	_, err := client.Index().
		Index("network_info").
		Id("network").
		BodyJson(result).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Debugln("InsertNetWorkInfo Success")
	return nil
}

func GetNetworkInfo(ctx context.Context, client *elastic.Client) (*types.NetworkInfoResult, error) {
	query := elastic.NewMatchAllQuery()
	res, err := client.Search().
		Index("network_info").
		Query(query).
		Size(1).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	if res.TotalHits() == 0 {
		return nil, errors.New("no data found")
	}
	var body types.NetworkInfoResult
	_ = json.Unmarshal(res.Hits.Hits[0].Source, &body)
	return &body, nil
}

func InsertValidator(ctx context.Context, client *elastic.Client, result types.ValidationStatusResult) error {
	_, err := client.Index().
		Index("validator").
		Id("validator").
		BodyJson(result).
		Do(ctx)
	if err != nil {
		return err
	}
	log.Debugln("InsertValidator Success")
	return nil
}

func QueryValidator(ctx context.Context, client *elastic.Client) (*types.ValidationStatusResult, error) {
	query := elastic.NewMatchAllQuery()
	res, err := client.Search().
		Index("validator").
		Query(query).
		Size(1).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var result types.ValidationStatusResult
	_ = json.Unmarshal(res.Hits.Hits[0].Source, &result)
	return &result, nil
}
