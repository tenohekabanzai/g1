package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gen := func(str string, c chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- str:
				// push content inside channel
			}
		}
	}

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})

	go gen("1", c1) // a constant stream of 1s till main finishes
	go gen("2", c2) // a constant stream of 2s till main finishes
	go gen("3", c3) // a constant stream of 3s till main finishes

	wg.Add(1)
	go f1(ctx, &wg, c1)

	wg.Add(1)
	go f2(ctx, &wg, c2)

	wg.Add(1)
	go f2(ctx, &wg, c3)

	wg.Wait()

}

func f1(ctx context.Context, pwg *sync.WaitGroup, chn chan interface{}) {
	defer pwg.Done()
	var wg sync.WaitGroup
	doWork := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-chn:
				if !ok {
					fmt.Println("channel closed")
				}
				fmt.Println(d)
			}
		}
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doWork(newCtx)
	}
	wg.Wait()
}

func f2(ctx context.Context, pwg *sync.WaitGroup, chn chan interface{}) {
	defer pwg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-chn:
			if !ok {
				fmt.Println("channel closed")
			}
			fmt.Println(d)
		}
	}
}
