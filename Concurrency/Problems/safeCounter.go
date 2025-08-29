package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter atomic.Int32

func count(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add(1)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 100_000; i++ {
		wg.Add(1)
		go count(&wg)
	}
	wg.Wait()
	fmt.Println(counter.Load())
}

/*
import (
	"fmt"
	"sync"
)

var counter int
var mtx sync.Mutex

func count(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.Lock()
	counter++
	mtx.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 100_000; i++ {
		wg.Add(1)
		go count(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

*/
