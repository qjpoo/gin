package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 上传文件

	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.MaxMultipartMemory = 100 << 20 // 100M


	// 加载模版文件
	r.LoadHTMLFiles("./04-文件上传/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file)
		dst := fmt.Sprintf("c:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded !", file.Filename),
		})
	})



	r.Run(":8000")


}