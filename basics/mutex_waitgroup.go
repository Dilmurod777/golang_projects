package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	i = i + 1
	mu.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		worker(wg, mu)
	}

	wg.Wait()
	fmt.Printf("i is %d\n", i)
}
