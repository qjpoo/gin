package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	// 组合查询示例

	// 连接数据库
	err := initDB()
	if err != nil {
		panic("connection DB failure ...")
	}

	r := gin.Default()

	r.GET("/search/:id", func(c *gin.Context) {
		// 得到id
		tmpId := c.Param("id")
		if len(tmpId) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求的数据非法...",
			})
			return
		}

		// string == > int64
		id, err := strconv.ParseInt(tmpId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求的数据非法...",
			})
			return
		}
		fmt.Println(id)

		// 得到的id, 去数据库里面查询
		b, e := unionSearchById(id)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		c.JSON(200, gin.H{
			"msg": "ok",
			"data": b,
		})

	})

	r.Run(":8000")

}
