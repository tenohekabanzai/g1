package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	go func() {
		c <- "hello from other thread"
	}()
	// no clumsy join required by sleeping the main thread , as channels handle all the synchronization
	msg := <-c
	fmt.Println(msg)
}
