package controller

import (
	"ecommerce-search/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

func QueryBySpuID(c *gin.Context) {
	// 直接干了,绑定一下参数

	req := entity.QueryBySpuIDReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error ", err)
		return
	}
	// 直接查es的数据

}