package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var mtx = sync.RWMutex{}

func main() {
	// runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mtx.RLock()
		go sayHello()
		mtx.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello there #%v\n", counter)
	mtx.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	mtx.Unlock()
	wg.Done()
}
