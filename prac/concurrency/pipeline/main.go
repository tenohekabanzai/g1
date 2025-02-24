package main

import "fmt"

func sliceToChannel(arr []int) chan int {

	go func() {

	}()
}

func main() {
	fmt.Println("Channels & Pipelines")

	nums = []int{2, 3, 4, 7, 1}

	ch1 := sliceToChannel(nums)
}
