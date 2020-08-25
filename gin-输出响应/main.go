package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gookit/color.v1"
	"net/http"
)

func main() {
	// gin-输出响应

	r := gin.Default()

	// 响应string
	r.GET("/g1", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// g.h 嵌套返回值
	r.GET("/g2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "123",
			"data": gin.H{
				"lala": "chiling",
				"age":  20,
			},
		})
	})

	// 返回第三方数据
	r.GET("/g3", func(c *gin.Context) {
		response, err := http.Get("https://www.baidu.com")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Dispostion": `attachment; filenname="xxoo"`,
		}
		//c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		color.Red.Println(http.StatusOK, contentType, contentLength, reader, extraHeaders)
		//fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
	})

	r.GET("/g4", func(c *gin.Context) {
		name := []string{"chiling", "qujian", "zhangs"}
		c.SecureJSON(200, name)
	})

	r.Run(":8000")
}
