package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line1, _ := reader.ReadString(' ')
	fmt.Println(line1)
	newSource := strings.NewReader("new source")
	reader.Reset(newSource)
	line2, _ := reader.ReadString(' ')
	fmt.Println(line2)
}
