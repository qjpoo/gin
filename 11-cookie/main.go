package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// cookie

	r := gin.Default()

	// 加载模版文件
	r.LoadHTMLGlob("./11-cookie/tmpl/*")

	r.Any("/login", func(c *gin.Context) {
		if c.Request.Method == "POST" {
			toPath := c.DefaultQuery("next", "/index")
			var u UserInfo
			err := c.ShouldBind(&u)
			if err != nil {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "用户名和密码不能为空!",
				})
			}

			if u.Username == "chiling" && u.Password == "123" {
				// 如果登陆成功,设置cookie
				// 1. 设置cookie
				c.SetCookie("username", u.Username, 20, "/", "localhost", false, true)

				// 正常的话跳到index页面
				c.Redirect(302, toPath)
			} else {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"err": "用户名或密码错误!",
				})
			}

		} else {
			c.HTML(200, "login.html", nil)
		}
	})

	r.GET("/index", func(c *gin.Context) {

		c.HTML(200, "index.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		// home页面只有登陆了之后,才能看.所以要在返回页面之前要先校验是否存在username的cookie
		// 获取cookie
		username, err := c.Cookie("username")
		if err != nil {
			// 说明没有登陆
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		c.HTML(200, "home.html", gin.H{
			"username": username,
		})
		//c.String(200, "xxoo")

	})

	r.GET("/path", func(c *gin.Context) {
		path := fmt.Sprintf("%s?next=%s", "/path", c.Request.URL.Path)
		c.JSON(200, gin.H{
			"data": path,
		})
	})

	r.Use(vipHandle)
	// vip页面
	r.GET("/vip", func(c *gin.Context) {
		/*
			c.HTML(200, "vip.html", gin.H{
				// username返回的是一个接口
				"username": c.MustGet("username").(string),
			})
		*/

		username, ok := c.Get("username")
		if !ok {
			// 如果取不到值 , 说明前面中间件出问题了
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		username, ok = username.(string)
		if !ok {
			// 类型断言失败
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		c.HTML(200, "vip.html", gin.H{
			// username返回的是一个接口
			"username": username,
		})
	})

	r.Run(":8000")

}

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func vipHandle(c *gin.Context) {
	// 获取cookie
	username, err := c.Cookie("username")
	if err != nil {
		//c.Redirect(http.StatusMovedPermanently, "/login")
		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		fmt.Println(toPath)
		c.Redirect(http.StatusMovedPermanently, toPath)
	}
	/*
		c.JSON(200, gin.H{
			"username": username,
		})
	*/
	// 用户已登陆了
	c.Set("username", username)
	// 执行后面的handle
	c.Next()

}
