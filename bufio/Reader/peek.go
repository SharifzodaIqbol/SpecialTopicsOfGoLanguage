package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		char, err := reader.Peek(2)
		if err != nil || string(char) == "ll" {
			break
		}
		reader.Discard(2)
	}
	str, err := io.ReadAll(reader)
	fmt.Println(string(str))
}
