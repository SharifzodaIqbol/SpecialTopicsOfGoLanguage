package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	source := strings.NewReader("Hello, World!")
	buffer := bytes.NewBuffer([]byte("Data from buffer"))

	n1, err := source.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nWritten %d bytes\n", n1)

	var buf bytes.Buffer
	n2, err := buffer.WriteTo(&buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written %d bytes to buffer: %s\n", n2, buf.String())
}
