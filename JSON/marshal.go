package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Ivan", Age: 30, Email: "ivan@example.com"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Ошибка при сериализации:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))
}
