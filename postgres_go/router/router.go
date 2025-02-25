package router

import (
	"postgres_go/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/stock/{id}", controller.GetStock).Methods("GET")
	r.HandleFunc("/api/stock", controller.GetAllStocks).Methods("GET")
	r.HandleFunc("/api/newstock", controller.AddStock).Methods("POST")
	r.HandleFunc("/api/stock/{id}", controller.UpdateStock).Methods("PUT")
	r.HandleFunc("/api/stock/{id}", controller.DeleteStock).Methods("DELETE")
	return r
}
