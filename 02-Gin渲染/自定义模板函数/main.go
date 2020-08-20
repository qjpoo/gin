package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	// 定义一个不转义相应内容的safe模板

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})

	r.LoadHTMLFiles("./02-Gin渲染/自定义模板函数/index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "<a href='https://www.baidu.com'>Click here</a>")
	})

	r.Run(":8000")
}
