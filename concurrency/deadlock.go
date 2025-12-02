package main

import (
	"sync"
	"time"
)

var mu1, mu2 sync.Mutex

func f1() {
	mu1.Lock()
	defer mu1.Unlock()
	mu2.Lock()
	defer mu2.Unlock()
}

func f2() {
	mu2.Lock()
	defer mu2.Unlock()
	mu1.Lock()
	defer mu1.Unlock()
}

func main() {
	go f1()
	go f2()
	time.Sleep(time.Second)

}
