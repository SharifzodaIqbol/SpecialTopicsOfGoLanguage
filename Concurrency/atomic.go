package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// var number int = 0
var number atomic.Int64

func plusOne(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		number.Add(1)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go plusOne(wg)
	go plusOne(wg)
	go plusOne(wg)
	go plusOne(wg)
	go plusOne(wg)
	wg.Wait()
	fmt.Println("number =", number.Load())
}
