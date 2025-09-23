package pkg4

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 一定不能忘记导入数据库驱动
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Employees struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

type Employ struct {
	Id         int
	Name       string
	Department string
}

func init() {
	fmt.Println("pkg4 init method invoked")
}

// 定义 Book 结构体，对应 books 表的字段
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func RunForInitOrder(db *gorm.DB) {
	fmt.Println("pkg4.PkgNameVar has been initialized")
	db.AutoMigrate(&Book{})
	Books := []Book{
		{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 78.99},
		{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 10.99},
		{ID: 3, Title: "1984", Author: "George Orwell", Price: 12.99},
		{ID: 4, Title: "The Cafe cate", Author: "Salinger", Price: 56.99},
		{ID: 5, Title: "To Kill a Mdfdfockingbird", Author: "Hafdrper Lee", Price: 10.99},
		{ID: 6, Title: "198444", Author: "Geoe Orwell", Price: 58.99},
	}
	db.Create(&Books)
	db2, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	} else {
		fmt.Println("yes")
	}
	defer db2.Close()
	// 执行复杂的查询：查询价格大于 50 元的书籍

	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	Books = []Book{}
	err = db2.Select(&Books, query, 50)
	if err != nil {
		log.Fatalf("Failed to query books: %v\n", err)
	} else {
		fmt.Println("The books with price greater than 50 are:")
		for _, book := range Books {
			fmt.Printf("%d %s\n", book.ID, book.Title)
		}
	}

}
