package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d. %s\n", i, name)
	}
	time.Sleep(50 * time.Millisecond)
}

func main() {
	wg := &sync.WaitGroup{}
	alex := "Alex"
	tom := "Tom"
	adam := "Adam"
	wg.Add(1)
	go myFunc(wg, alex)
	wg.Add(1)
	go myFunc(wg, tom)
	wg.Add(1)
	go myFunc(wg, adam)
	wg.Wait()
	fmt.Println("main завершился!")
}
