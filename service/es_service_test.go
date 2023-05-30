package service

import (
	"ecommerce-search/conf"
	"ecommerce-search/entity"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initEs() {

}

// TestGetInfoBySpuID 测试一下获取的
func TestGetInfoBySpuID(t *testing.T) {
	err := conf.NewEs()
	if err != nil {
		log.Fatal(err)
		return
	}

	spuID := "c11a8df3-5eb"
	err = GetInfoBySpuID(spuID)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(ret)
	assert.Nil(t, err)
	// assert.NotNil(t, ret)
}

// TestGetListByGoodsName 测试搜索
func TestGetListByGoodsName(t *testing.T) {
	err := conf.NewEs()
	if err != nil {
		log.Fatal(err)
		return
	}
	req := entity.QueryByGoodsName{
		GoodsName: "备里运中年联",
	}
	err = GetListByGoodsName(req)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(ret)
	assert.Nil(t, err)
	// assert.NotNil(t, ret)
}
