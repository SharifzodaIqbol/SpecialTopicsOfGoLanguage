package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	date map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		date: make(map[int]int),
	}
}

func (sm *SafeMap) Write(key, value int) {
	sm.Lock()
	defer sm.Unlock()
	sm.date[key] = value
}
func (sm *SafeMap) Read(key int) (int, bool) {
	sm.RLock()
	defer sm.RUnlock()
	value, exist := sm.date[key]
	return value, exist
}

func main() {
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}

	for i := 1; i <= 5000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeMap.Write(id, rand.Intn(100)+1)
		}(i)
	}
	wg.Wait()
	for i := 1; i <= 10; i++ {
		if value, exist := safeMap.Read(i); exist {
			fmt.Println(i, value)
		}
	}
}
