package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(ctx *gin.Context) {
		//返回数据
		//type any = interface{}
		//type H map[string]any
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})
	err := engine.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
