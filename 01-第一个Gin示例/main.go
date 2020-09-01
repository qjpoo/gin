package  main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func f(context *gin.Context)  {
	context.JSON(200, gin.H{
		"message": "ok, hi ...",
	})
}
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// GET请求 /hello
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello... world",
		})
	})

	// get /hi
	r.GET("/hi", f)


	// 启动
	r.Run(":8000")
}
