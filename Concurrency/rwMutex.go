package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 0
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100_000; i++ {
		mtx.Lock()
		count++
		mtx.Unlock()
	}
}
func getLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100_000; i++ {
		mtx.RLock()
		_ = count
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}
	wg.Wait()
	fmt.Println("Время выполнения программы:", time.Since(startTime))
}
