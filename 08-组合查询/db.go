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

func unionSearchById(id int64) (book *Book, err error) {
	sqlStr := `SELECT
b.id as book_id,
b.title,
b.price,
p.province,
p.city
from
	book AS b
INNER JOIN publisher AS p ON b.publisher_id = p.id
WHERE
	b.id = ?
`
	var b Book
	err = db.Get(&b, sqlStr, id)
	if err != nil {
		fmt.Println("通过id不能获取书籍的信息...")
		return
	}
	fmt.Println(b)
	return &b, nil

}
