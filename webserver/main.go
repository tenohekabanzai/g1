package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	fileService := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileService)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/form", formController)
	fmt.Println("Server running out at PORT 5004")

	http.ListenAndServe(":5004", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func formController(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error in parsing Form")
		return
	}

	fmt.Fprint(w, "POST request Success", http.StatusAccepted)
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Println(name)
	fmt.Println(address)
}
