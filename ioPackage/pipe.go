package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()
	buff := make([]byte, 20)
	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()
	if _, err := io.CopyBuffer(os.Stdout, r, buff); err != nil {
		log.Fatal(err)
	}
}
