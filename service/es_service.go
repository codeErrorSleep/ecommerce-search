package service

import (
	"context"
	"ecommerce-search/conf"
	"ecommerce-search/entity"
	"github.com/olivere/elastic/v7"
)

func GetInfoBySpuID(spuID string) error {
	ctx := context.TODO()

	matchEs := elastic.NewTermQuery("spu_id", spuID)
	// matchEs := elastic.("spu_id", spuID)
	esInfo, err := conf.EsClient.Search().
		Index("spu").
		Type("_doc").
		Query(matchEs).
		Do(ctx)

	if err != nil {
		return err
	}
	_ = esInfo.Error
	return nil
}

func GetListByGoodsName(req entity.QueryByGoodsName) error {
	ctx := context.TODO()

	matchEs := elastic.NewMatchQuery("goods_name", req.GoodsName)
	esList, err := conf.EsClient.Search().
		Index("spu").
		Type("_doc").
		Query(matchEs).
		Do(ctx)

	if err != nil {
		return err
	}

	_ = esList

	return nil
}
