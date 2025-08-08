package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "Name: " + h.name + " Age: " + strconv.Itoa(h.age) + " years, Contact: " + h.phone
}
func main() {
	Bob := Human{"Bob", 39, "000-777-XXX"}
	fmt.Println("This human is:", Bob)
}
