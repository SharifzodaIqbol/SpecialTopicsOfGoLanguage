package main

import "fmt"

type Men interface {
	SayHi()
	Sing(lyrics string)
}
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s\n", h.name, h.phone)
}
func (h Human) Sing(song string) {
	fmt.Println("La la la la...", song)
}
func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}
func main() {
	mike := Student{Human{name: "Mike", age: 25, phone: "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{name: "Paul", age: 19, phone: "333-333-XXX"}, "Politech", 0.00}
	sam := Employee{Human{name: "Sam", age: 25, phone: "444-444-XXX"}, "Golang Inc", 1000}
	tom := Employee{Human{name: "Tom", age: 27, phone: "555-222-XXX"}, "Things Ltd.", 10000}

	var i Men
	// Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// Employee
	i = tom
	fmt.Println("This is Tom,an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")

	x := make([]Men, 3)

	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

}
