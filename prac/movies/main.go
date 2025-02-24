package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	fmt.Println("Server running on PORT 5004")
	http.ListenAndServe(":5004", r)

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello From Movies API</h1>"))
}
