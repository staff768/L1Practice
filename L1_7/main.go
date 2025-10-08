package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeMap struct {
	mu  sync.Mutex
	set map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{set: make(map[string]int)}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	s.set[key] = value
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()
	v, ok := s.set[key]
	s.mu.Unlock()
	return v, ok
}

func demoMutexMap() {
	s := NewSafeMap()
	var wg sync.WaitGroup
	workers := 8
	perWorker := 1000

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
			for j := 0; j < perWorker; j++ {
				k := fmt.Sprintf("k_%d_%d", id, j)
				s.Set(k, r.Intn(1_000_000))
			}
		}(i)
	}

	wg.Wait()

	if v, ok := s.Get("k_0_0"); ok {
		fmt.Println("mutex map sample:", v)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	demoMutexMap()
	fmt.Println("done")
}
