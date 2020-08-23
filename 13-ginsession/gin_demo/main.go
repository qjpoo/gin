package main

import (
	ginsession "gin/13-ginsession"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.LoadHTMLGlob("./13-ginsession/gin_demo/tmpl/*")

	// 初始化全局的MgrObj对象
	ginsession.InitMgr()
	// session中间件
	r.Use(ginsession.SessionMiddleware(ginsession.MgrObj))
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", AuthMiddleware,vipHandler)

	// 没有匹配的路由
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})


	r.Run(":8000")

}

