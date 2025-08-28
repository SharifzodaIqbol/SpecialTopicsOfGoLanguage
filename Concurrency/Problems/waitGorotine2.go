package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter atomic.Int32

func plusOne(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)
}
func main() {
	var n int
	wg := sync.WaitGroup{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go plusOne(&wg)
	}
	wg.Wait()
	fmt.Println(counter.Load())
}
