package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	Router := gin.Default()
	Router.GET("/hello", func(ctx *gin.Context) {
		//返回数据
		//type any = interface{}
		//type H map[string]any
		ctx.JSON(200, gin.H{
			"name": "LongYorke",
		})
	})
	Router.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "POST")
	})
	Router.DELETE("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "DELETE")
	})
	Router.PUT("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "PUT")
	})
	Router.PATCH("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "PATCH")
	})
	Router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, "GET")
	})
	Router.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"method": "Any"})
	})

	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
