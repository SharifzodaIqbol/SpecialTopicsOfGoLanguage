package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := make(chan int)
	str := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		cnt <- 10
	}()
	go func() {
		str <- "hi"
	}()
	time.Sleep(50 * time.Millisecond)
	select {
	case digit := <-cnt:
		fmt.Println("digit:", digit)
	case s := <-str:
		fmt.Println("s:", s)
	default:
		fmt.Println("Никакой канал не готов")
	}
}
