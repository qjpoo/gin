package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
	querystring指的是URL中?后面携带的参数，例如：/user/search?username=小王子&address=沙河
	参数解析三种
	1. query string     c.query()  c.DefaultQuery("key", "DefaultValue")
	2. form param       c.PostForm("key")   c.DefaultPostForm("key", "DefaultValue")
	3. path param       /url/:key1/:key2    c.Param("key")
	 */

	// http://localhost:8000/user/search?address=xxoo&username=chiling
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		userName := c.DefaultQuery("username", "chiling")
		address := c.Query("address") // 查询不到就为空
		c.JSON(200, gin.H{
			"msg": "ok",
			"username": userName,
			"address": address,
		})
	})

	// 获取form参数
	r.POST("/post", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		username := c.DefaultPostForm("username","chiling")
		//username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(200, gin.H{
			"msg": "ok",
			"username": username,
			"address": address,
		})

	})

	//获取path参数
	// /usr/search/chiling/honghu
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(200, gin.H{
			"msg": "ok",
			"usernamme": username,
			"address": address,
		})
	})

	// 博客归档
	// /getTar/2020/08  获取2020年8月的文章
	r.GET("/getTar/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": map[string]string{
				"year": year,
				"month": month,
			},
		})
	})

	// 参数绑定
	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	r.POST("loginJson", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Println("login info: ", login)
			c.JSON(200, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//绑定form表单示例 (user=q1mi&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})


	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(200, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})







	r.Run(":8000")

}

type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
