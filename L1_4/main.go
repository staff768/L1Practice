package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Используем паттер Done channel, после поулчения sinint закрываем doneCH и с помощью wg дожидаемся что все горутины закончат работу 

func DoSome(doneCH <- chan struct{}, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case <- doneCH:
			fmt.Println("ShutDown of gorutine")
			return
		default:
			fmt.Println("Doing something ... ")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	
	wg.Add(1)
	go DoSome(done,wg)

	
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	
	<-signalChan
	close(done)
	wg.Wait()
	
	fmt.Println("ShutDown of work...")
}