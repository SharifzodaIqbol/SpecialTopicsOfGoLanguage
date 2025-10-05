package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func splitByComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
func main() {
	scanner := bufio.NewScanner(strings.NewReader("Hello,my name is Mat!"))
	scanner.Split(splitByComma)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
