package main

import (
	"fmt"
	"time"
)

type Message struct {
	Authtor string
	Text    string
}

func main() {
	msg1 := make(chan Message)
	msg2 := make(chan Message)

	go func() {
		for {
			msg1 <- Message{
				Authtor: "Author 1",
				Text:    "Привет!",
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			msg2 <- Message{
				Authtor: "Author 2",
				Text:    "Как дела!",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		select {
		case messageChan1 := <-msg1:
			fmt.Println("Сообщение от", messageChan1.Authtor, ":", messageChan1.Text)
		case messageChan2 := <-msg2:
			fmt.Println("Сообщение от", messageChan2.Authtor, ":", messageChan2.Text)
		}
	}
}
