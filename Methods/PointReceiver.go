package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) changeName(newName string) {
	p.name = newName
}

func main() {
	var oldName Person
	oldName.name = "oldName"
	fmt.Println("Before:", oldName.name)
	oldName.changeName("newName")
	fmt.Println("After:", oldName.name)
}
