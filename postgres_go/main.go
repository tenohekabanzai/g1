package main

import (
	"fmt"
	"net/http"
	"postgres_go/router"
)

func main() {
	r := router.Router()
	fmt.Println("server running on port 5004")
	http.ListenAndServe(":5004", r)
}
