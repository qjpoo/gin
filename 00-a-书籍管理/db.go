package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 跟数据库相关的操作
var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:qujian123@tcp(47.98.179.41:13360)/bms"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(10)
	return
}

// 查数据库的数据
func queryAllBook() (bookList []*Book, err error) {
	fmt.Printf("%T\n", bookList)
	sqlStr := `select id, title, price from book`
	//books := make([]*Book, 64)
	//err = db.Select(&bookList,sqlStr)
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bookList: ", bookList)
	return bookList, nil
}

// 插入数据
func insertBook(title string, price float64) (err error) {
	sqlStr := `insert into book(title, price) values (?, ?)`
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入书籍信息失败")
		return
	}
	return
}

// 删除数据
func deleteBook(id int64) (err error) {
	sqlStr := `delete from book where id = ?`
	_, err = db.Exec(sqlStr,id)
	if err != nil {
		fmt.Println("删除书籍失败 ...")
		return
	}
	return
}

func getDataById(id int64) (bookinfo *Book, err error )  {
	sqlStr := `select id, title, price from book where id = ?`
	var b Book
	err = db.Get(&b, sqlStr, id)
	if err != nil {
		fmt.Println("通过id不能获取书籍的信息...")
		return
	}
	return &b, nil

}


// 更新数据
func updateBook(id int64, title string, price float64) (err error) {
	//sqlStr := `insert into book(title, price) values (?, ?)`
	sqlStr := `update book set title = ?, price = ? where id = ?`
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		fmt.Println("插入书籍信息失败")
		return
	}
	return
}
