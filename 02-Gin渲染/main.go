package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 渲染

	r := gin.Default()

	// 加载模版文件
	r.LoadHTMLGlob("./02-Gin渲染/tmpl/*")

	// 设置静态文件的目录
	// args1: 第一个参数是代码里面使用的路径,就是url访问的路径
	// http://localhost:8000/ login/images/login.png   ./02-Gin渲染/static/images/login.png
	// args2:  实际保存静态文件的路径
	r.Static("/login", "./02-Gin渲染/static")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "hello, world ...", // 向login.html里面放的数据
		})
		
	})


	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "index ...", // 向login.html里面放的数据
		})

	})



	

	r.Run(":8000")

}
