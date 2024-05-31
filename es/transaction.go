package es

import (
	"context"
	"encoding/json"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertTxnStatus(ctx context.Context, client *elastic.Client, result types.TxnStoreResult) error {
	_, err := client.Index().
		Index("transaction").
		Id(result.Transaction.Hash).
		BodyJson(result).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Debugln("[InsertTransaction] Success")
	return nil
}

func QueryTxnStatusList(ctx context.Context, client *elastic.Client, pageNum, pageSize int) ([]types.TxnStoreResult, int64, error) {
	from := (pageNum - 1) * pageSize
	searchResult, err := client.Search().
		Index("transaction").
		From(from).
		Size(pageSize).
		Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	var txns []types.TxnStoreResult
	for _, hit := range searchResult.Hits.Hits {
		var res types.TxnStoreResult
		if err := json.Unmarshal(hit.Source, &res); err == nil {
			txns = append(txns, res)
		}
	}
	return txns, searchResult.TotalHits(), nil
}

// Query transaction status by hash
func QueryTxnStatusByHash(ctx context.Context, client *elastic.Client, hash string) (*types.TxnStoreResult, error) {
	result, err := client.Get().
		Index("transaction").
		Id(hash).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var txn types.TxnStoreResult
	if err := json.Unmarshal(result.Source, &txn); err != nil {
		return nil, err
	}
	return &txn, nil
}

// Query transaction status by block height
func QueryTxnByHeight(ctx context.Context, client *elastic.Client, height int64) (*types.TxnStoreResult, error) {
	termQuery := elastic.NewTermQuery("height", height)
	result, err := client.Search().
		Index("transaction").
		Query(termQuery).
		Size(1).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var txn types.TxnStoreResult
	if err := json.Unmarshal(result.Hits.Hits[0].Source, &txn); err != nil {
		return nil, err
	}
	return &txn, nil
}

func QueryAccountTxns(ctx context.Context, client *elastic.Client, pageNum, pageSize int, accId string) ([]types.TxnStoreResult, int64, error) {
	termQuery := elastic.NewBoolQuery().Should(
		elastic.NewTermQuery("transaction.receiver_id", accId),
		elastic.NewTermQuery("transaction.signer_id", accId),
	)
	from := (pageNum - 1) * pageSize
	searchResult, err := client.Search().
		Index("transaction").
		From(from).
		Size(pageSize).
		Query(termQuery).
		Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	var txns []types.TxnStoreResult
	for _, hit := range searchResult.Hits.Hits {
		var res types.TxnStoreResult
		if err := json.Unmarshal(hit.Source, &res); err == nil {
			txns = append(txns, res)
		}
	}
	return txns, searchResult.TotalHits(), nil
}
