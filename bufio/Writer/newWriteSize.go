package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "Hello world!")
	fmt.Fprintln(w, w.Available()-1)
	w.Flush()

	fmt.Println("\n--- С кастомным размером ---")
	x := bufio.NewWriterSize(os.Stdout, 5)
	fmt.Fprintln(x, "Hello world!")
	x.Flush()
}
