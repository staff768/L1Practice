package main

import (
	"fmt"
)

func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, x := range nums {
			out <- x
		}
		close(out)
	}()
	return out
}

func doubler(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x * 2
		}
		close(out)
	}()
	return out
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	final := doubler(generator(data))

	for v := range final {
		fmt.Println(v)
	}
}
