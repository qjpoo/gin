package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func main() {

	r := gin.Default()

	r.Any("/hi", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		// c.ShouldBind 使用了 c.Request.Body，不可重用。
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be formA -- json`)
		} else if errA := c.ShouldBindBodyWith(&objA, binding.XML); errA == nil {
			c.String(http.StatusOK, `the body should be formA -- xml`)
			// 因为如果使用ShouldBind,现在 c.Request.Body 是 EOF，所以这里会报错,不能用ShouldBind,要使用ShouldBindBodyWith。
		} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, `the body should be formB -- json`)
		} else if errB := c.ShouldBindBodyWith(&objB, binding.XML); errB == nil {
			c.String(http.StatusOK, `the body should be formB --- xml`)
		} else {
			c.JSON(200, gin.H{
				"msg": "not found",
			})
		}
	})
	r.Run(":8000")
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}
