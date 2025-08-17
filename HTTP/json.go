package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Addres      string `json:"addres"`
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Addres:", p.Addres)
}

type HttpResponse struct {
	Money          int       `json:"Money"`
	PaymentHistory []Payment `json:"PaymentHistory"`
}

var paymentHistory = make([]Payment, 0)
var mtx sync.Mutex
var money = 1000
var bank = 0

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := "fail to read HTTP body:" + err.Error()
			w.Write([]byte(msg))
			fmt.Println(msg)
			return
		}
		var payment Payment
		if err := json.Unmarshal(requestBody, &payment); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	*/

	payment.Println()
	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}
	paymentHistory = append(paymentHistory, payment)
	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}
	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("money:", money)
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)

	fmt.Println("Сервер работает...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Программа успешно завершена!")
}
