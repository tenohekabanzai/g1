package main

import "fmt"

func main() {

	// c := make(chan int)
	chars := []string{"a", "b", "c"}
	// buffered channels are
	c := make(chan string, 3)

	for _, val := range chars {
		c <- val
	}

	close(c)

	for res := range c {
		fmt.Println(res)
	}

}
