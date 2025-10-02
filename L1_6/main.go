package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 1) Остановка по условию (флаг)
func ConditionStop() {
	fmt.Println("-- ConditionStop")
	var stopFlag atomic.Bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if stopFlag.Load() {
				fmt.Println("goroutine: stop by condition")
				return
			}
			fmt.Println("goroutine: working...")
			time.Sleep(150 * time.Millisecond)
		}
	}()
	time.Sleep(400 * time.Millisecond)
	stopFlag.Store(true)
	wg.Wait()
}

// 2) Остановка через канал уведомления (done)
func DoneChannelStop() {
	fmt.Println("-- DoneChannelStop")
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("goroutine: stop via done channel")
				return
			default:
				fmt.Println("goroutine: working...")
				time.Sleep(120 * time.Millisecond)
			}
		}
	}()
	// Посылаем сигнал останова
	time.Sleep(360 * time.Millisecond)
	close(done)
	wg.Wait()
}

// 3) Остановка через контекст с ручной отменой
func ContextCancelStop() {
	fmt.Println("-- ContextCancelStop")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine: stop via context cancel")
				return
			default:
				fmt.Println("goroutine: working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	// Ручная отмена
	time.Sleep(300 * time.Millisecond)
	cancel()
	wg.Wait()
}

// 4) Остановка через контекст с тайм-аутом/дедлайном
func ContextTimeoutStop() {
	fmt.Println("-- ContextTimeoutStop")
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine: stop via context timeout")
				return
			default:
				fmt.Println("goroutine: working...")
				time.Sleep(90 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}

// 5) Остановка чтением закрытого рабочего канала
func CloseWorkChannelStop() {
	fmt.Println("-- CloseWorkChannelStop")
	work := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range work {
			fmt.Printf("goroutine: got %d\n", v)
		}
		fmt.Println("goroutine: channel closed -> stop")
	}()

	go func() {
		for i := 0; i < 3; i++ {
			work <- i
			time.Sleep(70 * time.Millisecond)
		}
		close(work)
	}()

	wg.Wait()
}

// 6) Принудительное завершение текущей горутины runtime.Goexit()
func RuntimeGoexitStop() {
	fmt.Println("-- RuntimeGoexitStop")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine: doing some work")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("goroutine: calling runtime.Goexit()")
		runtime.Goexit()
	}()
	wg.Wait()
}

func main() {
	ConditionStop()
	DoneChannelStop()
	ContextCancelStop()
	ContextTimeoutStop()
	CloseWorkChannelStop()
	RuntimeGoexitStop()
}
