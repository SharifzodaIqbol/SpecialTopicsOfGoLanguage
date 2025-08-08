package main

import "fmt"

func main() {
	var a interface{}
	var i int = 5
	s := "Hello World!"

	a = i
	fmt.Println("a =", a)

	a = s

	fmt.Println("a =", a)
}
