package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		select {
		case <-ctx.Done():
			return
		case stream <- dataItem:

		}
	}

	infApples := make(chan interface{})
	go generator("apple", infApples)

	infOranges := make(chan interface{})
	go generator("orange", infOranges)

	infBananas := make(chan interface{})
	go generator("banana", infBananas)
}
