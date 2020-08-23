package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// 中间件  相当于是一个钩子

	r := gin.Default() // 默认使用了二个中间件 Log和Recovery

	// 加载模版文件
	r.LoadHTMLGlob("./06-middleware-中间件/tmpl/*")

	// 一个可以匹配所有请求方法的Any方法
	r.Any("/login", func(c *gin.Context) {
		if c.Request.Method == "POST" {
			username := c.PostForm("username")
			pasword := c.PostForm("password")
			c.JSON(200, gin.H{
				"username": username,
				"password": pasword,
			})
		} else {
			// 如果是get方法就返回一个html的登陆页面
			c.HTML(200, "login.html", nil)

		}
	})

	// 为没有配置处理函数的路由添加处理程序
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})

	// 使用中间件
	r.Use(castTime)

	shoppingGroup := r.Group("/shopping")
	{
		// castTime先调用
		shoppingGroup.GET("/index", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
			fmt.Println(c.MustGet("name").(string)) //拿到一个interface的值,之后进行断言
			c.JSON(200, gin.H{
				"msg": "ok",
			})
		})
		shoppingGroup.GET("/home", func(c *gin.Context) {
			time.Sleep(3 * time.Second)
			fmt.Println("name")
			c.JSON(200, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(":8000")

}

func castTime(c *gin.Context) {
	start := time.Now()
	// 先进行变量的设置
	c.Set("name", "chiling")
	// 运行下一个handler函数
	c.Next()
	// 计算耗时
	cast := time.Since(start)
	log.Println("cast: ", cast)
	//fmt.Println("cast time is : ", cast)

}
