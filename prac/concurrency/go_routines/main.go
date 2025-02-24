package main

import (
	"fmt"
	"time"
)

func f(n string) {
	fmt.Println(n)
}

func main() {
	go f("Hello1") // fork
	go f("Hello2") // fork
	go f("Hello3") // fork

	// main thread delays exec for 2 secs to allow the other go routines to complete executions
	time.Sleep(2 * time.Second) // all go routines are joined after this point
	fmt.Println("hi")
}
