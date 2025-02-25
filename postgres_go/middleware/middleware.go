package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// type Response struct {
// 	Id      int64  `json:"id",omitempty`
// 	Message string `json:"message,omitempty"`
// }

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Postgres!!!")
	return db
}

func AddStock(w http.ResponseWriter, r *http.Request)     {}
func GetStock(w http.ResponseWriter, r *http.Request)     {}
func GetAllStocks(w http.ResponseWriter, r *http.Request) {}
func UpdateStock(w http.ResponseWriter, r *http.Request)  {}
func DeleteStock(w http.ResponseWriter, r *http.Request)  {}
