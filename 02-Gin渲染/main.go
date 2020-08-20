package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	// 渲染

	r := gin.Default()

	// 加载模版文件
	r.LoadHTMLGlob("./02-Gin渲染/tmpl/**/*")

	// 设置静态文件的目录
	// args1: 第一个参数是代码里面使用的路径,就是url访问的路径
	// http://localhost:8000/ login/images/login.png   ./02-Gin渲染/static/images/login.png
	// args2:  实际保存静态文件的路径
	r.Static("/login", "./02-Gin渲染/static")

	// 注意html文件里面的define
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "front/login.html", gin.H{
			"msg": "hello, world ...", // 向login.html里面放的数据
		})
		
	})


	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "front/index.html", gin.H{
			"msg": "index ...", // 向login.html里面放的数据
		})

	})




	fmt.Println("----------------------------------------------")

	// json
	// 自己拼接json
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"	})

	})
	// 使用结构体
	r.GET("/json2", func(c *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string `json:"msg"`

			Age int
		}
		msg.Name = "chiling"
		msg.Message = "ok"
		msg.Age = 48
		c.JSON(200, msg)
	})

	// xml
	// 自己拼接
	r.GET("/xml1", func(c *gin.Context) {
		c.XML(200, gin.H{
			"msg": "ok",
		})
	})

	// 使用结构体
	r.GET("/xml2", func(c *gin.Context) {
		type msgRecord struct {
			Name string
			Message string
			Age int
		}
		var msg msgRecord
		msg.Name = "qujian"
		msg.Age = 20
		msg.Message = "ok"
		c.XML(200, msg)
	})

	// 使用yaml
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"msg": "ok"})

	})

	// protobuf渲染
	r.GET("/proto", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "ok"
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(200, data)
	})

	r.Run(":8000")

}
