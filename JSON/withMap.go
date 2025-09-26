package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{"name" : "Ivan", "age" : 25, "skils" : ["Go", "Python", "JavaScript"]}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}
	fmt.Println("name:", data["name"])
	fmt.Println("age:", data["age"])
	fmt.Println("skils:", data["skils"])
}
