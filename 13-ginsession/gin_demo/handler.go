package main

import (
	"fmt"
	ginsession "gin/13-ginsession"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}


// 编写一个校验用户是否登陆的中间件
// 其实就是从上下文中取的session data, 从session data中取到isLogin字段
func AuthMiddleware(c *gin.Context)  {
	// 1. 从上下文中取的session data
	tmpSd, _ := c.Get(ginsession.SessionContextName)
	sd := tmpSd.(*ginsession.SessionData)

	// 2. 从session data取的isLogin
	value, err := sd.Get("isLogin")
	if err !=nil {
		// 取不到就是没有登陆,没有登陆就跳转
		c.Redirect(http.StatusFound, "/login")
		return
	}
	isLogin, ok := value.(bool)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	if !isLogin {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Next()
}





func loginHandler(c *gin.Context) {
	//c.HTML(200, "login.html", nil)
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名和密码不能为空!",
			})
		}
		fmt.Println("username: ",u.Username, "password: ", u.Password)
		if u.Username == "chiling" && u.Password == "123" {
			// 用cookie玩
			// 如果登陆成功,设置cookie
			// 1. 设置cookie
			//c.SetCookie("username", u.Username, 20, "/", "localhost", false, true)

			// 用session玩
			// 登陆成功, 在当前这个用户session data保存一个键值对 isLogin = true
			// 1. 先从上下文中获取session data
			tmpSd, ok := c.Get(ginsession.SessionContextName)
			if !ok {
				panic("session middleware failure ...")
				return
			}
			sd := tmpSd.(*ginsession.SessionData)

			// 2. 给session data设置islLogin
			sd.Set("isLogin", true)
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

}
func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)

}
func homeHandler(c *gin.Context) {
	c.HTML(200, "home.html", nil)

}
func vipHandler(c *gin.Context) {
	c.HTML(200, "vip.html", nil)

}
