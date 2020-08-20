package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	/*
		book management system
		书籍的信息:
			id			书的ID号
			bookName	书名
			price	    书的价格

		web页面版的增删改查


	*/

	// 连接数据库
	err := initDB()
	if err != nil {
		panic("connection DB failure ...")
	}

	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("./书籍管理/tmpl/**/*")

	// 查看所有的书数据
	r.GET("/book/list", func(c *gin.Context) {
		// 连接数据库
		// 查数据
		bookList, err := queryAllBook()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1, // 0 正确, 1 错误
				"message": err,
			})
			return
		}
		// 返回给浏览器,这里是json格式的
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 0,
		//	"data": bookList,

		// 以html格式返回数据
		c.HTML(200, "book/book_list.tmpl", gin.H{
			"code": 0,
			"data": bookList,
		})

	})

	// 插入数据

	// 返回一个页面给用户填写新增的书籍信息
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(200, "book/new.html", nil)
	})

	r.POST("/book/create", func(c *gin.Context) {
		// 创建书籍处理函数
		// 从form表单里面取数据
		title := c.PostForm("title")
		price := c.PostForm("price")
		//fmt.Printf("%T, %v\n", title,title) // string
		//fmt.Printf("%T, %v\n", price,price) // string
		priceValue, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%T, %v\n", priceValue, priceValue) // float64
		err = insertBook(title, priceValue)
		if err != nil {
			fmt.Println("插入数据失败 ...", err)
		}

		// 如果插入成功的话
		fmt.Println("插入数据成功 ...")
		c.Redirect(301, "/book/list")
	})

	/*
		r.POST("/book/insert", func(c *gin.Context) {
			// 连接数据库
			// 查数据
			err := insertBook("Jenkins使用", 63.8)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 1, // 0 正确, 1 错误
					"message": err,
				})
				return
			}
			// 返回给浏览器
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": "{}",
			})

		})
	*/

	//queryRows()

	r.Run(":8000")
}

// 多行
func queryRows() {
	sqlStr := `select id, title, price from book`
	var b []Book // 一定要是一个结构体的切片
	err := db.Select(&b, sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range b {
		fmt.Printf("id: %d, title: %s, price: %f\n", v.ID, v.Title, v.Price)
	}
}
