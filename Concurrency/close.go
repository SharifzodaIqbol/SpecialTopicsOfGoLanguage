package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channalInt := make(chan int)

	go func() {
		n := 3 + rand.Intn(4)
		fmt.Println("n:=", n)
		for i := 0; i < n; i++ {
			channalInt <- rand.Intn(20)
		}
		close(channalInt)
	}()
	/*
		for {
			read, ok := <-channalInt
			if !ok {
				break
			}
			fmt.Println(read)
		}
	*/
	for value := range channalInt {
		fmt.Println(value)
	}

}
