package pkg3

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
	fmt.Println("pkg3 init method invoked")
}

func RunForInitOrder(db *gorm.DB) {
	fmt.Println("pkg3.PkgNameVar has been initialized")
	db.AutoMigrate(&Employees{}, &Employ{})

	db2, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	} else {
		fmt.Println("yes")
	}
	defer db2.Close()

	// db.Create(&Employees{
	// 	Name:       "JNadsf",
	// 	Department: "技术部",
	// 	Salary:     50800,
	// })
	// db.Create(&Employees{
	// 	Name:       "Kadfd",
	// 	Department: "技术部",
	// 	Salary:     60100,
	// })
	// db.Create(&Employees{
	// 	Name:       "Dsfs",
	// 	Department: "营销部",
	// 	Salary:     70000,
	// })

	emp := []Employees{}
	err = db2.Select(&emp, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("使用sqlx查询技术部员工信息:", emp)
	/////////////////////////
	emp = []Employees{}
	err = db.Where("department = ?", "技术部").Find(&emp).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("使用gorm查询技术部员工信息:", emp)
	/////////////////////////
	emp2 := Employees{}
	err = db2.Get(&emp2, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("使用sqlx查询Salary最高的员工信息:", emp2)
	/////////////////////////
	emp2 = Employees{}
	err = db.Select("*").Order("salary desc").Limit(1).Find(&emp2).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("使用gorm查询Salary最高的员工信息:", emp2)

}
