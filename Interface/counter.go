package main

import "fmt"

type Counter interface {
	Add(increment int)
	Value() int
}
type Stats struct {
	value int
}
type AdvanceStats struct {
	value int
}

func (s *Stats) Add(data int) {
	s.value += data
}
func (s Stats) Value() int {
	return s.value
}
func (a *AdvanceStats) Add(data int) {
	a.value += data * 2
}
func (a AdvanceStats) Value() int {
	return a.value
}
func ProccesCounter(c Counter) {
	c.Add(10)
	fmt.Printf("%T Value: %d\n", c, c.Value())
}
func main() {
	stats := &Stats{value: 18}
	advance := &AdvanceStats{value: 20}
	ProccesCounter(stats)
	ProccesCounter(advance)
}
