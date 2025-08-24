package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic Second:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic First:", err)
			//panic("second error")
		}
	}()
	fmt.Println("some useful work")
	panic("first error")
}

func main() {
	deferTest()
	return
}
