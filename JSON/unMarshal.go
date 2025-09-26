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
	jsonString := `{"name" : "Ivan", "age" : 25, "email" : "ivan@example.com"}`
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}
	fmt.Println("Имя:", user.Name)
	fmt.Println("Возраст:", user.Age)
	fmt.Println("Email:", user.Email)
}
