package pkg1

import (
	"fmt"

	"gorm.io/gorm"
	//_ "github.com/learn/init_order/pkg2"
)

func init() {
	//fmt.Println("pkg1 init method invoked")
}

type Student struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Age   int
	Grade string
	gorm.Model
}

func RunForInitOrder(db *gorm.DB) {
	fmt.Println("pkg1.PkgNameVar has been initialized")

	errors := db.AutoMigrate(&Student{})
	fmt.Println(errors)
	students := []Student{
		{Name: "张三", Age: 23, Grade: "三年级"},
		{Name: "李四", Age: 14, Grade: "四年级"},
		{Name: "王五", Age: 15, Grade: "五年级"},
		{Name: "赵六", Age: 26, Grade: "六年级"},
		{Name: "李四", Age: 18, Grade: "五年级"},
		{Name: "张三", Age: 22, Grade: "四年级"},
	}
	db.Create(&students)

	//var studentsFromDB []Student
	//db.Model(&Student{}).Where("age > ?", 18).Find(&studentsFromDB)
	//fmt.Println("Students from DB:", studentsFromDB)

	//db.Model(&Student{}).Where("Name = ?", "张三").Update("Grade", "九年级")
	// stu := Student{}
	// db.Model(&Student{}).Where("Age < ?", 15).First(&stu)
	// fmt.Println(stu)

	//stu1 := Student{ID: 2, Name: "李四", Age: 20, Grade: "八年级"}
	//db.Create(&stu1)

	//db.Delete(&Student{}, 2)

	

}
