package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    string            `form:"address"`
	AddressMap map[string]string `form:"addressMap"`
}

//type User struct {
//	Address    []string          `json:"address"`
//	AddressMap map[string]string `json:"addressMap"`
//	ID         int64             `json:"id"`
//	Name       string            `json:"name"`
//}

//type AddressMap struct {
//	Company string `json:"company"`

type Person struct {
	// id限制string类型;必填而且必须是uuid格式
	ID string `uri:"id" binding:"required,uuid"`
	// age限制int类型必填
	Age int `uri:"age" binding:"required"`
	// name限制string类型必填
	Name string `uri:"name" binding:"required"`
}

func main() {
	Router := gin.Default()
	// http://localhost:8082/user/save
	// GET请求参数
	//Router.POST("/user/save", func(ctx *gin.Context) {
	//	var user User
	//	err := ctx.ShouldBind(&user) // 绑定表单
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	addressMap := ctx.PostFormMap("addressMap") // 单独获取map
	//	user.AddressMap = addressMap
	//
	//	ctx.JSON(http.StatusOK, user)
	//})

	// Post请求参数
	//Router.POST("/user/save", func(ctx *gin.Context) {
	//	var user User
	//	err := ctx.ShouldBindJSON(&user) // 绑定表单
	//	if err != nil {
	//		log.Println(err)
	//	}
	//
	//	ctx.JSON(http.StatusOK, user)
	//})

	// 路径参数
	//Router.POST("/find/:id/:age/:name", func(context *gin.Context) {
	//	var p Person
	//	if err := context.ShouldBindUri(&p); err != nil {
	//		context.JSON(http.StatusBadRequest, gin.H{
	//			"msg": err.Error(),
	//		})
	//		return
	//	}
	//	context.JSON(http.StatusOK, p)
	//})

	// 文件参数
	Router.POST("/user/save", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		//Value map[string][]string
		//File  map[string][]*FileHeader
		value := form.Value //表单中的除文件外所有数据
		files := form.File  //表单中的文件数据
		//遍历参数名称对应的每一个文件数组（同一个参数名称下可有多个文件）
		for _, fileArray := range files {
			//遍历文件数组中每一个FileHeader元素
			//type FileHeader struct {
			//	Filename string
			//	Header   textproto.MIMEHeader
			//	Size     int64
			//
			//	content []byte
			//	tmpfile string
			//}
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./"+v.Filename) //文件保存到当前目录
			}
		}
		ctx.JSON(http.StatusOK, value)
	})

	err := Router.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
