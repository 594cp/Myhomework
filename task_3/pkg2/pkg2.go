package pkg2

import (
	"fmt"

	"gorm.io/gorm"
)

type Accounts struct {
	Id      int
	Balance float64
}

type Transactions struct {
	Id            int
	FromAccountId string
	ToAccountId   string
	Amount        float64
}

func init() {
	fmt.Println("pkg2 init method invoked")
}

func RunForInitOrder(db *gorm.DB) {
	fmt.Println("pkg2.PkgNameVar has been initialized")
	ac := []Accounts{
		{Id: 1,
			Balance: 200,
		},
		{Id: 2,
			Balance: 100,
		},
	}
	db.AutoMigrate(&Accounts{})
	db.Create(&ac)
	fmt.Println("Accounts created successfully")
	ac = []Accounts{}
	db.Find(&ac)
	fmt.Println(ac)
	db.AutoMigrate(&Accounts{}, &Transactions{})
	tx := db.Begin()
	temp1 := Accounts{}
	if err := tx.Where("id = ?", 1).First(&temp1).Error; err != nil {
		fmt.Println("zero balance in account 1")
		fmt.Println(err)
	}
	temp2 := Accounts{}
	if err := tx.Where("id = ?", 2).First(&temp2).Error; err != nil {
		fmt.Println("zero balance in account 2")
		fmt.Println(err)
	}
	if temp1.Balance >= 100 {
		temp1.Balance -= 100
		tx.Save(&temp1)
		temp2.Balance += 100
		tx.Save(&temp2)
		Tra := Transactions{
			FromAccountId: "1",
			ToAccountId:   "2",
			Amount:        100,
		}
		if err := tx.Create(&Tra).Error; err != nil {
			fmt.Println("create failed")
			fmt.Println(err)
		}
	} else {

		fmt.Println("Insufficient balance in account 1")
		tx.Rollback() // Rollback the entire transaction
	}
	tx.Commit()
}
