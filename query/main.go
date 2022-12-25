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
	AddressMap map[string]string `form:"addressMap"`
}

func main() {
	Router := gin.Default()
	// http://localhost:8082/user/save?addressMap[college]=Beijing&addressMap[company]=Shanghai
	Router.GET("/user/save", func(ctx *gin.Context) {
		var user User
		addressMap, _ := ctx.GetQueryMap("addressMap") // 判断方式
		user.AddressMap = addressMap
		ctx.JSON(http.StatusOK, user)
	})
	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
