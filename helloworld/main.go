package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	// 带参数的url
	Router.GET("/user/:id", func(ctx *gin.Context) {
		params := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"get_user_id": params,
		})
	})
	// 模糊匹配
	Router.GET("/user/match/*path", func(ctx *gin.Context) {
		params := ctx.Param("path")
		ctx.JSON(200, gin.H{
			"your_path": params,
		})
	})
	// 混合使用
	Router.GET("/user/mix/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, "%s is %s", name, action)
	})

	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
