package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
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
		panic("Err connecting to DB")
	}
	fmt.Println("Connected to DB")
	err = db.AutoMigrate(&Lead{})
	if err != nil {
		panic("Failed to migrate to DB")
	}
}

func main() {
	initDB()
	fmt.Println("Fiber CRUD App with sqlite3")

	app := fiber.New()
	app.Listen(":5004")
}
