package es

import (
	"ecommerce-search/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	// 直接干了,绑定一下参数

	req := entity.QueryListReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error ", err)
		return
	}
	// 直接查es的数据
}