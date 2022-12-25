package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	Router := gin.Default()
	// http://localhost:8082/user/save?id=123456&name=LongYorke
	Router.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		address := ctx.DefaultQuery("address", "Beijing") // 默认方式
		mobile, mobileOk := ctx.GetQuery("mobile")        // 判断方式
		ctx.JSON(http.StatusOK, gin.H{
			"id":        id,
			"name":      name,
			"address":   address,
			"mobile":    mobile,
			"mobile_ok": mobileOk,
		})
	})
	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
