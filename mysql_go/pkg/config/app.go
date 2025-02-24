package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/jinzhu/gorm"

	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("DB_PASSWORD")
	conn_str := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local", password)

	d, err := gorm.Open("mysql", conn_str)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MySQL DB")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
