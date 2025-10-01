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
	buffer := make([]byte, 5)
	n, err := reader.Read(buffer)
	if err != nil || err == io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buffer[:n]))
}
