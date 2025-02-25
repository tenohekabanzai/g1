package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	done := make(chan interface{})
	defer close(done)

	dogs := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case dogs <- "woof":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumedogs(done, dogs)
	wg.Add(1)
	go consumepigs(done, pigs)

	wg.Wait()
}

func consumedogs(done chan interface{}, d chan interface{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case dog, ok := <-d:
			if !ok {
				fmt.Println("Channel Closed")
			}
			fmt.Println(dog)
		}
	}
}

func consumepigs(done chan interface{}, p chan interface{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case dog, ok := <-p:
			if !ok {
				fmt.Println("Channel Closed")
			}
			fmt.Println(dog)
		}
	}
}

func ordone(done, c chan interface{}) chan interface{} {
	c2 := make(chan interface{})
	go func() {
		defer close(c2)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					fmt.Println("channel closed")
				}
				fmt.Println(v)
			}
		}
	}()

	return c2
}
