package main

import (
	"fmt"
	"math/rand"
)

func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(10) + 1
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for num := range ch {
		fmt.Println(num)
	}
}
