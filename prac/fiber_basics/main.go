package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    int64   `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var items []Item = []Item{
	{1, "Item_1", 10.99},
	{2, "Item_2", 15.99},
	{3, "Item_3", 20.99},
}

func getAllItems(c *fiber.Ctx) error {
	return c.JSON(items)
}

func getItem(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 0, 0)
	for _, val := range items {
		if val.ID == id {
			return c.JSON(val)
		}
	}
	return c.Status(http.StatusNotFound).SendString("Could not find Item!!!")

}

func AddItem(c *fiber.Ctx) error {

	var item Item

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100000) + 1
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}
	item.ID = int64(randomNumber)
	items = append(items, item)
	return c.JSON(items)

}

func UpdateItem(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 0, 0)
	for idx, val := range items {
		if val.ID == id {
			items = append(items[:idx], items[idx+1:]...)
		}
	}
	var item Item
	err := c.BodyParser(&item)
	if err != nil {
		fmt.Println(err)
	}
	item.ID = id
	items = append(items, item)
	return c.JSON(items)
}

func DeleteItem(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 0, 0)
	for idx, val := range items {
		if val.ID == id {
			items = append(items[:idx], items[idx+1:]...)
		}
	}
	return c.JSON(items)
}

func main() {

	app := fiber.New(fiber.Config{
		AppName:           "Fiber Basics CRUD App",
		EnablePrintRoutes: true,
		ServerHeader:      "Fiber Basics CRUD App Response Header",
	})
	app.Static("/", "./static")
	app.Get("/items", getAllItems)
	app.Get("/items/:id", getItem)
	app.Post("/items", AddItem)
	app.Put("/items/:id", UpdateItem)
	app.Delete("/items/:id", DeleteItem)
	if err := app.Listen(":5004"); err != nil {
		fmt.Println(err)
	}

}
