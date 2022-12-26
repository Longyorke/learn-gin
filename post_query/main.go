package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    string            `form:"address" bind:"required"`
	AddressMap map[string]string `form:"addressMap"`
}

func main() {
	Router := gin.Default()
	// http://localhost:8082/user/save
	Router.POST("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBind(&user) // 绑定表单
		if err != nil {
			log.Println(err)
		}
		addressMap := ctx.PostFormMap("addressMap")
		user.AddressMap = addressMap

		ctx.JSON(http.StatusOK, user)
	})

	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
