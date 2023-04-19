package service

import (
	"ecommerce-search/conf"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
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

	spuID := "dfadsfasdfasdf"
	ret, err := GetInfoBySpuID(spuID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
	assert.Nil(t, err)
	assert.NotNil(t, ret)
}
