package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	N := 3
	timeout := time.After(time.Duration(N) * time.Second)

	dataChan := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-timeout:
				close(dataChan)
				return
			default:
				dataChan <- i
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for value := range dataChan {
		elapsed := time.Since(start)
		fmt.Printf("Прочитано - %d (Время -  %v)\n", value, elapsed.Truncate(time.Millisecond))
	}

	fmt.Printf("Конец. Общее время работы: %v\n", time.Since(start).Truncate(time.Millisecond))
}
