package main

import (
	"ecommerce-search/conf"
	"ecommerce-search/es"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := conf.NewEs()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := route()
	r.Run(":8099") // listen and serve on 0.0.0.0:8080
}

func route() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "request success")
	})

	// query es 直接查,后面直接分页
	r.POST("/query", es.QueryBySpuID)
	return r
}