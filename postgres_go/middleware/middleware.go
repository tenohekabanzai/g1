package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"postgres_go/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Response struct {
	ID      int64  `json:"id",omitempty`
	Message string `json:"message,omitempty"`
}

func Connect_to_PostGres() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load env file")
	}

	db, err := sql.Open("postgres", os.Getenv("postgres_url"))

	if err != nil {
		panic("Failed to Authenticate with Postgres!!")
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to connect with Postgres!!")
	}

	fmt.Println("Connected to Postgres DB!!")
	return db
}

func AddStock(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var stock model.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Unable to parse Request body", http.StatusBadRequest)
		return
	}

	db := Connect_to_PostGres()
	defer db.Close()

	q := "INSERT INTO stocks(name,price,company) VALUES ($1,$2,$3) RETURNING stockid"
	var id int64
	err = db.QueryRow(q, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := Response{
		ID:      id,
		Message: "Stock addded to DB successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		http.Error(w, "Give Proper Id in path", http.StatusBadRequest)
		return
	}

	var stock model.Stock
	db := Connect_to_PostGres()
	defer db.Close()

	q := "SELECT * FROM stocks where stockid =$1"
	row := db.QueryRow(q, id)
	err = row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No rows to show!", http.StatusOK)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stock)

}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {

	var stocks []model.Stock
	db := Connect_to_PostGres()
	defer db.Close()

	q := "SELECT * FROM stocks"
	rows, err := db.Query(q)
	if err != nil {
		http.Error(w, "Error fetching stocks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var st1 model.Stock
		err = rows.Scan(&st1.StockId, &st1.Name, &st1.Price, &st1.Company)
		if err != nil {
			http.Error(w, "Error fetching stocks", http.StatusInternalServerError)
			return
		}

		stocks = append(stocks, st1)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var stock model.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Unable to parse Request body", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		http.Error(w, "Give Proper Id in path", http.StatusBadRequest)
		return
	}

	db := Connect_to_PostGres()
	defer db.Close()
	q := "UPDATE stocks SET name=$2,price=$3,company=$4 where stockid=$1"
	res, err := db.Exec(q, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		http.Error(w, "Error updating stock", http.StatusInternalServerError)
		return
	}

	re, err := res.RowsAffected()
	ans := fmt.Sprintf("%v rows are affected", re)

	json.NewEncoder(w).Encode(ans)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 0, 0)
	if err != nil {
		http.Error(w, "Give Proper Id in path", http.StatusBadRequest)
		return
	}

	db := Connect_to_PostGres()
	defer db.Close()
	q := "DELETE FROM stocks where stockid=$1"
	res, err := db.Exec(q, id)
	if err != nil {
		http.Error(w, "Error deleting stock", http.StatusInternalServerError)
		return
	}

	re, err := res.RowsAffected()
	ans := fmt.Sprintf("%v rows are deleted", re)

	json.NewEncoder(w).Encode(ans)
}
