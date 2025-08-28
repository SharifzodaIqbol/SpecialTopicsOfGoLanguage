package main

import (
	"fmt"
	"sync"
)

func ping(pingCh <-chan string, pongCh chan<- string, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		msg := <-pingCh
		fmt.Println("Первый канал читает:", msg)
		pongCh <- fmt.Sprintf("Первый канал отправляет сообщение!")
	}
}
func pong(pongCh <-chan string, pingCh chan<- string, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		msg := <-pongCh
		fmt.Println("Второй канал читает:", msg)
		if i < n {
			pingCh <- fmt.Sprintf("Второй канал отправляет сообщение!")
		}
	}
}
func main() {
	const n = 10
	pingCh := make(chan string)
	pongCh := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go ping(pingCh, pongCh, &wg, n)
	go pong(pongCh, pingCh, &wg, n)
	pingCh <- "Начинаем игру!"
	wg.Wait()
	close(pingCh)
	close(pongCh)
}
