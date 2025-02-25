package router

import (
	"postgres_go/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET")
	r.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET")
	r.HandleFunc("/api/newstock", middleware.AddStock).Methods("POST")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT")
	r.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE")
	return r
}
