package main

import "time"

type RateLimiter struct {
	tokens chan bool
}

func NewRateLimiter(x int) *RateLimiter {
	return &RateLimiter{
		tokens: make(chan bool, x),
	}
}
func (r *RateLimiter) startRefill() {
	ticker := time.NewTicker(200 * time.Millisecond)

	for range ticker.C {
		select {
		case r.tokens <- true:
		default:
		}
	}
}
func (r *RateLimiter) Wait() {
	<-r.tokens
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func main() {
	limiter := NewRateLimiter(5) // 5 вызовов в секунду

	for i := 0; i < 10; i++ {
		go func(id int) {
			limiter.Wait()
			doSomething(id)
		}(i)
	}
}
