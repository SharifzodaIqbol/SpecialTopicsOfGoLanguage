package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	data := "Hello, world!\nThis is a ReadFrom example."
	reader := strings.NewReader(data)

	n, err := writer.ReadFrom(reader)
	if err != nil {
		fmt.Println("ReadFrom Error:", err)
		return
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Flush Error:", err)
		return
	}

	fmt.Println("Bytes written:", n)
	fmt.Println("Buffer contents:", buf.String())
}
