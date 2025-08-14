package main

import (
	"fmt"
	"sync"
)

var slice []int
var mtx sync.Mutex

func push_back(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go push_back(wg)
	go push_back(wg)
	go push_back(wg)
	go push_back(wg)
	go push_back(wg)
	wg.Wait()
	fmt.Println("silce len =", len(slice))
}
