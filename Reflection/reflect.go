package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	Comp := Company{ID: 1, Name: "Ozon"}
	t := reflect.TypeOf(Comp)
	v := reflect.ValueOf(Comp)
	fmt.Println(t.Name())
	var str []string
	fmt.Println(t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("Тип:", t.Field(i).Tag.Get("json"))
		fmt.Println("Val:", v.Field(i))
		str = append(str, t.Field(i).Tag.Get("json"))
	}
	fmt.Println(strings.Join(str, ", "))
}
