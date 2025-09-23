package main

import (
	"fmt"

	_ "github.com/learn/init_order/lesson01"
	_ "github.com/learn/init_order/lesson02"
	_ "github.com/learn/init_order/pkg1"
	_ "github.com/learn/init_order/pkg2"
	_ "github.com/learn/init_order/pkg3"
	_ "github.com/learn/init_order/pkg4"
	"github.com/learn/init_order/pkg5"
	_ "github.com/learn/init_order/pkg5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		fmt.Println("hear...", err)
	} else {
		fmt.Println("db is ok", db)
	}
	// fmt.Println("main method invoked!")

	//lesson01.Run(db)
	pkg5.RunForInitOrder(db)

}

//
