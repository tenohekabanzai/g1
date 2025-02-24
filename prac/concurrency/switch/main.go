package main

import "fmt"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "Hello1"
	}()

	go func() {
		c2 <- "Hello2"
	}()

	select {
	case a := <-c1:
		fmt.Println(a)
	case b := <-c2:
		fmt.Println(b)
	}

	fmt.Println("Hello_main")

}
