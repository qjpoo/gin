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
	r.LoadHTMLGlob("./00-a-书籍管理/tmpl/**/*")

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

	// 删除
	r.GET("/book/delete", func(c *gin.Context) {
		// 取query string参数
		id := c.Query("id")
		idValue, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// id数据没有问题
		err = deleteBook(idValue)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}

		// 删除成功之后跳转
		c.Redirect(301, "/book/list")
	})

	// 编辑, 这里的r.GET和r.POST可以合在一起写 r.any 里面写一个if post else 判断语句
	r.GET("/book/edit", func(c *gin.Context) {
		// 通过传来的id, 查找出 title, price
		// 获取id
		id := c.Query("id")
		// 字符串转化了Int64
		getId, _ := strconv.ParseInt(id, 10, 64)
		b, err := getDataById(getId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// 渲染edit页面
		c.HTML(200, "book/edit.html", gin.H{
			"code": 0,
			"data": b,
		})
		//fmt.Println(b)
	})

	// 编辑完成之后,在提交到数据库里面
	r.POST("/book/change", func(c *gin.Context) {
		// 从form表单里面取数据
		//如果id要从form表单里面获取的话,要在html多一个Input控件,可以把id参数放在url里面
		//id, _ := c.GetPostForm("id")
		//id := c.PostForm("id")
		id := c.Query("id")
		fmt.Println("------------------>id: ", id)
		// 如果id是无效的
		if len(id) == 0 {
			c.String(http.StatusBadRequest, "id err ...")
			return
		}

		// 把字符的id转化成int64
		idValue, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		title := c.PostForm("title")
		fmt.Println("------------------>title: ", title)
		price := c.PostForm("price")
		fmt.Println("------------------>price: ", price)
		//fmt.Printf("%T, %v\n", title,title) // string
		//fmt.Printf("%T, %v\n", price,price) // string
		// 把字符的price转化成float64
		priceValue, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Printf("%T, %v\n", priceValue, priceValue) // float64



		err = updateBook(idValue, title, priceValue)
		if err != nil {
			fmt.Println("更新数据失败 ...", err)
		}

		// 如果更新成功的话,就跳转到list页面
		fmt.Println("更新数据成功 ...")
		c.Redirect(301, "/book/list")
	})

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
