package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

/*
1 -> payHandler
2 -> saveHadler
*/

var mtx sync.Mutex
var money = 1000
var bank = 0

func payHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP body:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	value, err := strconv.Atoi(string(requestBody))
	if err != nil {
		msg := "formatting error:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	mtx.Lock()
	if money-value >= 0 {
		money -= value
		msg := "Ваш счет:" + strconv.Itoa(money)
		w.Write([]byte(msg))
		fmt.Println(msg)
		fmt.Println("Ваше средство в копилке:", bank)

	} else {
		msg := "Недостаточно средств:" + strconv.Itoa(money)
		w.Write([]byte(msg))
		fmt.Println(msg)
		fmt.Println("Ваш счет:", money)
	}
	mtx.Unlock()
}

func saveHadler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP body" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	value, err := strconv.Atoi(string(requestBody))

	if err != nil {
		msg := "formatting error:" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	mtx.Lock()
	if money-value >= 0 {

		money -= value
		bank += value

		msg := "Ваше средство в копилке:" + strconv.Itoa(bank) + "\nВаш счет:" + strconv.Itoa(money)
		w.Write([]byte(msg))
		fmt.Println(msg)

	} else {
		msg := "Недостаточно средств!\nВаш счет:" + strconv.Itoa(money)
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHadler)

	fmt.Println("Сервер работает...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Программа успешно завершена!")
}
