package main

import (
	"fmt"
	"mysql_go/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	r.HandleFunc("/home", serveHome)
	fmt.Println("Server running at PORT:5004")
	http.ListenAndServe(":5004", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Hello from MySQL RESTAPI</h1>`))
}
