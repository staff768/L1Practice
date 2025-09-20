package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func worker(ch <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d получил данные: %d\n", id, data)
	}
}


func main(){
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Fatalln(err)
	}
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ch, wg, i)
	}

	for i := 0; ; i++ {
		ch <- i
	}
}