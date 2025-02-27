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

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Err connecting to DB")
	}
	fmt.Println("Connected to DB")
	err = db.AutoMigrate(&Lead{})
	if err != nil {
		panic("Failed to migrate to DB")
	}

}

func Getleads(c *fiber.Ctx) error {
	var leads []Lead
	err := db.Find(&leads)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(leads)
	return nil
}
func Getlead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	err := db.Find(&lead, id)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(lead)
	return nil
}
func Addlead(c *fiber.Ctx) error {
	var lead Lead
	err := c.BodyParser(&lead)
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&lead)
	c.JSON(lead)
	return nil
}
func Deletelead(c *fiber.Ctx) error {

	id := c.Params("id")
	var lead Lead
	err := db.Find(&lead, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// if lead.Name == "" {
	// 	c.Status(500).Send([]byte("Cannot find lead with given id"))
	// 	return nil
	// }
	db.Delete(&lead)
	c.Send([]byte("deleted lead successfully"))
	return nil
}

func main() {

	initDB()
	fmt.Println("Fiber CRUD App with sqlite3")
	app := fiber.New()
	// routes
	app.Get("/api/leads", Getleads)
	app.Get("/api/lead/:id", Getlead)
	app.Post("/api/lead", Addlead)
	app.Delete("/api/lead/:id", Deletelead)
	app.Listen(":5004")

}
