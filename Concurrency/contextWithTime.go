package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Операция выполнена!")
	case <-ctx.Done():
		fmt.Println("Контекст завершен!", ctx.Err())
	}
}
