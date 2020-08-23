package main

import "github.com/gin-gonic/gin"

func main() {
	// 路由分组
	r := gin.Default()

	// 分组
	shoppingGroup := r.Group("/shopping")
	{
		// http://localhost:8000/shopping/index
		shoppingGroup.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// http://localhost:8000/shopping/home
		shoppingGroup.GET("/home", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	liveGroup := r.Group("/live")
	{
		// http://localhost:8000/live/index
		liveGroup.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// http://localhost:8000/live/home
		liveGroup.GET("/home", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}


	// 分组的嵌套
	v1 := r.Group("/v1")
	{
		v1Nginx := v1.Group("/nginx")
		{
			// http://localhost:8000/v1/nginx/n1
			v1Nginx.GET("/n1", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"Path": "v1/nginx/n1",
				})
			})

			// http://localhost:8000/v1/nginx/n2
			v1Nginx.GET("/n2", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"Path": "v1/nginx/n2",
				})
			})
		}
	}

	r.Run(":8000")
}
