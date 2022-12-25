package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Id      int64    `form:"id"` // 注意是反引号,内部值为字符串
	Name    string   `form:"name"`
	Address []string `form:"address" binding:"required"`
	//tag加上binding:"required"后.BindQuery()返回400，ShouldBindQuery()返回200
}

func main() {
	Router := gin.Default()
	// http://localhost:8082/user/save?address=Beijing&address=Shanghai
	Router.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindQuery(&user) // 绑定类型方式
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(http.StatusOK, user)
	})
	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
