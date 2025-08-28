package main

import (
	"fmt"
	"sync"
)

func print(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("Gorotine %d is Done!\n", num)
}
func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go print(&wg, i+1)
	}
	wg.Wait()
}
