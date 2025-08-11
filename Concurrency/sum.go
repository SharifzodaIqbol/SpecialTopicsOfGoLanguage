package main

import (
	"fmt"
)

func sum(nums []int, ch chan int) {
	result := 0
	for _, value := range nums {
		result += value
	}
	ch <- result
}
func main() {
	nums := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(nums[:len(nums)/2], ch)
	go sum(nums[len(nums)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

}
