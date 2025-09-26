package main

import (
	"encoding/json"
	"fmt"
)

type Adress struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Adress Adress `json:"adress"`
}

func main() {
	jsonString := `{"name" : "Ivan", "age" : 25, "adress" : {"city" : "Moskow", "state" : "Russia"}}`
	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("city:", person.Adress.City)
	fmt.Println("state:", person.Adress.State)
}
