package main

import (
	"fmt"
	"net/http"
	"postgres_go/middleware"
	"postgres_go/router"
)

func main() {
	r := router.Router()
	fmt.Println("server running on port 5004")
	middleware.Connect_to_PostGres()
	http.ListenAndServe(":5004", r)
}
