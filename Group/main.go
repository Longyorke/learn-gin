package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	Router := gin.Default()
	//	v2版本
	v2 := Router.Group("/v2")
	v2User := v2.Group("/user")
	{
		v2User.POST("/login", loginEndpoint)
		v2User.POST("/submit", submitEndpoint)
		v2User.POST("/read", readEndpoint)
	}

	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}

func loginEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "loginEndpoint")
}
func submitEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "submitEndpoint")
}
func readEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "readEndpoint")
}
