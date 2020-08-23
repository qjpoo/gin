package main

type Book struct {
	ID int64 `db:"book_id"`
	Title string `db:"title"`
	Price float64 `db:"price"`
	Publisher
}

// 书籍的出版社有关信息
type Publisher struct {
	ID int64 `db:"publisher_id"`
	Province string `db:"province"`
	City string `db:"city"`

}

