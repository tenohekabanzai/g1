package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Wait Groups")

	wl := []string{"https://go.dev", "https://google.com", "https://github.com"}

	fmt.Println("Main Started")

	wg.Add(len(wl))
	for _, val := range wl {
		go getStatusCode(val, &wg)
	}
	wg.Wait()

	fmt.Println("Main Finished Execution")
}

func getStatusCode(endpoint string, wg *sync.WaitGroup) {

	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
