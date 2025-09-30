package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 16)
	fmt.Printf("Изначально в буфере: %d байт\n", reader.Buffered())

	data, _ := reader.Peek(7)
	fmt.Printf("После Peek(7) в буфере: %d байт\n", reader.Buffered())
	fmt.Printf("Peek данные: %s\n", string(data))

	line, _ := reader.ReadString('\n')
	fmt.Printf("После ReadString в буфере: %d байт\n", reader.Buffered())
	fmt.Printf("Прочитанная строка: %s", line)
}
