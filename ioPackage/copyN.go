package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some text")
	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}
}
