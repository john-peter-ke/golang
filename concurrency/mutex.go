package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	mu.Lock()
	count++
	mu.Unlock()
	// Simulate work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Counter:", count)
}
