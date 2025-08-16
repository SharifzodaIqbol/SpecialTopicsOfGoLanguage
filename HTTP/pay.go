package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Я успешно оплатил!"
	sliceBytes := []byte(str)
	_, err := w.Write(sliceBytes)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}
}
func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	sliceBytes := []byte(str)
	_, err := w.Write(sliceBytes)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello World!"
	sliceBytes := []byte(str)
	_, err := w.Write(sliceBytes)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно обработал запрос!")
	}
}
func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Сервер работает...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}
}
