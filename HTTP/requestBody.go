package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

/*
1 -> payHandler
2 -> saveHadler
*/

var mtx sync.Mutex
var money atomic.Int64
var bank atomic.Int64

func payHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err)
		return
	}

	value, err := strconv.Atoi(string(requestBody))

	if err != nil {
		fmt.Println("formatting error:", err)
		return
	}
	mtx.Lock()
	if money.Load()-int64(value) >= 0 {

		money.Add(int64(-value))

		fmt.Println("Ваш счет:", money.Load())
		fmt.Println("Ваше средство в копилке:", bank.Load())

	} else {
		fmt.Println("Недостаточно средств!")
		fmt.Println("Ваш счет:", money.Load())
	}
	mtx.Unlock()
}

func saveHadler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err)
		return
	}
	value, err := strconv.Atoi(string(requestBody))

	if err != nil {
		fmt.Println("formatting error:", err)
		return
	}
	mtx.Lock()
	if money.Load()-int64(value) >= 0 {

		money.Add(int64(-value))
		bank.Add(int64(value))

		fmt.Println("Ваше средство в копилке:", bank.Load())
		fmt.Println("Ваш счет:", money.Load())

	} else {
		fmt.Println("Недостаточно средств!")
		fmt.Println("Ваш счет:", money.Load())
	}
	mtx.Unlock()
}

func main() {
	money.Add(1000)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHadler)

	fmt.Println("Сервер работает...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Программа успешно завершена!")
}
