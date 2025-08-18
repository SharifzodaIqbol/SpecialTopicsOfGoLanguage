package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	addres := r.URL.Query().Get("addres")
	fmt.Println("name:", name)
	fmt.Println("addres:", addres)
}
func main() {
	http.HandleFunc("/default", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error:", err)
	}
}
