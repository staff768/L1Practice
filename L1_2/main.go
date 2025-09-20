package main

import (
	"fmt"
	
)

func writer(arr []int) <- chan int {
	ch1 := make(chan int)
	go func() {
		for _,v := range arr {
			ch1 <- v
		}
		close(ch1)
	}()
	return ch1
}
func square(in <-chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		for value :=  range in{
			ch2 <- value * value
		}
		close(ch2)
	}()
	return ch2
}

func reader(in <-chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}


func main(){
	arr := []int{2,4,6,8,10}
	reader(square(writer(arr)))
	
}