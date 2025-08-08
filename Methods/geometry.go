package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Обычная функция
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.Y, q.Y-p.Y)
}

// то же самое, но как метод типа Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.Y, q.Y-p.Y)
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 4.47213595499958 вызов функции
	fmt.Println(p.Distance(q))  // 4.47213595499958 вызов метода
}
