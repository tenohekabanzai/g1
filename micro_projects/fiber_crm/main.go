package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Lead struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email" gorm:"unique"`
	Phone   int    `json:"phone"`
}

func initDB() {
	var db, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Err connecting to DB")
	}
	fmt.Println("Connected to DB")
	db.AutoMigrate(&Lead{})
}

func main() {
	initDB()
	fmt.Println("Fiber CRUD App with sqlite3")
}
