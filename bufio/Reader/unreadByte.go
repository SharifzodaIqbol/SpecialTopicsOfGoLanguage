package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	source := strings.NewReader("Hello")
	var delim byte = 'l'
	result := []byte{}
	for {
		char, err := source.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if char == delim {
			source.UnreadByte()
			break
		}
		result = append(result, char)
	}
	dataFile, err := io.ReadAll(source)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataFile))
}
